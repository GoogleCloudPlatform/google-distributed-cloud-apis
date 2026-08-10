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

import (
	"fmt"
	"net"

	corev1 "k8s.io/api/core/v1"
	k8snetworkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Validate returns an error if the project network policy spec is invalid.
func (pnp *ProjectNetworkPolicy) Validate() error {
	if pnp == nil {
		return fmt.Errorf("unexpected nil project network policy object")
	}
	errs := pnp.Spec.Validate()
	if len(errs) != 0 {
		return apierrors.NewInvalid(pnp.GroupVersionKind().GroupKind(), pnp.Name, errs)
	}
	return nil
}

func (pnpSpec ProjectNetworkPolicySpec) Validate() field.ErrorList {
	var errs field.ErrorList
	fPath := field.NewPath("spec")

	// Validate subject.
	switch pnpSpec.Subject.SubjectType {
	case PolicySubjectTypeUserWorkload:
		if pnpSpec.Subject.ManagedServices != nil {
			return append(errs, field.Forbidden(fPath.Child("subject", "managedServices"), "must not be specified for UserWorkload subject"))
		}
		if pnpSpec.Subject.WorkloadSelector != nil && pnpSpec.Subject.UserWorkloadSelector != nil {
			return append(errs, field.Forbidden(fPath.Child("subject"), "`workloadSelector` is deprecated and cannot be used with `userWorkloadSelector`"))
		}
		if pnpSpec.Subject.UserWorkloadSelector != nil {
			errs = append(errs, validateWorkloadSelector(fPath.Child("subject", "userWorkloadSelector"), *pnpSpec.Subject.UserWorkloadSelector, true)...)
		}
	case PolicySubjectTypeManagedService:
		if pnpSpec.Subject.WorkloadSelector != nil {
			return append(errs, field.Forbidden(fPath.Child("subject", "workloadSelector"), "must not be specified for ManagedService subject"))
		}
		if pnpSpec.Subject.UserWorkloadSelector != nil {
			return append(errs, field.Forbidden(fPath.Child("subject", "userWorkloadSelector"), "must not be specified for ManagedService subject"))
		}

		if pnpSpec.Subject.ManagedServices == nil {
			errs = append(errs, field.Required(fPath.Child("subject", "managedServices"), "must be specified for ManagedService subject"))
		} else if len(pnpSpec.Subject.ManagedServices.MatchTypes) != 1 {
			errs = append(errs, field.Invalid(fPath.Child("subject", "managedServices", "matchTypes"), pnpSpec.Subject.ManagedServices.MatchTypes, "must specify exactly one element"))
		}
	default:
		errs = append(errs, field.TypeInvalid(fPath.Child("subject", "subjectType"), pnpSpec.Subject.SubjectType, ""))
	}

	// Validate policy type.
	switch pnpSpec.PolicyType {
	case PolicyTypeIngress, PolicyTypeEgress:
	default:
		return append(errs, field.TypeInvalid(fPath.Child("policyType"), pnpSpec.PolicyType, ""))
	}

	// Validate Ingress rules.
	if pnpSpec.Ingress != nil {
		fPath := fPath.Child("ingress")
		// Return error if ingress rules are specified for an egress policy.
		if pnpSpec.PolicyType == PolicyTypeEgress {
			return append(errs, field.Forbidden(fPath, "must not be specified for Egress policy"))
		}
		if len(pnpSpec.Ingress) > 1 {
			return append(errs, field.TooMany(fPath, len(pnpSpec.Ingress), 1))
		}

		for i, ingRule := range pnpSpec.Ingress {
			if pnpSpec.Subject.SubjectType == PolicySubjectTypeManagedService && ingRule.Ports != nil {
				errs = append(errs, field.Forbidden(fPath.Index(i).Child("ports"), "must not be specified for ManagedService subject"))
				continue
			}
			if pnpSpec.Subject.SubjectType == PolicySubjectTypeManagedService && len(ingRule.Protocols) > 0 {
				errs = append(errs, field.Forbidden(fPath.Index(i).Child("protocols"), "must not be specified for ManagedService subject"))
				continue
			}
			if err := validateIngressRule(fPath.Index(i), ingRule); err != nil {
				errs = append(errs, err...)
			}
		}
	}

	// Validate Egress rules.
	if pnpSpec.Egress != nil {
		fPath := fPath.Child("egress")
		// Return error if egress rules are specified for an ingress policy.
		if pnpSpec.PolicyType == PolicyTypeIngress {
			return append(errs, field.Forbidden(fPath, "must not be specified for Ingress policy"))
		}
		if len(pnpSpec.Egress) > 1 {
			return append(errs, field.TooMany(fPath, len(pnpSpec.Egress), 1))
		}

		for i, egsRule := range pnpSpec.Egress {
			if pnpSpec.Subject.SubjectType == PolicySubjectTypeManagedService && egsRule.Ports != nil {
				errs = append(errs, field.Forbidden(fPath.Index(i).Child("ports"), "must not be specified for ManagedService subject"))
				continue
			}
			if pnpSpec.Subject.SubjectType == PolicySubjectTypeManagedService && len(egsRule.Protocols) > 0 {
				errs = append(errs, field.Forbidden(fPath.Index(i).Child("protocols"), "must not be specified for ManagedService subject"))
				continue
			}
			if err := validateEgressRule(fPath.Index(i), egsRule); err != nil {
				errs = append(errs, err...)
			}
		}
	}
	return errs
}

func validateIngressRule(fPath *field.Path, ingRule ProjectNetworkPolicyIngressRule) field.ErrorList {
	var errs field.ErrorList
	if len(ingRule.Ports) > 0 && len(ingRule.Protocols) > 0 {
		errs = append(errs, field.Forbidden(fPath.Child("ports"), "ports and protocols must not be used simultaneously"))
	}
	if ingRule.Action == ActionTypeDeny {
		if len(ingRule.From) == 0 {
			errs = append(errs, field.Required(fPath.Child("from"), "from must be specified for ingress Deny rules"))
		}
	}
	if len(ingRule.From) > 1 {
		errs = append(errs, field.TooMany(fPath.Child("from"), len(ingRule.From), 1))
	}
	for i, from := range ingRule.From {
		if err := validatePeer(fPath.Child("from").Index(i), from); err != nil {
			errs = append(errs, err...)
		}
	}

	for i, port := range ingRule.Ports {
		if err := validatePort(fPath.Child("ports").Index(i), port); err != nil {
			errs = append(errs, err...)
		}
	}

	for i, protocol := range ingRule.Protocols {
		if err := validateProtocol(fPath.Child("protocols").Index(i), protocol); err != nil {
			errs = append(errs, err...)
		}
	}
	return errs
}

func validateEgressRule(fPath *field.Path, egsRule ProjectNetworkPolicyEgressRule) field.ErrorList {
	var errs field.ErrorList
	if len(egsRule.Ports) > 0 && len(egsRule.Protocols) > 0 {
		errs = append(errs, field.Forbidden(fPath.Child("ports"), "ports and protocols must not be used simultaneously"))
	}
	if egsRule.Action == ActionTypeDeny {
		if len(egsRule.To) == 0 {
			errs = append(errs, field.Required(fPath.Child("to"), "to must be specified for egress Deny rules"))
		}
	}
	if len(egsRule.To) > 1 {
		errs = append(errs, field.TooMany(fPath.Child("to"), len(egsRule.To), 1))
	}
	for i, to := range egsRule.To {
		if err := validatePeer(fPath.Child("to").Index(i), to); err != nil {
			errs = append(errs, err...)
		}
	}

	for i, port := range egsRule.Ports {
		if err := validatePort(fPath.Child("ports").Index(i), port); err != nil {
			errs = append(errs, err...)
		}
	}

	for i, protocol := range egsRule.Protocols {
		if err := validateProtocol(fPath.Child("protocols").Index(i), protocol); err != nil {
			errs = append(errs, err...)
		}
	}
	return errs
}

func validatePeer(fPath *field.Path, peer ProjectNetworkPolicyPeer) field.ErrorList {
	specified := 0
	if peer.Projects != nil {
		specified++
	}
	if peer.ProjectSelector != nil {
		specified++
	}
	if peer.IPBlock != nil {
		specified++
	}
	if peer.IPBlocks != nil {
		specified++
	}
	if specified != 1 {
		return field.ErrorList{field.Invalid(fPath, peer, "must specify exactly one subfield")}
	}

	var errs field.ErrorList
	if peer.Projects != nil {
		if err := validateProjects(fPath.Child("projects"), *peer.Projects); err != nil {
			errs = append(errs, err...)
		}
	}

	if peer.ProjectSelector != nil {
		if err := validateProjectSelector(fPath.Child("projectSelector"), *peer.ProjectSelector); err != nil {
			errs = append(errs, err...)
		}
	}

	if peer.IPBlock != nil {
		if err := validateIPBlock(fPath.Child("ipBlock"), *peer.IPBlock); err != nil {
			errs = append(errs, err...)
		}
	}

	if peer.IPBlocks != nil {
		for i, ipBlock := range peer.IPBlocks {
			if err := validateIPBlock(fPath.Child("ipBlocks").Index(i), ipBlock); err != nil {
				errs = append(errs, err...)
			}
		}
	}
	return errs
}

func validatePort(fPath *field.Path, port ProjectNetworkPolicyPort) field.ErrorList {
	var errs field.ErrorList
	if port.Protocol != nil {
		switch *port.Protocol {
		case corev1.ProtocolTCP, corev1.ProtocolUDP, corev1.ProtocolSCTP:
		default:
			errs = append(errs, field.TypeInvalid(fPath.Child("protocol"), *port.Protocol, ""))
		}
	}

	if port.Port != nil {
		if port.Port.IntValue() < 0 || port.Port.IntValue() > 65535 {
			errs = append(errs, field.TypeInvalid(fPath.Child("port"), *port.Port, "must be in range [0, 65535]"))
		}
	}
	return errs
}

func validateProtocol(fPath *field.Path, protocol ProjectNetworkPolicyProtocol) field.ErrorList {
	var errs field.ErrorList

	specified := 0
	if protocol.TCP != nil {
		specified++
	}
	if protocol.UDP != nil {
		specified++
	}
	if protocol.SCTP != nil {
		specified++
	}
	if protocol.ICMP != nil {
		specified++
	}
	if protocol.DestinationNamedPort != "" {
		specified++
	}
	if specified != 1 {
		errs = append(errs, field.Invalid(fPath, protocol, "must specify exactly one of tcp, udp, sctp, icmp, or destinationNamedPort"))
	}

	if protocol.TCP != nil {
		if protocol.TCP.Port < 0 || protocol.TCP.Port > 65535 {
			errs = append(errs, field.Invalid(fPath.Child("tcp", "port"), protocol.TCP.Port, "must be in range [0, 65535]"))
		}
	}
	if protocol.UDP != nil {
		if protocol.UDP.Port < 0 || protocol.UDP.Port > 65535 {
			errs = append(errs, field.Invalid(fPath.Child("udp", "port"), protocol.UDP.Port, "must be in range [0, 65535]"))
		}
	}
	if protocol.SCTP != nil {
		if protocol.SCTP.Port < 0 || protocol.SCTP.Port > 65535 {
			errs = append(errs, field.Invalid(fPath.Child("sctp", "port"), protocol.SCTP.Port, "must be in range [0, 65535]"))
		}
	}
	return errs
}

func validateProjectSelector(fPath *field.Path, pjSelector ProjectSelector) field.ErrorList {
	var errs field.ErrorList

	if pjSelector.Projects == nil {
		errs = append(errs, field.Required(fPath.Child("projects"), "must not be nil"))
	} else {
		errs = append(errs, validateProjects(fPath.Child("projects"), *pjSelector.Projects)...)
	}

	if pjSelector.Workloads != nil && pjSelector.WorkloadSelector != nil {
		errs = append(errs, field.Forbidden(fPath.Child("projects"), "must not specify both `workloads` and `workloadSelector`"))
	}

	if pjSelector.WorkloadSelector != nil {
		path := fPath.Child("workloadSelector")
		if pjSelector.WorkloadSelector.LabelSelector == nil {
			errs = append(errs, field.Required(path.Child("labelSelector"), "must not be nil"))
		} else {
			errs = append(errs, validateWorkloadSelector(path, *pjSelector.WorkloadSelector, false)...)
		}
	}

	if pjSelector.Projects != nil && len(pjSelector.Projects.MatchNames) == 0 &&
		pjSelector.WorkloadSelector != nil &&
		pjSelector.WorkloadSelector.LabelSelector != nil &&
		pjSelector.WorkloadSelector.LabelSelector.Clusters != nil &&
		len(pjSelector.WorkloadSelector.LabelSelector.Clusters.MatchLabels) == 1 {
		errs = append(errs, field.Forbidden(fPath, "A project name must be specified in `projects.matchNames` when selecting a specific cluster"))
	}
	return errs
}

func validateProjects(fPath *field.Path, pjs PolicyProjects) field.ErrorList {
	if pjs.MatchNames == nil {
		return field.ErrorList{field.Required(fPath.Child("matchNames"), "must not be nil")}
	}
	if len(pjs.MatchNames) > 1 {
		return field.ErrorList{field.TooMany(fPath.Child("matchNames"), len(pjs.MatchNames), 1)}
	}
	return nil
}

func validateWorkloadSelector(fPath *field.Path, ws WorkloadSelector, isSubject bool) field.ErrorList {
	var errs field.ErrorList
	path := fPath.Child("labelSelector")
	if ws.LabelSelector == nil {
		return append(errs, field.Required(path, "must not be nil"))
	}
	ls := ws.LabelSelector

	if isSubject && ls.Clusters == nil && ls.Namespaces != nil {
		return append(errs, field.Forbidden(path, "`namespaces` must not be specified without `clusters`"))
	}

	// Validate clusters field.
	if ls.Clusters != nil {
		path := path.Child("clusters")
		if ls.Clusters.MatchExpressions != nil {
			errs = append(errs, field.Forbidden(path.Child("matchExpressions"), "must not be specified"))
		}
		path = path.Child("matchLabels")
		if len(ls.Clusters.MatchLabels) != 1 {
			return append(errs, field.Invalid(path, ls.Clusters.MatchLabels, "must have exactly one item"))
		}
		val, ok := ls.Clusters.MatchLabels["kubernetes.io/metadata.name"]
		if !ok {
			errs = append(errs, field.Invalid(path, ls.Clusters.MatchLabels, "must have the key `kubernetes.io/metadata.name`"))
		} else if len(val) == 0 {
			errs = append(errs, field.Invalid(path, ls.Clusters.MatchLabels, "value of the key `kubernetes.io/metadata.name` must not be empty"))
		}
	}
	return errs
}

func validateIPBlock(fPath *field.Path, ipBlock k8snetworkingv1.IPBlock) field.ErrorList {
	if ipBlock.CIDR == "" {
		return field.ErrorList{field.Required(fPath.Child("cidr"), "must specify valid CIDR")}
	}
	var errs field.ErrorList
	if _, _, err := net.ParseCIDR(ipBlock.CIDR); err != nil {
		errs = append(errs, field.Invalid(fPath.Child("cidr"), ipBlock.CIDR, err.Error()))
	}

	for i, except := range ipBlock.Except {
		if _, _, err := net.ParseCIDR(except); err != nil {
			errs = append(errs, field.Invalid(fPath.Child("except").Index(i), except, err.Error()))
		}
	}
	return errs
}
