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

package runes_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/abc-inc/goava/escape/runes"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
)

var testString = fmt.Sprintf("%cabyz\u0080\u0100\u0800\u1000ABYZ%c%c0189%c",
	MinCodePoint, MaxBMPCodePoint, SmallestSurrogate, LargestSurrogate)

const special = '☃'

// simpleEscaper escapes everything except [a-zA-Z0-9].
var simpleEscaper = NewEscaper(func(r rune) []rune {
	if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') {
		return nil
	}
	return []rune("[" + strconv.Itoa(int(r)) + "]")
})

func TestEscapeEmpty(t *testing.T) {
	e := NewEscaper(func(r rune) []rune { return []rune{special} })
	// Escapers operate on runes: no runes, no escaping.
	Equal(t, "", e.Escape(""))
}

func TestEscapeFunc(t *testing.T) {
	e := NewEscaper(func(r rune) []rune {
		if r == '@' {
			return []rune{'?'}
		}
		return nil
	})
	// '[' and '@' lie either side of [A-Z].
	Equal(t, "[FOO?BAR]", e.Escape("[FOO@BAR]"))
}

func TestEscapeUnicode(t *testing.T) {
	e := NewEscaper(func(r rune) []rune {
		return []rune{special, special}
	})
	Equal(t, string(special)+string(special), e.Escape(string(special)))
}

func TestNopEscaper(t *testing.T) {
	e := NewEscaper(func(r rune) []rune { return nil })
	Equal(t, testString, e.Escape(testString))
}

func TestSimpleEscaper(t *testing.T) {
	expected := fmt.Sprintf("[0]abyz[128][256][2048][4096]ABYZ[65535][%v]0189[%v]", SmallestSurrogate, LargestSurrogate)
	Equal(t, expected, simpleEscaper.Escape(testString))
}

func TestGrowBuffer(t *testing.T) {
	// Need to grow past an initial 1024 byte buffer
	input := strings.Builder{}
	expected := strings.Builder{}
	for i := int32(256); i < 1024; i++ {
		input.WriteRune(i)
		expected.WriteString(fmt.Sprintf("[%d]", i))
	}
	Equal(t, expected.String(), simpleEscaper.Escape(input.String()))
}

func TestSurrogatePairs(t *testing.T) {
	// Build up a range of surrogate pair characters to test
	const min = MinSupplementaryCodePoint
	const max = MaxCodePoint
	const r = max - min
	const s1 = min + (1*r)/4
	const s2 = min + (2*r)/4
	const s3 = min + (3*r)/4
	test := string([]rune{'x', min, s1, s2, s3, max, 'x'})

	// Get the expected result string
	expected := fmt.Sprintf("x[%d][%d][%d][%d][%d]x", min, s1, s2, s3, max)
	Equal(t, expected, simpleEscaper.Escape(test))
}
