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

import "fmt"

// AllEqual returns an ordering which treats all values as equal, indicating "no ordering."
// This includes even nil values.
func AllEqual[T any]() Ordering[T] {
	return &comp[T]{func(a, b T) int { return 0 }, "ordering.AllEqual()"}
}

// Explicit returns an Ordering that compares values according to the order in which they are given
// to this method. Only values present in the argument list may be compared.
// This Comparator imposes a "partial ordering" over the type T.
// nil values in the argument list are not supported.
// If the values contain duplicates, this function panics.
//
// The returned Ordering panics when it receives an input parameter that isn't among the provided
// values.
func Explicit[T comparable](es ...T) Ordering[T] {
	m := make(map[T]int)
	for i, e := range es {
		if v, ok := m[e]; ok {
			panic(fmt.Sprintf("multiple entries with same key: %v=%d and %v=%d", e, v, e, i))
		}
		m[e] = i
	}

	rank := func(e T) int {
		if r, ok := m[e]; ok {
			return r
		}
		panic(fmt.Sprint("cannot compare value: ", e))
	}

	cmp := func(a, b T) int {
		return rank(a) - rank(b) // safe because both are nonnegative
	}
	return &comp[T]{cmp, fmt.Sprintf("ordering.Explicit(%v)", es)}
}
