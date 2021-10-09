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
	"github.com/abc-inc/goava/base/precond"
	"sort"
)

// Reverse returns the reverse of the Ordering.
func Reverse[T any](o Ordering[T]) Ordering[T] {
	if r, ok := o.(interface{ Reverse() Ordering[T] }); ok {
		return r.Reverse()
	}
	return &reverse[T]{o}
}

// reverse is an Ordering that uses the reverse of a given order.
type reverse[T any] struct {
	o Ordering[T]
}

func (r reverse[T]) Compare(a, b T) int {
	return r.o.Compare(b, a)
}

func (r *reverse[T]) Reverse() Ordering[T] {
	return r.o
}

// Override the min/max methods to "hoist" delegation outside loops.

func (r reverse[T]) Min(a, b T) T {
	return r.o.Max(a, b)
}

func (r reverse[T]) MinElement(es ...T) (T, bool) {
	return r.o.MaxElement(es...)
}

func (r reverse[T]) Max(a, b T) T {
	return r.o.Min(a, b)
}

func (r reverse[T]) MaxElement(es ...T) (T, bool) {
	return r.o.MinElement(es...)
}

func (r reverse[T]) LeastOf(es []T, k int) []T {
	if _, err := precond.CheckNonnegative(k, "k"); err != nil {
		panic(err)
	}

	cs := r.SortedCopy(es)
	if len(cs) > k {
		return cs[:k]
	}
	return cs
}

func (r reverse[T]) GreatestOf(es []T, k int) []T {
	if _, err := precond.CheckNonnegative(k, "k"); err != nil {
		panic(err)
	}

	cs := Reverse[T](r).SortedCopy(es)
	if len(cs) > k {
		return cs[:k]
	}
	return cs
}

func (r reverse[T]) SortedCopy(es []T) []T {
	cs := make([]T, len(es))
	copy(cs, es)
	sort.SliceStable(cs, func(i, j int) bool { return r.Compare(cs[i], cs[j]) < 0 })
	return cs
}

func (r reverse[T]) IsOrdered(es []T) bool {
	if len(es) < 2 {
		return true
	}
	for i := 1; i < len(es); i++ {
		prev := es[i-1]
		next := es[i]
		if r.Compare(prev, next) > 0 {
			return false
		}
	}
	return true
}

func (r reverse[T]) IsStrictlyOrdered(es []T) bool {
	if len(es) < 2 {
		return true
	}
	for i := 1; i < len(es); i++ {
		prev := es[i-1]
		next := es[i]
		if r.Compare(prev, next) >= 0 {
			return false
		}
	}
	return true
}

func (r reverse[T]) String() string {
	return "ordering.Reverse(" + r.o.String() + ")"
}
