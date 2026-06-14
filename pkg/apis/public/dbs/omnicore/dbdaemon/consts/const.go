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

package consts

const (
	// ImportDumpFile is the file to store the downloaded dump file for import.
	ImportDumpFile = "import%s.dmp"
	ExportFileName = "%s.dmp"
	ExportLogName  = "%s.log"

	// Directory to hold ODS config files
	ConfigDir = "dbs_config"
	// Object storage config file
	StorageConfig = ConfigDir + "/storage.json"

	// Signal file for standbys
	StandbySignal = "standby.signal"

	BackupSignal   = "backup.signal"
	RestoreSignal  = "restore.signal"
	DateTimeFormat = "2006-01-02 15:04:05"
)
