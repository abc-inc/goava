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

import "unicode"

// Builder creates and initializes simple, fast escapers.
//
// Typically an escaper needs to deal with the escaping of high valued bytes or code points.
// In these cases it is necessary to extend either bytes.arrayBasedEscaper or runes.arrayBasedEscaper to provide the desired behavior.
// However this builder is suitable for creating escapers that replace a relative small set of runes.
type Builder struct {
	replMap          map[rune]string
	safeMin, safeMax rune
	unsafeRepl       *string
}

// NewBuilder returns a builder for creating simple, fast escapers.
// A builder instance can be reused and each escaper that is created will be a snapshot of the current builder state.
// Builders are not thread safe.
//
// The initial state of the builder is such that:
//
// • There are no replacement mappings
//
// • safeMin == 0
//
// • safeMax == unicode.MaxRune
//
// • unsafeRepl == nil
//
// For performance reasons escapers created by this builder will not validate the well-formedness of their input.
func NewBuilder() *Builder {
	return &Builder{
		replMap: map[rune]string{},
		safeMin: 0,
		safeMax: unicode.MaxRune,
	}
}

// SetSafeRange sets the safe range of runes for the escaper.
// Runes in this range that have no explicit replacement are considered 'safe' and remain unescaped in the output.
// If safeMax < safeMin then the safe range is empty.
func (b *Builder) SetSafeRange(safeMin, safeMax rune) *Builder {
	b.safeMin = safeMin
	b.safeMax = safeMax
	return b
}

// SetUnsafeReplacement sets the replacement string for any runes outside the 'safe' range that have no explicit replacement.
// If unsafeRepl is nil then no replacement will occur, if it is "" then the unsafe runes are removed from the output.
func (b *Builder) SetUnsafeReplacement(unsafeReplacement string) *Builder {
	b.unsafeRepl = &unsafeReplacement
	return b
}

// AddEscape adds a replacement string for the given input rune.
// The specified rune will be replaced by the given string whenever it occurs in the input,
// irrespective of whether it lies inside or outside the 'safe' range.
func (b *Builder) AddEscape(r rune, replacement string) *Builder {
	// This can replace an existing rune (the builder is re-usable).
	b.replMap[r] = replacement
	return b
}

// AddEscapes adds multiple mappings at once for a particular index.
func (b *Builder) AddEscapes(rs []rune, r string) *Builder {
	for _, c := range rs {
		b.AddEscape(c, r)
	}
	return b
}

// Build returns a new escaper based on the current state of the builder.
func (b *Builder) Build() *arrayBasedEscaper {
	var replRunes []rune
	if b.unsafeRepl != nil {
		replRunes = []rune(*b.unsafeRepl)
	}
	return NewArrayBasedFromMap(b.replMap, b.safeMin, b.safeMax, func(r rune) []rune { return replRunes })
}
