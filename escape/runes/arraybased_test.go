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

package runes_test

import (
	"fmt"
	. "github.com/abc-inc/goava/escape/runes"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
	"testing"
	"unicode"
)

var noReplacements = map[rune]string{}

var simpleReplacements = map[rune]string{
	'\n': "<newline>",
	'\t': "<tab>",
	'&':  "<and>",
}

func TestReplacements(t *testing.T) {
	// In reality this is not a very sensible escaper to have
	// (if you are only escaping elements from a map you would use a arrayBasedEscaper).
	escapeUnsafe := func(b rune) []rune { return []rune{} }
	e := NewArrayBasedFromMap(simpleReplacements, 0, unicode.MaxRune, escapeUnsafe)

	Equal(t, "<tab>Fish <and> Chips<newline>", e.Escape("\tFish & Chips\n"))

	// Verify that everything else is left unescaped.
	safeRunes := fmt.Sprintf("%c%c%c%c%c", MinCodePoint, '\u0100', MinHighSurrogate, MinLowSurrogate, MaxBMPCodePoint)
	Equal(t, safeRunes, e.Escape(safeRunes))
}

func TestSafeRange(t *testing.T) {
	escapeUnsafe := func(r rune) []rune { return append([]rune{}, '{', r, '}') }
	e := NewArrayBasedFromMap(noReplacements, 'A', 'Z', escapeUnsafe)

	AssertBasic(t, e)
	// '[' and '@' lie either side of [A-Z].
	Equal(t, "{[}FOO{@}BAR{]}", e.Escape("[FOO@BAR]"))
}

func TestSafeRangeMaxLessThanMin(t *testing.T) {
	// Basic escaping of unsafe runes (wrap them in {,}'s)
	escapeUnsafe := func(r rune) []rune { return append([]rune{}, '{', r, '}') }
	e := NewArrayBasedFromMap(noReplacements, 'Z', 'A', escapeUnsafe)

	AssertBasic(t, e)
	// escape everything.
	Equal(t, "{[}{F}{O}{O}{]}", e.Escape("[FOO]"))
}

func TestDeleteUnsafeRunes(t *testing.T) {
	escapeUnsafe := func(r rune) []rune { return []rune{} }
	e := NewArrayBasedFromMap(noReplacements, ' ', '~', escapeUnsafe)

	AssertBasic(t, e)
	Equal(t, "Everything outside the printable ASCII range is deleted.",
		e.Escape(fmt.Sprintf("\tEverything\u0000 outside the%c%c printable ASCII \uFFFFrange is \u007Fdeleted.\n", MinHighSurrogate, MinLowSurrogate)))
}

func TestReplacementPriority(t *testing.T) {
	escapeUnsafe := func(r rune) []rune { return []rune{'?'} }
	e := NewArrayBasedFromMap(simpleReplacements, ' ', '~', escapeUnsafe)

	AssertBasic(t, e)
	// Replacements are applied first regardless of whether the rune is in the safe range or not
	// ('&' is a safe rune while '\t' and '\n' are not).
	Equal(t, "<tab>Fish <and>? Chips?<newline>", e.Escape("\tFish &\u0000 Chips\r\n"))
}

func TestArrayBasedEscaperSurrogatePairs(t *testing.T) {
	escapeUnsafe := func(r rune) []rune { return []rune{'X'} }
	e := NewArrayBasedFromMap(noReplacements, 0, 0x20000, escapeUnsafe)

	// A surrogate pair defining a code point within the safe range.
	safeInput := fmt.Sprintf("%c%c", MinHighSurrogate, MinLowSurrogate) // 0x10000
	Equal(t, safeInput, e.Escape(safeInput))
}
