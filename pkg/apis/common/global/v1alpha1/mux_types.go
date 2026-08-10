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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NOTE: the getter names in the interfaces defined in this file do not conform
// to go/go-style/decisions#getters, to avoid naming conflicts between getters
// and exported fields.

// Lets you work with a global mux resource.
// Any global mux resoruce type must implement this interface for the global mux
// API machinery to roll out the resource spec and collect statutes.
type MuxResourceInterface interface {
	client.Object
	GetSpec() any
	SetSpec(spec any) error
	GetStatus() MuxStatusInterface
	SetStatus(status MuxStatusInterface) error
}

// Lets you work with a replica of a global mux resource.
// Any replica type of a global mux resource must implement this interface for
// the global mux API machinery to roll out the resource spec and collect
// statutes.
type MuxReplicaInterface interface {
	client.Object
	GetSpec() any
	SetSpec(spec any) error
	GetStatus() any
	SetStatus(status any) error
}
