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

import "constraints"

// Natural returns an Ordering that uses the natural order of the values.
func Natural[T constraints.Ordered]() Ordering[T] {
	return &comp[T]{compareOrdered[T], "ordering.Natural()"}
}

func compareOrdered[T constraints.Ordered](a, b T) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}
