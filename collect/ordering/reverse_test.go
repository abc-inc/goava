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
	"github.com/abc-inc/goava/primitives/ints"
	. "github.com/stretchr/testify/require"
	"math"
	"testing"
)

type wrapper struct {
	v string
}

func numCmp(a, b interface{}) int {
	return ints.Compare(a.(int), b.(int))
}

func TestReverseSortedCopy(t *testing.T) {
	o := Reverse(NilsLast(From(numCmp)))
	uis := []interface{}{5, 0, 3, nil, 0, 9}
	is := o.SortedCopy(uis)
	Equal(t, []interface{}{nil, 9, 5, 3, 0, 0}, is)
	Empty(t, o.SortedCopy([]interface{}{}))
}

func TestReverseIsOrdered(t *testing.T) {
	o := Reverse(Natural[int]())
	False(t, o.IsOrdered([]int{9, 0, 3, 5}))
	False(t, o.IsOrdered([]int{9, 3, 5, 0}))
	True(t, o.IsOrdered([]int{9, 5, 3, 0}))
	True(t, o.IsOrdered([]int{3, 3, 0, 0}))
	True(t, o.IsOrdered([]int{3, 0}))
	True(t, o.IsOrdered([]int{1}))
	True(t, o.IsOrdered([]int{}))
}

func TestReverseIsStrictlyOrdered(t *testing.T) {
	o := Reverse(Natural[int]())
	False(t, o.IsStrictlyOrdered([]int{9, 0, 3, 5}))
	False(t, o.IsStrictlyOrdered([]int{9, 3, 5, 0}))
	True(t, o.IsStrictlyOrdered([]int{9, 5, 3, 0}))
	False(t, o.IsStrictlyOrdered([]int{3, 3, 0, 0}))
	True(t, o.IsStrictlyOrdered([]int{3, 0}))
	True(t, o.IsStrictlyOrdered([]int{1}))
	True(t, o.IsStrictlyOrdered([]int{}))
}

func TestReverseLeastOf_empty_0(t *testing.T) {
	o := Reverse(Natural[float64]())
	Empty(t, o.LeastOf([]float64{}, 0))
}

func TestReverseLeastOf_empty_1(t *testing.T) {
	o := Reverse(Natural[float64]())
	Empty(t, o.LeastOf([]float64{}, 1))
}

func TestReverseLeastOf_simple_negativeOne(t *testing.T) {
	o := Reverse(Natural[float64]())
	PanicsWithError(t, "k cannot be negative but was: -1",
		func() { o.LeastOf([]float64{3, 4, 5, -1}, -1) })
}

func TestReverseLeastOf_singleton_0(t *testing.T) {
	o := Reverse(Natural[float64]())
	Empty(t, o.LeastOf([]float64{3}, 0))
}

func TestReverseLeastOf_simple_0(t *testing.T) {
	o := Reverse(Natural[float64]())
	Empty(t, o.LeastOf([]float64{3, 4, 5, -1}, 0))
}
func TestReverseLeastOf_simple_1(t *testing.T) {
	o := Reverse(Natural[float64]())
	Equal(t, []float64{5}, o.LeastOf([]float64{3, 4, 5, -1}, 1))
}

func TestReverseLeastOf_simple_nMinusOne_withNilElement(t *testing.T) {
	is := []interface{}{3, nil, 5, -1}
	res := Reverse(NilsLast(From(numCmp))).LeastOf(is, len(is)-1)
	Equal(t, []interface{}{nil, 5, 3}, res)
}

func TestReverseLeastOf_simple_nMinusOne(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Reverse(Natural[int]()).LeastOf(is, len(is)-1)
	Equal(t, []int{5, 4, 3}, res)
}

func TestReverseLeastOf_simple_n(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Reverse(Natural[int]()).LeastOf(is, len(is))
	Equal(t, []int{5, 4, 3, -1}, res)
}

func TestReverseLeastOf_simple_n_withNilElement(t *testing.T) {
	is := []interface{}{3, 4, 5, nil, -1}
	res := Reverse(NilsLast(From(numCmp))).LeastOf(is, len(is))
	Equal(t, []interface{}{nil, 5, 4, 3, -1}, res)
}

func TestReverseLeastOf_simple_nPlusOne(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Reverse(Natural[int]()).LeastOf(is, len(is)+1)
	Equal(t, []int{5, 4, 3, -1}, res)
}

func TestReverseLeastOf_ties(t *testing.T) {
	is := []int{3, math.MaxInt - 10, math.MaxInt - 10, -1}
	res := Reverse(Natural[int]()).LeastOf(is, len(is))
	Equal(t, []int{math.MaxInt - 10, math.MaxInt - 10, 3, -1}, res)
}

func TestReverseLeastOfLargeK(t *testing.T) {
	is := []int{4, 2, 3, 5, 1}
	Equal(t, []int{5, 4, 3, 2, 1}, Reverse(Natural[int]()).LeastOf(is, math.MaxInt))
}

func TestReverseGreatestOf_simple(t *testing.T) {
	/*
	 * If greatestOf() promised to be implemented as reverse().leastOf(), this
	 * test would be enough. It doesn't... but we'll cheat and act like it does
	 * anyway. There's a comment there to remind us to fix this if we change it.
	 */
	is := []int{3, 1, 3, 2, 4, 2, 4, 3}
	Equal(t, []int{1, 2, 2, 3}, Reverse(Natural[int]()).GreatestOf(is, 4))
}

func TestReverse_MinAndMax(t *testing.T) {
	is := []int{5, 3, 0, 9}
	o := Reverse(Natural[int]())

	m, ok := o.MaxElement(is...)
	Equal(t, 0, m)
	True(t, ok)

	m, ok = o.MinElement(is...)
	Equal(t, 9, m)
	True(t, ok)

	// when the values are the same, the first argument should be returned
	a, b := &wrapper{}, &wrapper{}
	ws := []interface{}{a, b, b}
	eqo := AllEqual[interface{}]()

	s, ok := eqo.MaxElement(ws...)
	Same(t, a, s)
	True(t, ok)

	s, ok = eqo.MinElement(ws...)
	Same(t, a, s)
	True(t, ok)
}

func TestReverse_MinAndMax2(t *testing.T) {
	o := Reverse(Natural[int]())

	Equal(t, -1, o.Max(-1, 0))
	Equal(t, -1, o.Max(0, -1))
	Equal(t, 0, o.Min(-1, 0))
	Equal(t, 0, o.Min(0, -1))

	// when the values are the same, the first argument should be returned
	a, b := &wrapper{}, &wrapper{}
	eqo := AllEqual[interface{}]()

	Same(t, a, eqo.Max(a, b))
	Same(t, a, eqo.Min(a, b))
}
