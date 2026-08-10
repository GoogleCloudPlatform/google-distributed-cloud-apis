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
	corev1 "k8s.io/api/core/v1"
)

const (
	DisableLoadBalancer   = "dbcluster.dbadmin.goog/disableLB"
	AnnotationRaasEnabled = "dbcluster.dbadmin.goog/raasEnabled"

	// AnnotationKeyAutoManageSettingTimestamp is used on ReplicationConfigs to
	// indicate the timestamp when the synchronous field in physical upstream config spec
	// is updated to use auto_managed setting.
	AnnotationKeyAutoManageSettingTimestamp = "replication.dbadmin.goog/automanage-timestamp"

	// LabelDBCluster is used to refer to the database cluster a
	// ReplicationConfig belongs to.
	LabelDBCluster = "dbs.dbadmin.goog/dbcluster"

	// LabelFeature is used to refer to which feature is managing this replication config.
	LabelFeature = "dbs.dbadmin.goog/feature"

	// LabelSecondInstance is used to refer to which instance at the other end of the connection.
	LabelSecondInstance = "dbs.dbadmin.goog/secondInstance"

	// ReplicationConfigFeatureHA is the value of the label LabelFeature on a ReplicationConfig if
	// HA is using this replication config
	ReplicationConfigFeatureHA = "HA"

	// ReplicationConfigFeatureReadPool is the value of the label LabelFeature on a ReplicationConfig if
	// ReadPool is using this replication config
	ReplicationConfigFeatureReadPool = "ReadPool"
)

//+kubebuilder:object:generate=true

// ReplicationSpec defines replication connections to other database instances.
type ReplicationSpec struct {
	// Profiles contains the collection of replication profiles.
	// +optional
	Profiles []ReplicationProfileSpec `json:"profiles,omitempty"`
}

//+kubebuilder:object:generate=true

// ReplicationProfileSpec is one replication connection to another database instance.
type ReplicationProfileSpec struct {
	// Name of the profile
	// +required
	Name string `json:"name"`

	// Host on the other side of the connection
	// +optional
	Host string `json:"host,omitempty"`

	// Port on the other side of the connection
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	// Username is the name of user to connect to another database instance
	// +optional
	Username string `json:"username,omitempty"`

	// Password is a reference to the secret that contains user password
	// +optional
	Password *corev1.SecretReference `json:"password,omitempty"`

	// PasswordResourceVersion specifies the password secret version
	// +optional
	PasswordResourceVersion string `json:"passwordResourceVersion,omitempty"`

	// CertificateReference refers to a secret to be used for TLS
	// +optional
	CertificateReference *CertificateRef `json:"certificateReference,omitempty"`

	// Role is the replication role of this instance to this replication connection.
	// +optional
	// +kubebuilder:validation:Enum=Upstream;Downstream
	Role ReplicationRole `json:"role,omitempty"`

	// IsSynchronous is true for synchronous replication connections
	// +required
	// +kubebuilder:default=false
	IsSynchronous bool `json:"isSynchronous,omitempty"`

	// Type is physical or logical
	// +required
	// +kubebuilder:validation:Enum=Logical;Physical
	Type ReplicationType `json:"type"`

	// IsActive is true for connections currently enabled, false pauses the connection
	// +optional
	IsActive bool `json:"isActive,omitempty"`
}

// ReplicationRole is the supported types of replication role
type ReplicationRole string

// ReplicationType is the supported types of replication
type ReplicationType string

const (
	ReplicationTypeLogical    ReplicationType = "Logical"
	ReplicationTypePhysical   ReplicationType = "Physical"
	ReplicationRoleUpstream   ReplicationRole = "Upstream"
	ReplicationRoleDownstream ReplicationRole = "Downstream"
)

//+kubebuilder:object:generate=true

// ReplicationStatus is the status of all replication connections on an instance.
type ReplicationStatus struct {
	EntityStatus `json:",inline"`
	Profiles     []ReplicationProfileStatus `json:"profiles,omitempty"`
}

//+kubebuilder:object:generate=true

// ReplicationProfileStatus is the status of one individual replication connection
type ReplicationProfileStatus struct {
	EntityStatus           `json:",inline"`
	ReplicationProfileSpec `json:"profile,omitempty"`
}

// FindProfile finds the profile with the matching name, otherwise returns nil
func (r ReplicationSpec) FindProfile(name string) *ReplicationProfileSpec {
	for i := 0; i < len(r.Profiles); i++ {
		if r.Profiles[i].Name == name {
			return &r.Profiles[i]
		}
	}
	return nil
}
