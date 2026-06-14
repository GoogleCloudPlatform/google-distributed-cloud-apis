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
	"fmt"
)

// Validate returns an error if the project network policy spec is invalid.
func (pnp *ProjectNetworkPolicy) Validate() error {
	if pnp == nil {
		return fmt.Errorf("unexpected nil project network policy object")
	}

	// Validate policy type.
	switch pnp.Spec.PolicyType {
	case PolicyTypeIngress:
	case PolicyTypeEgress:
	default:
		return fmt.Errorf("invalid policy type %q", pnp.Spec.PolicyType)
	}

	// Return error if egress rules are specified for an ingress policy.
	if pnp.Spec.PolicyType == PolicyTypeIngress && pnp.Spec.Egress != nil {
		return fmt.Errorf("spec.egress must not specified for Ingress policy")
	}

	// Return error if ingress rules are specified for an egress policy.
	if pnp.Spec.PolicyType == PolicyTypeEgress && pnp.Spec.Ingress != nil {
		return fmt.Errorf("spec.ingress must not specified for Egress policy")
	}

	return nil
}
