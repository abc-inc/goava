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

var wbLowerUnderscore = runematcher.Is('_')

// LowerUnderscore represents the C++ variable naming convention, e.g., "lower_underscore".
type LowerUnderscore struct{}

// wordBoundary matches the beginning of words.
func (c LowerUnderscore) wordBoundary() runematcher.Matcher {
	return wbLowerUnderscore
}

// wordSeparator is a (potentially empty) string between two words.
func (c LowerUnderscore) wordSeparator() string {
	return "_"
}

// normalizeWord formats a single word according to this case format.
func (c LowerUnderscore) normalizeWord(word string) string {
	return strings.ToLower(word)
}

// normalizeFirstWord formats a single word, which is the first of the string, according to this case format.
func (c LowerUnderscore) normalizeFirstWord(word string) string {
	return c.normalizeWord(word)
}

// To converts the specified string from this format to the specified format.
//
// A "best effort" approach is taken; if str does not conform to the assumed format, then the behavior is undefined
// but we make a reasonable effort at converting anyway.
func (c LowerUnderscore) To(tgtFmt CaseFormat, str string) string {
	if _, ok := tgtFmt.(LowerHyphen); ok {
		return strings.ReplaceAll(str, "_", "-")
	} else if _, ok := tgtFmt.(UpperUnderscore); ok {
		return strings.ToUpper(str)
	}
	return convert(c, tgtFmt, str)
}
