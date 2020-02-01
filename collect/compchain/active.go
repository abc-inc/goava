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

package compchain

import (
	"github.com/abc-inc/goava/primitives/bools"
	"github.com/abc-inc/goava/primitives/floats"
	"github.com/abc-inc/goava/primitives/ints"
	"github.com/abc-inc/goava/primitives/uints"
	"strings"
)

type active struct{}

func (c active) CompareFunc(left, right interface{}, cmp func(l, r interface{}) int) ComparisonChain {
	return classify(cmp(left, right))
}

func (c active) CompareInt(left, right int) ComparisonChain {
	return classify(ints.Compare(left, right))
}

func (c active) CompareInt8(left, right int8) ComparisonChain {
	return classify(ints.Compare8(left, right))
}

func (c active) CompareInt16(left, right int16) ComparisonChain {
	return classify(ints.Compare16(left, right))
}

func (c active) CompareInt32(left, right int32) ComparisonChain {
	return classify(ints.Compare32(left, right))
}

func (c active) CompareInt64(left, right int64) ComparisonChain {
	return classify(ints.Compare64(left, right))
}

func (c active) CompareUInt(left, right uint) ComparisonChain {
	return classify(uints.Compare(left, right))
}

func (c active) CompareUInt8(left, right uint8) ComparisonChain {
	return classify(uints.Compare8(left, right))
}

func (c active) CompareUInt16(left, right uint16) ComparisonChain {
	return classify(uints.Compare16(left, right))
}

func (c active) CompareUInt32(left, right uint32) ComparisonChain {
	return classify(uints.Compare32(left, right))
}

func (c active) CompareUInt64(left, right uint64) ComparisonChain {
	return classify(uints.Compare64(left, right))
}

func (c active) CompareUIntPtr(left, right uintptr) ComparisonChain {
	return classify(uints.ComparePtr(left, right))
}

func (c active) CompareFloat32(left, right float32) ComparisonChain {
	return classify(floats.Compare32(left, right))
}

func (c active) CompareFloat64(left, right float64) ComparisonChain {
	return classify(floats.Compare64(left, right))
}

func (c active) CompareTrueFirst(left, right bool) ComparisonChain {
	return classify(bools.Compare(right, left)) // reversed
}

func (c active) CompareFalseFirst(left, right bool) ComparisonChain {
	return classify(bools.Compare(left, right))
}

func (c active) CompareString(left, right string) ComparisonChain {
	return classify(strings.Compare(left, right))
}

func (c active) Result() int {
	return 0
}

func classify(result int) ComparisonChain {
	switch {
	case result < 0:
		return less
	case result > 0:
		return greater
	default:
		return a
	}
}
