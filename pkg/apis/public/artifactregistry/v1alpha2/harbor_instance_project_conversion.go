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

package v1alpha2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	arv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/artifactregistry/v1"
)

// convertTo converts this HarborInstanceProject (v1alpha2) to the Hub version (v1).
func (in *HarborInstanceProject) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*arv1.HarborInstanceProject)

	dst.ObjectMeta = in.ObjectMeta
	dst.Spec.ProjectName = in.Spec.ProjectName
	dst.Spec.HarborInstanceRef = in.Spec.HarborInstanceRef
	dst.Spec.CreatorSubject = in.Spec.CreatorSubject

	dst.Status.Conditions = in.Status.Conditions

	return nil
}

// convertFrom converts HarborInstanceProject from the Hub version (v1) to this (v1alpha2) version.
func (in *HarborInstanceProject) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*arv1.HarborInstanceProject)

	in.ObjectMeta = src.ObjectMeta
	in.Spec.ProjectName = src.Spec.ProjectName
	in.Spec.HarborInstanceRef = src.Spec.HarborInstanceRef
	in.Spec.CreatorSubject = src.Spec.CreatorSubject

	in.Status.Conditions = src.Status.Conditions

	return nil
}
