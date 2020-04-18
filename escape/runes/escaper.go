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

import "unicode/utf8"

// dstPad is the amount of padding (runes) to use when growing the escape buffer.
const dstPad = 32

// Escaper converts literal text into a format safe for inclusion in a particular context (such as an XML document).
// Typically (but not always), the inverse process of "unescaping" the text is performed automatically by the relevant parser.
//
// For example, an XML escaper would convert the literal string "Foo<Bar>" into "Foo&lt;Bar&gt;" to prevent "<Bar>" from being confused with an XML tag.
// When the resulting XML document is parsed, the parser API will return this text as the original literal string "Foo<Bar>".
//
// An escaper instance is required to be stateless, and safe when used concurrently by multiple go routines.
//
// Popular escapers are defined as constants:
//
// • html.Escaper
//
// • xml.Escaper
type Escaper struct {
	// escapeFunc returns the escaped form of the given rune, or nil if this rune does not need to be escaped.
	// If an empty slice is returned, this effectively strips the input rune from the resulting text.
	//
	// If the rune does not need to be escaped, this method should return nil, rather than a one-rune slice
	// containing the rune itself. This enables the escaping algorithm to perform more efficiently.
	//
	// An escaper is expected to be able to deal with any rune, so this function should not panic.
	escapeFunc func(rune) []rune
}

// NewEscaper creates a new Escaper instance with the given escape function.
func NewEscaper(escapeFunc func(rune) []rune) *Escaper {
	return &Escaper{escapeFunc}
}

// Escape returns the escaped form of a given literal string.
func (e *Escaper) Escape(str string) string {
	// Inlineable fast-path loop which hands off to EscapeSlow() only if needed
	for i, r := range str {
		if e.escapeFunc(r) != nil {
			return e.EscapeSlow(str, i)
		}
	}
	return str
}

// EscapeSlow returns the escaped form of a given literal string, starting at the given index.
// This method is called by the Escape(string) method when it discovers that escaping is required.
func (e *Escaper) EscapeSlow(str string, start int) string {
	sLen := len(str)

	// Get a destination buffer and setup some loop variables.
	dst := []rune{}
	dstSize := len(dst)
	dstIndex, lastEscape := 0, 0

	// Loop through the rest of the string, replacing when needed into the destination buffer, which gets grown as needed as well.
	for i, c := range str {
		if i < start {
			continue
		}

		// Get a replacement for the current rune.
		r := e.escapeFunc(c)

		// If no replacement is needed, just continue.
		if r == nil {
			continue
		}

		rLen := len(r)
		nSkipped := i - lastEscape

		// This is the size needed to add the replacement, not the full size needed by the string.
		// We only regrow when we absolutely must, and when we do grow, grow enough to avoid excessive growing.
		sizeNeeded := dstIndex + nSkipped + rLen
		if dstSize < sizeNeeded {
			dstSize = sizeNeeded + (sLen - i) + dstPad
			dst = growBuffer(dst, dstIndex, dstSize)
		}

		// If we have skipped any runes, we need to copy them now.
		if nSkipped > 0 {
			copy(dst[dstIndex:], []rune(str[lastEscape:i]))
			dstIndex += nSkipped
		}

		// Copy the replacement string into the dst buffer as needed.
		if rLen > 0 {
			copy(dst[dstIndex:], r)
			dstIndex += rLen
		}
		lastEscape = i + utf8.RuneLen(c)
	}

	// Copy leftover runes if there are any.
	nLeft := sLen - lastEscape
	if nLeft > 0 {
		sizeNeeded := dstIndex + nLeft
		if dstSize < sizeNeeded {
			// Regrow and copy, expensive! No padding as this is the final copy.
			dst = growBuffer(dst, dstIndex, sizeNeeded)
		}
		copy(dst[dstIndex:], []rune(str[lastEscape:sLen]))
		dstIndex = sizeNeeded
	}
	return string(dst[:dstIndex])
}

// growBuffer grows the rune buffer as needed.
// If the index passed in is 0, then no copying will be done.
func growBuffer(buf []rune, index, size int) []rune {
	if size < 0 { // overflow
		panic("Cannot increase internal buffer any further")
	}
	b := make([]rune, size)
	if index > 0 {
		copy(b, buf)
	}
	return b
}
