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

	logmonv1alpha1 "sigs.k8s.io/cluster-operators/logmon-operator/api/v1alpha1"

	v1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/logging/v1"
)

// ConvertTo converts this LoggingTarget to the Hub version (v1).
func (lt *LoggingTarget) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.LoggingTarget)
	dst.ObjectMeta = lt.ObjectMeta

	// Copy the spec
	dst.Spec.ServiceName = lt.Spec.ServiceName
	dst.Spec.LogAccessLevel = v1.LogAccessLevel(string(lt.Spec.LogAccessLevel))
	dst.Spec.Parser = logmonv1alpha1.OperationalLogParser(string(lt.Spec.Parser))

	dst.Spec.AdditionalFields = make(map[string]string)
	for k, v := range lt.Spec.AdditionalLabels {
		dst.Spec.AdditionalFields[k] = v
	}

	var dstLoggingTargetSelectors v1.LoggingTargetSelectors
	if lt.Spec.Selector.MatchClusters != nil {
		dstLoggingTargetSelectors.MatchClusters = append(dstLoggingTargetSelectors.MatchClusters, lt.Spec.Selector.MatchClusters...)
	}
	if lt.Spec.Selector.MatchContainerNames != nil {
		dstLoggingTargetSelectors.MatchContainerNames = append(dstLoggingTargetSelectors.MatchContainerNames, lt.Spec.Selector.MatchContainerNames...)
	}
	if lt.Spec.Selector.MatchPodNames != nil {
		dstLoggingTargetSelectors.MatchPodNames = append(dstLoggingTargetSelectors.MatchPodNames, lt.Spec.Selector.MatchPodNames...)
	}
	dst.Spec.Selector = dstLoggingTargetSelectors

	// Copy the status
	dst.Status.Conditions = append(dst.Status.Conditions, lt.Status.Conditions...)
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (lt *LoggingTarget) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.LoggingTarget)
	lt.ObjectMeta = src.ObjectMeta

	// Copy the spec.
	lt.Spec.ServiceName = src.Spec.ServiceName
	lt.Spec.LogAccessLevel = LogAccessLevel(string(src.Spec.LogAccessLevel))
	lt.Spec.Parser = logmonv1alpha1.OperationalLogParser(string(src.Spec.Parser))

	lt.Spec.AdditionalLabels = make(map[string]string)
	for k, v := range src.Spec.AdditionalFields {
		lt.Spec.AdditionalLabels[k] = v
	}

	var ltLoggingTargetSelectors LoggingTargetSelectors
	if src.Spec.Selector.MatchClusters != nil {
		ltLoggingTargetSelectors.MatchClusters = append(ltLoggingTargetSelectors.MatchClusters, src.Spec.Selector.MatchClusters...)
	}
	if src.Spec.Selector.MatchContainerNames != nil {
		ltLoggingTargetSelectors.MatchContainerNames = append(ltLoggingTargetSelectors.MatchContainerNames, src.Spec.Selector.MatchContainerNames...)
	}
	if src.Spec.Selector.MatchPodNames != nil {
		ltLoggingTargetSelectors.MatchPodNames = append(ltLoggingTargetSelectors.MatchPodNames, src.Spec.Selector.MatchPodNames...)
	}
	lt.Spec.Selector = ltLoggingTargetSelectors

	// Copy the status
	lt.Status.Conditions = append(lt.Status.Conditions, src.Status.Conditions...)
	return nil
}
