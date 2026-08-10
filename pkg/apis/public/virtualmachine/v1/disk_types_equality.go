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
	"fmt"

	"k8s.io/apimachinery/pkg/api/equality"
)

func init() {
	// controller-runtime's CreateOrUpdate/CreateOrPatch functions are not happy when there is an unexported field in the API struct (https://github.com/kubernetes/kubernetes/issues/124154#issuecomment-2317904995).
	// Therefore we need to register a customized equality function that compares every field of the VirtualMachineDiskSpec except the unexported one.
	if err := equality.Semantic.AddFunc(compareVirtualMachineDiskSpec); err != nil {
		panic(fmt.Errorf("failed to register VirtualMachineDiskSpec compare function: %w", err))
	}
}

func compareVirtualMachineDiskSpec(a, b VirtualMachineDiskSpec) bool {
	return equality.Semantic.DeepEqual(a.Source, b.Source) &&
		equality.Semantic.DeepEqual(a.Size, b.Size) &&
		equality.Semantic.DeepEqual(a.Type, b.Type)
}
