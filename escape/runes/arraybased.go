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

package runes

import "math"

// arrayBasedEscaper uses an array to quickly look up replacement rune for a given rune value.
// An additional safe range is provided that determines whether rune values without specific replacements are to be
// considered safe and left unescaped or should be escaped in a general way.
//
// A good example of usage of this type is for source code escaping where the replacement array contains information
// about special ASCII characters such as \t and \n while escapeUnsafe(rune) handles general escaping of the form \uxxxx.
//
// The size of the data structure used by arrayBasedEscaper is proportional to the highest valued rune that requires escaping.
// If you need to create multiple escaper instances that have the same replacement mapping consider using EscaperMap.
type arrayBasedEscaper struct {
	*Escaper

	// The replacement array (see EscaperMap).
	repl [][]rune
	// The number of elements in the replacement array.
	replLen int32
	// The first rune in the safe range.
	safeMin rune
	// The last rune in the safe range.
	safeMax rune

	// escapeUnsafe escapes a rune value that has no direct explicit value in the replacement array and lies outside the stated safe range.
	//
	// Note that arrays returned by this method must not be modified once they have been returned.
	// However it is acceptable to return the same array multiple times (even for different input runes).
	escapeUnsafe func(rune) []rune
}

// NewArrayBasedFromMap creates a new arrayBasedEscaper instance with the given replacement map and specified safe range.
// If safeMax < safeMin then no runes are considered safe.
//
// If a rune has no mapped replacement then it is checked against the safe range.
// If it lies outside that, then escapeUnsafe(rune) is called, otherwise no escaping is performed.
func NewArrayBasedFromMap(replMap map[rune]string, safeMin, safeMax rune, escapeUnsafe func(rune) []rune) *arrayBasedEscaper {
	return NewArrayBased(NewEscaperMap(replMap).GetReplacements(), safeMin, safeMax, escapeUnsafe)
}

// NewArrayBased creates a new arrayBasedEscaper instance with the given replacement map and specified safe range.
// If safeMax < safeMin then no runes are considered safe.
// This initializer is useful when explicit instances of EscaperMap are used to allow the sharing of large replacement mappings.
//
// If a rune has no mapped replacement then it is checked against the safe range.
// If it lies outside that, then escapeUnsafe(rune) is called, otherwise no escaping is performed.
func NewArrayBased(replArray [][]rune, safeMin, safeMax rune, escapeUnsafe func(rune) []rune) (e *arrayBasedEscaper) {
	if safeMax < safeMin {
		// If the safe range is empty, set the range limits to opposite extremes
		// to ensure the first test of either value will (almost certainly) fail.
		safeMax = 0
		safeMin = math.MaxUint8
	}

	e = &arrayBasedEscaper{
		Escaper:      NewEscaper(func(r rune) []rune { return e.escapeRune(r) }),
		repl:         replArray,
		replLen:      int32(len(replArray)),
		safeMin:      safeMin,
		safeMax:      safeMax,
		escapeUnsafe: escapeUnsafe,
	}
	return e
}

// Escape returns the escaped form of a given literal string.
func (e *arrayBasedEscaper) Escape(str string) string {
	for i, r := range str {
		if (r < e.replLen && e.repl[r] != nil) || r > e.safeMax || r < e.safeMin {
			return e.EscapeSlow(str, i)
		}
	}
	return str
}

// escapeRune escapes a single rune using the replacement array and safe range values.
// If the given rune does not have an explicit replacement and lies outside the safe range then escapeUnsafe(rune) is called.
func (e *arrayBasedEscaper) escapeRune(r rune) []rune {
	if r < e.replLen {
		if runes := e.repl[r]; runes != nil {
			return runes
		}
	}
	if r >= e.safeMin && r <= e.safeMax {
		return nil
	}
	return e.escapeUnsafe(r)
}
