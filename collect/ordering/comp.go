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

type comp[T any] struct {
	cmp  Comparator[T]
	name string
}

func (o comp[T]) Compare(a, b T) int {
	return o.cmp(a, b)
}

func (o *comp[T]) Reverse() Ordering[T] {
	return &reverse[T]{o}
}

func (o comp[T]) Min(a, b T) T {
	if o.Compare(a, b) <= 0 {
		return a
	}
	return b
}

func (o comp[T]) MinElement(es ...T) (min T, ok bool) {
	if len(es) == 0 {
		return min, false
	}
	min = es[0]
	for i := 1; i < len(es); i++ {
		min = o.Min(min, es[i])
	}
	return min, true
}

func (o comp[T]) Max(a, b T) T {
	if o.Compare(a, b) >= 0 {
		return a
	}
	return b
}

func (o comp[T]) MaxElement(es ...T) (max T, ok bool) {
	if len(es) == 0 {
		return max, false
	}
	max = es[0]
	for i := 1; i < len(es); i++ {
		max = o.Max(max, es[i])
	}
	return max, true
}

func (o comp[T]) LeastOf(es []T, k int) []T {
	if _, err := precond.CheckNonnegative(k, "k"); err != nil {
		panic(err)
	}

	cs := o.SortedCopy(es)
	if len(cs) > k {
		return cs[:k]
	}
	return cs
}

func (o comp[T]) GreatestOf(es []T, k int) []T {
	if _, err := precond.CheckNonnegative(k, "k"); err != nil {
		panic(err)
	}

	cs := Reverse[T](o).SortedCopy(es)
	if len(cs) > k {
		return cs[:k]
	}
	return cs
}

func (o comp[T]) SortedCopy(es []T) []T {
	cs := make([]T, len(es))
	copy(cs, es)
	sort.SliceStable(cs, func(i, j int) bool { return o.Compare(cs[i], cs[j]) < 0 })
	return cs
}

func (o comp[T]) IsOrdered(es []T) bool {
	if len(es) < 2 {
		return true
	}
	for i := 1; i < len(es); i++ {
		prev := es[i-1]
		next := es[i]
		if o.Compare(prev, next) > 0 {
			return false
		}
	}
	return true
}

func (o comp[T]) IsStrictlyOrdered(es []T) bool {
	if len(es) < 2 {
		return true
	}
	for i := 1; i < len(es); i++ {
		prev := es[i-1]
		next := es[i]
		if o.Compare(prev, next) >= 0 {
			return false
		}
	}
	return true
}

func (o comp[T]) String() string {
	return o.name
}
