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

package domain_test

import (
	"math"
	"testing"

	"github.com/abc-inc/goava/collect/domain"
	. "github.com/stretchr/testify/require"
)

func TestInt64_Offset(t *testing.T) {
	d := domain.Int64{}

	offset, err := d.Offset(0, 1)
	NoError(t, err)
	Equal(t, int64(1), offset)

	offset, err = d.Offset(math.MinInt64, math.MaxInt64)
	NoError(t, err)
	Equal(t, int64(-1), offset)
}

func TestInt64_OffsetErrors(t *testing.T) {
	d := domain.Int64{}

	_, err := d.Offset(math.MaxInt64, 1)
	EqualError(t, err, "int64 overflow")

	v, err := d.Offset(math.MinInt64, math.MaxInt64)
	NoError(t, err)
	Equal(t, int64(-1), v)

	v, err = d.Offset(0, math.MaxUint64)
	EqualError(t, err, "distance cannot be negative but was: -1")
}

func TestInt64_Next(t *testing.T) {
	d := domain.Int64{}

	Equal(t, int64(3), fst64(d.Next(fst64(d.Next(fst64(d.Next(0)))))))

	_, err := d.Next(math.MaxInt64)
	EqualError(t, err, "int64 overflow")
}

func TestInt64_Previous(t *testing.T) {
	d := domain.Int64{}

	Equal(t, int64(-3), fst64(d.Previous(fst64(d.Previous(fst64(d.Previous(0)))))))

	_, err := d.Previous(math.MinInt64)
	EqualError(t, err, "int64 underflow")
}

func TestInt64_Distance(t *testing.T) {
	d := domain.Int64{}

	v, err := d.Distance(math.MaxInt64, math.MaxInt64)
	NoError(t, err)
	Equal(t, int64(0), v)

	_, err = d.Distance(math.MaxInt64, math.MinInt64)
	EqualError(t, err, "int64 underflow")

	_, err = d.Distance(math.MinInt64, math.MaxInt64)
	EqualError(t, err, "int64 overflow")
}

func TestInt64_String(t *testing.T) {
	Equal(t, "domain.Int64", domain.Int64{}.String())
}

func fst64(v int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return v
}
