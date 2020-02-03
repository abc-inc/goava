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
)

// Int64 is the discrete domain for values of type int64.
type Int64 struct {
}

// Offset returns, conceptually, "origin + distance", or equivalently, the result of calling Next origin distance times.
//
// Note that the calculation is done in int64 i.e., distance>math.MaxInt64 leads to an overflow error.
func (d Int64) Offset(origin int64, distance uint64) (int64, error) {
	if _, err := precond.CheckNonnegative64(int64(distance), "distance"); err != nil {
		return 0, err
	}
	v := origin + int64(distance)
	if err := precond.CheckStatef(v >= origin, "int64 overflow"); err != nil {
		return 0, err
	}
	return v, nil
}

// Next returns the unique least value of type int64 that is greater than value, or an error if none exists.
//
// Inverse operation to Previous.
func (d Int64) Next(v int64) (int64, error) {
	if err := precond.CheckStatef(v < d.MaxValue(), "int64 overflow"); err != nil {
		return 0, err
	}
	return v + 1, nil
}

// Previous returns the unique greatest value of type int64 that is less than value, or an error if none exists.
//
// Inverse operation to Next.
func (d Int64) Previous(v int64) (int64, error) {
	if err := precond.CheckStatef(v > d.MinValue(), "int64 underflow"); err != nil {
		return 0, err
	}
	return v - 1, nil
}

// Distance returns a signed value indicating how many nested invocations of Next (if positive) or Previous
// (if negative) are needed to reach end starting from start.
//
// For example, if end = Next(Next(Next(start))), then Distance(start, end) == 3 and Distance(end, start) == -3.
// As well, Distance(a, a) is always zero.
func (d Int64) Distance(start, end int64) (int64, error) {
	r := end - start
	if err := precond.CheckStatef(end <= start || r >= 0, "int64 overflow"); err != nil {
		return 0, err
	}
	if err := precond.CheckStatef(end >= start || r <= 0, "int64 underflow"); err != nil {
		return 0, err
	}
	return r, nil
}

// MinValue returns the minimum value of type int64.
//
// The minimum value m is the unique value for which o<m never returns a true for any value o of type int64.
func (d Int64) MinValue() int64 {
	return math.MinInt64
}

// MaxValue returns the maximum value of type int64.
//
// The maximum value m is the unique value for which o>m never returns a true for any value o of type int64.
func (d Int64) MaxValue() int64 {
	return math.MaxInt64
}

// String returns a string representation of this domain.
func (d Int64) String() string {
	return "domain.Int64"
}
