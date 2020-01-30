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

package runematcher_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	. "github.com/abc-inc/goava/base/runematcher"
	. "github.com/stretchr/testify/require"
)

func TestAnyAndNone_logicalOps(t *testing.T) {
	// These are testing behavior that's never promised by the API, but since
	// we're lucky enough that these do pass, it saves us from having to write
	// more excruciating tests! Hooray!

	var Whatever = Is('^')

	Equal(t, Any(), None().Negate())
	Equal(t, None(), Any().Negate())

	Equal(t, Whatever, Any().And(Whatever))
	Equal(t, Any(), Any().Or(Whatever))

	Equal(t, None(), None().And(Whatever))
	Equal(t, Whatever, None().Or(Whatever))
}

func TestWhitespaceBreakingWhitespaceSubset(t *testing.T) {
	for c := rune(0); c <= unicode.MaxLatin1; c++ {
		if BreakingWhitespace().Matches(c) {
			Truef(t, Whitespace().Matches(c), "\\u"+fmt.Sprintf("%U", c)[2:])
		}
	}
}

func TestEmpty(t *testing.T) {
	matchers := []Matcher{
		Any(),
		None(),
		Is('a'),
		IsNot('a'),
		AnyOf(""),
		AnyOf("x"),
		AnyOf("xy"),
		AnyOf("Matcher"),
		AnyOf("Matcher"),
		InRange('n', 'q'),
		ForPredicate(func(r rune) bool { return r == 'c' }),
	}

	doTestEmpty := func(t *testing.T, m Matcher) {
		Equal(t, -1, m.IndexIn("", 0))
		Equal(t, -1, m.IndexIn("", 1))
		Equal(t, -1, m.IndexIn("", -1))
		Equal(t, -1, m.LastIndexIn(""))
		False(t, m.MatchesAnyOf(""))
		True(t, m.MatchesAllOf(""))
		True(t, m.MatchesNoneOf(""))
		Equal(t, "", m.RemoveFrom(""))
		Equal(t, "", m.ReplaceFromRune("", 'z'))
		Equal(t, "", m.ReplaceFrom("", "ZZ"))
		Equal(t, "", m.TrimFrom(""))
		Equal(t, 0, m.CountIn(""))
	}

	for _, m := range matchers {
		doTestEmpty(t, m)
		doTestEmpty(t, m.Negate())
	}
}

func TestNoMatches(t *testing.T) {
	doTestNoMatches(t, None(), "blah")
	doTestNoMatches(t, Is('a'), "bcde")
	doTestNoMatches(t, IsNot('a'), "aaaa")
	doTestNoMatches(t, AnyOf(""), "abcd")
	doTestNoMatches(t, AnyOf("x"), "abcd")
	doTestNoMatches(t, AnyOf("xy"), "abcd")
	doTestNoMatches(t, AnyOf("RuneMatcher"), "zxqy")
	doTestNoMatches(t, NoneOf("RuneMatcher"), "RuMa")
	doTestNoMatches(t, InRange('p', 'x'), "mom")
	doTestNoMatches(t, ForPredicate(func(r rune) bool { return r == 'c' }), "abe")
	doTestNoMatches(t, InRange('A', 'Z').And(InRange('F', 'K').Negate()), "F1a")
	doTestNoMatches(t, Digit(), "\tAz()")
	doTestNoMatches(t, Digit().And(ASCII()), "\tAz()")
	doTestNoMatches(t, SingleWidth(), "\u05bf\u3000")
}

func doTestNoMatches(t *testing.T, m Matcher, s string) {
	reallyTestNoMatches(t, m, s)
	reallyTestAllMatches(t, m.Negate(), s)
}

func TestAllMatches(t *testing.T) {
	doTestAllMatches(t, Any(), "blah")
	doTestAllMatches(t, IsNot('a'), "bcde")
	doTestAllMatches(t, Is('a'), "aaaa")
	doTestAllMatches(t, NoneOf("Matcher"), "zxqy")
	doTestAllMatches(t, AnyOf("x"), "xxxx")
	doTestAllMatches(t, AnyOf("xy"), "xyyx")
	doTestAllMatches(t, AnyOf("CharMatcher"), "ChMa")
	doTestAllMatches(t, InRange('m', 'p'), "mom")
	doTestAllMatches(t, ForPredicate(func(r rune) bool { return r == 'c' }), "ccc")
	doTestAllMatches(t, Digit(), "0123456789\u0ED0\u1B59")
	doTestAllMatches(t, Digit().And(ASCII()), "0123456789")
	doTestAllMatches(t, SingleWidth(), "\t0123ABCdef~\u007F\u007F")
}

func doTestAllMatches(t *testing.T, m Matcher, s string) {
	reallyTestAllMatches(t, m, s)
	reallyTestNoMatches(t, m.Negate(), s)
}

func reallyTestNoMatches(t *testing.T, m Matcher, s string) {
	r, _ := utf8.DecodeRuneInString(s)
	False(t, m.Matches(r))
	Equal(t, -1, m.IndexIn(s, 0))
	Equal(t, -1, m.IndexIn(s, 1))
	Equal(t, -1, m.IndexIn(s, len(s)))
	Equal(t, -1, m.IndexIn(s, len(s)+1))
	Equal(t, -1, m.IndexIn(s, -1))
	Equal(t, -1, m.IndexInRunes([]rune(s), 0))
	Equal(t, -1, m.LastIndexIn(s))
	False(t, m.MatchesAnyOf(s))
	False(t, m.MatchesAnyOf(s))
	True(t, m.MatchesNoneOf(s))
	Equal(t, s, m.RemoveFrom(s))
	Equal(t, s, m.ReplaceFromRune(s, 'z'))
	Equal(t, s, m.ReplaceFrom(s, "ZZ"))
	Equal(t, s, m.TrimFrom(s))
	Equal(t, 0, m.CountIn(s))
}

func reallyTestAllMatches(t *testing.T, m Matcher, s string) {
	r, _ := utf8.DecodeRuneInString(s)
	True(t, m.Matches(r))
	Equal(t, 0, m.IndexIn(s, 0))
	Equal(t, 1, m.IndexIn(s, 1))
	Equal(t, -1, m.IndexIn(s, len(s)))
	Equal(t, utf8.RuneCountInString(s)-1, m.LastIndexIn(s))
	True(t, m.MatchesAnyOf(s))
	True(t, m.MatchesAnyOf(s))
	False(t, m.MatchesNoneOf(s))
	Equal(t, "", m.RemoveFrom(s))
	Equal(t, strings.Repeat("z", utf8.RuneCountInString(s)), m.ReplaceFromRune(s, 'z'))
	Equal(t, strings.Repeat("ZZ", utf8.RuneCountInString(s)), m.ReplaceFrom(s, "ZZ"))
	Equal(t, "", m.TrimFrom(s))
	Equal(t, utf8.RuneCountInString(s), m.CountIn(s))
}

func TestGeneral(t *testing.T) {
	doTestGeneral(t, Is('a'), 'a', 'b')
	doTestGeneral(t, IsNot('a'), 'b', 'a')
	doTestGeneral(t, AnyOf("x"), 'x', 'z')
	doTestGeneral(t, AnyOf("xy"), 'y', 'z')
	doTestGeneral(t, AnyOf("RuneMatcher"), 'R', 'z')
	doTestGeneral(t, NoneOf("RuneMatcher"), 'z', 'R')
	doTestGeneral(t, InRange('p', 'x'), 'q', 'z')
}

func doTestGeneral(t *testing.T, m Matcher, match rune, noMatch rune) {
	doTestOneCharMatch(t, m, string(match))
	doTestOneCharNoMatch(t, m, string(noMatch))
	doTestMatchThenNoMatch(t, m, string(match)+string(noMatch))
	doTestNoMatchThenMatch(t, m, string(noMatch)+string(match))
}

func doTestOneCharMatch(t *testing.T, m Matcher, s string) {
	reallyTestOneCharMatch(t, m, s)
	reallyTestOneCharNoMatch(t, m.Negate(), s)
}

func doTestOneCharNoMatch(t *testing.T, m Matcher, s string) {
	reallyTestOneCharNoMatch(t, m, s)
	reallyTestOneCharMatch(t, m.Negate(), s)
}

func doTestMatchThenNoMatch(t *testing.T, m Matcher, s string) {
	reallyTestMatchThenNoMatch(t, m, s)
	reallyTestNoMatchThenMatch(t, m.Negate(), s)
}

func doTestNoMatchThenMatch(t *testing.T, m Matcher, s string) {
	reallyTestNoMatchThenMatch(t, m, s)
	reallyTestMatchThenNoMatch(t, m.Negate(), s)
}

func reallyTestOneCharMatch(t *testing.T, m Matcher, s string) {
	r, _ := utf8.DecodeRuneInString(s)
	True(t, m.Matches(r))
	Equal(t, 0, m.IndexIn(s, 0))
	Equal(t, -1, m.IndexIn(s, 1))
	Equal(t, 0, m.LastIndexIn(s))
	True(t, m.MatchesAnyOf(s))
	True(t, m.MatchesAnyOf(s))
	False(t, m.MatchesNoneOf(s))
	Equal(t, "", m.RemoveFrom(s))
	Equal(t, "z", m.ReplaceFromRune(s, 'z'))
	Equal(t, "ZZ", m.ReplaceFrom(s, "ZZ"))
	Equal(t, "", m.TrimFrom(s))
	Equal(t, 1, m.CountIn(s))
}

func reallyTestOneCharNoMatch(t *testing.T, m Matcher, s string) {
	r, _ := utf8.DecodeRuneInString(s)
	False(t, m.Matches(r))
	Equal(t, -1, m.IndexIn(s, 0))
	Equal(t, -1, m.IndexIn(s, 1))
	Equal(t, -1, m.LastIndexIn(s))
	False(t, m.MatchesAnyOf(s))
	False(t, m.MatchesAnyOf(s))
	True(t, m.MatchesNoneOf(s))

	Equal(t, s, m.RemoveFrom(s))
	Equal(t, s, m.ReplaceFromRune(s, 'z'))
	Equal(t, s, m.ReplaceFrom(s, "ZZ"))
	Equal(t, s, m.TrimFrom(s))
	Equal(t, 0, m.CountIn(s))
}

func reallyTestMatchThenNoMatch(t *testing.T, m Matcher, s string) {
	Equal(t, 0, m.IndexIn(s, 0))
	Equal(t, -1, m.IndexIn(s, 1))
	Equal(t, -1, m.IndexIn(s, 2))
	Equal(t, 0, m.LastIndexIn(s))
	True(t, m.MatchesAnyOf(s))
	False(t, m.MatchesAllOf(s))
	False(t, m.MatchesNoneOf(s))
	Equal(t, s[1:], m.RemoveFrom(s))
	Equal(t, "z"+s[1:], m.ReplaceFromRune(s, 'z'))
	Equal(t, "ZZ"+s[1:], m.ReplaceFrom(s, "ZZ"))
	Equal(t, s[1:], m.TrimFrom(s))
	Equal(t, 1, m.CountIn(s))
}

func reallyTestNoMatchThenMatch(t *testing.T, m Matcher, s string) {
	Equal(t, 1, m.IndexIn(s, 0))
	Equal(t, 1, m.IndexIn(s, 1))
	Equal(t, -1, m.IndexIn(s, 2))
	Equal(t, 1, m.LastIndexIn(s))
	True(t, m.MatchesAnyOf(s))
	False(t, m.MatchesAllOf(s))
	False(t, m.MatchesNoneOf(s))
	Equal(t, s[0:1], m.RemoveFrom(s))
	Equal(t, s[0:1]+"z", m.ReplaceFromRune(s, 'z'))
	Equal(t, s[0:1]+"ZZ", m.ReplaceFrom(s, "ZZ"))
	Equal(t, s[0:1], m.TrimFrom(s))
	Equal(t, 1, m.CountIn(s))
}

func TestCollapse(t *testing.T) {
	// collapsing groups of '-' into '_' or '-'
	doTestCollapse(t, "-", "_")
	doTestCollapse(t, "x-", "x_")
	doTestCollapse(t, "-x", "_x")
	doTestCollapse(t, "--", "_")
	doTestCollapse(t, "x--", "x_")
	doTestCollapse(t, "--x", "_x")
	doTestCollapse(t, "-x-", "_x_")
	doTestCollapse(t, "x-x", "x_x")
	doTestCollapse(t, "---", "_")
	doTestCollapse(t, "--x-", "_x_")
	doTestCollapse(t, "--xx", "_xx")
	doTestCollapse(t, "-x--", "_x_")
	doTestCollapse(t, "-x-x", "_x_x")
	doTestCollapse(t, "-xx-", "_xx_")
	doTestCollapse(t, "x--x", "x_x")
	doTestCollapse(t, "x-x-", "x_x_")
	doTestCollapse(t, "x-xx", "x_xx")
	doTestCollapse(t, "x-x--xx---x----x", "x_x_xx_x_x")
	doTestCollapseWithNoChange(t, "")
	doTestCollapseWithNoChange(t, "x")
	doTestCollapseWithNoChange(t, "xx")
}

func doTestCollapse(t *testing.T, in, out string) {
	// Try a few different matchers which all match '-' and not 'x'
	// Try replacement chars that both do and do not change the value.
	for _, replacement := range []rune{'_', '-'} {
		expected := strings.ReplaceAll(out, string('_'), string(replacement))
		Equal(t, expected, Is('-').CollapseFrom(in, replacement))
		Equal(t, expected, Is('-').CollapseFrom(in, replacement))
		Equal(t, expected, Is('-').Or(Is('#')).CollapseFrom(in, replacement))
		Equal(t, expected, IsNot('x').CollapseFrom(in, replacement))
		Equal(t, expected, Is('x').Negate().CollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-").CollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-#").CollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-#123").CollapseFrom(in, replacement))
	}
}

func doTestCollapseWithNoChange(t *testing.T, inout string) {
	Equal(t, inout, Is('-').CollapseFrom(inout, '_'))
	Equal(t, inout, Is('-').Or(Is('#')).CollapseFrom(inout, '_'))
	Equal(t, inout, IsNot('x').CollapseFrom(inout, '_'))
	Equal(t, inout, Is('x').Negate().CollapseFrom(inout, '_'))
	Equal(t, inout, AnyOf("-").CollapseFrom(inout, '_'))
	Equal(t, inout, AnyOf("-#").CollapseFrom(inout, '_'))
	Equal(t, inout, AnyOf("-#123").CollapseFrom(inout, '_'))
	Equal(t, inout, None().CollapseFrom(inout, '_'))
}

func TestCollapse_any(t *testing.T) {
	Equal(t, "", Any().CollapseFrom("", '_'))
	Equal(t, "_", Any().CollapseFrom("a", '_'))
	Equal(t, "_", Any().CollapseFrom("ab", '_'))
	Equal(t, "_", Any().CollapseFrom("abcd", '_'))
}

func TestTrimFrom(t *testing.T) {
	// trimming -
	doTestTrimFrom(t, "-", "")
	doTestTrimFrom(t, "x-", "x")
	doTestTrimFrom(t, "-x", "x")
	doTestTrimFrom(t, "--", "")
	doTestTrimFrom(t, "x--", "x")
	doTestTrimFrom(t, "--x", "x")
	doTestTrimFrom(t, "-x-", "x")
	doTestTrimFrom(t, "x-x", "x-x")
	doTestTrimFrom(t, "---", "")
	doTestTrimFrom(t, "--x-", "x")
	doTestTrimFrom(t, "--xx", "xx")
	doTestTrimFrom(t, "-x--", "x")
	doTestTrimFrom(t, "-x-x", "x-x")
	doTestTrimFrom(t, "-xx-", "xx")
	doTestTrimFrom(t, "x--x", "x--x")
	doTestTrimFrom(t, "x-x-", "x-x")
	doTestTrimFrom(t, "x-xx", "x-xx")
	doTestTrimFrom(t, "x-x--xx---x----x", "x-x--xx---x----x")
	// additional testing using the doc example
	Equal(t, "cat", AnyOf("ab").TrimFrom("abacatbab"))
}

func doTestTrimFrom(t *testing.T, in, out string) {
	// Try a few different matchers which all match '-' and not 'x'
	Equal(t, out, Is('-').TrimFrom(in))
	Equal(t, out, Is('-').Or(Is('#')).TrimFrom(in))
	Equal(t, out, IsNot('x').TrimFrom(in))
	Equal(t, out, Is('x').Negate().TrimFrom(in))
	Equal(t, out, AnyOf("-").TrimFrom(in))
	Equal(t, out, AnyOf("-#").TrimFrom(in))
	Equal(t, out, AnyOf("-#123").TrimFrom(in))
}

func TestTrimLeadingFrom(t *testing.T) {
	// trimming -
	doTestTrimLeadingFrom(t, "-", "")
	doTestTrimLeadingFrom(t, "x-", "x-")
	doTestTrimLeadingFrom(t, "-x", "x")
	doTestTrimLeadingFrom(t, "--", "")
	doTestTrimLeadingFrom(t, "x--", "x--")
	doTestTrimLeadingFrom(t, "--x", "x")
	doTestTrimLeadingFrom(t, "-x-", "x-")
	doTestTrimLeadingFrom(t, "x-x", "x-x")
	doTestTrimLeadingFrom(t, "---", "")
	doTestTrimLeadingFrom(t, "--x-", "x-")
	doTestTrimLeadingFrom(t, "--xx", "xx")
	doTestTrimLeadingFrom(t, "-x--", "x--")
	doTestTrimLeadingFrom(t, "-x-x", "x-x")
	doTestTrimLeadingFrom(t, "-xx-", "xx-")
	doTestTrimLeadingFrom(t, "x--x", "x--x")
	doTestTrimLeadingFrom(t, "x-x-", "x-x-")
	doTestTrimLeadingFrom(t, "x-xx", "x-xx")
	doTestTrimLeadingFrom(t, "x-x--xx---x----x", "x-x--xx---x----x")
	// additional testing using the doc example
	Equal(t, "catbab", AnyOf("ab").TrimLeadingFrom("abacatbab"))
}

func doTestTrimLeadingFrom(t *testing.T, in, out string) {
	// Try a few different matchers which all match '-' and not 'x'
	Equal(t, out, Is('-').TrimLeadingFrom(in))
	Equal(t, out, Is('-').Or(Is('#')).TrimLeadingFrom(in))
	Equal(t, out, IsNot('x').TrimLeadingFrom(in))
	Equal(t, out, Is('x').Negate().TrimLeadingFrom(in))
	Equal(t, out, AnyOf("-#").TrimLeadingFrom(in))
	Equal(t, out, AnyOf("-#123").TrimLeadingFrom(in))
}

func TestTrimTrailingFrom(t *testing.T) {
	// trimming -
	doTestTrimTrailingFrom(t, "-", "")
	doTestTrimTrailingFrom(t, "x-", "x")
	doTestTrimTrailingFrom(t, "-x", "-x")
	doTestTrimTrailingFrom(t, "--", "")
	doTestTrimTrailingFrom(t, "x--", "x")
	doTestTrimTrailingFrom(t, "--x", "--x")
	doTestTrimTrailingFrom(t, "-x-", "-x")
	doTestTrimTrailingFrom(t, "x-x", "x-x")
	doTestTrimTrailingFrom(t, "---", "")
	doTestTrimTrailingFrom(t, "--x-", "--x")
	doTestTrimTrailingFrom(t, "--xx", "--xx")
	doTestTrimTrailingFrom(t, "-x--", "-x")
	doTestTrimTrailingFrom(t, "-x-x", "-x-x")
	doTestTrimTrailingFrom(t, "-xx-", "-xx")
	doTestTrimTrailingFrom(t, "x--x", "x--x")
	doTestTrimTrailingFrom(t, "x-x-", "x-x")
	doTestTrimTrailingFrom(t, "x-xx", "x-xx")
	doTestTrimTrailingFrom(t, "x-x--xx---x----x", "x-x--xx---x----x")
	// additional testing using the doc example
	Equal(t, "abacat", AnyOf("ab").TrimTrailingFrom("abacatbab"))
}

func doTestTrimTrailingFrom(t *testing.T, in, out string) {
	// Try a few different matchers which all match '-' and not 'x'
	Equal(t, out, Is('-').TrimTrailingFrom(in))
	Equal(t, out, Is('-').Or(Is('#')).TrimTrailingFrom(in))
	Equal(t, out, IsNot('x').TrimTrailingFrom(in))
	Equal(t, out, Is('x').Negate().TrimTrailingFrom(in))
	Equal(t, out, AnyOf("-#").TrimTrailingFrom(in))
	Equal(t, out, AnyOf("-#123").TrimTrailingFrom(in))
}

func TestTrimAndCollapse(t *testing.T) {
	// collapsing groups of '-' into '_' or '-'
	doTestTrimAndCollapse(t, "", "")
	doTestTrimAndCollapse(t, "x", "x")
	doTestTrimAndCollapse(t, "-", "")
	doTestTrimAndCollapse(t, "x-", "x")
	doTestTrimAndCollapse(t, "-x", "x")
	doTestTrimAndCollapse(t, "--", "")
	doTestTrimAndCollapse(t, "x--", "x")
	doTestTrimAndCollapse(t, "--x", "x")
	doTestTrimAndCollapse(t, "-x-", "x")
	doTestTrimAndCollapse(t, "x-x", "x_x")
	doTestTrimAndCollapse(t, "---", "")
	doTestTrimAndCollapse(t, "--x-", "x")
	doTestTrimAndCollapse(t, "--xx", "xx")
	doTestTrimAndCollapse(t, "-x--", "x")
	doTestTrimAndCollapse(t, "-x-x", "x_x")
	doTestTrimAndCollapse(t, "-xx-", "xx")
	doTestTrimAndCollapse(t, "x--x", "x_x")
	doTestTrimAndCollapse(t, "x-x-", "x_x")
	doTestTrimAndCollapse(t, "x-xx", "x_xx")
	doTestTrimAndCollapse(t, "x-x--xx---x----x", "x_x_xx_x_x")
}

func doTestTrimAndCollapse(t *testing.T, in, out string) {
	// Try a few different matchers which all match '-' and not 'x'
	for _, replacement := range []rune{'_', '-'} {
		expected := strings.ReplaceAll(out, string('_'), string(replacement))
		Equal(t, expected, Is('-').TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, Is('-').Or(Is('#')).TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, IsNot('x').TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, Is('x').Negate().TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-").TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-#").TrimAndCollapseFrom(in, replacement))
		Equal(t, expected, AnyOf("-#123").TrimAndCollapseFrom(in, replacement))
	}
}

func TestReplaceFrom(t *testing.T) {
	Equal(t, "yoho", Is('a').ReplaceFromRune("yaha", 'o'))
	Equal(t, "yh", Is('a').ReplaceFrom("yaha", ""))
	Equal(t, "yoho", Is('a').ReplaceFrom("yaha", "o"))
	Equal(t, "yoohoo", Is('a').ReplaceFrom("yaha", "oo"))
	Equal(t, "12 &gt; 5", Is('>').ReplaceFrom("12 > 5", "&gt;"))
}

func TestString(t *testing.T) {
	assertToStringWorks(t, "Matcher.none()", AnyOf(""))
	assertToStringWorks(t, "Matcher.is('\\u0031')", AnyOf("1"))
	assertToStringWorks(t, "Matcher.isNot('\\u0031')", IsNot('1'))
	assertToStringWorks(t, "Matcher.anyOf(\"\\u0031\\u0032\")", AnyOf("12"))
	assertToStringWorks(t, "Matcher.anyOf(\"\\u0031\\u0032\\u0033\")", AnyOf("321"))
	assertToStringWorks(t, "Matcher.inRange('\\u0031', '\\u0033')", InRange('1', '3'))
}

func assertToStringWorks(t *testing.T, expected string, m Matcher) {
	Equal(t, expected, m.String())
	Equal(t, expected, m.Negate().Negate().String())
}
