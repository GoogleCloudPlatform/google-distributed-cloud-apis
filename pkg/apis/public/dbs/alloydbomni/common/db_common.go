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

package common

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/mod/semver"
	corev1 "k8s.io/api/core/v1"

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

const (
	DBEngine           = "alloydbomni"
	defaultStorageSize = "1Gi"
	defaultObsDiskSize = "80Gi"
	AlloydbOmniPort    = 5432
	DefaultDatabase    = "postgres"
	DataMount          = "/mnt/disks/pgsql"

	PostStartupSuccess = "/tmp/post_startup_success"

	AlloydbomniScriptDir = "/alloydbomni-scripts"

	ObsMount = "/obs"

	L2OperatorTLSVolume                 = "l2-operator-tls"
	DBDaemonTLSVolume                   = "dbd-tls"
	AlloydbOmniOperatorCertResourceName = "l2-operator"
	AlloydbOmniOperatorNsName           = "dbs-alloydbomni-system"
	MLAgentContainerName                = "mlagent"
	MLAgentVertexAIKeyName              = "vertex-ai-key"
	MLAgentVertexAiKeyMountPath         = "/var/alloydb/config"
	GDCHRootCertMountPath               = "/var/alloydb/certs"

	VertexAIRegionKey                = "CLOUD_ML_REGION"
	MLAgentVertexAIRegion            = "us-central1"
	VertexAISecretLabel              = "vertexaijsonkey/secret"
	VertexAIOriginalSecretLabelValue = "vertex-ai-original-secret"
	VertexAISyncedSecretLabelValue   = "vertex-ai-sync-secret"
	VertexAISecretNamePrefix         = "vertex-ai-key-"
	GDCHRootCertVolume               = "gdch-root-cert"
	GDCHAlloyDBOmniMLAgentEnvVar     = "GDCH_ALLOYDB_OMNI"
)

var (
	OneOmniUID = int64(999)
	OneOmniGID = int64(999)

	MultiOmniUID = int64(2345)
	MultiOmniGID = int64(2345)

	OneOmniModeShortName   = "one"
	MultiOmniModeShortName = "multi"

	// OneOmniStartingVersion starts at 15.5.5
	// TODO(b/418016590): Handle major version dynamically.
	OneOmniStartingVersion = "15.5.5"
	OneOmniPgServicePath   = "/usr/lib/postgresql/15/bin"

	MultiOmniPgServicePath = "/bin"

	DataDisk = occoreapi.DiskSpec{
		Name: occoreapi.DataDisk,
		Size: defaultStorageSize,
	}

	ObsDisk = occoreapi.DiskSpec{
		Name: occoreapi.ObsDisk,
		Size: defaultObsDiskSize,
	}

	SmurfScriptVM = corev1.VolumeMount{
		Name:      "scripts-repo",
		MountPath: AlloydbomniScriptDir,
	}

	SmurfScriptV = corev1.Volume{
		Name:         "scripts-repo",
		VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
	}

	// Store dbdaemon
	AgentVM = corev1.VolumeMount{
		Name:      "agents-repo",
		MountPath: "/scripts",
	}

	AgentV = corev1.Volume{
		Name:         "agents-repo",
		VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
	}

	VarAlloyVM = corev1.VolumeMount{
		Name:      "var-alloydb",
		MountPath: "/var/alloydb",
	}

	VarAlloyV = corev1.Volume{
		Name:         "var-alloydb",
		VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
	}

	TLSVM = corev1.VolumeMount{
		MountPath: "/tls",
		Name:      "tls",
	}

	L2operatorTLSVM = corev1.VolumeMount{
		MountPath: fmt.Sprintf("/%s", L2OperatorTLSVolume),
		Name:      L2OperatorTLSVolume,
	}

	DBDaemonTLSVM = corev1.VolumeMount{
		MountPath: fmt.Sprintf("/%s", DBDaemonTLSVolume),
		Name:      DBDaemonTLSVolume,
	}
	MLAgentVM = corev1.VolumeMount{
		MountPath: MLAgentVertexAiKeyMountPath,
		Name:      MLAgentVertexAIKeyName,
	}

	GDCHRootCertVM = corev1.VolumeMount{
		MountPath: GDCHRootCertMountPath,
		Name:      GDCHRootCertVolume,
	}

	// DiagnosticDir is the directory holding database diagnostic logs
	DiagnosticDir = filepath.Join(ObsMount, "diagnostic")
	// DiagnosticLogsPattern is the path pattern for database diagnostic logs
	// The database log_filename in dbdaemon should be kept in sync with this value.
	DiagnosticLogsPattern = filepath.Join(DiagnosticDir, "postgresql.log")
	// LogRetention specifies how long the log will be purged, if a
	// log file’s modTime is earlier than LogRetention from now, it will be removed from the disk.
	LogRetention = 11 * time.Minute

	AuditLogsPattern = []string{filepath.Join(DiagnosticDir, "postgresql.audit")}

	InternalLogsPattern = []string{filepath.Join(DiagnosticDir, "postgresql.internal")}

	LogArchiveDir = filepath.Join(DiagnosticDir, "archive")
)

func IsOneOmniFromEnv() bool {
	return os.Getenv("OMNI_MODE") == "one"
}

func GetDefaultUIDFromEnv() int64 {
	if IsOneOmniFromEnv() {
		return OneOmniUID
	}
	return MultiOmniUID
}

func GetDefaultGIDFromEnv() int64 {
	if IsOneOmniFromEnv() {
		return OneOmniGID
	}
	return MultiOmniGID
}

func GetDefaultUIDFromOneOmniFlag(oneOmni bool) int64 {
	if oneOmni {
		return OneOmniUID
	}
	return MultiOmniUID
}

func GetDefaultGIDFromOneOmniFlag(oneOmni bool) int64 {
	if oneOmni {
		return OneOmniGID
	}
	return MultiOmniGID
}

func IsOneOmni(version string) bool {
	return semver.Compare(ToSemVerFormat(version), ToSemVerFormat(OneOmniStartingVersion)) >= 0
}

func GetPgServiceExecutablePathFromOneOmniFlag(oneOmni bool) string {
	if oneOmni {
		return OneOmniPgServicePath
	}
	return MultiOmniPgServicePath
}

func GetPgServiceExecutablePathFromEnv() string {
	return GetPgServiceExecutablePathFromOneOmniFlag(IsOneOmniFromEnv())
}

func ToSemVerFormat(version string) string {
	return "v" + version
}
