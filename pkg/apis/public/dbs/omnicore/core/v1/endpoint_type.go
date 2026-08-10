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

type EndpointName string

const (
	ReadWriteEndpoint         EndpointName = "Read-Write"
	ReadOnlyEndpoint          EndpointName = "Read-Only"
	MigrationOutgoingEndpoint EndpointName = "Migration-Outgoing"
)

// Endpoint represents a access point through which user can access the database.
type Endpoint struct {
	// Name contains the name of the endpoint
	// +required
	// +kubebuilder:validation:Required
	Name EndpointName `json:"name"`
	// Value contains the endpoint information.
	Value string `json:"value,omitempty"`
}

// UpsertEndpoint inserts or replace an endpoint from within a list, searching by its Name
func UpsertEndpoint(es *[]Endpoint, name EndpointName, value string) {
	for i := range *es {
		if (*es)[i].Name == name {
			(*es)[i].Value = value
			return
		}
	}
	*es = append(*es, Endpoint{name, value})
}

// DeleteEndpoint deletes an Endpoint from a list of Endpoints, searching by its Name.
func DeleteEndpoint(es *[]Endpoint, name EndpointName) {
	for i := range *es {
		if (*es)[i].Name == name {
			*es = append((*es)[:i], (*es)[i+1:]...)
			return
		}
	}
}

func FindEndpoint(es []Endpoint, name EndpointName) (string, bool) {
	for _, e := range es {
		if e.Name == name {
			return e.Value, true
		}
	}
	return "", false
}
