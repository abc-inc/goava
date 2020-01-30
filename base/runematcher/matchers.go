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

import "sort"

// Any matches any character.
func Any() Matcher {
	return anyMatcher{}
}

// None matches no characters.
func None() Matcher {
	return noneMatcher{}
}

// Whitespace determines whether a character is whitespace according to the latest Unicode standard.
func Whitespace() Matcher {
	return whitespaceMatcher{}
}

// BreakingWhitespace determines whether a character is a breaking whitespace (that is, a whitespace which can be
// interpreted as a break between words for formatting purposes).
func BreakingWhitespace() Matcher {
	return breakingWhitespaceMatcher{}
}

// ASCII determines whether a character is ASCII, meaning that its code point is less than 128.
func ASCII() Matcher {
	return asciiMatcher{}
}

// Digit determines whether a character is a digit according to Unicode.
// If you only care to match ASCII digits, you can use InRange('0', '9').
func Digit() Matcher {
	return digitMatcher{}
}

// Invisible determines whether a character is invisible; that is, if its Unicode category is any of
// SPACE_SEPARATOR, LINE_SEPARATOR, PARAGRAPH_SEPARATOR, CONTROL, FORMAT, SURROGATE, and PRIVATE_USE.
func Invisible() Matcher {
	return invisibleMatcher{}
}

// SingleWidth determines whether a character is single-width (not double-width).
//
// When in doubt, this matcher errs on the side of returning false (that is, it tends to assume a character is
// double-width).
func SingleWidth() Matcher {
	return singleWidthMatcher{}
}

// Is returns a char matcher that matches only one specified character.
func Is(r rune) Matcher {
	return isMatcher{r}
}

// IsNot returns a char matcher that matches any character except the character specified.
//
// To negate another Matcher, use Negate().
func IsNot(r rune) Matcher {
	return isNotMatcher{r}
}

// AnyOf returns a char matcher that matches any character present in the given character sequence.
func AnyOf(match string) Matcher {
	runes := []rune(match)

	switch len(runes) {
	case 0:
		return None()
	case 1:
		return Is(runes[0])
	case 2:
		return isEitherMatcher{runes[0], runes[1]}
	default:
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		return anyOfMatcher{string(runes)}
	}
}

// NoneOf returns a char matcher that matches any character not present in the given character sequence.
func NoneOf(str string) Matcher {
	return AnyOf(str).Negate()
}

// InRange returns a char matcher that matches any character in a given range (both endpoints are inclusive).
//
// For example, to match any lowercase letter of the English alphabet, use InRange('a', 'z').
func InRange(startIncl, endIncl rune) Matcher {
	return inRangeMatcher{startIncl, endIncl}
}

// ForPredicate returns a matcher with identical behavior to the given character-based predicate.
func ForPredicate(p func(rune) bool) Matcher {
	return forPredicateMatcher{p}
}
