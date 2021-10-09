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

package ordering_test

import (
	"fmt"
	. "github.com/abc-inc/goava/collect/ordering"
	"github.com/abc-inc/goava/primitives/ints"
	. "github.com/stretchr/testify/require"
	"math"
	"math/rand"
	"testing"
)

type wrapper struct {
	v string
}

func numCmp(a, b interface{}) int {
	return ints.Compare(a.(int), b.(int))
}

func TestAllEqual(t *testing.T) {
	var o Ordering[interface{}] = AllEqual[interface{}]()
	Same(t, o, Reverse(Reverse(o)))

	Equal(t, 0, o.Compare(nil, nil))
	Equal(t, 0, o.Compare(wrapper{}, wrapper{}))
	Equal(t, 0, o.Compare("apples", "oranges"))
	Equal(t, "ordering.AllEqual()", o.String())

	ss := []interface{}{"b", "a", "d", "c"}
	Equal(t, ss, o.SortedCopy(ss))
}

// From https://github.com/google/guava/issues/1342
func TestComplicatedOrderingExample(t *testing.T) {
	var nilInt *int
	o := NilsLast(Reverse(Lexicographical(
		Reverse(NilsFirst(From(numCmp))).Compare,
	)))

	l1 := []interface{}{}
	l2 := []interface{}{1}
	l3 := []interface{}{1, 1}
	l4 := []interface{}{1, 2}
	l5 := []interface{}{1, nil, 2}
	l6 := []interface{}{2}
	l7 := []interface{}{nilInt}
	l8 := []interface{}{nilInt, nilInt}
	ls := [][]interface{}{l1, l2, l3, l4, l5, l6, l7, l8, nil}

	sorted := o.SortedCopy(ls)

	// [[nil, nil], [nil], [1, nil, 2], [1, 1], [1, 2], [1], [2], [], nil]
	Equal(t, [][]interface{}{
		[]interface{}{nilInt, nilInt},
		[]interface{}{nilInt},
		[]interface{}{1, nil, 2},
		[]interface{}{1, 1},
		[]interface{}{1, 2},
		[]interface{}{1},
		[]interface{}{2},
		[]interface{}{},
		nil},
		sorted)
}

func TestNatural(t *testing.T) {
	o := From[interface{}](numCmp)
	testComparator(t, o.Compare, math.MinInt, -1, 0, 1, math.MaxInt)
	Panics(t, func() { o.Compare(1, nil) })
	Panics(t, func() { o.Compare(nil, 2) })
	Panics(t, func() { o.Compare(nil, nil) })
	// Equal(t, "ordering.From(ordering.numCmp)", o.String())
}

func TestFrom(t *testing.T) {
	o := From(CaseInsensitive())
	Equal(t, 0, o.Compare("A", "a"))
	Less(t, o.Compare("a", "B"), 0)
	Greater(t, o.Compare("B", "a"), 0)
	// anonymous functions are called funcN
	Equal(t, "ordering.From(ordering.CaseInsensitive.func1)", o.String())
}

func TestExplicit_none(t *testing.T) {
	o := Explicit[interface{}](nil)
	PanicsWithValue(t, "cannot compare value: 0", func() { o.Compare(0, 0) })
	Equal(t, "ordering.Explicit([<nil>])", o.String())
}

func TestExplicit_one(t *testing.T) {
	o := Explicit(0)
	Equal(t, 0, o.Compare(0, 0))
	PanicsWithValue(t, "cannot compare value: 1", func() { o.Compare(0, 1) })
	Equal(t, "ordering.Explicit([0])", o.String())
}

func TestExplicit_two(t *testing.T) {
	o := Explicit(42, 5)
	Equal(t, 0, o.Compare(5, 5))
	True(t, o.Compare(5, 42) > 0)
	True(t, o.Compare(42, 5) < 0)
	PanicsWithValue(t, "cannot compare value: 666", func() { o.Compare(5, 666) })
}

func TestExplicit_sortingExample(t *testing.T) {
	o := Explicit(2, 8, 6, 1, 7, 5, 3, 4, 0, 9)
	is := []int{0, 3, 5, 6, 7, 8, 9}
	cs := o.SortedCopy(is)
	Equal(t, []int{8, 6, 7, 5, 3, 0, 9}, cs)
}

func TestExplicit_withDuplicates(t *testing.T) {
	PanicsWithValue(t, "multiple entries with same key: 2=1 and 2=4",
		func() { Explicit(1, 2, 3, 4, 2) })
}

func TestUsingString(t *testing.T) {
	o := UsingString[int]()
	testComparator[int](t, o.Compare, 1, 12, 124, 2)
	Equal(t, "ordering.UsingString()", o.String())
}

func byCharAt(index int) Ordering[string] {
	return OnResultOf[string, byte](
		func(v string) byte { return v[index] },
		Natural[byte]())
}

func TestCompound_static(t *testing.T) {
	o := Compound(
		byCharAt(0).Compare,
		byCharAt(1).Compare,
		byCharAt(2).Compare,
		byCharAt(3).Compare,
		byCharAt(4).Compare,
		byCharAt(5).Compare,
	)

	testComparator(t,
		o.Compare,
		[]string{
			"applesauce",
			"apricot",
			"artichoke",
			"banality",
			"banana",
			"banquet",
			"tangelo",
			"tangerine"}...)
}

func TestCompound_instance(t *testing.T) {
	o := Compound(byCharAt(1).Compare, byCharAt(0).Compare)
	testComparator(t,
		o.Compare,
		[]string{
			"red",
			"yellow",
			"violet",
			"blue",
			"indigo",
			"green",
			"orange",
		}...)
}

func TestReverse(t *testing.T) {
	r := Reverse(Natural[int]())
	testComparator(t, r.Compare, math.MaxInt, 1, 0, -1, math.MinInt)
}

func TestReverseOfReverseSameAsForward(t *testing.T) {
	// Not guaranteed by spec, but it works, and saves us from testing exhaustively
	o := Natural[int]()
	Same(t, o, Reverse(Reverse(o)))
}

func TestOnResultOf_natural(t *testing.T) {
	o := OnResultOf[string, int](
		func(v string) int { return len(v) },
		Natural[int](),
	)
	Equal(t, 0, o.Compare("to", "be"))
	Equal(t, -1, o.Compare("or", "not"))
	Equal(t, 1, o.Compare("that", "to"))
	Equal(t, "ordering.OnResultOf(ordering.Natural())", o.String())
}

func TestOnResultOf_chained(t *testing.T) {
	o := OnResultOf[string, int](
		func(v string) int { return len(v) },
		Reverse(Natural[int]()),
	)
	Equal(t, 0, o.Compare("to", "be"))
	Equal(t, 1, o.Compare("or", "not"))
	Equal(t, -1, o.Compare("that", "to"))
	Equal(t, "ordering.OnResultOf(ordering.Reverse(ordering.Natural()))", o.String())
}

func TestLexicographical(t *testing.T) {
	o := Lexicographical(Natural[string]().Compare)

	empty := []string{}
	a := []string{"a"}
	aa := []string{"a", "a"}
	ab := []string{"a", "b"}
	b := []string{"b"}

	testComparator(t, o.Compare, empty, a, aa, ab, b)
}

func TestNilsFirst(t *testing.T) {
	o := NilsFirst(From(numCmp))
	testComparator(t, o.Compare, nil, math.MinInt, 0, 1)
}

func TestNilsLast(t *testing.T) {
	o := NilsLast(From(numCmp))
	testComparator(t, o.Compare, 0, 1, math.MaxInt, nil)
}

func TestSortedCopy(t *testing.T) {
	o := NilsLast(From(numCmp))
	uis := []interface{}{5, 0, 3, nil, 0, 9}
	is := o.SortedCopy(uis)
	Equal(t, []interface{}{0, 0, 3, 5, 9, nil}, is)
	Empty(t, o.SortedCopy([]interface{}{}))
}

func TestIsOrdered(t *testing.T) {
	o := Natural[int]()
	False(t, o.IsOrdered([]int{5, 3, 0, 9}))
	False(t, o.IsOrdered([]int{0, 5, 3, 9}))
	True(t, o.IsOrdered([]int{0, 3, 5, 9}))
	True(t, o.IsOrdered([]int{0, 0, 3, 3}))
	True(t, o.IsOrdered([]int{0, 3}))
	True(t, o.IsOrdered([]int{1}))
	True(t, o.IsOrdered([]int{}))
}

func TestIsStrictlyOrdered(t *testing.T) {
	o := Natural[int]()
	False(t, o.IsStrictlyOrdered([]int{5, 3, 0, 9}))
	False(t, o.IsStrictlyOrdered([]int{0, 5, 3, 9}))
	True(t, o.IsStrictlyOrdered([]int{0, 3, 5, 9}))
	False(t, o.IsStrictlyOrdered([]int{0, 0, 3, 3}))
	True(t, o.IsStrictlyOrdered([]int{0, 3}))
	True(t, o.IsStrictlyOrdered([]int{1}))
	True(t, o.IsStrictlyOrdered([]int{}))
}

func TestLeastOf_empty_0(t *testing.T) {
	o := Natural[float64]()
	Empty(t, o.LeastOf([]float64{}, 0))
}

func TestLeastOf_empty_1(t *testing.T) {
	o := Natural[float64]()
	Empty(t, o.LeastOf([]float64{}, 1))
}

func TestLeastOf_simple_negativeOne(t *testing.T) {
	o := Natural[float64]()
	PanicsWithError(t, "k cannot be negative but was: -1",
		func() { o.LeastOf([]float64{3, 4, 5, -1}, -1) })
}

func TestLeastOf_singleton_0(t *testing.T) {
	o := Natural[float64]()
	Empty(t, o.LeastOf([]float64{3}, 0))
}

func TestLeastOf_simple_0(t *testing.T) {
	o := Natural[float64]()
	Empty(t, o.LeastOf([]float64{3, 4, 5, -1}, 0))
}
func TestLeastOf_simple_1(t *testing.T) {
	o := Natural[float64]()
	Equal(t, []float64{-1}, o.LeastOf([]float64{3, 4, 5, -1}, 1))
}

func TestLeastOf_simple_nMinusOne_withNilElement(t *testing.T) {
	is := []interface{}{3, nil, 5, -1}
	res := NilsLast(From(numCmp)).LeastOf(is, len(is)-1)
	Equal(t, []interface{}{-1, 3, 5}, res)
}

func TestLeastOf_simple_nMinusOne(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Natural[int]().LeastOf(is, len(is)-1)
	Equal(t, []int{-1, 3, 4}, res)
}

func TestLeastOf_simple_n(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Natural[int]().LeastOf(is, len(is))
	Equal(t, []int{-1, 3, 4, 5}, res)
}

func TestLeastOf_simple_n_withNilElement(t *testing.T) {
	is := []interface{}{3, 4, 5, nil, -1}
	res := NilsLast(From(numCmp)).LeastOf(is, len(is))
	Equal(t, []interface{}{-1, 3, 4, 5, nil}, res)
}

func TestLeastOf_simple_nPlusOne(t *testing.T) {
	is := []int{3, 4, 5, -1}
	res := Natural[int]().LeastOf(is, len(is)+1)
	Equal(t, []int{-1, 3, 4, 5}, res)
}

func TestLeastOf_ties(t *testing.T) {
	is := []int{3, math.MaxInt - 10, math.MaxInt - 10, -1}
	res := Natural[int]().LeastOf(is, len(is))
	Equal(t, []int{-1, 3, math.MaxInt - 10, math.MaxInt - 10}, res)
}

func TestLeastOf_reconcileAgainstSortAndSublist(t *testing.T) {
	runLeastOfComparison(t, 1000, 300, 20)
}

func TestLeastOf_reconcileAgainstSortAndSublistSmall(t *testing.T) {
	runLeastOfComparison(t, 10, 30, 2)
}

func runLeastOfComparison(t *testing.T, iter, elements, seeds int) {
	o := Natural[int]()

	for i := 0; i < iter; i++ {
		is := []int{}
		for j := 0; j < elements; j++ {
			is = append(is, rand.Intn(10*i+j+1))
		}

		for seed := 1; seed < seeds; seed++ {
			k := rand.Intn(10 * seed)
			Equal(t, o.SortedCopy(is)[0:k], o.LeastOf(is, k))
		}
	}
}

func TestLeastOfLargeK(t *testing.T) {
	is := []int{4, 2, 3, 5, 1}
	Equal(t, []int{1, 2, 3, 4, 5}, Natural[int]().LeastOf(is, math.MaxInt))
}

func TestGreatestOf_simple(t *testing.T) {
	/*
	 * If greatestOf() promised to be implemented as reverse().leastOf(), this
	 * test would be enough. It doesn't... but we'll cheat and act like it does
	 * anyway. There's a comment there to remind us to fix this if we change it.
	 */
	is := []int{3, 1, 3, 2, 4, 2, 4, 3}
	Equal(t, []int{4, 4, 3, 3}, Natural[int]().GreatestOf(is, 4))
}

func TestMinAndMax(t *testing.T) {
	is := []int{5, 3, 0, 9}
	o := Natural[int]()

	m, ok := o.MaxElement(is...)
	Equal(t, 9, m)
	True(t, ok)

	m, ok = o.MinElement(is...)
	Equal(t, 0, m)
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

func TestMinAndMax_empty(t *testing.T) {
	o := Natural[int]()

	m, ok := o.MaxElement()
	Equal(t, 0, m)
	False(t, ok)

	m, ok = o.MinElement()
	Equal(t, 0, m)
	False(t, ok)
}

func TestMinAndMax_nil(t *testing.T) {
	o := AllEqual[interface{}]()

	m, ok := o.MaxElement(nil)
	Nil(t, m)
	True(t, ok)

	m, ok = o.MinElement(nil)
	Nil(t, m)
	True(t, ok)
}

func testComparator[T any](t *testing.T, cmp Comparator[T], es ...T) {
	// This does an O(n^2) test of all pairs of values in both orders
	for i := 0; i < len(es); i++ {
		e := es[i]

		for j := 0; j < i; j++ {
			lesser := es[j]
			Less(t, cmp(lesser, e), 0,
				fmt.Sprintf("%v.compare(%v, %v)", cmp, lesser, e))
		}

		Equal(t, 0, cmp(e, e),
			fmt.Sprintf("%v.compare(%v, %v)", cmp, e, e))

		for j := i + 1; j < len(es); j++ {
			greater := es[j]
			Greater(t, cmp(greater, e), 0,
				fmt.Sprintf("%v.compare(%v, %v) %d %d", cmp, greater, e, i, j))
		}
	}
}
