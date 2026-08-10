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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespacedName refers an objects by namespace and name. It is the same as
// types.NamespacedName, but with json struct tags for controller-gen to
// generate CRD manifests.
type NamespacedName struct {
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

// UniversalObjectReference represents an object reference that can refer to
// either a local or a remote object. If it is a remote object, the `Cluster`
// field needs to be set to indicate which cluster this reference assumes.
type UniversalObjectReference struct {
	// Name of the reference. This field is required.
	Name string `json:"name"`

	// Namespace of the reference. This field is required.
	Namespace string `json:"namespace"`

	// Cluster holds the reference to the Cluster object if it is a remote reference.
	Cluster *corev1.ObjectReference `json:"cluster,omitempty"`
}

// BaseError represents a base error type. Fields are immutable.
type BaseError struct {
	// The code for the error. It includes a prefix with letters followed by
	// a four-digit numeric code.
	// +kubebuilder:validation:Pattern:=`^[A-Z]{2,8}[0-9]{4}$`
	Code string `json:"code"`

	// The human-readable error message.
	Message string `json:"message"`
}

// +kubebuilder:validation:Pattern=`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]{0,61}[A-Za-z0-9])$`
//
// DomainName represents a domain name.
// A valid domain name consists of one or more labels separated by dots (`.`)
// where each label:
// - has 1-63 characters
// - begins and ends with an alphanumeric character ([a-zA-Z0-9])
// - can contain dashes
// Note: due to regex's limitation, this permits names longer than 255
// characters, which isn't valid per RFC1035
type DomainName string

// HostAddress represents an IPv4, IPv6 address, or a domain name.
// (e.g. 172.17.0.1, 198a:a94b:98dc:4769:50ad:7281:6ba3:0306, 198A:A94B:98DC:4769:50AD:7281:6BA3:0306, or www.google.com)
// Regex validation is intentionally not used because the regex pattern would be too complex to verify by a developer.
// Instead, clients writing to fields of this type should use ToHostAddress to convert a string to this HostAddress.
type HostAddress string

// MachineType describes the type of nodes in a cluster.
// Only one of the following machine types may be specified.
type MachineType string

const (
	MachineTypeBaremetal MachineType = "Baremetal"
	MachineTypeVirtual   MachineType = "Virtual"
)

// ErrorStatus represents a list of error type with last update time.
type ErrorStatus struct {
	// Errors is a list of errors.
	Errors []BaseError `json:"errors,omitempty"`

	// LastUpdateTime is the timestamp of when the errors were last updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
}

// AddErrorCode adds an err to ErrorStatus.
func AddErrorCode(es **ErrorStatus, err BaseError) {
	if *es == nil {
		*es = &ErrorStatus{}
	}
	(*es).Errors = append((*es).Errors, err)
	(*es).LastUpdateTime = metav1.Now()
}

// SetErrorCodes clears ErrorStatus and adds errs.
func SetErrorCodes(es **ErrorStatus, errs []BaseError) {
	*es = &ErrorStatus{}
	(*es).Errors = append((*es).Errors, errs...)
	(*es).LastUpdateTime = metav1.Now()
}

// ClearErrorCodes clears ErrorStatus.
func ClearErrorCodes(es **ErrorStatus) {
	*es = nil
}
