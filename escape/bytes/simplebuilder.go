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

package bytes

import (
	"math"
)

// SimpleBuilder builds a "sparse" array of replacement mappings based on the indexes that were added to it.
// The array will be from 0 to the maximum index given.
// All non-set indexes will contain nil (so it's not really a sparse array, just a pseudo sparse array).
type SimpleBuilder struct {
	// replMap holds the replacement mappings.
	replMap map[byte]string
	// max is the highest index we've seen so far.
	max byte
}

// NewSimpleBuilder constructs a new sparse array builder.
func NewSimpleBuilder() *SimpleBuilder {
	return &SimpleBuilder{replMap: map[byte]string{}}
}

// AddEscape adds a new mapping from an index to a string to the escaping.
func (b *SimpleBuilder) AddEscape(c byte, r string) *SimpleBuilder {
	b.replMap[c] = r
	if c > b.max {
		b.max = c
	}
	return b
}

// AddEscapes adds multiple mappings at once for a particular index.
func (b *SimpleBuilder) AddEscapes(cs []byte, r string) *SimpleBuilder {
	for _, c := range cs {
		b.AddEscape(c, r)
	}
	return b
}

// ToArray converts this builder into a [][]byte where the maximum index is the value of the highest byte that has been seen.
// The array will be sparse in the sense that any unseen index will default to nil.
func (b *SimpleBuilder) ToArray() [][]byte {
	result := make([][]byte, b.max+1)
	for k, v := range b.replMap {
		result[k] = []byte(v)
	}
	return result
}

// ToEscaper converts this SimpleBuilder into an escaper, which is just a decorator around the underlying array of replacement [][]byte.
func (b *SimpleBuilder) ToEscaper() *byteArrayDecorator {
	return NewByteArrayEscaper(b.ToArray())
}

// byteArrayDecorator is a simple decorator that turns an array of replacement [][]bytes into an escaper, this results in a very fast escape method.
type byteArrayDecorator struct {
	*Escaper
	repl    [][]byte
	replLen byte
}

// NewByteArrayEscaper returns an escaper that escapes based on the underlying array.
func NewByteArrayEscaper(repl [][]byte) (e *byteArrayDecorator) {
	replLen := len(repl)
	if replLen > math.MaxUint8 {
		repl := repl[:math.MaxUint8]
		replLen = len(repl)
	}

	e = &byteArrayDecorator{
		Escaper: NewEscaper(func(c byte) []byte { return e.escapeByte(c) }),
		repl:    repl,
		replLen: uint8(replLen),
	}
	return e
}

func (e *byteArrayDecorator) Escape(str string) string {
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c < e.replLen && e.repl[c] != nil {
			return e.EscapeSlow(str, i)
		}
	}
	return str
}

func (e *byteArrayDecorator) escapeByte(c byte) []byte {
	if c < e.replLen {
		return e.repl[c]
	}
	return nil
}
