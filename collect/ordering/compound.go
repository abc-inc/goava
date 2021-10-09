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

// Compound returns an Ordering, which tries each given Comparator in order until a non-zero result
// is found, returning that result, and returning zero only if all Comparators return zero.
func Compound[T any](cs ...Comparator[T]) Ordering[T] {
	cmp := func(a, b T) int {
		for _, c := range cs {
			if r := c(a, b); r != 0 {
				return r
			}
		}
		return 0
	}

	return &comp[T]{cmp, fmt.Sprintf("ordering.Compound(%v)", cs)}
}
