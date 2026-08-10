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

const (
	// KeyUIDAnnotation is set by controllers on Secrets to identify the UID of
	// the access key currently stored in the Secret.
	KeyUIDAnnotation = "object.gdc.goog/key-uid"

	// SubjectAnnotation is set by controllers on Secrets to identify the
	// subject that a resource corresponds to.
	SubjectAnnotation = "object.gdc.goog/subject"

	// SubjectAnnotationIndexField is used as a field name on which controller manager
	// adds index.
	SubjectAnnotationIndexField = ".metadata.annotations." + SubjectAnnotation

	// AlphaBucketQuotaAnnotation is used to configure a StorageGRID bucket quota limit.
	// Note this is a temporary API and will be remove in the future.
	AlphaBucketQuotaAnnotation = "object.gdc.goog/alpha-bucket-quota"

	// AlphaBucketQuotaAppliedAnnotation is set by the reconciler when the quota has been successfully applied.
	AlphaBucketQuotaAppliedAnnotation = "object.gdc.goog/alpha-bucket-quota-applied"

	// AlphaBucketQuotaClearSetting is a value that causes the quota to be cleared.
	AlphaBucketQuotaClearSetting = "clear"

	// AlphaBucketQuotaUnsupportedSetting is a value set when the backend does not support quotas.
	AlphaBucketQuotaUnsupportedSetting = "unsupported"

	// AlphaBucketQuotaErrorSetting is a value set when applying the quota encountered an error.
	AlphaBucketQuotaErrorSetting = "error"
)
