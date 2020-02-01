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

package ints_test

import (
	. "github.com/abc-inc/goava/primitives/ints"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestCompare(t *testing.T) {
	Equal(t, 0, Compare(0, 0))
	Equal(t, -1, Compare(0, 1))
	Equal(t, +1, Compare(1, 0))
}

func TestCompare8(t *testing.T) {
	Equal(t, 0, Compare8(0, 0))
	Equal(t, -1, Compare8(0, 1))
	Equal(t, +1, Compare8(1, 0))
}

func TestCompare16(t *testing.T) {
	Equal(t, 0, Compare16(0, 0))
	Equal(t, -1, Compare16(0, 1))
	Equal(t, +1, Compare16(1, 0))
}

func TestCompare32(t *testing.T) {
	Equal(t, 0, Compare32(0, 0))
	Equal(t, -1, Compare32(0, 1))
	Equal(t, +1, Compare32(1, 0))
}

func TestCompare64(t *testing.T) {
	Equal(t, 0, Compare64(0, 0))
	Equal(t, -1, Compare64(0, 1))
	Equal(t, +1, Compare64(1, 0))
}
