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
	. "github.com/abc-inc/goava/escape/bytes"
	. "github.com/abc-inc/goava/testing/escapetest"
	. "github.com/stretchr/testify/require"
	"testing"
	"unicode"
)

func TestNewEscaperMapNil(t *testing.T) {
	Empty(t, NewEscaperMap(nil).GetReplacements())
}

func TestNewEscaperMapEmpty(t *testing.T) {
	Empty(t, NewEscaperMap(map[byte]string{}).GetReplacements())
}

func TestEscaperMapGetReplacements(t *testing.T) {
	m := map[byte]string{'a': "first", 'z': "last"}
	em := NewEscaperMap(m)
	// Array length is highest byte value + 1
	Equal(t, int('z'+1), len(em.GetReplacements()))
}

func TestEscaperMapMapping(t *testing.T) {
	m := map[byte]string{byte(MinCodePoint): "zero", 'a': "first", 'b': "second", 'z': "last", unicode.MaxLatin1: "biggest"}
	em := NewEscaperMap(m)
	repl := em.GetReplacements()
	// Array length is highest byte value + 1
	Equal(t, 256, len(repl))
	// The final element should always be non null.
	NotNil(t, repl[len(repl)-1])
	// Exhaustively check all mappings (an int index avoids wrapping).
	replLen := byte(len(repl))
	for n := byte(0); n < replLen; n++ {
		if repl[n] != nil {
			Equal(t, m[n], repl[n])
		} else {
			_, ok := m[n]
			False(t, ok)
		}
	}
}
