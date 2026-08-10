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

// Protocol identifies the transport layer protocol.
//
// +kubebuilder:validation:Enum=TCP;UDP
type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

// A given port (or range) with the protocol.
type QualifiedPortRange struct {
	// Protocol for the service port.
	Protocol Protocol `json:"protocol"`

	// The port(s) for this service.
	Ports PortRange `json:"ports"`
}

// Can be a specific port, or a range specified with a dash (`-`).
// A range is inclusive of both start and end values.
//
// Examples: "80", "80-100"
//
// Note: Does not validate port number is within valid range.
// +kubebuilder:validation:Pattern=`^\d+(-\d+){0,1}$`
type PortRange string

const (
	PortRangeAny             PortRange = "1-65535"
	DNSPortRange             PortRange = "53"
	DNSoverTLSPortRange      PortRange = "853"
	ObjectStorageS3PortRange PortRange = "10443-10842"
)

// IPv4 address, range, or subnet.
//
// Examples: "10.120.0.0", "10.120.0.0-10.120.0.10", "10.120.0.0/26"
//
// Note: Does not validate numbers are within valid range, just that they are digits.
// +kubebuilder:validation:Pattern=`^((\d{1,3}.){3}\d{1,3})(-((\d{1,3}.){3}\d{1,3})|\/\d{1,2}){0,1}$`
type IPv4Range string

// IPv6 address, range, or subnet.
//
// Examples:
// "0000:0000:0000:0000:0000:0000:0000:aaaa",
// ":::::::aaaa"
// "0000:0000:0000:0000:0000:0000:0000:aaaa-0000:0000:0000:0000:0000:0000:0000:bbbb"
// "0000:0000:0000:0000:0000:0000:0000:aaaa/128"
//
// Note: Does not validate octents are within valid range, just that they are digits or letters.
// +kubebuilder:validation:Pattern=`(([\d\w|]{4}){0,1}:){7}([\d\w|]{4})(-(([\d\w|]{4}){0,1}:){7}([\d\w|]{4})|\/\d{1,3}){0,1}`
type IPv6Range string
