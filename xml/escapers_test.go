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

package xml_test

import (
	"fmt"
	"github.com/abc-inc/goava/escape"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/abc-inc/goava/xml"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestContentEscaper(t *testing.T) {
	assertBasicXMLEscaper(t, ContentEscaper, false, false)
	// Test quotes are not escaped.
	Equal(t, "\"test\"", ContentEscaper.Escape("\"test\""))
	Equal(t, "'test'", ContentEscaper.Escape("'test'"))
}

func TestAttributeEscaper(t *testing.T) {
	assertBasicXMLEscaper(t, AttributeEscaper, true, true)
	// Test quotes are escaped.
	Equal(t, "&quot;test&quot;", AttributeEscaper.Escape("\"test\""))
	Equal(t, "&apos;test&apos;", AttributeEscaper.Escape("'test'"))
	// Test all escapes.
	Equal(t, "a&quot;b&lt;c&gt;d&amp;e&quot;f&apos;", AttributeEscaper.Escape("a\"b<c>d&e\"f'"))
	// Test '\t', '\n' and '\r' are escaped.
	Equal(t, "a&#x9;b&#xA;c&#xD;d", AttributeEscaper.Escape("a\tb\nc\rd"))
}

// Helper to assert common properties of xml escapers.
func assertBasicXMLEscaper(t *testing.T, e escape.Escaper, shouldEscapeQuotes, shouldEscapeWhitespaces bool) {
	// Simple examples (smoke tests)
	Equal(t, "xxx", e.Escape("xxx"))
	Equal(t, "test &amp; test &amp; test", e.Escape("test & test & test"))
	Equal(t, "test &lt;&lt; 1", e.Escape("test << 1"))
	Equal(t, "test &gt;&gt; 1", e.Escape("test >> 1"))
	Equal(t, "&lt;tab&gt;", e.Escape("<tab>"))

	// Test all non-escaped ASCII characters.
	s := "!@#$%^*()_+=-/?\\|]}[{,.;:" +
		"abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"1234567890"
	Equal(t, s, e.Escape(s))

	// Test ASCII control characters.
	for c := byte(0); c < 0x20; c++ {
		if c == '\t' || c == '\n' || c == '\r' {
			// Only these whitespace chars are permitted in XML,
			if shouldEscapeWhitespaces {
				AssertEscapingByte(t, e, fmt.Sprintf("&#x%X;", c), c)
			} else {
				AssertUnescapedByte(t, e, c)
			}
		} else {
			// and everything else is replaced with FFFD.
			AssertEscapingByte(t, e, "\uFFFD", c)
		}
	}

	// Test _all_ allowed characters (including surrogate values).
	for c := rune(0x20); c <= 0xFFFD; c++ {
		// There are a small number of cases to consider, so just do it manually.
		if c == '&' {
			AssertEscapingRune(t, e, "&amp;", c)
		} else if c == '<' {
			AssertEscapingRune(t, e, "&lt;", c)
		} else if c == '>' {
			AssertEscapingRune(t, e, "&gt;", c)
		} else if shouldEscapeQuotes && c == '\'' {
			AssertEscapingRune(t, e, "&apos;", c)
		} else if shouldEscapeQuotes && c == '"' {
			AssertEscapingRune(t, e, "&quot;", c)
		} else {
			input := string(c)
			escaped := e.Escape(input)
			Equal(t, input, escaped, "char 0x%X should not be escaped", c)
		}
	}

	// Test that 0xFFFE and 0xFFFF are replaced with 0xFFFD.
	AssertEscapingRune(t, e, "\uFFFD", '\uFFFE')
	AssertEscapingRune(t, e, "\uFFFD", '\uFFFF')

	Equal(t, "[\uFFFD]", e.Escape("[\ufffe]"), "0xFFFE is forbidden and should be replaced during escaping")
	Equal(t, "[\uFFFD]", e.Escape("[\uffff]"), "0xFFFF is forbidden and should be replaced during escaping")
}
