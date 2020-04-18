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

package url

import (
	"github.com/abc-inc/goava/escape/runes"
	"regexp"
	"strings"
)

// In some escapers spaces are escaped to '+'.
var plusSign = []rune{'+'}

// Percent escapers output upper case hex digits (URI escapers require this).
var upperHexDigits = []rune("0123456789ABCDEF")

// PercentEscaper escapes some set of characters using a UTF-8 based percent encoding scheme.
// The set of safe characters (those which remain unescaped) can be specified on construction.
//
// This type is primarily used for creating URI escapers in url.Escapers but can be used directly if required.
// While URI escapers impose specific semantics on which characters are considered 'safe', this type has a minimal set of restrictions.
//
// When escaping a string, the following rules apply:
//
// • All specified safe characters remain unchanged.
//
// • If plusForSpace was specified, the space character " " is converted into a plus sign "+".
//
// • All other characters are converted into one or more bytes using UTF-8 encoding and each byte is then represented
// by the 3-character string "%XX", where "XX" is the two-digit, uppercase, hexadecimal representation of the byte value.
//
// For performance reasons the only currently supported character encoding of this type is UTF-8.
//
// Note: This escaper produces uppercase hexadecimal sequences.
type PercentEscaper struct {
	*runes.Escaper

	// An array of flags where for any rune r, if safeOctets[r] is true then r should remain unmodified in the output.
	// If r >= len(safeOctets) then it should be escaped.
	safeOctets    []bool
	safeOctetsLen int32

	// If true we should convert space to the + character.
	plusForSpace bool
}

// NewPercentEscaper constructs an escaper with the specified safe characters and optional handling of the space character.
//
// Not that it is allowed, but not necessarily desirable to specify % as a safe character.
// This has the effect of creating an escaper which has no well defined inverse but it can be useful when escaping additional characters.
func NewPercentEscaper(safeChars string, plusForSpace bool) (e *PercentEscaper) {
	// Avoid any misunderstandings about the behavior of this escaper
	if ok, err := regexp.MatchString(".*[0-9A-Za-z].*", safeChars); ok || err != nil {
		panic("Alphanumeric characters are always 'safe' and should not be explicitly specified")
	}
	safeChars += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Avoid ambiguous parameters. Safe characters are never modified so if space is a safe character then setting plusForSpace is meaningless.
	if plusForSpace && strings.Contains(safeChars, " ") {
		panic("plusForSpace cannot be specified when space is a 'safe' character")
	}
	safeOctets := createSafeOctets(safeChars)

	e = &PercentEscaper{
		Escaper:       runes.NewEscaper(func(r rune) []rune { return e.escapeRune(r) }),
		plusForSpace:  plusForSpace,
		safeOctets:    safeOctets,
		safeOctetsLen: int32(len(safeOctets)),
	}
	return e
}

// Escape returns the escaped form of a given literal string.
//
// Overridden for performance. For unescaped strings this improved the performance of the URL escaper significantly.
func (e *PercentEscaper) Escape(str string) string {
	for i, r := range str {
		if r >= e.safeOctetsLen || !e.safeOctets[r] {
			return e.EscapeSlow(str, i)
		}
	}
	return str
}

// createSafeOctets creates a boolean slice with entries corresponding to the rune values specified in safeChars set to true.
// The array is as small as is required to hold the given character information.
func createSafeOctets(safeRunes string) []bool {
	maxRune := rune(-1)
	for _, c := range safeRunes {
		if c > maxRune {
			maxRune = c
		}
	}

	octets := make([]bool, maxRune+1)
	for _, c := range safeRunes {
		octets[c] = true
	}
	return octets
}

// escapeRune escapes the given Unicode code point in UTF-8.
func (e *PercentEscaper) escapeRune(r rune) []rune {
	// We should never get negative values here but if we do it will panic, so at least it will get spotted.
	if r < e.safeOctetsLen && e.safeOctets[r] {
		return nil
	} else if r == ' ' && e.plusForSpace {
		return plusSign
	} else if r <= 0x7F {
		// Single byte UTF-8 characters
		// Start with "%--" and fill in the blanks
		dest := make([]rune, 3)
		dest[0] = '%'
		dest[2] = upperHexDigits[r&0xF]
		dest[1] = upperHexDigits[r>>4]
		return dest
	} else if r <= 0x7ff {
		// Two byte UTF-8 characters [r >= 0x80 && r <= 0x7ff]
		// Start with "%--%--" and fill in the blanks
		dest := make([]rune, 6)
		dest[0] = '%'
		dest[3] = '%'
		dest[5] = upperHexDigits[r&0xF]
		r >>= 4
		dest[4] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[2] = upperHexDigits[r&0xF]
		r >>= 4
		dest[1] = upperHexDigits[0xC|r]
		return dest
	} else if r <= 0xffff {
		// Three byte UTF-8 characters [r >= 0x800 && r <= 0xffff]
		// Start with "%E-%--%--" and fill in the blanks
		dest := make([]rune, 9)
		dest[0] = '%'
		dest[1] = 'E'
		dest[3] = '%'
		dest[6] = '%'
		dest[8] = upperHexDigits[r&0xF]
		r >>= 4
		dest[7] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[5] = upperHexDigits[r&0xF]
		r >>= 4
		dest[4] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[2] = upperHexDigits[r]
		return dest
	} else if r <= 0x10ffff {
		dest := make([]rune, 12)
		// Four byte UTF-8 characters [r >= 0xffff && r <= 0x10ffff]
		// Start with "%F-%--%--%--" and fill in the blanks
		dest[0] = '%'
		dest[1] = 'F'
		dest[3] = '%'
		dest[6] = '%'
		dest[9] = '%'
		dest[11] = upperHexDigits[r&0xF]
		r >>= 4
		dest[10] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[8] = upperHexDigits[r&0xF]
		r >>= 4
		dest[7] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[5] = upperHexDigits[r&0xF]
		r >>= 4
		dest[4] = upperHexDigits[0x8|(r&0x3)]
		r >>= 2
		dest[2] = upperHexDigits[r&0x7]
		return dest
	} else {
		// If this ever happens it is due to bug in UnicodeEscaper, not bad input.
		panic("Invalid unicode character value " + string(r))
	}
}
