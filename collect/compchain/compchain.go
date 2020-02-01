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

// ComparisonChain performs a chained comparison statement.
//
// The value of this expression will have the same sign as the first nonzero comparison result in the chain, or will be
// zero if every comparison result was zero.
//
// Note: ComparisonChain instances are immutable. For this utility to work correctly, calls must be chained as
// illustrated in the example.
//
// Performance note: Even though the ComparisonChain caller always invokes its compare functions unconditionally, the
// ComparisonChain implementation stops comparing its inputs and as soon as one of them returns a nonzero result.
// This optimization is typically important only in the presence of expensive comparisons.
type ComparisonChain interface {

	// Compares two objects using a comparator, if the result of this comparison chain has not already been determined.
	CompareFunc(left, right interface{}, cmp func(l, r interface{}) int) ComparisonChain

	// Compares two int values, if the result of this comparison chain has not already been determined.
	CompareInt(left, right int) ComparisonChain

	// Compares two int8 values, if the result of this comparison chain has not already been determined.
	CompareInt8(left, right int8) ComparisonChain

	// Compares two int16 values, if the result of this comparison chain has not already been determined.
	CompareInt16(left, right int16) ComparisonChain

	// Compares two int32 values, if the result of this comparison chain has not already been determined.
	CompareInt32(left, right int32) ComparisonChain

	// Compares two int64 values, if the result of this comparison chain has not already been determined.
	CompareInt64(left, right int64) ComparisonChain

	// Compares two uint values, if the result of this comparison chain has not already been determined.
	CompareUInt(left, right uint) ComparisonChain

	// Compares two uint8 values, if the result of this comparison chain has not already been determined.
	CompareUInt8(left, right uint8) ComparisonChain

	// Compares two uint16 values, if the result of this comparison chain has not already been determined.
	CompareUInt16(left, right uint16) ComparisonChain

	// Compares two uint32 values, if the result of this comparison chain has not already been determined.
	CompareUInt32(left, right uint32) ComparisonChain

	// Compares two uint64 values, if the result of this comparison chain has not already been determined.
	CompareUInt64(left, right uint64) ComparisonChain

	// Compares two uintptr values, if the result of this comparison chain has not already been determined.
	CompareUIntPtr(left, right uintptr) ComparisonChain

	// Compares two float32 values, if the result of this comparison chain has not already been determined.
	CompareFloat32(left, right float32) ComparisonChain

	// Compares two float64 values, if the result of this comparison chain has not already been determined.
	CompareFloat64(left, right float64) ComparisonChain

	// Compares two bool, considering true to be less than false, if the result of this comparison chain has not
	// already been determined.
	CompareTrueFirst(left, right bool) ComparisonChain

	// Compares two bool values, considering false to be less than true, if the result of this comparison chain has not
	// already been determined.
	CompareFalseFirst(left, right bool) ComparisonChain

	// Compares two string values, if the result of this comparison chain has not already been determined.
	CompareString(left, right string) ComparisonChain

	// Ends this comparison chain and returns its result: a value having the same sign as the first nonzero comparison
	// result in the chain, or zero if every result was zero.
	Result() int
}

var a = active{}

// Start begins a new chained comparison statement.
func Start() ComparisonChain {
	return a
}
