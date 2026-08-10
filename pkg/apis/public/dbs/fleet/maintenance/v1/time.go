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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Monday    DayOfWeek = "Monday"
	Tuesday   DayOfWeek = "Tuesday"
	Wednesday DayOfWeek = "Wednesday"
	Thursday  DayOfWeek = "Thursday"
	Friday    DayOfWeek = "Friday"
	Saturday  DayOfWeek = "Saturday"
	Sunday    DayOfWeek = "Sunday"
)

var weekdayMap = map[DayOfWeek]time.Weekday{
	Monday:    time.Monday,
	Tuesday:   time.Tuesday,
	Wednesday: time.Wednesday,
	Thursday:  time.Thursday,
	Friday:    time.Friday,
	Saturday:  time.Saturday,
	Sunday:    time.Sunday,
}

// +kubebuilder:validation:Enum=Monday;Tuesday;Wednesday;Thursday;Friday;Saturday;Sunday
type DayOfWeek string

func (d DayOfWeek) Weekday() time.Weekday {
	return weekdayMap[d]
}

// TimeOfDay is the time in the text format "HH:mm"
// +kubebuilder:validation:Pattern=`^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`
type TimeOfDay string

type DateTimeRange struct {
	// Deny period start date.
	// Date matching this period will have to be the same or after the start.
	Start metav1.Time `json:"startDateTime"`
	// Deny period end date.
	// Date matching this period will have to be before the end.
	End metav1.Time `json:"endDateTime"`
}
