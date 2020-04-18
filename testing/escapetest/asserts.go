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

package escapetest

import (
	"fmt"
	"github.com/abc-inc/goava/escape"
	"github.com/stretchr/testify/require"
	"testing"
)

// AssertBasic asserts common expected behaviour of escapers.
func AssertBasic(t *testing.T, escaper escape.Escaper) {
	// Escapers operate on bytes: no bytes, no escaping.
	require.Equal(t, "", escaper.Escape(""))
}

// AssertBasicURLEscaperExceptPercent asserts common expected behaviour of URI escapers.
// You should call assertBasicURLEscaper() unless the escaper explicitly does not escape '%'.
func AssertBasicURLEscaperExceptPercent(t *testing.T, e escape.Escaper) {
	// All URL escapers should leave 0-9, A-Z, a-z unescaped
	AssertUnescapedByte(t, e, 'a')
	AssertUnescapedByte(t, e, 'z')
	AssertUnescapedByte(t, e, 'A')
	AssertUnescapedByte(t, e, 'Z')
	AssertUnescapedByte(t, e, '0')
	AssertUnescapedByte(t, e, '9')

	// Unreserved characters
	AssertUnescapedByte(t, e, '-')
	AssertUnescapedByte(t, e, '_')
	AssertUnescapedByte(t, e, '.')
	AssertUnescapedByte(t, e, '*')

	AssertEscapingByte(t, e, "%00", '\u0000')       // nul
	AssertEscapingByte(t, e, "%7F", '\u007f')       // del
	AssertEscapingByte(t, e, "%C2%80", '\u0080')    // xx-00010,x-000000
	AssertEscapingRune(t, e, "%DF%BF", '\u07ff')    // xx-11111,x-111111
	AssertEscapingRune(t, e, "%E0%A0%80", '\u0800') // xxx-0000,x-100000,x-00,0000
	AssertEscapingRune(t, e, "%EF%BF%BF", '\uffff') // xxx-1111,x-111111,x-11,1111
	AssertEscapingUnicode(t, e, "%F0%90%80%80", MinHighSurrogate, MinLowSurrogate)
	AssertEscapingUnicode(t, e, "%F4%8F%BF%BF", MaxHighSurrogate, MaxLowSurrogate)

	require.Equal(t, "", e.Escape(""))
	require.Equal(t, "safestring", e.Escape("safestring"))
	require.Equal(t, "embedded%00null", e.Escape("embedded\u0000null"))
	require.Equal(t, "max%EF%BF%BFchar", e.Escape("max\uffffchar"))
}

// AssertBasicURLEscaper asserts common expected behaviour of URI escapers.
func AssertBasicURLEscaper(t *testing.T, e escape.Escaper) {
	AssertBasicURLEscaperExceptPercent(t, e)
	// The escape character must always be escaped
	AssertEscapingByte(t, e, "%25", '%')
}

// AssertEscapingByte asserts that an escaper escapes the given character into the expected string.
func AssertEscapingByte(t *testing.T, e escape.Escaper, want string, c byte) {
	require.Equal(t, want, computeReplacement(e, c))
}

// AssertEscapingRune asserts that an runes.Escaper escapes the given rune into the expected string.
func AssertEscapingRune(t *testing.T, e escape.Escaper, want string, r rune) {
	require.Equal(t, want, computeReplacement(e, r))
}

// AssertUnescapedByte asserts that an escaper does not escape the given character.
func AssertUnescapedByte(t *testing.T, e escape.Escaper, c byte) {
	require.Equal(t, string(c), computeReplacement(e, c))
}

// AssertUnescapedRune asserts that an escaper does not escape the given character.
func AssertUnescapedRune(t *testing.T, e escape.Escaper, r rune) {
	require.Equal(t, string(r), computeReplacement(e, r))
}

// AssertEscapingUnicode asserts that a Unicode escaper escapes the given hi/lo surrogate pair into the expected string.
func AssertEscapingUnicode(t *testing.T, e escape.Escaper, want string, hi, lo rune) {
	cp := toCodePoint(hi, lo)
	require.Equal(t, want, computeReplacement(e, cp))
}

func toCodePoint(high, low rune) rune {
	return (high << 10) + low + -56613888
}

func computeReplacement(e escape.Escaper, c interface{}) string {
	return e.Escape(fmt.Sprintf("%c", c))
}
