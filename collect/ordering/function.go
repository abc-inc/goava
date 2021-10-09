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

// OnResultOf returns a new Ordering on elements of type F, which orders elements by first applying
// a function to them, then comparing those results.
// For example, to compare values by their string forms, in a case-insensitive manner, use:
//
//  toString := func(val interface{}) string { return fmt.Sprint(val) }
//  ci := ordering.From(ordering.CaseInsensitive())
//  o := ordering.OnResultOf(toString, ci)
func OnResultOf[F any, T any](f func(val F) T, o Ordering[T]) Ordering[F] {
	cmp := func(a, b F) int {
		return o.Compare(f(a), f(b))
	}

	return &comp[F]{cmp, fmt.Sprintf("ordering.OnResultOf(%v)", o)}
}
