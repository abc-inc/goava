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

// Package set provides a data structure that contains no duplicates and models the mathematical set abstraction.
package set

var present struct{}

// Set is a collection that contains no duplicate elements.
// More formally, sets contain no pair of elements e1 and e2 such that e1 == e2.
// As implied by its name, this type models the mathematical set abstraction.
//
// Note: Great care must be exercised if mutable objects are used as set elements.
// The behavior of a set is not specified if the value of an object is changed in a manner that affects equality
// comparisons while the object is an element in the set.
type Set struct {
	m map[interface{}]struct{}
}

// Size returns the number of elements in this set (its cardinality).
func (s Set) Size() int {
	return len(s.m)
}

// IsEmpty returns true if this set contains no elements.
func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

// Contains returns true if this set contains the specified element.
func (s Set) Contains(e interface{}) bool {
	v, exists := s.m[e]
	return exists && v == present
}

// ToArray returns a slice containing all the elements in this set.
//
// This method does not make any guarantees as to that order its elements are returned.
// The returned slice will be "safe" in that no references to it are maintained by this set.
// The caller is thus free to modify the returned slice.
func (s Set) ToArray() []interface{} {
	var es []interface{}
	for e := range s.m {
		es = append(es, e)
	}
	return es
}

// Add adds the specified element to this set if it is not already present (optional operation).
//
// More formally, adds the specified element e to this set if the set contains no element e2 such that e == e2.
// If this set already contains the element, the call leaves the set unchanged and returns false.
// This ensures that sets never contain duplicate elements.
//
// The stipulation above does not imply that sets must accept all elements;
// sets may refuse to add any particular element, including nil, and throw an error.
func (s Set) Add(e interface{}) bool {
	if contains := s.Contains(e); contains {
		return false
	}
	s.m[e] = present
	return true
}

// Remove removes the specified element from this set if it is present (optional operation).
//
// More formally, removes an element e2 such that e == e2, if this set contains such an element.
// Returns true if this set contained the element (or equivalently, if this set changed as a result of the call).
// (This set will not contain the element once the call returns.)
func (s Set) Remove(e interface{}) bool {
	if contains := s.Contains(e); !contains {
		return false
	}
	delete(s.m, e)
	return true
}

// ContainsAll returns true if this set contains all the elements of the other set.
//
// In other words, this method returns true if the other set is a subset of this set.
func (s Set) ContainsAll(other Set) bool {
	for e := range other.m {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// AddAll adds all the elements in the other set to this set if they're not already present (optional operation).
//
// The AddAll operation effectively modifies this set so that its value is the union of the two sets.
func (s Set) AddAll(other Set) bool {
	modified := false
	for e := range other.m {
		if s.Add(e) {
			modified = true
		}
	}
	return modified
}

// RetainAll retains only the elements in this set that are contained in the other set (optional operation).
//
// In other words, removes from this set all of its elements that are not contained in the other set.
// This operation effectively modifies this set so that its value is the intersection of the two sets.
func (s Set) RetainAll(other Set) bool {
	modified := false
	for e := range s.m {
		if !other.Contains(e) {
			s.Remove(e)
			modified = true
		}
	}
	return modified
}

// RemoveAll removes from this set all of its elements that are contained in the other set (optional operation).
//
// This operation effectively modifies this set so that its value is the asymmetric set difference of the two sets.
func (s Set) RemoveAll(other Set) bool {
	modified := false
	for e := range other.m {
		if s.Contains(e) {
			s.Remove(e)
			modified = true
		}
	}
	return modified
}

// Clear removes all the elements from this set (optional operation).
// The set will be empty after this call returns.
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}
