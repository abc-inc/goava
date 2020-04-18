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

package url_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/abc-inc/goava/base/precond"
	. "github.com/abc-inc/goava/net/url"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
)

// TestPercentEscaperSimple tests that the simple escaper treats 0-9, a-z and A-Z as safe.
func TestPercentEscaperSimple(t *testing.T) {
	e := NewPercentEscaper("", false)
	for c := byte(0); c < 128; c++ {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			AssertUnescapedByte(t, e, c)
		} else {
			AssertEscapingByte(t, e, escapeASCII(c), c)
		}
	}

	// testing multibyte escape sequences
	AssertEscapingRune(t, e, "%00", '\u0000')       // nul
	AssertEscapingRune(t, e, "%7F", '\u007f')       // del
	AssertEscapingRune(t, e, "%C2%80", '\u0080')    // xx-00010,x-000000
	AssertEscapingRune(t, e, "%DF%BF", '\u07ff')    // xx-11111,x-111111
	AssertEscapingRune(t, e, "%E0%A0%80", '\u0800') // xxx-0000,x-100000,x-00,0000
	AssertEscapingRune(t, e, "%EF%BF%BF", '\uffff') // xxx-1111,x-111111,x-11,1111
	AssertEscapingUnicode(t, e, "%F0%90%80%80", MinHighSurrogate, MinLowSurrogate)
	AssertEscapingUnicode(t, e, "%F4%8F%BF%BF", MaxHighSurrogate, MaxLowSurrogate)

	// simple string tests
	Equal(t, "", e.Escape(""))
	Equal(t, "safestring", e.Escape("safestring"))
	Equal(t, "embedded%00null", e.Escape("embedded\u0000null"))
	Equal(t, "max%EF%BF%BFchar", e.Escape("max\uffffchar"))
}

// TestPlusForSpace tests the various ways that the space character can be handled.
func TestPlusForSpace(t *testing.T) {
	basicEscaper := NewPercentEscaper("", false)
	plusForSpaceEscaper := NewPercentEscaper("", true)
	spaceEscaper := NewPercentEscaper(" ", false)

	Equal(t, "string%20with%20spaces", basicEscaper.Escape("string with spaces"))
	Equal(t, "string+with+spaces", plusForSpaceEscaper.Escape("string with spaces"))
	Equal(t, "string with spaces", spaceEscaper.Escape("string with spaces"))
}

// TestCustomEscaper tests that if we add extra 'safe' characters they remain unescaped.
func TestCustomEscaper(t *testing.T) {
	e := NewPercentEscaper("+*/-", false)
	for c := byte(0); c < 128; c++ {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || strings.IndexByte("+*/-", c) >= 0 {
			AssertUnescapedByte(t, e, c)
		} else {
			AssertEscapingByte(t, e, escapeASCII(c), c)
		}
	}
}

// TestCustomEscaperPercent tests that if specify '%' as safe the result is an idempotent escaper.
func TestCustomEscaperPercent(t *testing.T) {
	e := NewPercentEscaper("%", false)
	Equal(t, "foo%7Cbar", e.Escape("foo|bar"))
	Equal(t, "foo%7Cbar", e.Escape("foo%7Cbar")) // idempotent
}

// TestBadArgumentsBadSafe tests that specifying any alphanumeric characters as 'safe' causes a panic.
func TestBadArgumentsBadSafe(t *testing.T) {
	msg := "Alphanumeric characters are always 'safe' and should not be explicitly specified"
	Panics(t, func() { NewPercentEscaper("-+#abc.!", false) }, msg)
}

// TestBadArgumentsPlusForSpace tests that if space is a safe character you cannot also specify 'plusForSpace' (panics).
func TestBadArgumentsPlusForSpace(t *testing.T) {
	NewPercentEscaper(" ", false)

	msg := "plusForSpace cannot be specified when space is a 'safe' character"
	Panics(t, func() { NewPercentEscaper(" ", true) }, msg)
}

// escapeASCII manually escapes a 7-bit ASCII character
func escapeASCII(c byte) string {
	if err := precond.CheckArgument(c < 128); err != nil {
		panic(err)
	}
	hex := "0123456789ABCDEF"
	return fmt.Sprintf("%%%c%c", hex[((c>>4)&0xf)], hex[c&0xf])
}
