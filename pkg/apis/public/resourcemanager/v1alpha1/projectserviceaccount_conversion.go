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

// ConvertTo converts this ProjectServiceAccount to the Hub version (v1).
func (psa *ProjectServiceAccount) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.ProjectServiceAccount)

	dst.ObjectMeta = psa.ObjectMeta

	if psa.Spec.Keys != nil {
		dst.Spec.Keys = []rmv1.ProjectServiceAccountKey{}
		for _, k := range psa.Spec.Keys {
			cpy := rmv1.ProjectServiceAccountKey{
				Algorithm:   rmv1.ProjectServiceAccountKeyAlgorithm(k.Algorithm),
				ID:          k.ID,
				Key:         k.Key,
				ValidBefore: k.ValidBefore,
				ValidAfter:  k.ValidAfter,
			}
			dst.Spec.Keys = append(dst.Spec.Keys, cpy)
		}
	}

	dst.Status.PropagatedName = psa.Status.PropagatedName
	dst.Status.Conditions = psa.Status.Conditions
	dst.Status.Clusters = ConvertClusterStatusToV1(psa.Status.Clusters)

	return nil
}

// ConvertFrom converts this ProjectServiceAccount from the Hub version (v1) to this version.
func (psa *ProjectServiceAccount) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.ProjectServiceAccount)

	psa.ObjectMeta = src.ObjectMeta
	if src.Spec.Keys != nil {
		psa.Spec.Keys = []ProjectServiceAccountKey{}
		for _, k := range src.Spec.Keys {
			cpy := ProjectServiceAccountKey{
				Algorithm:   ProjectServiceAccountKeyAlgorithm(k.Algorithm),
				ID:          k.ID,
				Key:         k.Key,
				ValidBefore: k.ValidBefore,
				ValidAfter:  k.ValidAfter,
			}
			psa.Spec.Keys = append(psa.Spec.Keys, cpy)
		}
	}

	psa.Status.PropagatedName = src.Status.PropagatedName
	psa.Status.Conditions = src.Status.Conditions
	psa.Status.Clusters = ConvertClusterStatusFromV1(src.Status.Clusters)
	return nil
}
