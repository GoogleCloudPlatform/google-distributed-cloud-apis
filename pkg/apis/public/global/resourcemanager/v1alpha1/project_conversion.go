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

	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/global/resourcemanager/v1"
)

// ConvertTo converts this Project to the Hub version (v1).
func (p *Project) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*rmv1.Project)

	dst.ObjectMeta = *p.ObjectMeta.DeepCopy()

	dst.Status.MuxStatus = p.Status.MuxStatus
	dst.Status.Zones = convertZonesToHub(p.Status.Zones)

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (p *Project) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*rmv1.Project)

	p.ObjectMeta = *src.ObjectMeta.DeepCopy()

	p.Status.MuxStatus = src.Status.MuxStatus
	p.Status.Zones = convertZonesFromHub(src.Status.Zones)

	return nil
}

func convertZonesToHub(src []ProjectZoneStatus) []rmv1.ProjectZoneStatus {
	var zones []rmv1.ProjectZoneStatus
	for _, zone := range src {
		cpy := rmv1.ProjectZoneStatus{
			ZoneStatus: zone.ZoneStatus,
			ReplicaStatus: rmv1.ProjectReplicaStatus{
				Conditions:        zone.ReplicaStatus.Conditions,
				AvailableClusters: zone.ReplicaStatus.AvailableClusters,
				ErrorStatus:       zone.ReplicaStatus.ErrorStatus,
			},
		}

		zones = append(zones, cpy)
	}
	return zones
}

func convertZonesFromHub(src []rmv1.ProjectZoneStatus) []ProjectZoneStatus {
	var zones []ProjectZoneStatus
	for _, zone := range src {
		cpy := ProjectZoneStatus{
			ZoneStatus: zone.ZoneStatus,
			ReplicaStatus: ProjectReplicaStatus{
				Conditions:        zone.ReplicaStatus.Conditions,
				AvailableClusters: zone.ReplicaStatus.AvailableClusters,
				ErrorStatus:       zone.ReplicaStatus.ErrorStatus,
			},
		}

		zones = append(zones, cpy)
	}
	return zones
}
