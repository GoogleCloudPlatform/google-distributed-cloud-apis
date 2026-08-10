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

	rmv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/global/resourcemanager/v1"
)

// ConvertTo converts this ProjectReplica to the Hub version (v1).
func (p *ProjectReplica) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.ProjectReplica)

	dst.ObjectMeta = *p.ObjectMeta.DeepCopy()

	dst.Status.Conditions = p.Status.Conditions
	dst.Status.AvailableClusters = p.Status.AvailableClusters
	dst.Status.ErrorStatus = p.Status.ErrorStatus

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (p *ProjectReplica) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.ProjectReplica)

	p.ObjectMeta = *src.ObjectMeta.DeepCopy()

	p.Status.Conditions = src.Status.Conditions
	p.Status.AvailableClusters = src.Status.AvailableClusters
	p.Status.ErrorStatus = src.Status.ErrorStatus
	return nil
}
