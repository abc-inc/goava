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

package casefmt

import (
	"github.com/abc-inc/goava/base/runematcher"
	"strings"
)

// CaseFormat converts strings between various ASCII case formats.
//
// Behavior is undefined for non-ASCII input.
type CaseFormat interface {
	// wordBoundary matches the beginning of words.
	wordBoundary() runematcher.Matcher

	// wordSeparator is a (potentially empty) string between two words.
	wordSeparator() string

	// normalizeWord formats a single word according to this case format.
	normalizeWord(word string) string

	// normalizeFirstWord formats a single word, which is the first of the string, according to this case format.
	normalizeFirstWord(word string) string

	// To converts the specified string from this format to the specified format.
	//
	// A "best effort" approach is taken; if str does not conform to the assumed format, then the behavior is undefined
	// but we make a reasonable effort at converting anyway.
	To(tgtFmt CaseFormat, str string) string
}

// convert converts the specified string from this format to the specified format.
//
// A "best effort" approach is taken; if s does not conform to the assumed format,
// then the behavior of this method is undefined but we make a reasonable effort at converting anyway.
func convert(this CaseFormat, format CaseFormat, s string) string {
	if this == format || len(s) == 0 {
		return s
	}

	runes := []rune(s)
	var out strings.Builder
	i := 0
	for j := this.wordBoundary().IndexInRunes(runes, 0); j != -1; j = this.wordBoundary().IndexInRunes(runes, j+1) {
		if i == 0 {
			out = strings.Builder{}
			out.Grow(len(s) + 4*len(format.wordSeparator()))
			out.WriteString(format.normalizeFirstWord(string(runes[i:j])))
		} else {
			out.WriteString(format.normalizeWord(s[i:j]))
		}
		out.WriteString(format.wordSeparator())
		i = j + len(this.wordSeparator())
	}

	if i == 0 {
		return format.normalizeFirstWord(s)
	}

	out.WriteString(format.normalizeWord(s[i:]))
	return out.String()
}

func firstCharOnlyToUpper(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
}
