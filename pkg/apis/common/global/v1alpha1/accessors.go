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
	"fmt"
)

// Takes a pointer to a slice of zone statuses and returns a
// ZoneStatusListAccessor.
// This function provides default implementations for the GetZones and SetZones
// methods.
// The pointer must be non-nil.
func NewZoneStatusListAccessor[T any, P interface {
	*T
	ZoneStatusInterface
}](zones *[]T) ZoneStatusListAccessor[T, P] {
	return ZoneStatusListAccessor[T, P]{
		zones: (*zoneStatusList[T, P])(zones),
	}
}

type ZoneStatusListAccessor[T any, P interface {
	*T
	ZoneStatusInterface
}] struct {
	zones *zoneStatusList[T, P]
}

func (a ZoneStatusListAccessor[T, P]) GetZones() ListMap[string, ZoneStatusInterface] {
	return a.zones
}

func (a ZoneStatusListAccessor[T, P]) SetZones(zones ListMap[string, ZoneStatusInterface]) error {
	// If the input is of the same type, no item-by-item copy is needed.
	if v, ok := zones.(*zoneStatusList[T, P]); ok {
		*a.zones = *v
		return nil
	}

	zoneStatuses := zones.Values()
	if zoneStatuses == nil {
		*a.zones = nil
		return nil
	}

	*a.zones = make([]T, len(zoneStatuses))
	for i, zoneStatus := range zoneStatuses {
		// If the item is of the same type, no field-by-field copy is needed.
		if v, ok := zoneStatus.(P); ok {
			(*a.zones)[i] = *v
			continue
		}

		P(&(*a.zones)[i]).SetName(zoneStatus.GetName())
		P(&(*a.zones)[i]).SetRolloutStatus(zoneStatus.GetRolloutStatus())
		if err := P(&(*a.zones)[i]).SetReplicaStatus(zoneStatus.GetReplicaStatus()); err != nil {
			return fmt.Errorf("set replica status for zone %q: %v", zoneStatus.GetName(), err)
		}
	}
	return nil
}

type zoneStatusList[T any, P interface {
	*T
	ZoneStatusInterface
}] []T

func (l *zoneStatusList[T, P]) Keys() []string {
	if *l == nil {
		return nil
	}

	keys := make([]string, len(*l))
	for i := range *l {
		keys[i] = P(&(*l)[i]).GetName()
	}
	return keys
}

func (l *zoneStatusList[T, P]) Values() []ZoneStatusInterface {
	if *l == nil {
		return nil
	}

	values := make([]ZoneStatusInterface, len(*l))
	for i := range *l {
		values[i] = P(&(*l)[i])
	}
	return values
}

func (l *zoneStatusList[T, P]) Get(name string) ZoneStatusInterface {
	for i := range *l {
		if P(&(*l)[i]).GetName() == name {
			return P(&(*l)[i])
		}
	}
	return nil
}

func (l *zoneStatusList[T, P]) Set(zoneStatus ZoneStatusInterface) error {
	var t *T
	for i := range *l {
		if P(&(*l)[i]).GetName() == zoneStatus.GetName() {
			t = &(*l)[i]
			break
		}
	}
	if t == nil {
		*l = append(*l, *new(T))
		t = &(*l)[len(*l)-1]
	}

	// If the input is of the same type, no field-by-field copy is needed.
	if v, ok := zoneStatus.(P); ok {
		*t = *v
		return nil
	}

	P(t).SetName(zoneStatus.GetName())
	P(t).SetRolloutStatus(zoneStatus.GetRolloutStatus())
	if err := P(t).SetReplicaStatus(zoneStatus.GetReplicaStatus()); err != nil {
		return fmt.Errorf("set replica status: %v", err)
	}
	return nil
}

func (l *zoneStatusList[T, P]) Clear() {
	if *l != nil {
		*l = (*l)[:0]
	}
}

// Lets you work with any typed API field.
// The setter takes any typed or untyped value and performs necessary
// conversions to set the field.
type TypedFieldAccessor[T any] struct {
	field *T
}

// Takes a pointer to an API field and returns a TypedFieldAccessor.
// This function provides default implementations for generic getters and
// setters.
// The pointer must be non-nil.
func NewTypedFieldAccessor[T any](field *T) TypedFieldAccessor[T] {
	return TypedFieldAccessor[T]{
		field: field,
	}
}

func (a TypedFieldAccessor[T]) GetField() T {
	return *a.field
}

func (a TypedFieldAccessor[T]) SetField(v any) error {
	typed, ok := v.(T)
	if !ok {
		// TODO(b/309005955): support untyped assignment.
		return fmt.Errorf("cannot set %T to %T", v, typed)
	}
	*a.field = typed
	return nil
}

// Takes a pointer to the .status field of a global mux resource and returns
// a MuxStatusAccessor.
// This function provides default implementations for the GetStatus and
// SetStatus methods.
// The pointer must be non-nil.
func NewMuxStatusAccessor[T any, P interface {
	*T
	MuxStatusInterface
}](status *T) MuxStatusAccessor[T, P] {
	return MuxStatusAccessor[T, P]{
		status: status,
	}
}

type MuxStatusAccessor[T any, P interface {
	*T
	MuxStatusInterface
}] struct {
	status *T
}

func (a MuxStatusAccessor[T, P]) GetStatus() MuxStatusInterface {
	return P(a.status)
}

func (a MuxStatusAccessor[T, P]) SetStatus(status MuxStatusInterface) error {
	// If the input is of the same type, no field-by-field copy is needed.
	if v, ok := status.(P); ok {
		*a.status = *v
		return nil
	}

	P(a.status).SetConditions(status.GetConditions())
	P(a.status).SetRollout(status.GetRollout())
	if err := P(a.status).SetZones(status.GetZones()); err != nil {
		return fmt.Errorf("set zones: %v", err)
	}
	return nil
}
