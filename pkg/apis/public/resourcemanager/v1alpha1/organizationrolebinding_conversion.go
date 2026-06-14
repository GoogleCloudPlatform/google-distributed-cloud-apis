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

// ConvertTo converts this OrganizationRoleBinding to the Hub version (v1).
func (orb *OrganizationRoleBinding) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.OrganizationRoleBinding)

	dst.ObjectMeta = orb.ObjectMeta
	dst.Spec = rmv1.OrganizationRoleBindingSpec(orb.Spec)

	dst.Status.PropagatedName = orb.Status.PropagatedName
	dst.Status.ErrorStatus = orb.Status.ErrorStatus
	dst.Status.Conditions = orb.Status.Conditions
	dst.Status.Clusters = ConvertClusterStatusToV1(orb.Status.Clusters)
	return nil
}

// ConvertFrom converts this OrganizationRoleBinding from the Hub version (v1) to this version.
func (orb *OrganizationRoleBinding) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.OrganizationRoleBinding)

	orb.ObjectMeta = src.ObjectMeta
	orb.Spec = OrganizationRoleBindingSpec(src.Spec)

	orb.Status.PropagatedName = src.Status.PropagatedName
	orb.Status.ErrorStatus = src.Status.ErrorStatus
	orb.Status.Conditions = src.Status.Conditions
	orb.Status.Clusters = ConvertClusterStatusFromV1(src.Status.Clusters)
	return nil
}
