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
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// CreateZeroObjFromSameGroupVersion creates a new object of the given Kind by
// looking up the Kind name with the same GroupVersion as `fromObj` in the
// provided scheme.
func CreateZeroObjFromSameGroupVersion[T runtime.Object](scheme *runtime.Scheme, gvObj runtime.Object, kind string) (T, error) {
	var emptyObj T

	// Get the GVK of gvObj
	gvk, err := apiutil.GVKForObject(gvObj, scheme)
	if err != nil {
		return emptyObj, fmt.Errorf("getting gvk for object: %w", err)
	}

	// Use the sample GroupVersion but replace the Kind with the resource we want to create
	gvk.Kind = kind
	obj, err := scheme.New(gvk)

	// This is a hack for the function to work with the mockapis since the Kinds
	// defined there have a `Mock` prefix.
	if runtime.IsNotRegisteredError(err) && strings.HasPrefix(gvk.Group, "test.") {
		gvk.Kind = fmt.Sprintf("Mock%s", gvk.Kind)
		obj, err = scheme.New(gvk)
	}

	if err != nil {
		return emptyObj, err
	}
	typedObj, ok := obj.(T)
	if !ok {
		return emptyObj, fmt.Errorf("wrong object kind received from schema, wanted %T but got %T", emptyObj, obj)
	}
	return typedObj, nil
}

// LabelPrefixFor returns the prefix that should be used for labels set on
// the object. The prefix value will be the first segment of the Group of the
// object.
func LabelPrefixFor(obj runtime.Object, scheme *runtime.Scheme) (string, error) {
	var err error
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Group == "" {
		gvk, err = apiutil.GVKForObject(obj, scheme)
		if err != nil {
			return "", err
		}
	}
	return strings.Split(gvk.Group, ".")[0], nil
}
