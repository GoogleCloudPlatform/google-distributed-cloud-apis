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

	v1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/observability/v1"
)

// ConvertTo converts this Dashboard to the Hub version (v1).
func (src *Dashboard) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Dashboard)
	dst.TypeMeta = src.TypeMeta
	dst.ObjectMeta = src.ObjectMeta
	//Spec
	dst.Spec.ConfigMapRef.Name = src.Spec.ConfigMapRef.Name
	dst.Spec.ConfigMapRef.Namespace = src.Spec.ConfigMapRef.Namespace
	dst.Spec.ConfigMapRef.Key = src.Spec.ConfigMapRef.Key
	dst.Spec.Foldername = src.Spec.Foldername

	//Status
	dst.Status.Conditions = append(dst.Status.Conditions, src.Status.Conditions...)
	dst.Status.Dashboards = make(map[string]v1.DashboardInfo)
	for key, dashInfo := range src.Status.Dashboards {
		srcDashInfo := v1.DashboardInfo{
			UID: dashInfo.UID,
			URL: dashInfo.URL,
		}
		dst.Status.Dashboards[key] = srcDashInfo
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (src *Dashboard) ConvertFrom(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Dashboard)
	src.TypeMeta = dst.TypeMeta
	src.ObjectMeta = dst.ObjectMeta
	//Spec
	src.Spec.ConfigMapRef.Name = dst.Spec.ConfigMapRef.Name
	src.Spec.ConfigMapRef.Namespace = dst.Spec.ConfigMapRef.Namespace
	src.Spec.ConfigMapRef.Key = dst.Spec.ConfigMapRef.Key
	src.Spec.Foldername = dst.Spec.Foldername

	//Status
	src.Status.Conditions = append(src.Status.Conditions, dst.Status.Conditions...)
	src.Status.Dashboards = make(map[string]DashboardInfo)
	for key, dashInfo := range dst.Status.Dashboards {
		srcDashInfo := DashboardInfo{
			UID: dashInfo.UID,
			URL: dashInfo.URL,
		}
		src.Status.Dashboards[key] = srcDashInfo
	}

	return nil
}
