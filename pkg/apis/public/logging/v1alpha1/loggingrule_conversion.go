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
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/logging/v1"
)

// ConvertTo converts this LoggingRule to the Hub version (v1).
func (lr *LoggingRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.LoggingRule)
	dst.ObjectMeta = lr.ObjectMeta

	// Copy the spec
	dst.Spec.Source = v1.Operational // default source
	if lr.Spec.Source == "audit" {
		dst.Spec.Source = v1.Audit
	}
	dst.Spec.Interval = lr.Spec.Interval
	dst.Spec.Limit = lr.Spec.Limit

	if lr.Spec.RecordRules != nil {
		for _, r := range lr.Spec.RecordRules {
			var dstRecordRule v1.RecordRule
			dstRecordRule.Record = r.Record
			dstRecordRule.Expr = r.Expr
			dstRecordRule.Labels = make(map[string]string)
			for k, v := range r.Labels {
				dstRecordRule.Labels[k] = v
			}
			dst.Spec.RecordRules = append(dst.Spec.RecordRules, dstRecordRule)
		}
	}

	if lr.Spec.AlertRules != nil {
		for _, r := range lr.Spec.AlertRules {
			var dstAlertRule v1.AlertRule
			dstAlertRule.Alert = r.Alert
			dstAlertRule.Expr = r.Expr
			dstAlertRule.For = r.For
			dstAlertRule.Labels = make(map[string]string)
			for k, v := range r.Labels {
				dstAlertRule.Labels[k] = v
			}
			dstAlertRule.Annotations = make(map[string]string)
			for k, v := range r.Annotations {
				dstAlertRule.Annotations[k] = v
			}
			dst.Spec.AlertRules = append(dst.Spec.AlertRules, dstAlertRule)
		}
	}

	// Copy the status
	dst.Status.Conditions = append(dst.Status.Conditions, lr.Status.Conditions...)
	dst.Status.LokiInstance = lr.Status.LokiInstance

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (lr *LoggingRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.LoggingRule)
	lr.ObjectMeta = src.ObjectMeta

	// Copy the spec
	if src.Spec.Source == v1.Audit {
		lr.Spec.Source = "audit"
	} else {
		lr.Spec.Source = "operational"
	}

	lr.Spec.Interval = src.Spec.Interval
	lr.Spec.Limit = src.Spec.Limit

	if src.Spec.RecordRules != nil {
		for _, r := range src.Spec.RecordRules {
			var srcRecordRule RecordRule
			srcRecordRule.Record = r.Record
			srcRecordRule.Expr = r.Expr
			srcRecordRule.Labels = make(map[string]string)
			for k, v := range r.Labels {
				srcRecordRule.Labels[k] = v
			}
			lr.Spec.RecordRules = append(lr.Spec.RecordRules, srcRecordRule)
		}
	}

	if src.Spec.AlertRules != nil {
		for _, r := range src.Spec.AlertRules {
			var srcAlertRule AlertRule
			srcAlertRule.Alert = r.Alert
			srcAlertRule.Expr = r.Expr
			srcAlertRule.For = r.For
			srcAlertRule.Labels = make(map[string]string)
			for k, v := range r.Labels {
				srcAlertRule.Labels[k] = v
			}
			srcAlertRule.Annotations = make(map[string]string)
			for k, v := range r.Annotations {
				srcAlertRule.Annotations[k] = v
			}
			lr.Spec.AlertRules = append(lr.Spec.AlertRules, srcAlertRule)
		}
	}

	// Copy the status
	lr.Status.Conditions = append(lr.Status.Conditions, src.Status.Conditions...)
	lr.Status.LokiInstance = src.Status.LokiInstance

	return nil
}
