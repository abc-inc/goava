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

package set

// Empty returns a set containing zero elements.
func Empty() Set {
	return Set{make(map[interface{}]struct{})}
}

// Singleton returns a set containing one element.
func Singleton(e interface{}) Set {
	set := Empty()
	set.Add(e)
	return set
}

// Of returns a set containing an arbitrary number of elements.
func Of(es ...interface{}) Set {
	set := Empty()
	for _, e := range es {
		set.Add(e)
	}
	return set
}

// CopyOf returns a set containing the elements of the given set.
func CopyOf(s Set) Set {
	var newMap = make(map[interface{}]struct{})
	for k, v := range s.m {
		newMap[k] = v
	}
	return Set{newMap}
}
