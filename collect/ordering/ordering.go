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

// Package ordering provides a comparison utility for sorting and accessing the least or greatest
// elements of a finite set of values of arbitrary types T.
// Custom ordering can be implemented for types that are not comparable out of the box.
package ordering

import "fmt"

const leftIsGreater = 1
const rightIsGreater = -1
const equal = 0

// Ordering is a comparator, with additional methods to support common operations.
//
// There are three types of methods present: methods for acquiring, chaining and using.
//
// Acquiring: The common ways to get an instance of Ordering are
//
// • implement a new type of Ordering instead of implementing Comparator directly
//
// • pass a pre-existing Comparator instance to From()
//
// • use the natural ordering, Natural()
//
// Chaining: then you can use the chaining methods to get an altered version of that Ordering
//
//  • Reverse()
//
//  • Compound()
//
//  • OnResultOf()
//
//  • NilsFirst() / NilsLast()
//
// Using: Finally, use the resulting Ordering anywhere a Comparator is required,
// or use any of its special operations, such as
//
//  • SortedCopy()
//
//  • IsOrdered() / IsStrictlyOrdered()
//
//  • Min() / Max()
type Ordering[T any] interface {
	// Compare compares its two arguments for order.
	// The same contract applies as for Comparator.
	Compare(a, b T) int

	// MinElement returns the least of the specified values according to this Ordering.
	// If there are multiple least values, the first of those is returned.
	MinElement(es ...T) (T, bool)

	// Min returns the lesser of the two values according to this Ordering.
	// If the values compare as 0, the first is returned.
	Min(a, b T) T

	// MaxElement returns the greatest of the specified values according to this Ordering.
	// If there are multiple greatest values, the first of those is returned.
	MaxElement(es ...T) (T, bool)

	// Max returns the greater of the two values according to this Ordering.
	// If the values compare as 0, the first is returned.
	Max(a, b T) T

	// LeastOf returns the k least elements of the given slice according to this Ordering,
	// in order from least to greatest.
	// If there are fewer than k elements present, all will be included.
	// The returned Ordering panics when it receives a negative k.
	//
	// The implementation does not necessarily use a stable sorting algorithm;
	// when multiple elements are equivalent, it is undefined which will come first.
	LeastOf(es []T, k int) []T

	// GreatestOf returns the k greatest elements of the given slice according to this Ordering,
	// in order from greatest to least.
	// If there are fewer than k elements present, all will be included.
	// The returned Ordering panics when it receives a negative k.
	//
	// The implementation does not necessarily use a stable sorting algorithm;
	// when multiple elements are equivalent, it is undefined which will come first.
	GreatestOf(es []T, k int) []T

	// SortedCopy returns a copy containing elements sorted by this Ordering.
	// The input is not modified.
	//
	// The sort performed is stable, meaning that such elements will appear in the returned slice in
	// the same order they appeared in es.
	SortedCopy(es []T) []T

	// IsOrdered returns true if each element in es after the first is greater than or equal to
	// the element that preceded it, according to this Ordering.
	// Note that this is always true when the slice has fewer than two elements.
	IsOrdered(es []T) bool

	// IsStrictlyOrdered returns true if each element in es after the first is strictly greater than
	// the element that preceded it, according to this Ordering.
	// Note that this is always true when the slice has fewer than two elements.
	IsStrictlyOrdered(es []T) bool

	fmt.Stringer
}
