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

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func negate(m Matcher) Matcher {
	return negatedMatcher{m}
}

func and(first, second Matcher) Matcher {
	return andMatcher{first, second}
}

func or(first, second Matcher) Matcher {
	return orMatcher{first, second}
}

func matchesAnyOf(m Matcher, str string) bool {
	return !m.MatchesNoneOf(str)
}

func matchesAllOf(m Matcher, str string) bool {
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		if !m.Matches(runes[i]) {
			return false
		}
	}
	return true
}

func matchesNoneOf(m Matcher, str string) bool {
	return m.IndexIn(str, 0) == -1
}

func indexInRunes(m Matcher, runes []rune, start int) int {
	l := len(runes)
	if start < 0 || start >= l {
		return -1
	}

	for i := start; i < l; i++ {
		if m.Matches(runes[i]) {
			return i
		}
	}
	return -1
}

func indexIn(m Matcher, str string, start int) int {
	return m.IndexInRunes([]rune(str), start)
}

func lastIndexIn(m Matcher, str string) int {
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		if m.Matches(runes[i]) {
			return i
		}
	}
	return -1
}

func countIn(m Matcher, str string) int {
	count := 0
	for _, r := range str {
		if m.Matches(r) {
			count++
		}
	}
	return count
}

func removeFrom(m Matcher, str string) string {
	runes := []rune(str)
	pos := m.IndexInRunes(runes, 0)
	if pos == -1 {
		return str
	}

	spread := 1
	for {
		pos++
		for {
			if pos == len(runes) {
				return string(runes[0 : pos-spread])
			}
			if m.Matches(runes[pos]) {
				break
			}
			runes[pos-spread] = runes[pos]
			pos++
		}
		spread++
	}
}

func retainFrom(m Matcher, str string) string {
	return m.Negate().RemoveFrom(str)
}

func replaceFromRune(m Matcher, str string, replacement rune) string {
	runes := []rune(str)

	pos := m.IndexInRunes(runes, 0)
	if pos == -1 {
		return str
	}

	runes[pos] = replacement
	for i := pos + 1; i < len(runes); i++ {
		if m.Matches(runes[i]) {
			runes[i] = replacement
		}
	}
	return string(runes)
}

func replaceFrom(m Matcher, str string, replacement string) string {
	replacementLen := len(replacement)
	if replacementLen == 0 {
		return m.RemoveFrom(str)
	}
	if replacementLen == 1 {
		r, _ := utf8.DecodeRuneInString(replacement)
		return m.ReplaceFromRune(str, r)
	}

	runes := []rune(str)
	pos := m.IndexInRunes(runes, 0)
	if pos == -1 {
		return str
	}

	buf := strings.Builder{}
	buf.Grow((len(str) * 3 / 2) + 16)

	oldpos := 0
	for {
		buf.WriteString(string(runes[oldpos:pos]))
		buf.WriteString(replacement)
		oldpos = pos + 1
		pos = m.IndexInRunes(runes, oldpos)
		if pos == -1 {
			break
		}
	}

	buf.WriteString(string(runes[oldpos:]))
	return buf.String()
}

func trimFrom(m Matcher, str string) string {
	runes := []rune(str)
	length := len(runes)

	var first int
	var last int

	for first = 0; first < length; first++ {
		if !m.Matches(runes[first]) {
			break
		}
	}
	for last = length - 1; last > first; last-- {
		if !m.Matches(runes[last]) {
			break
		}
	}

	return string(runes[first : last+1])
}

func trimLeadingFrom(m Matcher, str string) string {
	runes := []rune(str)
	length := len(runes)

	for first := 0; first < length; first++ {
		if !m.Matches(runes[first]) {
			return string(runes[first:length])
		}
	}
	return ""
}

func trimTrailingFrom(m Matcher, str string) string {
	runes := []rune(str)
	for last := len(runes) - 1; last >= 0; last-- {
		if !m.Matches(runes[last]) {
			return string(runes[0 : last+1])
		}
	}
	return ""
}

func collapseFrom(m Matcher, str string, replacement rune) string {
	runes := []rune(str)

	length := len(runes)
	for i := 0; i < length; i++ {
		c := runes[i]
		if m.Matches(c) {
			if c == replacement && (i == length-1 || !m.Matches(runes[i+1])) {
				// a no-op replacement
				i++
			} else {
				builder := &strings.Builder{}
				builder.Grow(length)
				builder.WriteString(string(runes[0:i]))
				builder.WriteRune(replacement)
				return finishCollapseFrom(m, runes, i+1, length, replacement, builder, true)
			}
		}
	}
	// no replacement needed
	return str
}

func trimAndCollapseFrom(m Matcher, str string, replacement rune) string {
	runes := []rune(str)

	length := len(runes)
	first := 0
	last := length - 1

	for first < length && m.Matches(runes[first]) {
		first++
	}

	for last > first && m.Matches(runes[last]) {
		last--
	}

	if first == 0 && last == length-1 {
		return collapseFrom(m, str, replacement)
	}
	builder := &strings.Builder{}
	builder.Grow(last + 1 - first)
	return finishCollapseFrom(m, runes, first, last+1, replacement, builder, false)

}

func finishCollapseFrom(m Matcher, runes []rune, start, end int, repl rune, builder *strings.Builder, inMatchingGroup bool) string {
	for i := start; i < end; i++ {
		c := runes[i]
		if m.Matches(c) {
			if !inMatchingGroup {
				builder.WriteRune(repl)
				inMatchingGroup = true
			}
		} else {
			builder.WriteRune(c)
			inMatchingGroup = false
		}
	}
	return builder.String()
}

func showCharacter(r rune) string {
	return "\\u" + fmt.Sprintf("%U", r)[2:]
}
