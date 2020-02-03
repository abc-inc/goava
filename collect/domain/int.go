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
	"github.com/abc-inc/goava/base/precond"
	"golang.org/x/tools/container/intsets"
)

// Int is the discrete domain for values of type int.
type Int struct {
}

// Offset returns, conceptually, "origin + distance", or equivalently, the result of calling Next origin distance times.
//
// Note that the calculation is done in int64 i.e., distance>math.MaxInt64 leads to an overflow error.
func (d Int) Offset(origin int, distance uint64) (int64, error) {
	if _, err := precond.CheckNonnegative64(int64(distance), "distance"); err != nil {
		return 0, err
	}
	v := int64(origin) + int64(distance)
	if err := precond.CheckStatef(v >= int64(origin), "int overflow"); err != nil {
		return 0, err
	}
	return v, nil
}

// Next returns the unique least value of type int that is greater than value, or an error if none exists.
//
// Inverse operation to Previous.
func (d Int) Next(v int) (int, error) {
	if err := precond.CheckStatef(v < d.MaxValue(), "int overflow"); err != nil {
		return 0, err
	}
	return v + 1, nil
}

// Previous returns the unique greatest value of type int that is less than value, or an error if none exists.
//
// Inverse operation to Next.
func (d Int) Previous(v int) (int, error) {
	if err := precond.CheckStatef(v > d.MinValue(), "int underflow"); err != nil {
		return 0, err
	}
	return v - 1, nil
}

// Distance returns a signed value indicating how many nested invocations of Next (if positive) or Previous
// (if negative) are needed to reach end starting from start.
//
// For example, if end = Next(Next(Next(start))), then Distance(start, end) == 3 and Distance(end, start) == -3.
// As well, Distance(a, a) is always zero.
func (d Int) Distance(start, end int) (int64, error) {
	r := int64(end) - int64(start)
	if err := precond.CheckStatef(end <= start || r >= 0, "int overflow"); err != nil {
		return 0, err
	}
	if err := precond.CheckStatef(end >= start || r <= 0, "int underflow"); err != nil {
		return 0, err
	}
	return r, nil
}

// MinValue returns the minimum value of type int.
//
// The minimum value m is the unique value for which o<m never returns a true for any value o of type int.
func (d Int) MinValue() int {
	return intsets.MinInt
}

// MaxValue returns the maximum value of type int.
//
// The maximum value m is the unique value for which o>m never returns a true for any value o of type int.
func (d Int) MaxValue() int {
	return intsets.MaxInt
}

// String returns a string representation of this domain.
func (d Int) String() string {
	return "domain.Int"
}
