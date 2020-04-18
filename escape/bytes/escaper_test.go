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

package bytes_test

import (
	"fmt"
	. "github.com/abc-inc/goava/escape/bytes"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
	"testing"
)

var testString = fmt.Sprintf("%cabyz\u0080\u0100\u0800\u1000ABYZ", MinCodePoint)

const special = '~'

func TestEscapeEmpty(t *testing.T) {
	e := NewEscaper(func(r byte) []byte { return []byte{special} })
	// Escapers operate on runes: no runes, no escaping.
	Equal(t, "", e.Escape(""))
}

func TestEscapeFunc(t *testing.T) {
	e := NewEscaper(func(c byte) []byte {
		if c == '@' {
			return []byte{'?'}
		}
		return nil
	})
	// '[' and '@' lie either side of [A-Z].
	Equal(t, "[FOO?BAR]", e.Escape("[FOO@BAR]"))
}

func TestEscapeUnicode(t *testing.T) {
	e := NewEscaper(func(c byte) []byte {
		return []byte{special, special}
	})
	Equal(t, string(special)+string(special), e.Escape(string(special)))
}

func TestNopEscaper(t *testing.T) {
	e := NewEscaper(func(c byte) []byte { return nil })
	Equal(t, testString, e.Escape(testString))
}
