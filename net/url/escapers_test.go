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
	"testing"

	"github.com/abc-inc/goava/escape"
	. "github.com/abc-inc/goava/net/url"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
)

func TestFormParameterEscaper(t *testing.T) {
	e := FormParameterEscaper
	// Verify that these are the same escaper (as documented)
	Same(t, e, FormParameterEscaper)
	AssertBasicURLEscaper(t, e)

	// Specified as safe by RFC 2396
	AssertUnescapedByte(t, e, '!')
	AssertUnescapedByte(t, e, '(')
	AssertUnescapedByte(t, e, ')')
	AssertUnescapedByte(t, e, '~')
	AssertUnescapedByte(t, e, '\'')

	// Plus for spaces
	AssertEscapingByte(t, e, "+", ' ')
	AssertEscapingByte(t, e, "%2B", '+')

	Equal(t, "safe+with+spaces", e.Escape("safe with spaces"))
	Equal(t, "foo%40bar.com", e.Escape("foo@bar.com"))
}

func TestPathSegmentEscaper(t *testing.T) {
	e := PathSegmentEscaper
	assertPathEscaper(t, e)
	AssertUnescapedByte(t, e, '+')
}

func TestFragmentEscaper(t *testing.T) {
	e := FragmentEscaper
	AssertUnescapedByte(t, e, '+')
	AssertUnescapedByte(t, e, '/')
	AssertUnescapedByte(t, e, '?')

	assertPathEscaper(t, e)
}

func assertPathEscaper(t *testing.T, e escape.Escaper) {
	AssertBasicURLEscaper(t, e)

	AssertUnescapedByte(t, e, '!')
	AssertUnescapedByte(t, e, '\'')
	AssertUnescapedByte(t, e, '(')
	AssertUnescapedByte(t, e, ')')
	AssertUnescapedByte(t, e, '~')
	AssertUnescapedByte(t, e, ':')
	AssertUnescapedByte(t, e, '@')

	// Don't use plus for spaces
	AssertEscapingByte(t, e, "%20", ' ')

	Equal(t, "safe%20with%20spaces", e.Escape("safe with spaces"))
	Equal(t, "foo@bar.com", e.Escape("foo@bar.com"))
}
