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

/*
Copyright 2022.
*/
package v1

import "strings"

// DatabaseType is type of database that can serve as a migration source database.
// Currently only postgresql and alloydbomni are supported.
type DatabaseType string

var DatabaseTypeToEngineType = map[DatabaseType]EngineType{
	PostgreSQL_DatabaseType:  PostgreSQL,
	AlloyDBOmni_DatabaseType: AlloyDBOmni,
}

const (
	PostgreSQL_DatabaseType  DatabaseType = "postgresql"
	AlloyDBOmni_DatabaseType DatabaseType = "alloydbomni"
)

func DatabaseEngineType(t string) EngineType {
	return DatabaseTypeToEngineType[DatabaseType(strings.ToLower(t))]
}
