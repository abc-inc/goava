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

package runematcher

//go:generate go run gen.go invisible

import (
	"unicode"
)

// invisibleMatcher determines whether a character is invisible; that is, if its Unicode category is any of
// SPACE_SEPARATOR, LINE_SEPARATOR, PARAGRAPH_SEPARATOR, CONTROL, FORMAT, SURROGATE, and PRIVATE_USE.
type invisibleMatcher struct {
}

// Matches determines a true or false value for the given character.
func (m invisibleMatcher) Matches(r rune) bool {
	return unicode.IsPrint(r)
}

// IndexInRunes returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m invisibleMatcher) IndexInRunes(runes []rune, start int) int {
	return indexInRunes(m, runes, start)
}

// IndexIn returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m invisibleMatcher) IndexIn(str string, start int) int {
	return indexIn(m, str, start)
}

// String returns a string representation of this Matcher.
func (m invisibleMatcher) String() string {
	return "Matcher.invisible()"
}

// Code generated. DO NOT EDIT.

// Negate returns a matcher that matches any character not matched by this matcher.
func (m invisibleMatcher) Negate() Matcher {
	return negate(m)
}

// And returns a matcher that matches any character matched by both this matcher and other.
func (m invisibleMatcher) And(other Matcher) Matcher {
	return and(m, other)
}

// Or returns a matcher that matches any character matched by either this matcher or other.
func (m invisibleMatcher) Or(other Matcher) Matcher {
	return or(m, other)
}

// MatchesAnyOf returns true if a character sequence contains at least one matching character.
//
// Equivalent to !MatchesNoneOf(sequence)
func (m invisibleMatcher) MatchesAnyOf(str string) bool {
	return matchesAnyOf(m, str)
}

// MatchesAllOf returns true if a character sequence contains only matching characters.
func (m invisibleMatcher) MatchesAllOf(str string) bool {
	return matchesAllOf(m, str)
}

// MatchesNoneOf returns true if a character sequence contains no matching characters.
//
// Equivalent to !MatchesAnyOf(sequence).
func (m invisibleMatcher) MatchesNoneOf(str string) bool {
	return matchesNoneOf(m, str)
}

// LastIndexIn returns the index of the last matching character in a character sequence,
// or -1 if no matching character is present.
func (m invisibleMatcher) LastIndexIn(str string) int {
	return lastIndexIn(m, str)
}

// CountIn returns the number of matching characters found in a character sequence.
func (m invisibleMatcher) CountIn(str string) int {
	return countIn(m, str)
}

// RemoveFrom returns a string containing all non-matching characters of a character sequence, in order.
func (m invisibleMatcher) RemoveFrom(str string) string {
	return removeFrom(m, str)
}

// RetainFrom returns a string containing all matching characters of a character sequence, in order.
func (m invisibleMatcher) RetainFrom(str string) string {
	return retainFrom(m, str)
}

// ReplaceFromRune returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement character.
func (m invisibleMatcher) ReplaceFromRune(str string, replacement rune) string {
	return replaceFromRune(m, str, replacement)
}

// ReplaceFrom returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement sequence.
func (m invisibleMatcher) ReplaceFrom(str string, replacement string) string {
	return replaceFrom(m, str, replacement)
}

// TrimFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning and from the end of the string.
func (m invisibleMatcher) TrimFrom(str string) string {
	return trimFrom(m, str)
}

// TrimLeadingFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning of the string.
func (m invisibleMatcher) TrimLeadingFrom(str string) string {
	return trimLeadingFrom(m, str)
}

// TrimTrailingFrom returns a substring of the input character sequence that omits all matching characters
// from the end of the string.
func (m invisibleMatcher) TrimTrailingFrom(str string) string {
	return trimTrailingFrom(m, str)
}

// CollapseFrom returns a string copy of the input character sequence, with each group of consecutive matching
// characters replaced by a single replacement character.
func (m invisibleMatcher) CollapseFrom(str string, replacement rune) string {
	return collapseFrom(m, str, replacement)
}

// TrimAndCollapseFrom collapses groups of matching characters exactly as CollapseFrom(str, replacement) does,
// except that groups of matching characters at the start or end of the sequence are removed without replacement.
func (m invisibleMatcher) TrimAndCollapseFrom(str string, replacement rune) string {
	return trimAndCollapseFrom(m, str, replacement)
}
