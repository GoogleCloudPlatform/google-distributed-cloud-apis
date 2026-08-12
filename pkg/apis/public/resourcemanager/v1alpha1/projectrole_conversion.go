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

package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1"
)

// ConvertTo converts this ProjectRole to the Hub version (v1).
func (pr *ProjectRole) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.ProjectRole)

	dst.ObjectMeta = pr.ObjectMeta
	dst.Spec = rmv1.ProjectRoleSpec(pr.Spec)

	dst.Status.PropagatedName = pr.Status.PropagatedName
	dst.Status.ErrorStatus = pr.Status.ErrorStatus
	dst.Status.Conditions = pr.Status.Conditions
	dst.Status.Clusters = ConvertClusterStatusToV1(pr.Status.Clusters)
	return nil
}

// ConvertFrom converts this ProjectRole from the Hub version (v1) to this version.
func (pr *ProjectRole) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.ProjectRole)

	pr.ObjectMeta = src.ObjectMeta
	pr.Spec = ProjectRoleSpec(src.Spec)

	pr.Status.PropagatedName = src.Status.PropagatedName
	pr.Status.ErrorStatus = src.Status.ErrorStatus
	pr.Status.Conditions = src.Status.Conditions
	pr.Status.Clusters = ConvertClusterStatusFromV1(src.Status.Clusters)
	return nil
}
