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

// ConvertTo converts this OrganizationRole to the Hub version (v1).
func (or *OrganizationRole) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.OrganizationRole)

	dst.ObjectMeta = or.ObjectMeta
	dst.Spec = rmv1.OrganizationRoleSpec(or.Spec)

	dst.Status.PropagatedName = or.Status.PropagatedName
	dst.Status.ErrorStatus = or.Status.ErrorStatus
	dst.Status.Conditions = or.Status.Conditions
	dst.Status.Clusters = ConvertClusterStatusToV1(or.Status.Clusters)

	return nil
}

// ConvertFrom converts this OrganizationRole from the Hub version (v1) to this version.
func (or *OrganizationRole) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.OrganizationRole)

	or.ObjectMeta = src.ObjectMeta
	or.Spec = OrganizationRoleSpec(src.Spec)

	or.Status.PropagatedName = src.Status.PropagatedName
	or.Status.ErrorStatus = src.Status.ErrorStatus
	or.Status.Conditions = src.Status.Conditions
	or.Status.Clusters = ConvertClusterStatusFromV1(src.Status.Clusters)
	return nil
}
