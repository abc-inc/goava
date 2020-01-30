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

//go:generate go run gen.go none

// noneMatcher matches no characters.
type noneMatcher struct {
}

// Matches determines a true or false value for the given character.
func (m noneMatcher) Matches(r rune) bool {
	return false
}

// IndexInRunes returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m noneMatcher) IndexInRunes(runes []rune, start int) int {
	return -1
}

// IndexIn returns the index of the first matching character in a character sequence,
// starting from a given position, or -1 if no character matches after that position.
func (m noneMatcher) IndexIn(str string, start int) int {
	return -1
}

// LastIndexIn returns the index of the last matching character in a character sequence,
// or -1 if no matching character is present.
func (m noneMatcher) LastIndexIn(str string) int {
	return -1
}

// MatchesAllOf returns true if a character sequence contains only matching characters.
func (m noneMatcher) MatchesAllOf(str string) bool {
	return len(str) == 0
}

// MatchesNoneOf returns true if a character sequence contains no matching characters.
//
// Equivalent to !MatchesAnyOf(sequence).
func (m noneMatcher) MatchesNoneOf(str string) bool {
	return true
}

// RemoveFrom returns a string containing all non-matching characters of a character sequence, in order.
func (m noneMatcher) RemoveFrom(str string) string {
	return str
}

// ReplaceFromRune returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement character.
func (m noneMatcher) ReplaceFromRune(str string, replacement rune) string {
	return str
}

// ReplaceFrom returns a string copy of the input character sequence, with each matching character
// replaced by a given replacement sequence.
func (m noneMatcher) ReplaceFrom(str, replacement string) string {
	return str
}

// CollapseFrom returns a string copy of the input character sequence, with each group of consecutive matching
// characters replaced by a single replacement character.
func (m noneMatcher) CollapseFrom(str string, replacement rune) string {
	return str
}

// TrimFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning and from the end of the string.
func (m noneMatcher) TrimFrom(str string) string {
	return str
}

// CountIn returns the number of matching characters found in a character sequence.
func (m noneMatcher) CountIn(str string) int {
	return 0
}

// And returns a matcher that matches any character matched by both this matcher and other.
func (m noneMatcher) And(other Matcher) Matcher {
	return m
}

// Or returns a matcher that matches any character matched by either this matcher or other.
func (m noneMatcher) Or(other Matcher) Matcher {
	return other
}

// Negate returns a matcher that matches any character not matched by this matcher.
func (m noneMatcher) Negate() Matcher {
	return Any()
}

// String returns a string representation of this Matcher.
func (m noneMatcher) String() string {
	return "Matcher.none()"
}

// Code generated. DO NOT EDIT.

// MatchesAnyOf returns true if a character sequence contains at least one matching character.
//
// Equivalent to !MatchesNoneOf(sequence)
func (m noneMatcher) MatchesAnyOf(str string) bool {
	return matchesAnyOf(m, str)
}

// RetainFrom returns a string containing all matching characters of a character sequence, in order.
func (m noneMatcher) RetainFrom(str string) string {
	return retainFrom(m, str)
}

// TrimLeadingFrom returns a substring of the input character sequence that omits all matching characters
// from the beginning of the string.
func (m noneMatcher) TrimLeadingFrom(str string) string {
	return trimLeadingFrom(m, str)
}

// TrimTrailingFrom returns a substring of the input character sequence that omits all matching characters
// from the end of the string.
func (m noneMatcher) TrimTrailingFrom(str string) string {
	return trimTrailingFrom(m, str)
}

// TrimAndCollapseFrom collapses groups of matching characters exactly as CollapseFrom(str, replacement) does,
// except that groups of matching characters at the start or end of the sequence are removed without replacement.
func (m noneMatcher) TrimAndCollapseFrom(str string, replacement rune) string {
	return trimAndCollapseFrom(m, str, replacement)
}
