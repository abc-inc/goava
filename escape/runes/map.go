// Copyright 2020 The Goava authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runes

// EscaperMap is an implementation-specific parameter type suitable for initializing arrayBasedEscaper instances.
// This struct should be used when more than one escaper is created using the same replacement mapping to
// allow the underlying (implementation specific) data structures to be shared.
//
// The size of the data structure used by arrayBasedEscaper is proportional to the highest valued rune that has a replacement.
type EscaperMap struct {
	// The underlying replacement array we can share between multiple escaper instances.
	replArray [][]rune
}

// NewEscaperMap returns a new EscaperMap for creating arrayBasedEscaper instances.
func NewEscaperMap(replMap map[rune]string) *EscaperMap {
	return &EscaperMap{createReplArray(replMap)}
}

// GetReplacements returns the non-nil array of replacements for fast lookup.
func (em *EscaperMap) GetReplacements() [][]rune {
	return em.replArray
}

// createReplArray creates a replacement array from the given map.
// The returned array is a linear lookup table of replacement sequences indexed by the original rune value.
func createReplArray(m map[rune]string) [][]rune {
	if len(m) == 0 {
		return [][]rune{}
	}

	max := rune(0)
	for k := range m {
		if k > max {
			max = k
		}
	}

	replArray := make([][]rune, max+1)
	for k, v := range m {
		replArray[k] = []rune(v)
	}
	return replArray
}
