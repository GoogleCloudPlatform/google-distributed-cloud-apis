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

// Defines the quota for the service.
type Quota struct {
	// The quota name that identifies the quota within the service. Once a quota is in use, its name should be immutable. Unique in the service.
	Name string `json:"name"`

	// The name of the metric this quota applies to. The quota metric must be defined in the metrics list.
	MetricName string `json:"metricName"`

	// The scope of the quota:
	// "Organization" - The quota targets at the scope of an organization.
	// "Project" - The quota targets at the scope of a project.
	Scope QuotaScope `json:"scope"`

	// The dimensions the quota is defined on.
	// A dimension can be predefined or service specific.
	// The following dimensions are predefined: region and zone. The value of the predefined dimensions will be validated.
	// Service specific dimensions should document their cardinalities.
	// Quota is counted and limited on each dimension value combination. e.g. if the quota is defined on dimensions [“region”, “gpu-family”], then the quota is counted and limited on each combination of region value and gpu-family value, such as N200 gpu-family in region gdc-region1.
	Dimensions []string `json:"dimensions,omitempty"`

	// The time period the quota is refreshed. Used only for rate quota.
	// Currently supported refresh periods: "Day", "Minute".
	// +optional
	RefreshPeriod *RefreshPeriodType `json:"refreshPeriod,omitempty"`

	// The type that indicates whether the quota is precisely enforced or approximately enforced. Current enum values
	// Precise - the quota will be precisely enforced.
	// Imprecise - the quota will be approximately enforced, with around 1% skew rate above or below the quota value.
	Precise PreciseType `json:"precise"`

	// The type of the quota. Currently supported the following types:
	// Quota - the quota value can be adjusted by customers.
	// SystemLimit - the quota value cannot be adjusted by customers.
	Type QuotaType `json:"type"`
}

// +kubebuilder:validation:Enum=Project;Organization
type QuotaScope string

const (
	QuotaScopeProject      QuotaScope = "Project"
	QuotaScopeOrganization QuotaScope = "Organization"
)

// +kubebuilder:validation:Enum=Day;Minute
type RefreshPeriodType string

const (
	RefreshPeriodTypeDay    RefreshPeriodType = "Day"
	RefreshPeriodTypeMinute RefreshPeriodType = "Minute"
)

// +kubebuilder:validation:Enum=Precise;Imprecise
type PreciseType string

const (
	PreciseTypePrecise   PreciseType = "Precise"
	PreciseTypeImprecise PreciseType = "Imprecise"
)

// +kubebuilder:validation:Enum=Quota;SystemLimit
type QuotaType string

const (
	QuotaTypeQuota       QuotaType = "Quota"
	QuotaTypeSystemLimit QuotaType = "SystemLimit"
)
