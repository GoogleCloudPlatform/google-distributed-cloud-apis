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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LifecycleRuleStatus string

const (
	Enabled  LifecycleRuleStatus = "Enabled"
	Disabled LifecycleRuleStatus = "Disabled"
)

type LifecyclePolicy struct {
	// Defines whether the user wants to enable custom lifecycle policy on the bucket.
	// +kubebuilder:default:=false
	Enable bool `json:"enable,omitempty"`
	// Consists of one or more lifecycle configuration rules that can take expiration action on
	// objects in the bucket.
	LifecycleRules []*LifecycleRule `json:"lifecycleRules,omitempty"`
}

type LifecycleRule struct {
	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID *string `json:"id"`
	// Status of the lifecycle rule. Indicate whether this rule takes action.
	// The status can always be changed, and only Enabled rule would take effect.
	Status *LifecycleRuleStatus `json:"status"`
	// Expiration behavior for objects in current version.
	// Either Expiration or NoncurrentExpiration, or both should be provided.
	// Rule without expiration behavior would be ignored.
	Expiration *LifecycleExpiration `json:"expiration,omitempty"`
	// Expiration behavior for objects noncurrent.
	NoncurrentExpiration *LifecycleNoncurrentExpiration `json:"noncurrentExpiration,omitempty"`
	// Filter for the rule. Empty filer means the rule apply to all objects in the bucket.
	Filter *LifecycleRuleFilter `json:"filter,omitempty"`
	// Indicates whether the lifecycle rule will remove an expired object delete marker.
	// In a versioned bucket, an object delete marker is considered expired if it is the
	// current and only version of the object, with no remaining noncurrent versions.
	// When set to true, the lifecycle policy automatically removes these expired delete markers.
	// This field cannot be specified together with Expiration.Days or Expiration.Date.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty"`
}

type LifecycleExpiration struct {
	// Specific date when the objects should be deleted.
	Date *metav1.Time `json:"date,omitempty"`
	// Number of days this object is subject to the rule.
	Days *int64 `json:"days,omitempty"`
}

type LifecycleNoncurrentExpiration struct {
	// Number of days an object is noncurrent before lifecycle rules takes the action.
	NoncurrentDays *int64 `json:"noncurrentDays"`
}

// Defines the filter that can be set on lifecycle rule.
type LifecycleRuleFilter struct {
	// Match objects with provided prefix.
	// Empty string of prefix means apply to all objects in the bucket.
	PrefixFilter *string `json:"prefixFilter,omitempty"`
	// Match objects with provided tags.
	// The tag on object must match both the key and value exactly.
	TagFilters []*Tag `json:"tagFilters,omitempty"`
}

// Defines the tags on object that can be used by lifecycle filter.
type Tag struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}
