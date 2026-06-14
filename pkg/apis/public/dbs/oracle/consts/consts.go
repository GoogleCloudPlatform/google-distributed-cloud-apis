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

// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package consts provides common Oracle constants across the entire Data Plane.
package consts

import occonst "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/dbdaemon/consts"

// Listener is an Oracle listener struct
type Listener struct {
	LType    string
	Port     int32
	Local    bool
	Protocol string
}

// OraDir wraps up the host and oracle names for a specific directory. e.g.
// /u03/app/oracle/<dbname>/dpdump and PDB_DATA_PUMP_DIR
type OraDir struct {
	Linux  string
	Oracle string
}

const (
	// DefaultHealthAgentPort is Health Agent's default port number.
	DefaultHealthAgentPort = 3201

	// DefaultConfigAgentPort is Config Agent's default port number.
	DefaultConfigAgentPort = 3202

	// DefaultMonitoringAgentPort is the default port where the oracle exporter runs
	DefaultMonitoringAgentPort = 9161

	// DefaultDBDaemonMetricsPort is the DBDaemon's default metrics port number.
	DefaultDBDaemonMetricsPort = 9187

	// Localhost is a general localhost name.
	Localhost = "localhost"

	// DomainSocketFile is meant for the agents to communicate to the Database Daemon.
	DomainSocketFile = "/var/tmp/dbdaemon.sock"

	// ProxyDomainSocketFile is meant for the database daemon to communicate to the database daemon proxy.
	ProxyDomainSocketFile = "/var/tmp/dbdaemon_proxy.sock"

	// SecureListenerPort is a secure listener port number.
	SecureListenerPort = 6021
	// SSLListenerPort is an SSL listener port number.
	SSLListenerPort = 3307

	// OpenPluggableDatabaseSQL is used to open pluggable database (e.g. after CDB start).
	OpenPluggableDatabaseSQL = "alter pluggable database all open"

	// ListPDBsSQL lists pluggable databases excluding a root container.
	ListPDBsSQL = "select name from v$containers where name !='CDB$ROOT'"

	// ListPluggableDatabaseExcludeSeedSQL is used to list pluggable databases exclude PDB$SEED
	ListPluggableDatabaseExcludeSeedSQL = "select pdb_name from dba_pdbs where pdb_name!='PDB$SEED'"

	// GetDataAmountInBytes is used to get the bulk of the amount of stored db data on the disk
	GetDataAmountInBytes = "select ( select sum(bytes) data_size from dba_data_files ) + ( select sum(BLOCK_SIZE*FILE_SIZE_BLKS) controlfile_size from v$controlfile ) \"" + GetDataAmountInBytesKey + "\" from dual"

	// GetDataAmountInBytesKey is the column name from the GetDataAmountInBytes query
	GetDataAmountInBytesKey = "BYTES"

	// Oracle18c Version for Oracle 18c XE
	Oracle18c = "18c"

	// SourceDatabaseHost is the hostname used during image build process.
	SourceDatabaseHost = "ol7-db12201-gi-cdb-docker-template-vm"

	// CharSet is the supported character set for Oracle databases.
	CharSet = "AL32UTF8"

	// SecurityUser is the user for lockdown triggers.
	SecurityUser = "gcsql$security"

	// PDBLoaderUser is the user for impdb/expdp operations by end users.
	PDBLoaderUser = "gcsql$pdbloader"

	//  HealthcheckUser is the user for the ods liveness probe healthcheck.
	HealthcheckUser = "gcsql$healthcheck"

	// MonitoringAgentName is the container name for the monitoring agent.
	MonitoringAgentName = "oracle-monitoring"

	// DefaultExitErrorCode is default exit code
	DefaultExitErrorCode = 128

	// RMANBackup is the oracle rman command for taking backups.
	RMANBackup = "backup"

	// DefaultPdb is the name of the PDB which gets created during instance provisioning.
	DefaultPdb = "GPDB"

	// Default CDB name
	DefaultCdb = "GCLOUD"

	// LockDownProfile is used to enforce an additional layer of security in event of privilege escalation.
	LockDownProfile = "sec_profile"

	// FedRAMPPolicy is the audit policy for FedRAMP compliance.
	FedRAMPPolicy = "FEDRAMP_POLICY"

	// StartupParamFileName is file stores database dbs custom startup parameter values.
	StartupParamFileName = "30startup.txt"

	// DynamicStartupParamFileName is file stores database dbs custom startup parameter values that are variable at re-initialization.
	DynamicStartupParamFileName = "35dynamicstartup.txt"

	// StartupParamFileName is file stores user specified database parameter values.
	UserParamFileName = "50user.txt"
)

var (
	// ProvisioningDoneFile is a flag at the end of provisioning.
	ProvisioningDoneFile = "/tmp/provisioning_successful"

	// ProvisioningDiskFile is a flag created at the end of provisioning.
	// this is placed on the PD storage so that on recreate, the whole bootstrap doesn't re-run.
	ProvisioningDiskFile = "/u02/app/oracle/provisioning_disk"

	// PendingRestartFile is a flag created at after static database parameters are set.
	// PendingRestartFile indicate current database parameters in the memory mismatches the spfile files.
	// This flag is removed once database startup.
	PendingRestartFile = "/u02/app/oracle/pending_restart"

	// SeededImageFile indicates that a CDB exists in the image or one of the volumes mounted to it
	SeededImageFile = "/tmp/seeded_image"

	// UnseededImageFile indicates that a CDB does not exist in the image, nor in any volume mounted to it
	UnseededImageFile = "/tmp/unseeded_image"

	// UnifiedAuditing indicates that the Unified Auditing option is enabled
	// The file is saved in the PD storage so that on recreate, the check on unified auditing doesn't re-run
	UnifiedAuditing = "/u02/app/oracle/unified_auditing"

	// PrestopShellFile is the command that runs on the Oracle database container before it stops
	PrestopShellFile = "/" + DataMount + "/prestop.sh"

	// DeferralFile saves the data of the grace period between provisioning and read write readiness
	DeferralFile = "/tmp/read_write_grace_period.json"

	// RecoveryPendingFile is a flag created after bootstrapping a CDB.
	RecoveryPendingFile = "/tmp/recovery_pending"

	// SECURE is the name of the secure tns listener
	SECURE = "SECURE"

	// SSL is the name of the SSL tns listener
	SSL = "SSL"

	// ListenerNames is the list of listeners
	ListenerNames = map[string]*Listener{
		"SECURE": {
			LType:    SECURE,
			Port:     SecureListenerPort,
			Local:    true,
			Protocol: "TCP",
		},
		"SSL": {
			LType:    SSL,
			Port:     SSLListenerPort,
			Protocol: "TCPS",
		},
	}

	// DpdumpDir is the Impdp/Expdp directory and oracle directory name.
	// Linux is relative to the PDB PATH_PREFIX.
	DpdumpDir = OraDir{Linux: "dmp", Oracle: "PDB_DATA_PUMP_DIR"}

	// OraGroup is the group that owns the database software.
	OraGroup = []string{"dba", "oinstall"}

	// OraTab is the oratab file path.
	OraTab = "/etc/oratab"

	// OraUser is the owner of the database and database software.
	OraUser = "oracle"

	// OracleBase is the Oracle base path.
	OracleBase = "/u02/app/oracle"

	// DataDir is the directory where datafiles exists.
	DataDir = "/%s/app/oracle/oradata/%s"

	// PDBDataDir is the directory where PDB datafiles exists.
	PDBDataDir = DataDir + "/%s/data"

	// PDBSeedDir is the directory where the SEED datafiles exists.
	PDBSeedDir = DataDir + "/pdbseed"

	// PDBPathPrefix is the directory where PDB data directory exists.
	PDBPathPrefix = DataDir + "/%s"

	// ConfigDir is where the spfile, pfile and pwd file are persisted.
	ConfigDir = "/%s/app/oracle/oraconfig/%s"

	// ParamBackUpDir is where the historical database parameters config dir are stored
	ParamBackUpDir = "/%s/app/oracle/oraconfig/%s/parambackup"

	// ActiveParamConfigDir stores current database parameters config files
	ActiveParamConfigDir = "/%s/app/oracle/oraconfig/%s/ora.conf.d"

	// RecoveryAreaDir is where the flash recovery area will be.
	RecoveryAreaDir = "/%s/app/oracle/fast_recovery_area/%s"

	// DataMount is the PD mount where the data is persisted.
	DataMount = "u02"

	// LogMount is the PD mount where the logs are persisted.
	LogMount = "u03"

	ObsMount = "/obs"

	// ListenerDir is the listener directory.
	ListenerDir = "/%s/app/oracle/oraconfig/network"

	// ScriptDir is where the scripts are located on the container image.
	ScriptDir = "/agents"

	// WalletDir is where the SSL Certs are stored.
	WalletDir = "/u02/app/oracle/wallet"

	// ServerCert is the file name of the server certificate.
	ServerCert = "tls.crt"

	// ServerPrivateKey is the file name of the server's private key.
	ServerPrivateKey = "tls.key"

	// ServerCACert is the file name of the server's CA certificate.
	ServerCACert = "ca.crt"

	// ServerEncryptedPrivateKey is the file name of the password encrypted private key required for Oracle wallet.
	ServerEncryptedPrivateKey = "server_enc_pkey.pem"

	IntermediateCerts = "intermediate.crt"

	// OracleDir is where the env file is located
	OracleDir = "/home/oracle"

	// DefaultRMANDir sets the default rman backup directory
	DefaultRMANDir = "/u03/app/oracle/rman"

	// RMANStagingDir sets the staging directory for rman backup to GCS.
	RMANStagingDir = "/u03/app/oracle/rmanstaging"

	// CertDir is where the SSL Certs are stored.
	CertDir = "/tls"

	// AuditFileLoc is where the audit logs gets extracted.
	AuditFileLoc = "/obs/app/oracle/audit/audit.log"

	// RedoLogArchiveDir is the redo log backup directory.
	RedoLogArchiveDir = "/u03/app/oracle/fast_recovery_area/%s/%s/archivelog"

	// ControlFileArchive is the control file backup path.
	ControlFileArchive = "/u03/app/oracle/fast_recovery_area/%s/%s/control/control.ctl"

	// PITRMetricsFile is where some constants for metrics are stored to be persisted
	PITRMetricsFile = "/obs/app/oracle/pitrmetrics/pitrmetrics.json"

	// OracleStorageConfigFile is the remote storage config file path.
	OracleStorageConfigFile = "/u02/" + occonst.StorageConfig

	// ParameterConfigFileByPriority stores the file path of each database parameter fragment and its priority.
	ParameterConfigFileByPriority = map[int]string{
		30: StartupParamFileName,
		35: DynamicStartupParamFileName,
		50: UserParamFileName,
	}

	// RMANRecoverDateCmd is an RMAN command to perform recovery to a point in time
	RMANRecoverDateCmd = `recover database until time "to_date('%s', 'YYYY-MM-DD HH24:MI:SS')";`

	// RMANRecoverSCNCmd is an RMAN command to perform recovery to a redo log system change number
	RMANRecoverSCNCmd = `recover database until scn = %s;`

	// RMANRestoreTemplate is an RMAN script used to restore the database.
	RMANRestoreTemplate = `
	run {
		restore controlfile from '%s';
		startup mount;
		catalog start with '%s' noprompt;
	}
	`

	// RMANRecoverTemplate is an RMAN script used to recover the database to a point in time.
	RMANRecoverTemplate = `
	run {
		%s
		alter database open resetlogs;
		alter pluggable database all open;
	}
	`
)
