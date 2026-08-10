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
	"sort"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	schema "k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Validate returns an error if the logging rule spec is invalid.
func (loggingrule *LoggingRule) Validate() error {
	if loggingrule == nil {
		return fmt.Errorf("unexpected nil logging rule object")
	}

	var allErrs field.ErrorList
	var specPath = field.NewPath("spec").Child("alertRules")

	severityLabel := "severity"
	codeLabel := "code"
	errorcodeLabel := "errorcode"
	resourceLabel := "resource"
	gdchProjectLabel := "_gdch_project"

	if loggingrule.Spec.Source != Audit &&
		loggingrule.Spec.Source != Operational {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("source"), loggingrule.Spec.Source, "accepted values are audit and operational"))
	}

	for i, alertRule := range loggingrule.Spec.AlertRules {
		if severityValue, ok := alertRule.Labels[severityLabel]; ok {
			if !validSeveritiesMap[severityValue] {
				allErrs = append(allErrs, field.Invalid(specPath.Index(i).Key(severityLabel), loggingrule.Spec.AlertRules[i].Labels[severityLabel], fmt.Sprintf("accepted severity values are %s", validSeveritiesString())))
			}
		} else {
			allErrs = append(allErrs, field.Required(specPath.Index(i).Key(severityLabel), fmt.Sprintf("accepted severity values are %s", validSeveritiesString())))
		}

		_, codeOK := alertRule.Labels[codeLabel]
		_, errorcodeOK := alertRule.Labels[errorcodeLabel]
		if !codeOK && !errorcodeOK {
			allErrs = append(allErrs, field.Required(specPath.Index(i).Key(errorcodeLabel), ""))
		}

		//TODO(b/441246561) After OC adher to alerting standards obs-alert-01 make _gdch_project required label
		_, resourceOK := alertRule.Labels[resourceLabel]
		_, gdchProjectOK := alertRule.Labels[gdchProjectLabel]
		if !resourceOK && !gdchProjectOK {
			allErrs = append(allErrs, field.Required(specPath.Index(i).Key(resourceLabel), ""))
		}
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(schema.GroupKind{
		Group: GroupVersion.Group,
		Kind:  "LoggingRule",
	}, loggingrule.Name, allErrs)
}

func validSeveritiesString() string {
	var severities []string
	for s := range validSeveritiesMap {
		severities = append(severities, s)
	}
	sort.Strings(severities)
	return strings.Join(severities, ", ")
}

var validSeveritiesMap = map[string]bool{
	"critical": true,
	"high":     true,
	"moderate": true,
	"low":      true,
	"error":    true, // Deprecated: Kept for backward compatibility.
	"warning":  true, // Deprecated: Kept for backward compatibility.
	"info":     true, // Deprecated: Kept for backward compatibility.
}
