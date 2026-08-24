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
	v1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/observability/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (als *AdditionalLogSink) ConvertTo(dst *v1.AdditionalLogSink) {
	if srcExcludedClusterList := als.ClusterSelector.ExcludedClusterList; srcExcludedClusterList != nil {
		dstExcludedClusterList := make([]string, len(srcExcludedClusterList))
		copy(dstExcludedClusterList, srcExcludedClusterList)
		dst.ClusterSelector.ExcludedClusterList = dstExcludedClusterList
	}
	if srcFluentBitConfigMaps := als.FluentBitConfigMaps; srcFluentBitConfigMaps != nil {
		dstFluentBitConfigMaps := make([]string, len(srcFluentBitConfigMaps))
		copy(dstFluentBitConfigMaps, srcFluentBitConfigMaps)
		dst.FluentBitConfigMaps = dstFluentBitConfigMaps
	}
	if srcLoggingVolumes := als.Volumes; srcLoggingVolumes != nil {
		dstLoggingVolumes := make([]corev1.Volume, len(srcLoggingVolumes))
		copy(dstLoggingVolumes, srcLoggingVolumes)
		dst.Volumes = dstLoggingVolumes
	}
	if srcLoggingVolumeMounts := als.VolumeMounts; srcLoggingVolumeMounts != nil {
		dstLoggingVolumeMounts := make([]corev1.VolumeMount, len(srcLoggingVolumeMounts))
		copy(dstLoggingVolumeMounts, srcLoggingVolumeMounts)
		dst.VolumeMounts = dstLoggingVolumeMounts
	}
}

func (als *AdditionalLogSink) ConvertFrom(src *v1.AdditionalLogSink) {
	if srcExcludedClusterList := src.ClusterSelector.ExcludedClusterList; srcExcludedClusterList != nil {
		dstExcludedClusterList := make([]string, len(srcExcludedClusterList))
		copy(dstExcludedClusterList, srcExcludedClusterList)
		als.ClusterSelector.ExcludedClusterList = dstExcludedClusterList
	}
	if srcFluentBitConfigMaps := src.FluentBitConfigMaps; srcFluentBitConfigMaps != nil {
		dstFluentBitConfigMaps := make([]string, len(srcFluentBitConfigMaps))
		copy(dstFluentBitConfigMaps, srcFluentBitConfigMaps)
		als.FluentBitConfigMaps = dstFluentBitConfigMaps
	}
	if srcLoggingVolumes := src.Volumes; srcLoggingVolumes != nil {
		dstLoggingVolumes := make([]corev1.Volume, len(srcLoggingVolumes))
		copy(dstLoggingVolumes, srcLoggingVolumes)
		als.Volumes = dstLoggingVolumes
	}
	if srcLoggingVolumeMounts := src.VolumeMounts; srcLoggingVolumeMounts != nil {
		dstLoggingVolumeMounts := make([]corev1.VolumeMount, len(srcLoggingVolumeMounts))
		copy(dstLoggingVolumeMounts, srcLoggingVolumeMounts)
		als.VolumeMounts = dstLoggingVolumeMounts
	}
}

// ConvertTo converts this ObservabilityPipeline to the Hub version (v1).
func (op *ObservabilityPipeline) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.ObservabilityPipeline)
	dst.ObjectMeta = op.ObjectMeta

	// Copy the spec
	dst.Spec.Enabled = op.Spec.Enabled

	dst.Spec.Monitoring.RetentionTime = op.Spec.Monitoring.RetentionTime
	dst.Spec.Monitoring.LocalStorageSize = op.Spec.Monitoring.LocalStorageSize
	dst.Spec.Monitoring.Sink = op.Spec.Monitoring.Sink
	dst.Spec.Monitoring.Grafana = v1.Grafana(op.Spec.Monitoring.Grafana)

	dst.Spec.Alerting = v1.ObservabilityAlerting(op.Spec.Alerting)

	dst.Spec.Logging.RetentionTime = op.Spec.Logging.RetentionTime
	dst.Spec.Logging.LocalStorageSize = op.Spec.Logging.LocalStorageSize
	dst.Spec.Logging.Sink = op.Spec.Logging.Sink
	op.Spec.Logging.AdditionalSink.ConvertTo(&dst.Spec.Logging.AdditionalSink)
	if srcLoggingDynamicAdditionalSinks := op.Spec.Logging.DynamicAdditionalSinks; srcLoggingDynamicAdditionalSinks != nil {
		dstLoggingDynamicAdditionalSinks := make([]string, len(srcLoggingDynamicAdditionalSinks))
		copy(dstLoggingDynamicAdditionalSinks, srcLoggingDynamicAdditionalSinks)
		dst.Spec.Logging.DynamicAdditionalSinks = dstLoggingDynamicAdditionalSinks
	}

	dst.Spec.AuditLogging.RetentionTime = op.Spec.AuditLogging.RetentionTime
	dst.Spec.AuditLogging.LocalStorageSize = op.Spec.AuditLogging.LocalStorageSize
	op.Spec.AuditLogging.AdditionalSink.ConvertTo(&dst.Spec.AuditLogging.AdditionalSink)
	if srcAuditLoggingDynamicAdditionalSinks := op.Spec.AuditLogging.DynamicAdditionalSinks; srcAuditLoggingDynamicAdditionalSinks != nil {
		dstAuditLoggingDynamicAdditionalSinks := make([]string, len(srcAuditLoggingDynamicAdditionalSinks))
		copy(dstAuditLoggingDynamicAdditionalSinks, srcAuditLoggingDynamicAdditionalSinks)
		dst.Spec.AuditLogging.DynamicAdditionalSinks = dstAuditLoggingDynamicAdditionalSinks
	}

	dst.Spec.SecurityLogging.RetentionTime = op.Spec.SecurityLogging.RetentionTime
	dst.Spec.SecurityLogging.LocalStorageSize = op.Spec.SecurityLogging.LocalStorageSize
	op.Spec.SecurityLogging.AdditionalSink.ConvertTo(&dst.Spec.SecurityLogging.AdditionalSink)
	if srcSecurityLoggingDynamicAdditionalSinks := op.Spec.SecurityLogging.DynamicAdditionalSinks; srcSecurityLoggingDynamicAdditionalSinks != nil {
		dstSecurityLoggingDynamicAdditionalSinks := make([]string, len(srcSecurityLoggingDynamicAdditionalSinks))
		copy(dstSecurityLoggingDynamicAdditionalSinks, srcSecurityLoggingDynamicAdditionalSinks)
		dst.Spec.SecurityLogging.DynamicAdditionalSinks = dstSecurityLoggingDynamicAdditionalSinks
	}

	// Copy the status
	dst.Status.Version = op.Status.Version
	dst.Status.Conditions = append(dst.Status.Conditions, op.Status.Conditions...)
	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (op *ObservabilityPipeline) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.ObservabilityPipeline)
	op.ObjectMeta = src.ObjectMeta

	// Copy the spec.
	op.Spec.Enabled = src.Spec.Enabled

	op.Spec.Monitoring.RetentionTime = src.Spec.Monitoring.RetentionTime
	op.Spec.Monitoring.LocalStorageSize = src.Spec.Monitoring.LocalStorageSize
	op.Spec.Monitoring.Sink = src.Spec.Monitoring.Sink
	op.Spec.Monitoring.Grafana = Grafana(src.Spec.Monitoring.Grafana)

	op.Spec.Alerting = ObservabilityAlerting(src.Spec.Alerting)
	op.Spec.Logging.RetentionTime = src.Spec.Logging.RetentionTime

	op.Spec.Logging.LocalStorageSize = src.Spec.Logging.LocalStorageSize
	op.Spec.Logging.Sink = src.Spec.Logging.Sink
	op.Spec.Logging.AdditionalSink.ConvertFrom(&src.Spec.Logging.AdditionalSink)
	if srcLoggingDynamicAdditionalSinks := src.Spec.Logging.DynamicAdditionalSinks; srcLoggingDynamicAdditionalSinks != nil {
		dstLoggingDynamicAdditionalSinks := make([]string, len(srcLoggingDynamicAdditionalSinks))
		copy(dstLoggingDynamicAdditionalSinks, srcLoggingDynamicAdditionalSinks)
		op.Spec.Logging.DynamicAdditionalSinks = dstLoggingDynamicAdditionalSinks
	}

	op.Spec.AuditLogging.RetentionTime = src.Spec.AuditLogging.RetentionTime
	op.Spec.AuditLogging.LocalStorageSize = src.Spec.AuditLogging.LocalStorageSize
	op.Spec.AuditLogging.AdditionalSink.ConvertFrom(&src.Spec.AuditLogging.AdditionalSink)
	if srcAuditLoggingDynamicAdditionalSinks := src.Spec.AuditLogging.DynamicAdditionalSinks; srcAuditLoggingDynamicAdditionalSinks != nil {
		dstAuditLoggingDynamicAdditionalSinks := make([]string, len(srcAuditLoggingDynamicAdditionalSinks))
		copy(dstAuditLoggingDynamicAdditionalSinks, srcAuditLoggingDynamicAdditionalSinks)
		op.Spec.AuditLogging.DynamicAdditionalSinks = dstAuditLoggingDynamicAdditionalSinks
	}

	op.Spec.SecurityLogging.RetentionTime = src.Spec.SecurityLogging.RetentionTime
	op.Spec.SecurityLogging.LocalStorageSize = src.Spec.SecurityLogging.LocalStorageSize
	op.Spec.SecurityLogging.AdditionalSink.ConvertFrom(&src.Spec.SecurityLogging.AdditionalSink)
	if srcSecurityLoggingDynamicAdditionalSinks := src.Spec.SecurityLogging.DynamicAdditionalSinks; srcSecurityLoggingDynamicAdditionalSinks != nil {
		dstSecurityLoggingDynamicAdditionalSinks := make([]string, len(srcSecurityLoggingDynamicAdditionalSinks))
		copy(dstSecurityLoggingDynamicAdditionalSinks, srcSecurityLoggingDynamicAdditionalSinks)
		op.Spec.SecurityLogging.DynamicAdditionalSinks = dstSecurityLoggingDynamicAdditionalSinks
	}

	// Copy the status
	op.Status.Version = src.Status.Version
	op.Status.Conditions = append(op.Status.Conditions, src.Status.Conditions...)
	return nil
}
