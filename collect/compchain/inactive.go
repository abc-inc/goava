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

var less = inactive{-1}

var greater = inactive{1}

type inactive struct {
	result int
}

func (c inactive) CompareFunc(left, right interface{}, cmp func(l, r interface{}) int) ComparisonChain {
	return c
}

func (c inactive) CompareInt(left, right int) ComparisonChain {
	return c
}

func (c inactive) CompareInt8(left, right int8) ComparisonChain {
	return c
}

func (c inactive) CompareInt16(left, right int16) ComparisonChain {
	return c
}

func (c inactive) CompareInt32(left, right int32) ComparisonChain {
	return c
}

func (c inactive) CompareInt64(left, right int64) ComparisonChain {
	return c
}

func (c inactive) CompareUInt(left, right uint) ComparisonChain {
	return c
}

func (c inactive) CompareUInt8(left, right uint8) ComparisonChain {
	return c
}

func (c inactive) CompareUInt16(left, right uint16) ComparisonChain {
	return c
}

func (c inactive) CompareUInt32(left, right uint32) ComparisonChain {
	return c
}

func (c inactive) CompareUInt64(left, right uint64) ComparisonChain {
	return c
}

func (c inactive) CompareUIntPtr(left, right uintptr) ComparisonChain {
	return c
}

func (c inactive) CompareFloat32(left, right float32) ComparisonChain {
	return c
}

func (c inactive) CompareFloat64(left, right float64) ComparisonChain {
	return c
}

func (c inactive) CompareTrueFirst(left, right bool) ComparisonChain {
	return c
}

func (c inactive) CompareFalseFirst(left, right bool) ComparisonChain {
	return c
}

func (c inactive) CompareString(left, right string) ComparisonChain {
	return c
}

func (c inactive) Result() int {
	return c.result
}
