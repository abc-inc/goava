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
	. "github.com/abc-inc/goava/base/runematcher"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestCollapseFrom(t *testing.T) {
	m := Whitespace()
	Equal(t, " ", m.CollapseFrom(" ", ' '))
	Equal(t, "babab", m.CollapseFrom("  a  a  ", 'b'))
	Equal(t, " a a ", m.CollapseFrom("  a  a  ", ' '))
}

func TestRemoveFrom(t *testing.T) {
	m := Whitespace()
	Equal(t, "aa", m.RemoveFrom("  a  a  "))
}

func TestReplaceFromEmpty(t *testing.T) {
	m := Whitespace()
	Equal(t, "aa", m.ReplaceFrom("  a  a  ", ""))
}

func TestTrimAndCollapseTwo(t *testing.T) {
	m := Whitespace()
	Equal(t, "aba", m.TrimAndCollapseFrom("  a  a  ", 'b'))
}
