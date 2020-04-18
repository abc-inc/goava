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

package html_test

import (
	. "github.com/abc-inc/goava/html"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestEscaper(t *testing.T) {
	Equal(t, "xxx", Escaper.Escape("xxx"))
	Equal(t, "&quot;test&quot;", Escaper.Escape("\"test\""))
	Equal(t, "&#39;test&#39;", Escaper.Escape("'test'"))
	Equal(t, "test &amp; test &amp; test", Escaper.Escape("test & test & test"))
	Equal(t, "test &lt;&lt; 1", Escaper.Escape("test << 1"))
	Equal(t, "test &gt;&gt; 1", Escaper.Escape("test >> 1"))
	Equal(t, "&lt;tab&gt;", Escaper.Escape("<tab>"))

	// Test simple escape of '&'.
	Equal(t, "foo&amp;bar", Escaper.Escape("foo&bar"))

	// If the string contains no escapes, it should return the arg.
	s := "blah blah farhvergnugen"
	Equal(t, s, Escaper.Escape(s))

	// Tests escapes at begin and end of string.
	Equal(t, "&lt;p&gt;", Escaper.Escape("<p>"))

	// Test all escapes.
	Equal(t, "a&quot;b&lt;c&gt;d&amp;", Escaper.Escape("a\"b<c>d&"))

	// Test two escapes in a row.
	Equal(t, "foo&amp;&amp;bar", Escaper.Escape("foo&&bar"))

	// Test many non-escaped characters.
	s = "!@#$%^*()_+=-/?\\|]}[{,.;:" +
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"1234567890"
	Equal(t, s, Escaper.Escape(s))
}
