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

package ints_test

import (
	"math"
	"testing"

	. "github.com/abc-inc/goava/primitives/ints"
	. "github.com/stretchr/testify/require"
	"golang.org/x/tools/container/intsets"
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

func TestCheckedCast(t *testing.T) {
	v, err := CheckedCast(int64(intsets.MaxInt))
	NoError(t, err)
	Equal(t, intsets.MaxInt, v)
}

func TestCheckedCast8(t *testing.T) {
	v, err := CheckedCast8(math.MaxInt8)
	NoError(t, err)
	Equal(t, int8(math.MaxInt8), v)

	_, err = CheckedCast8(math.MaxInt8 + 1)
	EqualError(t, err, "out of range: 128")
}

func TestCheckedCast16(t *testing.T) {
	v, err := CheckedCast16(math.MaxInt16)
	NoError(t, err)
	Equal(t, int16(math.MaxInt16), v)

	_, err = CheckedCast16(math.MaxInt16 + 1)
	EqualError(t, err, "out of range: 32768")
}

func TestCheckedCast32(t *testing.T) {
	v, err := CheckedCast32(math.MaxInt32)
	NoError(t, err)
	Equal(t, int32(math.MaxInt32), v)

	_, err = CheckedCast32(math.MaxInt32 + 1)
	EqualError(t, err, "out of range: 2147483648")
}
