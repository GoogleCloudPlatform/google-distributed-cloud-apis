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

// Contains API Schema definitions for the Marketplace API group.
// +kubebuilder:object:generate=true
// +groupName=marketplaceview.gdc.goog
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public"
)

const (
	Group   = "marketplaceview." + public.Group
	Version = "v1"
)

var (
	// The group version used to register the given objects.
	GroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = GroupVersion

	// The builder used to add Golang types to the `GroupVersionKind` scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// The mechanism used to add the types in this group-version to the given
	// scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}
