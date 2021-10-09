/*
 Copyright 2021 The Goava authors

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package ordering

import (
	"fmt"
	"reflect"
)

// NilsLast returns an Ordering that treats nil as less than all other values and uses the
// given Ordering to compare non-nil values.
func NilsFirst[T any](o Ordering[T]) Ordering[T] {
	cmp := func(a, b T) int {
		aVal, bVal := reflect.ValueOf(a), reflect.ValueOf(b)
		if isNil(aVal) {
			if isNil(bVal) {
				return equal
			}
			return rightIsGreater
		}
		if isNil(bVal) {
			return leftIsGreater
		}
		return o.Compare(a, b)
	}

	return &comp[T]{cmp, fmt.Sprintf("ordering.NilsFirst(%v)", o)}
}

// NilsLast returns an Ordering that treats nil as greater than all other values and uses the
// given Ordering to compare non-nil values.
func NilsLast[T any](o Ordering[T]) Ordering[T] {
	cmp := func(a, b T) int {
		aVal, bVal := reflect.ValueOf(a), reflect.ValueOf(b)
		if isNil(aVal) {
			if isNil(bVal) {
				return equal
			}
			return leftIsGreater
		}
		if isNil(bVal) {
			return rightIsGreater
		}
		return o.Compare(a, b)
	}

	return &comp[T]{cmp, fmt.Sprintf("ordering.NilsLast(%v)", o)}
}

// isNil reports whether its argument v is a nil pointer.
//
// Implementation note: chan, func, map and slice are not fully supported.
// If v was created by calling ValueOf with an uninitialized
// interface variable i, isNil returns true.
func isNil(v reflect.Value) bool {
	if v.Kind() == reflect.Invalid {
		return true // a nil interface{}
	}
	// if the type is known, check whether it's a nil pointer
	return v.Kind() == reflect.Ptr && v.IsNil()
}
