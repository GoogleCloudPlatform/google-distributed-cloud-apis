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

	arv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/artifactregistry/v1"
)

// convertTo converts this HarborInstance (v1alpha2) to the Hub version (v1).
func (in *HarborInstance) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*arv1.HarborInstance)

	dst.ObjectMeta = in.ObjectMeta
	// no field since v1alpha2 has an empty Spec
	dst.Spec = arv1.HarborInstanceSpec{}

	dst.Status.Conditions = in.Status.Conditions
	dst.Status.Version = in.Status.Version
	dst.Status.URL = in.Status.URL

	return nil
}

// convertFrom converts HarborInstance from the Hub version (v1) to this (v1alpha2) version.
func (in *HarborInstance) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*arv1.HarborInstance)

	in.ObjectMeta = src.ObjectMeta
	// no field since v1alpha2 has an empty Spec
	in.Spec = HarborInstanceSpec{}

	in.Status.Conditions = src.Status.Conditions
	in.Status.Version = src.Status.Version
	in.Status.URL = src.Status.URL

	return nil
}
