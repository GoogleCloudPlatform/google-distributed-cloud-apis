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

	v1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/monitoring/v1"
)

func (m *MonitoringRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.MonitoringRule)
	dst.ObjectMeta = m.ObjectMeta
	//Spec
	dst.Spec.Interval = m.Spec.Interval
	dst.Spec.Limit = m.Spec.Limit
	dst.Spec.RecordRules = make([]v1.RecordRule, 0)
	for _, rec := range m.Spec.RecordRules {
		v1_rec := v1.RecordRule{
			Record: rec.Record,
			Expr:   rec.Expr,
			Labels: map[string]string{},
		}
		for k, v := range rec.Labels {
			v1_rec.Labels[k] = v
		}

		dst.Spec.RecordRules = append(dst.Spec.RecordRules, v1_rec)
	}
	dst.Spec.AlertRules = make([]v1.AlertRule, 0)
	for _, rule := range m.Spec.AlertRules {
		v1_rule := v1.AlertRule{
			Alert:       rule.Alert,
			Expr:        rule.Expr,
			For:         rule.For,
			Labels:      map[string]string{},
			Annotations: map[string]string{},
		}
		for k, v := range rule.Labels {
			v1_rule.Labels[k] = v
		}
		for k, v := range rule.Annotations {
			v1_rule.Annotations[k] = v
		}
		if len(rule.GroupbyLabels) > 0 {
			v1_rule.GroupbyLabels = append([]string(nil), rule.GroupbyLabels...)
		}

		dst.Spec.AlertRules = append(dst.Spec.AlertRules, v1_rule)
	}
	//Status
	dst.Status.Conditions = append(dst.Status.Conditions, m.Status.Conditions...)

	return nil
}

func (m *MonitoringRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.MonitoringRule)
	m.ObjectMeta = src.ObjectMeta
	//Spec
	m.Spec.Interval = src.Spec.Interval
	m.Spec.Limit = src.Spec.Limit
	m.Spec.RecordRules = make([]RecordRule, 0)
	for _, rec := range src.Spec.RecordRules {
		a1_rec := RecordRule{
			Record: rec.Record,
			Expr:   rec.Expr,
			Labels: map[string]string{},
		}
		for k, v := range rec.Labels {
			a1_rec.Labels[k] = v
		}

		m.Spec.RecordRules = append(m.Spec.RecordRules, a1_rec)
	}
	m.Spec.AlertRules = make([]AlertRule, 0)
	for _, v1_rule := range src.Spec.AlertRules {
		rule := AlertRule{
			Alert:         v1_rule.Alert,
			Expr:          v1_rule.Expr,
			For:           v1_rule.For,
			Labels:        map[string]string{},
			Annotations:   map[string]string{},
			GroupbyLabels: append([]string(nil), v1_rule.GroupbyLabels...),
		}
		for k, v := range v1_rule.Labels {
			rule.Labels[k] = v
		}
		for k, v := range v1_rule.Annotations {
			rule.Annotations[k] = v
		}
		m.Spec.AlertRules = append(m.Spec.AlertRules, rule)
	}
	//Status
	m.Status.Conditions = append(m.Status.Conditions, src.Status.Conditions...)

	return nil
}
