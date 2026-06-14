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

package common

import (
	"fmt"
	"hash/fnv"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func FNV64Hash(input string) string {
	hasher := fnv.New64()
	hasher.Write([]byte(input))
	return fmt.Sprintf("%x", hasher.Sum64())
}

func GetDBClusterInternalName(dbcName string) string {
	if len(dbcName) <= 40 {
		return dbcName
	}
	return dbcName[:DbClusterNamePrefixLengthInInternalName] + FNV64Hash(dbcName)
}

func IsOperatorManagedResource(obj client.Object) bool {
	if obj.GetLabels() == nil {
		return false
	}

	labels := obj.GetLabels()
	enabled, ok := labels[ManagedByLabel]
	return ok && enabled == ManagedByOperator
}
