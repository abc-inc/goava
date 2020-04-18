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

package bytes_test

import (
	"fmt"
	"testing"

	. "github.com/abc-inc/goava/escape/bytes"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
)

var noReplacements = map[byte]string{}

var simpleReplacements = map[byte]string{
	'\n': "<newline>",
	'\t': "<tab>",
	'&':  "<and>",
}

func TestSafeRange(t *testing.T) {
	escapeUnsafe := func(c byte) []byte { return append([]byte{}, '{', c, '}') }
	e := NewArrayBasedFromMap(noReplacements, 'A', 'Z', escapeUnsafe)

	AssertBasic(t, e)
	// '[' and '@' lie either side of [A-Z].
	Equal(t, "{[}FOO{@}BAR{]}", e.Escape("[FOO@BAR]"))
}

func TestSafeRangeMaxLessThanMin(t *testing.T) {
	// Basic escaping of unsafe bytes (wrap them in {,}'s)
	escapeUnsafe := func(c byte) []byte { return append([]byte{}, '{', c, '}') }
	e := NewArrayBasedFromMap(noReplacements, 'Z', 'A', escapeUnsafe)

	AssertBasic(t, e)
	// escape everything.
	Equal(t, "{[}{F}{O}{O}{]}", e.Escape("[FOO]"))
}

func TestDeleteUnsafeBytes(t *testing.T) {
	escapeUnsafe := func(c byte) []byte { return []byte{} }
	e := NewArrayBasedFromMap(noReplacements, ' ', '~', escapeUnsafe)

	AssertBasic(t, e)
	Equal(t, "Everything outside the printable ASCII range is deleted.",
		e.Escape(fmt.Sprintf("\tEverything\u0000 outside the%c%c printable ASCII \uFFFFrange is \u007Fdeleted.\n", MinHighSurrogate, MinLowSurrogate)))
}

func TestReplacementPriority(t *testing.T) {
	escapeUnsafe := func(c byte) []byte { return []byte{'?'} }
	e := NewArrayBasedFromMap(simpleReplacements, ' ', '~', escapeUnsafe)

	AssertBasic(t, e)
	// Replacements are applied first regardless of whether the byte is in the safe range or not
	// ('&' is a safe byte while '\t' and '\n' are not).
	Equal(t, "<tab>Fish <and>? Chips?<newline>", e.Escape("\tFish &\u0000 Chips\r\n"))
}
