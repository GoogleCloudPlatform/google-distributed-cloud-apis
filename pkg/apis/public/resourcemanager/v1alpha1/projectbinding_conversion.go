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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1"
)

// ConvertTo converts this ProjectBinding to the Hub version (v1).
func (pb *ProjectBinding) ConvertTo(dstRaw conversion.Hub) error {
	v1pb := dstRaw.(*rmv1.ProjectBinding)
	v1pb.ObjectMeta = *pb.ObjectMeta.DeepCopy()
	v1pb.Spec.ClusterRef = rmv1.ProjectBindingClusterRef(pb.Spec.ClusterRef)

	if value, ok := pb.Annotations[v1pbSelectorAnnotation]; ok {
		if err := json.Unmarshal([]byte(value), &v1pb.Spec.Selector); err != nil {
			return fmt.Errorf("failed to unmarshall ProjectBinding Selector: %v", err)
		}
		delete(v1pb.Annotations, v1pbSelectorAnnotation)
		if len(v1pb.Annotations) == 0 {
			v1pb.Annotations = nil
		}
		return nil
	}

	if len(pb.Spec.Selector.MatchNames) == 0 {
		v1pb.Spec.Selector = rmv1.ProjectBindingSelector{}
	} else {
		v1pb.Spec.Selector = rmv1.ProjectBindingSelector{
			NameSelector: &rmv1.NameSelector{
				MatchNames: pb.Spec.Selector.MatchNames,
			},
		}
	}

	return nil
}

// ConvertFrom converts this ProjectBinding from the Hub version (v1) to this version.
func (pb *ProjectBinding) ConvertFrom(srcRaw conversion.Hub) error {
	v1pb := srcRaw.(*rmv1.ProjectBinding)

	pb.ObjectMeta = *v1pb.ObjectMeta.DeepCopy()
	pb.Spec.ClusterRef = ClusterRef(v1pb.Spec.ClusterRef)
	// Since all controllers use v1 pb, we just put v1 pb selector into annotation for round trip conversion.
	pb.Spec.Selector = Selector{MatchNames: []string{}}

	data, err := json.Marshal(v1pb.Spec.Selector)
	if err != nil {
		return fmt.Errorf("failed to marshall v1 ProjectBinding ProjectBindingSelector")
	}
	metav1.SetMetaDataAnnotation(&pb.ObjectMeta, v1pbSelectorAnnotation, string(data))
	return nil
}
