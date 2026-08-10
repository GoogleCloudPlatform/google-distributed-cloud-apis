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

import "sigs.k8s.io/controller-runtime/pkg/client"

type LockStore interface {
	client.Object
	LockData() *LockData
}

type LockData struct {
	// OwnerUID is the UID of the cr obj that holds the lock
	// +optional
	OwnerUID string `json:"ownerUID"`

	// InheritorUID is the UID of the cr obj that has the current most senior claim to the lock
	// +optional
	InheritorUID string `json:"inheritorUID"`

	// Priority is the numerical priority of the inheritor of the lock, lower values are higher priority
	// +optional
	Priority int `json:"priority"`
}

type LockStorePlugin interface {
	NewLockData() LockStore
}
