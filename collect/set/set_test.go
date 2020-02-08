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

package set_test

import (
	"github.com/abc-inc/goava/collect/set"
	. "github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := set.Empty()
	True(t, s.Add(""))
	True(t, s.Add(0))
	False(t, s.Add(0))
	Equal(t, 2, s.Size())
}

func TestSet_Clear(t *testing.T) {
	s := set.Singleton(false)
	Equal(t, 1, s.Size())

	s.Clear()
	Equal(t, 0, s.Size())

	s.Clear()
	Equal(t, 0, s.Size())
}

func TestSet_Contains(t *testing.T) {
	s := set.Of("a", "b")
	True(t, s.Contains("a"))
	False(t, s.Contains("c"))
}

func TestSet_ContainsAll(t *testing.T) {
	s := set.Of("a", "b")
	True(t, set.Empty().ContainsAll(set.Empty()))
	True(t, s.ContainsAll(set.Empty()))
	True(t, s.ContainsAll(s))
	True(t, s.ContainsAll(set.Of("b", "a")))
	False(t, s.ContainsAll(set.Of("b", "c")))
}

func TestSet_AddAll(t *testing.T) {
	s := set.Empty()
	s.AddAll(set.Empty())
	True(t, s.IsEmpty())

	s.AddAll(set.Of("a", "b"))
	s.AddAll(set.Of("a", "b"))
	Equal(t, 2, s.Size())
}

func TestSet_Remove(t *testing.T) {
	s := set.Singleton("0")

	False(t, s.Remove(0))
	False(t, s.IsEmpty())

	True(t, s.Remove("0"))
	True(t, s.IsEmpty())
}

func TestSet_RemoveAll(t *testing.T) {
	s := set.Of("a", "b", "c")
	False(t, s.RemoveAll(set.Empty()))
	True(t, s.RemoveAll(set.Of("b", "a")))
	Equal(t, s, set.Of("c"))
	True(t, s.RemoveAll(s))
	True(t, s.IsEmpty())
}

func TestSet_RetainAll(t *testing.T) {
	s := set.Of("a", "b", "c")
	False(t, s.RetainAll(s))
	True(t, s.RetainAll(set.Of("a")))
	True(t, s.RetainAll(set.Empty()))
	True(t, s.IsEmpty())
}

func TestSet_ToArray(t *testing.T) {
	Empty(t, set.Of().ToArray())

	s := set.Of(0, 1)
	is := s.ToArray()
	Equal(t, 2, len(is))
	sort.Slice(is, func(i, j int) bool { return is[i].(int) < is[j].(int) })
	Equal(t, 1, is[1])
}

func TestCopyOf(t *testing.T) {
	s := set.Of("a", "b", "c")
	c := set.CopyOf(s)
	Equal(t, s, c)
}
