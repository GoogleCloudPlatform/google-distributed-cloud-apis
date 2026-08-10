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

/*
Copyright 2022.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// DatabaseParameters stores database parameters which are allowed to be modified by the users
type DatabaseParameters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DatabaseParametersSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:generate=true

type DatabaseParametersSpec struct {
	// The database engine name of the modifiable parameter list, should be one of [PostgreSQL, Oracle,AlloyDBOmni]
	// +required
	// +kubebuilder:validation:Enum=PostgreSQL;Oracle;AlloyDBOmni
	Engine EngineType `json:"engine,omitempty"`

	// The major database version of the modifiable parameter list
	// +required
	MajorVersion string `json:"majorversion,omitempty"`

	// Key would be name of the parameter e.g. "max_connections" and
	// value would be the metadata of the parameter.
	Parameters map[string]DatabaseParameter `json:"parameters"`
}

//+kubebuilder:object:generate=true

type DatabaseParameter struct {
	// Type of value of this parameter.
	// +required
	// +kubebuilder:validation:Enum=Integer;Boolean;Float;String;Enum;RepeatedEnum
	DataType DatabaseParameterDataType `json:"datatype,omitempty"`

	// Allowed values for this parameter.
	// +optional
	AllowedValues DatabaseParameterAllowedValues `json:"allowedvalues,omitempty"`

	// If ApplyType is Static, the parameter will be applied only after database is restarted.
	// Otherwise the parameter can be applied immediately.
	// +required
	// +kubebuilder:validation:Enum=Static;Dynamic
	ApplyType DatabaseParameterApplyType `json:"applytype,omitempty"`
}

type DatabaseParameterDataType string

const (
	Integer      DatabaseParameterDataType = "Integer"
	Boolean      DatabaseParameterDataType = "Boolean"
	Float        DatabaseParameterDataType = "Float"
	String       DatabaseParameterDataType = "String"
	Enum         DatabaseParameterDataType = "Enum"
	RepeatedEnum DatabaseParameterDataType = "RepeatedEnum"
)

type DatabaseParameterApplyType string

const (
	Static  DatabaseParameterApplyType = "Static"
	Dynamic DatabaseParameterApplyType = "Dynamic"
)

//+kubebuilder:object:generate=true

type DatabaseParameterAllowedValues struct {
	// MinIntValue and MaxIntValue specifies the ineteger value boundary of
	// Integer type parameters. These values are inclusive.
	MinIntValue int64 `json:"minintvalue,omitempty"`
	MaxIntValue int64 `json:"maxintvalue,omitempty"`

	// MinFloatValue and MaxFloatValue specifies the float value boundary of
	// Float type parameters. These values are inclusive.
	MinFloatValue resource.Quantity `json:"minfloatvalue,omitempty"`
	MaxFloatValue resource.Quantity `json:"maxfloatvalue,omitempty"`

	// EnumValues will store the list of allowed values for both Enum and RepeatedEnum parameters.
	// This field will also store the true/false text for Boolean type parameters.
	// E.g. “on”/”off” for PostgreSQL and “true”/”false” for Oracle
	EnumValues []string `json:"enumvalues,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseParametersList contains a list of database parameters.
type DatabaseParametersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseParameters `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseParameters{}, &DatabaseParametersList{})
}
