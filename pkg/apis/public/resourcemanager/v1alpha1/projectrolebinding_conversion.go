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

	rmv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/resourcemanager/v1"
)

// ConvertTo converts this ProjectRoleBinding to the Hub version (v1).
func (prb *ProjectRoleBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.ProjectRoleBinding)

	dst.ObjectMeta = prb.ObjectMeta
	dst.Spec = rmv1.ProjectRoleBindingSpec(prb.Spec)

	dst.Status.PropagatedName = prb.Status.PropagatedName
	dst.Status.ErrorStatus = prb.Status.ErrorStatus
	dst.Status.Conditions = prb.Status.Conditions
	dst.Status.Clusters = ConvertClusterStatusToV1(prb.Status.Clusters)

	return nil
}

// ConvertFrom converts this ProjectRoleBinding from the Hub version (v1) to this version.
func (prb *ProjectRoleBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.ProjectRoleBinding)

	prb.ObjectMeta = src.ObjectMeta
	prb.Spec = ProjectRoleBindingSpec(src.Spec)

	prb.Status.PropagatedName = src.Status.PropagatedName
	prb.Status.ErrorStatus = src.Status.ErrorStatus
	prb.Status.Conditions = src.Status.Conditions
	prb.Status.Clusters = ConvertClusterStatusFromV1(src.Status.Clusters)
	return nil
}
