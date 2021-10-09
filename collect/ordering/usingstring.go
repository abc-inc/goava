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
	"strings"
)

// UsingString returns an Ordering that compares objects by the natural ordering of their string
// representations.
func UsingString[T any]() Ordering[T] {
	// Optimize performance for common cases.
	toString := func(val interface{}) string {
		switch v := val.(type) {
		case fmt.Stringer:
			return v.String()
		case string:
			return v
		default:
			return fmt.Sprint(v)
		}
	}

	cmp := func(a, b T) int {
		return strings.Compare(toString(a), toString(b))
	}
	return &comp[T]{cmp, "ordering.UsingString()"}
}
