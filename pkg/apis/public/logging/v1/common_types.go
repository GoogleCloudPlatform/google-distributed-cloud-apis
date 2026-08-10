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

// OperationalLogParser defines the predefined parser for log entries.
// +kubebuilder:validation:Enum=klog_text;klog_json;klogr;gdch_json;json
type OperationalLogParser string

const (
	// LogParserKlogText is the parser for klog text logs.
	LogParserKlogText OperationalLogParser = "klog_text"
	// LogParserKlogJson is the parser for klog json logs.
	LogParserKlogJson OperationalLogParser = "klog_json"
	// LogParserKlogr is the parser for klogr logs.
	LogParserKlogr OperationalLogParser = "klogr"
	// LogParserGdchJson is the parser for gdch_json logs.
	LogParserGdchJson OperationalLogParser = "gdch_json"
	// LogParserJson is the parser for json logs.
	LogParserJson OperationalLogParser = "json"
)
