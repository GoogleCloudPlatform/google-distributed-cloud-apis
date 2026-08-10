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
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

const (
	Group   = "backup.gdc.goog"
	Version = "v1"
)

var (
	SchemeGroupVersion = schema.GroupVersion{
		Group:   Group,
		Version: Version,
	}

	GroupVersion = SchemeGroupVersion

	baseSchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	allSchemeBuilder  = &scheme.Builder{GroupVersion: SchemeGroupVersion}
	SchemeBuilder     = &dualBuilder{filtered: baseSchemeBuilder, all: allSchemeBuilder}

	AddToScheme    = allSchemeBuilder.AddToScheme
	AddToSchemeAll = allSchemeBuilder.AddToScheme
)

// Resource takes an unqualified resource and returns a Group qualified
// GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

type dualBuilder struct {
	filtered *scheme.Builder
	all      *scheme.Builder
}

func (d *dualBuilder) Register(objs ...runtime.Object) *dualBuilder {
	d.all.Register(objs...)
	var filtered []runtime.Object
	for _, obj := range objs {
		t := reflect.TypeOf(obj)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		name := t.Name()
		if isConflictingType(name) {
			continue
		}
		filtered = append(filtered, obj)
	}
	if len(filtered) > 0 {
		d.filtered.Register(filtered...)
	}
	return d
}

func (d *dualBuilder) AddToScheme(s *runtime.Scheme) error {
	return d.all.AddToScheme(s)
}

// AddNonConflictingToScheme registers only non-conflicting types into the scheme.
// This is used by controllers in Phase 1 that need to register both vendor and public backup schemas.
func AddNonConflictingToScheme(s *runtime.Scheme) error {
	return baseSchemeBuilder.AddToScheme(s)
}

func isConflictingType(name string) bool {
	switch name {
	case "Backup", "BackupList",
		"BackupPlan", "BackupPlanList",
		"BackupRepository", "BackupRepositoryList",
		"DeleteBackupRequest", "DeleteBackupRequestList",
		"ManualBackupRequest", "ManualBackupRequestList",
		"ManualRestoreRequest", "ManualRestoreRequestList",
		"ObjectStorageBackupConfig",
		"Restore", "RestoreList",
		"RestorePlan", "RestorePlanList",
		"VMBackupJob", "VMBackupJobList",
		"VolumeBackup", "VolumeBackupList",
		"VolumeRestore", "VolumeRestoreList":
		return true
	}
	return false
}
