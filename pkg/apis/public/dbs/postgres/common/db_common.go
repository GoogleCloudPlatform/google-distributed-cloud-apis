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
	"path/filepath"
	"time"

	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	DBEngine           = "postgresql"
	defaultStorageSize = "1Gi"
	// TODO(b/294089230): investigate better approach for obs disk sizing
	DefaultObsDiskSize = "80Gi"
	PostgresPort       = 5432

	// 2345 is the postgres uid/gid
	DefaultUID = int64(2345)
	DefaultGID = int64(2345)

	ObsMount  = "/obs"
	DataMount = "/pgdata"

	L2OperatorTLSVolume              = "l2-operator-tls"
	DBDaemonTLSVolume                = "dbd-tls"
	PostgresOperatorCertResourceName = "l2-operator"
	PostgresOperatorNsName           = "dbs-postgres-system"

	// A database is created by default on each postgresql server
	DefaultDatabase = "postgres"

	HAProbePeriod           = 30 * time.Second
	HAHealthcheckSyncPeriod = 30 * time.Second
)

var (
	DataDisk = occoreapi.DiskSpec{
		Name: occoreapi.DataDisk,
		Size: defaultStorageSize,
	}

	ObsDisk = occoreapi.DiskSpec{
		Name: occoreapi.ObsDisk,
		Size: DefaultObsDiskSize,
	}

	// DiagnosticDir is the directory holding database diagnostic logs
	DiagnosticDir = filepath.Join(ObsMount, "diagnostic")
	// DiagnosticLogsPattern is the path pattern for database diagnostic logs
	// The database log_filename in dbdaemon should be kept in sync with this value.
	DiagnosticLogsPattern = filepath.Join(DiagnosticDir, "postgresql.log")
	// LogRetention specifies how long the log will be purged, if a
	// log file’s modTime is earlier than LogRetention from now, it will be removed from the disk.
	LogRetention = 24 * time.Hour

	AuditLogsPattern = []string{filepath.Join(DiagnosticDir, "postgresql.audit")}

	LogArchiveDir = filepath.Join(DiagnosticDir, "archive")

	TLSVM = corev1.VolumeMount{
		MountPath: "/tls",
		Name:      "tls",
	}
)
