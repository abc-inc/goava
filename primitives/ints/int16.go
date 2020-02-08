/*
 * Copyright 2020 The Goava authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ints

import "github.com/abc-inc/goava/base/precond"

// Compare16 compares the two specified int16 values.
//
// It returns a negative value if a is less than b; a positive value if a is greater than b; or zero if they are equal.
func Compare16(a, b int16) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}

// CheckedCast16 returns the int16 value that is equal to value, if possible.
func CheckedCast16(v int64) (int16, error) {
	r := int16(v)
	if err := precond.CheckArgumentf(int64(r) == v, "out of range: %d", v); err != nil {
		return 0, err
	}
	return r, nil
}
