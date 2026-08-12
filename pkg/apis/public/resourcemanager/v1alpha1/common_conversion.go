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
	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1"
)

func ConvertClusterStatusFromV1(src []rmv1.ClusterStatus) []ClusterStatus {
	cpy := []ClusterStatus{}
	for _, cs := range src {
		cpy = append(cpy, ClusterStatus(cs))
	}
	return cpy
}

func ConvertClusterStatusToV1(src []ClusterStatus) []rmv1.ClusterStatus {
	cpy := []rmv1.ClusterStatus{}
	for _, cs := range src {
		cpy = append(cpy, rmv1.ClusterStatus(cs))
	}
	return cpy
}
