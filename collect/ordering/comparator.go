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
	"reflect"
	"runtime"
	"strings"
)

// Comparator compares its two arguments.
// It returns a negative integer, zero, or a positive integer as the first argument is less than,
// equal to, or greater than the second.
//
// In the foregoing description, the notation sgn(expression) designates the mathematical signum
// function, which is defined to return one of -1, 0, or 1 according to whether the value of
// expression is negative, zero or positive.
//
// The implementor must ensure that
//  sgn(Compare(x, y)) == -sgn(Compare(y, x))
// for all x and y.
//
// The implementor must also ensure that the relation is transitive:
//  ((Compare(x, y)>0) && (Compare(y, z)>0)) implies Compare(x, z)>0.
//
// Finally, the implementor must ensure that
//  Compare(x, y)==0 implies that sgn(Compare(x, z))==sgn(Compare(y, z))
// for all z.
//
// It is generally the case, but not strictly required that
//  (Compare(x, y) == 0) == (x == y)).
// Generally speaking, any Comparator that violates this condition should clearly indicate it.
type Comparator[T any] func(a, b T) int

// CaseInsensitive returns a Comparator that compares strings ignoring their case.
func CaseInsensitive() Comparator[string] {
	return func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	}
}

// From returns an Ordering based on an existing Comparator instance.
func From[T any](cmp Comparator[T]) Ordering[T] {
	return &comp[T]{cmp, "ordering.From(" + funcName(cmp) + ")"}
}

func funcName(i interface{}) string {
	n := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return n[strings.LastIndexByte(n, '/')+1:]
}
