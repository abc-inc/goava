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

package runematcher

//go:generate go run gen.go any

import "strings"

// anyMatcher matches any character.
type anyMatcher struct {
}

// Matches determines a true or false value for the given character.
func (m anyMatcher) Matches(r rune) bool {
	return true
}

// IndexIn returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m anyMatcher) IndexIn(str string, start int) int {
	return m.IndexInRunes([]rune(str), start)
}

// IndexInRunes returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m anyMatcher) IndexInRunes(runes []rune, start int) int {
	if start < 0 || start >= len(runes) {
		return -1
	}
	return start
}

// LastIndexIn returns the index of the last matching character in a character sequence,
// or -1 if no matching character is present.
func (m anyMatcher) LastIndexIn(str string) int {
	return len(str) - 1
}

// MatchesAllOf returns true if a character sequence contains only matching characters.
func (m anyMatcher) MatchesAllOf(str string) bool {
	return true
}

// MatchesNoneOf returns true if a character sequence contains no matching characters.
//
// Equivalent to !MatchesAnyOf(sequence).
func (m anyMatcher) MatchesNoneOf(str string) bool {
	return len(str) == 0
}

// RemoveFrom returns a string containing all non-matching characters of a character sequence, in order.
func (m anyMatcher) RemoveFrom(str string) string {
	return ""
}

// ReplaceFromRune returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement character.
func (m anyMatcher) ReplaceFromRune(str string, replacement rune) string {
	return strings.Repeat(string(replacement), len(str))
}

// ReplaceFrom returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement sequence.
func (m anyMatcher) ReplaceFrom(str, replacement string) string {
	result := strings.Builder{}
	result.Grow(len(str) * len(replacement))
	for i := 0; i < len(str); i++ {
		result.WriteString(replacement)
	}
	return result.String()
}

// CollapseFrom returns a string copy of the input character sequence, with each group of consecutive matching
// characters replaced by a single replacement character.
func (m anyMatcher) CollapseFrom(str string, replacement rune) string {
	if len(str) == 0 {
		return ""
	}
	return string(replacement)
}

// TrimFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning and from the end of the string.
func (m anyMatcher) TrimFrom(str string) string {
	return ""
}

// CountIn returns the number of matching characters found in a character sequence.
func (m anyMatcher) CountIn(str string) int {
	return len(str)
}

// And returns a matcher that matches any character matched by both this matcher and other.
func (m anyMatcher) And(other Matcher) Matcher {
	return other
}

// Or returns a matcher that matches any character matched by either this matcher or other.
func (m anyMatcher) Or(other Matcher) Matcher {
	return m
}

// Negate returns a matcher that matches any character not matched by this matcher.
func (m anyMatcher) Negate() Matcher {
	return None()
}

// String returns a string representation of this Matcher.
func (m anyMatcher) String() string {
	return "Matcher.any()"
}

// Code generated. DO NOT EDIT.

// MatchesAnyOf returns true if a character sequence contains at least one matching character.
//
// Equivalent to !MatchesNoneOf(sequence)
func (m anyMatcher) MatchesAnyOf(str string) bool {
	return matchesAnyOf(m, str)
}

// RetainFrom returns a string containing all matching characters of a character sequence, in order.
func (m anyMatcher) RetainFrom(str string) string {
	return retainFrom(m, str)
}

// TrimLeadingFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning of the string.
func (m anyMatcher) TrimLeadingFrom(str string) string {
	return trimLeadingFrom(m, str)
}

// TrimTrailingFrom returns a substring of the input character sequence that omits all matching characters
// from the end of the string.
func (m anyMatcher) TrimTrailingFrom(str string) string {
	return trimTrailingFrom(m, str)
}

// TrimAndCollapseFrom collapses groups of matching characters exactly as CollapseFrom(str, replacement) does,
// except that groups of matching characters at the start or end of the sequence are removed without replacement.
func (m anyMatcher) TrimAndCollapseFrom(str string, replacement rune) string {
	return trimAndCollapseFrom(m, str, replacement)
}
