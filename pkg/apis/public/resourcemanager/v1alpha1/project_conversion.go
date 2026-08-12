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
	"encoding/json"
	"fmt"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1"
)

// ConvertTo converts this Project to the Hub version (v1).
func (p *Project) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.Project)

	dst.ObjectMeta = *p.ObjectMeta.DeepCopy()

	// If the project has a cluster selector that does not conform to v1
	// default, marshal it into an annotation for backward compatibility.
	if !reflect.DeepEqual(p.Spec.ClusterSelector, &ClusterSelector{}) {
		data, err := json.Marshal(p.Spec.ClusterSelector)
		if err != nil {
			return fmt.Errorf("failed to marshall cluster selector: %v", err)
		}
		metav1.SetMetaDataAnnotation(&dst.ObjectMeta, rmv1.ClusterSelectorAnnotation, string(data))
	}

	dst.Status.PropagatedName = p.Status.PropagatedName
	dst.Status.Conditions = p.Status.Conditions
	dst.Status.ErrorStatus = p.Status.ErrorStatus
	for _, pcs := range p.Status.Clusters {
		tmpPCS := rmv1.ProjectClusterStatus{
			ClusterStatus:      rmv1.ClusterStatus(pcs.ClusterStatus),
			EgressNATIPAddress: pcs.EgressNATIPAddress,
		}
		dst.Status.Clusters = append(dst.Status.Clusters, tmpPCS)
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (p *Project) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.Project)

	p.ObjectMeta = *src.ObjectMeta.DeepCopy()

	if value, ok := src.Annotations[rmv1.ClusterSelectorAnnotation]; ok {
		if err := json.Unmarshal([]byte(value), &p.Spec.ClusterSelector); err != nil {
			return fmt.Errorf("failed to unmarshall cluster selector: %v", err)
		}
		delete(p.Annotations, rmv1.ClusterSelectorAnnotation)
		if len(p.Annotations) == 0 {
			p.Annotations = nil
		}
	} else {
		// By default, a v1 project should propagate to no cluster.
		p.Spec.ClusterSelector = &ClusterSelector{}
	}

	p.Status.PropagatedName = src.Status.PropagatedName
	p.Status.Conditions = src.Status.Conditions
	p.Status.ErrorStatus = src.Status.ErrorStatus

	for _, pcs := range src.Status.Clusters {
		tmpPCS := ProjectClusterStatus{
			ClusterStatus:      ClusterStatus(pcs.ClusterStatus),
			EgressNATIPAddress: pcs.EgressNATIPAddress,
		}
		p.Status.Clusters = append(p.Status.Clusters, tmpPCS)
	}

	return nil
}
