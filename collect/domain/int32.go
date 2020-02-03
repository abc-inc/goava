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

package domain

import (
	"math"

	"github.com/abc-inc/goava/base/precond"
	"github.com/abc-inc/goava/primitives/ints"
)

// Int32 is the discrete domain for values of type int32.
type Int32 struct {
}

// Offset returns, conceptually, "origin + distance", or equivalently, the result of calling Next origin distance times.
//
// Note that the calculation is done in int64 i.e., distance>math.MaxInt64 leads to an overflow error.
func (d Int32) Offset(origin int32, distance uint64) (int32, error) {
	if _, err := precond.CheckNonnegative64(int64(distance), "distance"); err != nil {
		return 0, err
	}
	v, err := ints.CheckedCast32(int64(origin) + int64(distance))
	if err != nil {
		return 0, err
	}
	return v, nil
}

// Next returns the unique least value of type int32 that is greater than value, or an error if none exists.
//
// Inverse operation to Previous.
func (d Int32) Next(v int32) (int32, error) {
	if err := precond.CheckStatef(v < d.MaxValue(), "int32 overflow"); err != nil {
		return 0, err
	}
	return v + 1, nil
}

// Previous returns the unique greatest value of type int32 that is less than value, or an error if none exists.
//
// Inverse operation to Next.
func (d Int32) Previous(v int32) (int32, error) {
	if err := precond.CheckStatef(v > d.MinValue(), "int32 underflow"); err != nil {
		return 0, err
	}
	return v - 1, nil
}

// Distance returns a signed value indicating how many nested invocations of Next (if positive) or Previous
// (if negative) are needed to reach end starting from start.
//
// For example, if end = Next(Next(Next(start))), then Distance(start, end) == 3 and Distance(end, start) == -3.
// As well, Distance(a, a) is always zero.
func (d Int32) Distance(start, end int32) (int64, error) {
	return int64(end) - int64(start), nil
}

// MinValue returns the minimum value of type int32.
//
// The minimum value m is the unique value for which o<m never returns a true for any value o of type int32.
func (d Int32) MinValue() int32 {
	return math.MinInt32
}

// MaxValue returns the maximum value of type int32.
//
// The maximum value m is the unique value for which o>m never returns a true for any value o of type int32.
func (d Int32) MaxValue() int32 {
	return math.MaxInt32
}

// String returns a string representation of this domain.
func (d Int32) String() string {
	return "domain.Int32"
}
