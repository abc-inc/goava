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

package compchain_test

import (
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/abc-inc/goava/collect/compchain"
	. "github.com/stretchr/testify/require"
)

var dontCompareMe struct{}
var incomparable = func(l, r interface{}) int { panic("don't call me") }

func TestDegenerate(t *testing.T) {
	// kinda bogus, but who cares?
	Equal(t, 0, compchain.Start().Result())
}

func TestOneEqual(t *testing.T) {
	Equal(t, 0, compchain.Start().CompareString("a", "a").Result())
}

func TestOneEqualUsingComparator(t *testing.T) {
	ci := func(l, r interface{}) int {
		return strings.Compare(strings.ToLower(l.(string)), strings.ToLower(r.(string)))
	}
	Equal(t, 0, compchain.Start().CompareFunc("a", "A", ci).Result())
}

func TestManyEqual(t *testing.T) {
	cmpIntStr := func(l, r interface{}) int { return strings.Compare(strconv.Itoa(l.(int)), r.(string)) }

	Equal(t, 0,
		compchain.Start().
			CompareFunc(0, "0", cmpIntStr).
			CompareInt(0, 0).
			CompareInt8(math.MinInt8, math.MinInt8).
			CompareInt16(math.MinInt16, math.MinInt16).
			CompareInt32(math.MinInt32, math.MinInt32).
			CompareInt64(math.MinInt64, math.MinInt64).
			CompareUInt(0, 0).
			CompareUInt8(math.MaxUint8, math.MaxUint8).
			CompareUInt16(math.MaxUint16, math.MaxUint16).
			CompareUInt32(math.MaxUint32, math.MaxUint32).
			CompareUInt64(math.MaxUint64, math.MaxUint64).
			CompareUIntPtr(math.MaxUint64, math.MaxUint64).
			CompareFalseFirst(true, true).
			CompareTrueFirst(true, true).
			CompareFloat32(1.0, 1.0).
			CompareFloat64(1.0, 1.0).
			CompareString("a", "a").
			Result())
}

func TestShortCircuitLess(t *testing.T) {
	Greater(t, 0, compchain.Start().
		CompareString("a", "b").
		CompareFunc(dontCompareMe, dontCompareMe, incomparable).
		Result())
}

func TestShortCircuitGreater(t *testing.T) {
	Less(t, 0, compchain.Start().
		CompareString("b", "a").
		CompareFunc(dontCompareMe, dontCompareMe, incomparable).
		Result())
}

func TestShortCircuitSecondStep(t *testing.T) {
	Greater(t, 0, compchain.Start().
		CompareString("a", "a").
		CompareString("a", "b").
		CompareFunc(dontCompareMe, dontCompareMe, incomparable).
		CompareInt(0, 0).
		CompareInt8(math.MaxInt8, math.MinInt8).
		CompareInt16(math.MaxInt16, math.MinInt16).
		CompareInt32(math.MaxInt32, math.MinInt32).
		CompareInt64(math.MaxInt64, math.MinInt64).
		CompareUInt(0, 0).
		CompareUInt8(math.MaxUint8, 0).
		CompareUInt16(math.MaxUint16, 0).
		CompareUInt32(math.MaxUint32, 0).
		CompareUInt64(math.MaxUint64, 0).
		CompareUIntPtr(math.MaxUint64, 0).
		CompareFalseFirst(true, false).
		CompareTrueFirst(false, true).
		CompareFloat32(2.0, 1.0).
		CompareFloat64(2.0, 1.0).
		CompareString("b", "a").
		Result())
}

func TestCompareFalseFirst(t *testing.T) {
	True(t, compchain.Start().CompareFalseFirst(true, true).Result() == 0)
	True(t, compchain.Start().CompareFalseFirst(true, false).Result() > 0)
	True(t, compchain.Start().CompareFalseFirst(false, true).Result() < 0)
	True(t, compchain.Start().CompareFalseFirst(false, false).Result() == 0)
}

func TestCompareTrueFirst(t *testing.T) {
	True(t, compchain.Start().CompareTrueFirst(true, true).Result() == 0)
	True(t, compchain.Start().CompareTrueFirst(true, false).Result() < 0)
	True(t, compchain.Start().CompareTrueFirst(false, true).Result() > 0)
	True(t, compchain.Start().CompareTrueFirst(false, false).Result() == 0)
}
