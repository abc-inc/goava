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

// Lexicographical returns a new Ordering, which sorts iterables by comparing corresponding elements
// pairwise until a nonzero result is found; imposes "dictionary order".
// If the end of slice is reached, but not the other, the shorter slice is considered to be less
// than the longer one. For example, a lexicographical natural ordering over integers considers
//  [] < [1] < [1, 1] < [1, 2] < [2]
func Lexicographical[T any](elCmp Comparator[T]) Ordering[[]T] {
	cmp := func(as, bs []T) int {
		for i, a := range as {
			if i >= len(bs) {
				return leftIsGreater // because it's longer
			}
			if r := elCmp(a, bs[i]); r != 0 {
				return r
			}
		}
		if len(bs) > len(as) {
			return rightIsGreater // because it's longer
		}
		return equal
	}

	return &comp[[]T]{cmp, fmt.Sprintf("ordering.Lexicographical(%v)", funcName(elCmp))}
}
