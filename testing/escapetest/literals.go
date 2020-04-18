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

package escapetest

// Common unicode literals.
const (
	MinHighSurrogate          = rune(0xd800)
	MaxHighSurrogate          = rune(0xdbff)
	MinLowSurrogate           = rune(0xdc00)
	MaxLowSurrogate           = rune(0xdfff)
	SmallestSurrogate         = rune(0xd800 + 0xdc00)
	LargestSurrogate          = rune(0xdbff + 0xdfff)
	MaxBMPCodePoint           = rune(65535)
	MinSupplementaryCodePoint = rune(65536)
	MinCodePoint              = rune(0)
	MaxCodePoint              = rune(1114111)
)
