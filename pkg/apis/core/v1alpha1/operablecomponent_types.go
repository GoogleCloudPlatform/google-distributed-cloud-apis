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

// OperableComponent represents the name of the operable component.
type OperableComponent string

// Operable components are defined in go/gdch-operable-components.
const (
	// OperableComponentAR represents the operable component artifact repository.
	OperableComponentAR OperableComponent = "AR"
	// OperableComponentBIL represents the operable component Billing/Metering System.
	OperableComponentBIL OperableComponent = "BIL"
	// OperableComponentKUB represents the operable component user clusters.
	OperableComponentKUB OperableComponent = "KUB"
	// OperableComponentPNET represents the operable component Physical Network.
	OperableComponentPNET OperableComponent = "PNET"
	// OperableComponentFW represents the operable component Firewall.
	OperableComponentFW OperableComponent = "FW"
	// OperableComponentHSM represents the operable component HSM.
	OperableComponentHSM OperableComponent = "HSM"
	// OperableComponentOS represents the operable component OS.
	OperableComponentOS OperableComponent = "OS"
	// OperableComponentOBJ represents the operable component Object Storage.
	OperableComponentOBJ OperableComponent = "OBJ"
	// OperableComponentFILE represents the operable component File and Block Storage.
	OperableComponentFILE OperableComponent = "FILE"
	// OperableComponentFILE represents the operable component Vertex Workbench.
	OperableComponentVTXWB OperableComponent = "VTXWB"
	// OperableComponentIAM represents the operable component IAM.
	OperableComponentIAM OperableComponent = "IAM"
	// OperableComponentASM represents the operable component Anthos Service Mesh/ Istio.
	OperableComponentASM OperableComponent = "ASM"
	// OperableComponentUPORC represents the operable component UPORC.
	OperableComponentUPORC OperableComponent = "UPORC"
	// OperableComponentOPA represents the operable component OPA.
	OperableComponentOPA OperableComponent = "OPA"
)
