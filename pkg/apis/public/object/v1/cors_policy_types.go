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

// Defines the information related with CORS rule.
type CorsRule struct {
	// Headers that are specified in the Access-Control-Request-Headers header.
	AllowedHeaders []*string `json:"allowedHeaders"`
	// HTTP methods that are permitted to be executed by an allowed origin.
	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods"`
	// Origins that can access the bucket.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins"`
	// Headers in the response that can be accessed.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty"`
	// Unique identifier for the rule. The value can not be longer than 255 characters.
	// +kubebuilder:validation:Required
	ID *string `json:"id"`
}

// Defines the strategy for setting up custom CORS policy on the bucket.
type CorsPolicy struct {
	// Defines whether the user wants the custom policy to take effect on the bucket.
	// If yes, the custom policy defined in CorsDetail will be read.
	// Otherwise, CorsDetail would not be used even if it has custom policy set up.
	// +optional
	// +kubebuilder:default:=false
	EnableCorsPolicy bool `json:"enableCorsPolicy,omitempty"`
	// Detail of the custom CORS policy being set.
	// +optional
	CorsDetail []*CorsRule `json:"corsDetail,omitempty"`
}
