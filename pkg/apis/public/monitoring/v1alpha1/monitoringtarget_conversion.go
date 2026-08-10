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

// ConvertTo converts this MonitoringTarget to the Hub version (v1).
func (mt *MonitoringTarget) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.MonitoringTarget)
	dst.ObjectMeta = mt.ObjectMeta
	//Spec
	//Spec-Selector
	dst.Spec.Selector.MatchClusters = make([]string, 0)
	dst.Spec.Selector.MatchClusters = append(dst.Spec.Selector.MatchClusters, mt.Spec.Selector.MatchClusters...)
	dst.Spec.Selector.MatchLabels = make(map[string]string)
	for key, label := range mt.Spec.Selector.MatchLabels {
		dst.Spec.Selector.MatchLabels[key] = label
	}
	dst.Spec.Selector.MatchAnnotations = make(map[string]string)
	for key, annotation := range mt.Spec.Selector.MatchAnnotations {
		dst.Spec.Selector.MatchAnnotations[key] = annotation
	}
	//Spec-PodMetricEndpoints
	dst.Spec.PodMetricsEndpoints.Port.Value = mt.Spec.PodMetricsEndpoints.Port.Value
	dst.Spec.PodMetricsEndpoints.Port.Annotation = mt.Spec.PodMetricsEndpoints.Port.Annotation

	dst.Spec.PodMetricsEndpoints.Path.Value = mt.Spec.PodMetricsEndpoints.Path.Value
	dst.Spec.PodMetricsEndpoints.Path.Annotation = mt.Spec.PodMetricsEndpoints.Path.Annotation

	dst.Spec.PodMetricsEndpoints.Scheme.Value = mt.Spec.PodMetricsEndpoints.Scheme.Value
	dst.Spec.PodMetricsEndpoints.Scheme.Annotation = mt.Spec.PodMetricsEndpoints.Scheme.Annotation

	dst.Spec.PodMetricsEndpoints.Params = make(map[string][]string)
	for key, param := range mt.Spec.PodMetricsEndpoints.Params {
		dst.Spec.PodMetricsEndpoints.Params[key] = append(dst.Spec.PodMetricsEndpoints.Params[key], param...)
	}

	dst.Spec.PodMetricsEndpoints.ScrapeInterval = mt.Spec.PodMetricsEndpoints.ScrapeInterval
	dst.Spec.PodMetricsEndpoints.ScrapeTimeout = mt.Spec.PodMetricsEndpoints.ScrapeTimeout

	dst.Spec.PodMetricsEndpoints.MetricsRelabelings = make([]v1.MonitoringTargetMetricsRelabeling, 0)
	for _, mr := range mt.Spec.PodMetricsEndpoints.MetricsRelabelings {
		v1_mr := v1.MonitoringTargetMetricsRelabeling{
			SourceLabels: []string{},
			Separator:    mr.Separator,
			Regex:        mr.Regex,
			Action:       mr.Action,
			TargetLabel:  mr.TargetLabel,
			Replacement:  mr.Replacement,
		}
		v1_mr.SourceLabels = append(v1_mr.SourceLabels, mr.SourceLabels...)

		dst.Spec.PodMetricsEndpoints.MetricsRelabelings = append(dst.Spec.PodMetricsEndpoints.MetricsRelabelings, v1_mr)
	}

	//Status
	dst.Status.Conditions = append(dst.Status.Conditions, mt.Status.Conditions...)
	dst.Status.ClusterStatuses = make([]v1.ClusterStatus, 0)

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (mt *MonitoringTarget) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.MonitoringTarget)
	mt.ObjectMeta = src.ObjectMeta
	//Spec
	//Spec-Selector
	mt.Spec.Selector.MatchClusters = make([]string, 0)
	mt.Spec.Selector.MatchClusters = append(mt.Spec.Selector.MatchClusters, src.Spec.Selector.MatchClusters...)
	mt.Spec.Selector.MatchLabels = make(map[string]string)
	for key, v1Label := range src.Spec.Selector.MatchLabels {
		mt.Spec.Selector.MatchLabels[key] = v1Label
	}
	mt.Spec.Selector.MatchAnnotations = make(map[string]string)
	for key, v1Label := range src.Spec.Selector.MatchAnnotations {
		mt.Spec.Selector.MatchAnnotations[key] = v1Label
	}
	//Spec-PodMetricEndpoints
	mt.Spec.PodMetricsEndpoints.Port.Value = src.Spec.PodMetricsEndpoints.Port.Value
	mt.Spec.PodMetricsEndpoints.Port.Annotation = src.Spec.PodMetricsEndpoints.Port.Annotation

	mt.Spec.PodMetricsEndpoints.Path.Value = src.Spec.PodMetricsEndpoints.Path.Value
	mt.Spec.PodMetricsEndpoints.Path.Annotation = src.Spec.PodMetricsEndpoints.Path.Annotation

	mt.Spec.PodMetricsEndpoints.Scheme.Value = src.Spec.PodMetricsEndpoints.Scheme.Value
	mt.Spec.PodMetricsEndpoints.Scheme.Annotation = src.Spec.PodMetricsEndpoints.Scheme.Annotation

	mt.Spec.PodMetricsEndpoints.Params = make(map[string][]string)
	for key, param := range src.Spec.PodMetricsEndpoints.Params {
		mt.Spec.PodMetricsEndpoints.Params[key] = append(mt.Spec.PodMetricsEndpoints.Params[key], param...)
	}

	mt.Spec.PodMetricsEndpoints.ScrapeInterval = src.Spec.PodMetricsEndpoints.ScrapeInterval
	mt.Spec.PodMetricsEndpoints.ScrapeTimeout = src.Spec.PodMetricsEndpoints.ScrapeTimeout

	mt.Spec.PodMetricsEndpoints.MetricsRelabelings = make([]MonitoringTargetMetricsRelabeling, 0)
	for _, srcMetricsRelabeling := range src.Spec.PodMetricsEndpoints.MetricsRelabelings {
		mr := MonitoringTargetMetricsRelabeling{
			SourceLabels: []string{},
			Separator:    srcMetricsRelabeling.Separator,
			Regex:        srcMetricsRelabeling.Regex,
			Action:       srcMetricsRelabeling.Action,
			TargetLabel:  srcMetricsRelabeling.TargetLabel,
			Replacement:  srcMetricsRelabeling.Replacement,
		}
		mr.SourceLabels = append(mr.SourceLabels, srcMetricsRelabeling.SourceLabels...)
		mt.Spec.PodMetricsEndpoints.MetricsRelabelings = append(mt.Spec.PodMetricsEndpoints.MetricsRelabelings, mr)
	}
	mt.Status.Conditions = append(mt.Status.Conditions, src.Status.Conditions...)

	return nil
}
