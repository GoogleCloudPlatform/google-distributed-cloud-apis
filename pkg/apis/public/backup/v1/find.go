// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	objectv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/object/v1"
)

// GetBackupRepositoryName will return the name of the backup repository to which the provided object
// belongs and a boolean describing if the name was successfully determined. There are certain import
// scenarios where it will not be possible to find the backup repository name for the object, so a return
// value of false does not necessarily constitute an error.
func GetBackupRepositoryName(ctx context.Context, object client.Object, c client.Reader) (string, bool) {
	if object == nil {
		return "", false
	}

	switch obj := object.(type) {
	case *Backup:
		return obj.Spec.BackupConfig.BackupRepository, true
	case *BackupPlan:
		return obj.Spec.BackupConfig.BackupRepository, true
	case *VolumeBackup:
		template := &Backup{}
		err := getObject(ctx, c, template, obj.Spec.BackupName, obj.GetNamespace())
		if err != nil {
			return "", false
		}

		// Determine the backup repository using the "backup" parent object
		return GetBackupRepositoryName(ctx, template, c)
	case *Restore:
		return getRestoreBackupRepositoryName(ctx, obj, c)
	case *RestorePlan:
		return getRestoreBackupRepositoryName(ctx, obj, c)
	case *VolumeRestore:
		template := &Restore{}
		err := getObject(ctx, c, template, obj.Spec.RestoreName, obj.GetNamespace())
		if err != nil {
			return "", false
		}
		// Determine the backup repository using the "restore" parent object
		return GetBackupRepositoryName(ctx, template, c)
	case *DeleteBackupRequest:
		template := &Backup{}
		err := getObject(ctx, c, template, obj.Spec.BackupName, obj.GetNamespace())
		if err != nil {
			return "", false
		}
		// Determine the backup repository using the "backup" parent object
		return GetBackupRepositoryName(ctx, template, c)
	case *ManualBackupRequest:
		template := &BackupPlan{}
		// Determine the backup repository using the "backup plan" parent object
		err := getObject(ctx, c, template, obj.Spec.BackupPlanName, obj.GetNamespace())
		if err != nil {
			return "", false
		}

		return GetBackupRepositoryName(ctx, template, c)
	case *VMBackupJob:
		template := &Backup{}
		// Determine the backup repository using the "backup" parent object
		err := getObject(ctx, c, template, obj.Spec.BackupRef.Name, obj.GetNamespace())
		if err != nil {
			return "", false
		}
		return GetBackupRepositoryName(ctx, template, c)
	default:
		return "", false
	}
}

// getRestoreBackupRepositoryName determines the backup repository for a restore/restore plan via the following logic
//  1. If a backup repository is specified in the restore config then that is returned.
//  2. If the backup/backup plan that we are performing the restore on points to a read/write repository in the current cluster
//     we will return its name.
//  3. If the backup/backup plan that we are performing a restore on points to a read/only repository we will pick any
//     read/write repository from the cluster and use that.
//  4. If no read/write repositories exist then an empty string is returned.
func getRestoreBackupRepositoryName(ctx context.Context, obj client.Object, c client.Reader) (string, bool) {
	if obj == nil {
		return "", false
	}

	logger := log.FromContext(ctx)
	var name string
	var backupObj client.Object

	switch v := obj.(type) {
	case *RestorePlan:
		if v.Spec.RestoreConfig.BackupRepository != "" {
			return v.Spec.RestoreConfig.BackupRepository, true
		}
		name = v.Spec.BackupPlanName
		backupObj = &BackupPlan{}
	case *Restore:
		if v.Spec.RestoreConfig.BackupRepository != "" {
			return v.Spec.RestoreConfig.BackupRepository, true
		}
		backupObj = &Backup{}
		name = v.Spec.BackupName
	default:
		return "", false
	}

	err := c.Get(ctx, types.NamespacedName{Namespace: obj.GetNamespace(), Name: name}, backupObj)
	if err != nil {
		logger.Error(err, "Unable to retrieve associated backup object")
		return "", false
	}

	br, err := FindBackupRepository(ctx, backupObj, c)
	if err != nil {
		return "", false
	}
	if br.Spec.ImportPolicy == ReadWrite {
		return br.Name, true
	}

	// TODO(b/223866400): Enhance this logic so that we tag a certain RW repo as "primary" and always select that.
	brList := &BackupRepositoryList{}
	if err := c.List(ctx, brList); err != nil {
		logger.Error(err, "Could not fetch backup repository list.")
		return "", false
	}
	for _, repo := range brList.Items {
		if repo.Spec.ImportPolicy == ReadWrite {
			return repo.Name, true
		}
	}
	return "", false
}

// FindBackupRepository will return the backup repository that is linked to the provided object. If no backup repository
// is found then an error will be returned.
func FindBackupRepository(ctx context.Context, object client.Object, c client.Reader) (*BackupRepository, error) {
	backupRepository := &BackupRepository{}
	backupRepoName, found := GetBackupRepositoryName(ctx, object, c)
	if !found || len(backupRepoName) == 0 {
		return backupRepository, fmt.Errorf("the backup repository name could not be found for resource '%s'", object.GetName())
	}
	key := client.ObjectKey{Name: backupRepoName}
	err := c.Get(ctx, key, backupRepository)
	if err != nil {
		if apiStatus, ok := err.(apierrors.APIStatus); ok {
			return backupRepository, &BackupRepositoryNotFoundError{Name: backupRepoName, ErrorStatus: apiStatus.Status()}
		}
		return backupRepository, err
	}
	return backupRepository, nil
}

// FindSecret will return the secret that is linked to the provided BackupRepository.
func FindSecret(ctx context.Context, r *BackupRepository, c client.Client) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := c.Get(ctx, types.NamespacedName{Name: r.Spec.SecretReference.Name, Namespace: r.Spec.SecretReference.Namespace}, secret)
	return secret, err
}

// FindBucket will return the bucket that is linked to the provided BackupRepository.
func FindBucket(ctx context.Context, r *BackupRepository, c client.Client) *objectv1.Bucket {
	logger := log.FromContext(ctx)
	namespace := r.Spec.SecretReference.Namespace
	bucketFQDN := r.Spec.S3Options.Bucket

	bucketList := &objectv1.BucketList{}
	err := c.List(ctx, bucketList, client.InNamespace(namespace))
	if err != nil {
		logger.Error(err, "failed to list buckets", "namespace", namespace)
		return nil
	}

	for _, b := range bucketList.Items {
		if b.Status.FullyQualifiedName == bucketFQDN {
			return &b
		}
	}

	logger.V(1).Info("Bucket not found in local namespace; it might be external", "bucketFQDN", bucketFQDN, "namespace", namespace)
	return nil
}

func getObject(ctx context.Context, c client.Reader, object client.Object, name, namespace string) error {
	key := client.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}

	return c.Get(ctx, key, object)
}
