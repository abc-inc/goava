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

package runematcher_test

import (
	"reflect"
	"testing"

	. "github.com/abc-inc/goava/base/runematcher"
	. "github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	ms := []Matcher{
		Any(), None(), Whitespace(), BreakingWhitespace(), ASCII(), Digit(),
		Invisible(), SingleWidth(), Is('^'), IsNot('^'),
		AnyOf(""), AnyOf("."), AnyOf("01"), AnyOf("abc"),
		NoneOf(" "), InRange('a', 'z'), ForPredicate(func(rune) bool { return true }),
		Whitespace().And(BreakingWhitespace()),
		Whitespace().Or(BreakingWhitespace()),
		Whitespace().Negate(),
	}
	for _, m := range ms {
		t.Run(reflect.TypeOf(m).Name(), func(t *testing.T) {
			False(t, m.And(None()).Matches('^'))
			Equal(t, expVal(m, true, false), m.And(Any()).Matches('X'))
			Equal(t, expVal(m, true, false), m.Or(None()).Matches('X'))
			Equal(t, true, m.Or(Any()).Matches('X'))
			if reflect.TypeOf(m).Name() != "forPredicateMatcher" {
				Equal(t, m, m.Negate().Negate())
			}

			Equal(t, m != None(), m.MatchesAnyOf("^0 a."))
			Equal(t, expVal(m, true, false), m.MatchesAllOf("?"))

			Equal(t, expVal(m, false, true), m.MatchesNoneOf("?"))
			LessOrEqual(t, -1, m.IndexIn("^0 a.", -1))
			LessOrEqual(t, expVal(m, 0, -1), m.IndexIn("^0 a.", 0))
			LessOrEqual(t, expVal(m, 0, -1), m.LastIndexIn("^0 a."))
			LessOrEqual(t, expVal(m, 0, -1), m.CountIn("^0 a."))

			Equal(t, expVal(m, "", "X"), m.RemoveFrom("X"))
			Equal(t, expVal(m, "X", ""), m.RetainFrom("X"))
			Equal(t, expVal(m, "Y", "X"), m.ReplaceFromRune("X", 'Y'))
			Equal(t, expVal(m, "Y", "X"), m.ReplaceFrom("X", "Y"))

			Equal(t, expVal(m, "", "X"), m.TrimFrom("X"))
			Equal(t, expVal(m, "", "X"), m.TrimLeadingFrom("X"))
			Equal(t, expVal(m, "", "X"), m.TrimTrailingFrom("X"))
			Equal(t, expVal(m, "", ""), m.CollapseFrom("", 'Y'))
			Equal(t, expVal(m, "Y", "X"), m.CollapseFrom("X", 'Y'))
			Equal(t, expVal(m, "", "X"), m.TrimAndCollapseFrom("X", 'Y'))

			NotEmpty(t, m.String())
		})
	}
}

func expVal(m Matcher, matchVal interface{}, noMatchVal interface{}) interface{} {
	if m.Matches('X') {
		return matchVal
	}
	return noMatchVal
}
