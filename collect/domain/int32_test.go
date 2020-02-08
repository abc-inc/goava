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

package domain_test

import (
	"github.com/abc-inc/goava/collect/domain"
	. "github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

func TestInt32_Offset(t *testing.T) {
	d := domain.Int32{}

	offset, err := d.Offset(0, 1)
	NoError(t, err)
	Equal(t, int32(1), offset)

	offset, err = d.Offset(math.MinInt32, math.MaxInt32)
	NoError(t, err)
	Equal(t, int32(-1), offset)
}

func TestInt32_OffsetErrors(t *testing.T) {
	d := domain.Int32{}

	_, err := d.Offset(math.MaxInt32, 1)
	EqualError(t, err, "out of range: "+strconv.FormatInt(math.MaxInt32+1, 10))

	_, err = d.Offset(math.MinInt32, math.MaxInt64+1)
	EqualError(t, err, "distance cannot be negative but was: "+strconv.FormatInt(math.MinInt64, 10))
}

func TestInt32_Next(t *testing.T) {
	d := domain.Int32{}

	Equal(t, int32(3), fst32(d.Next(fst32(d.Next(fst32(d.Next(0)))))))

	_, err := d.Next(math.MaxInt32)
	EqualError(t, err, "int32 overflow")
}

func TestInt32_Previous(t *testing.T) {
	d := domain.Int32{}

	Equal(t, int32(-3), fst32(d.Previous(fst32(d.Previous(fst32(d.Previous(0)))))))

	_, err := d.Previous(math.MinInt32)
	EqualError(t, err, "int32 underflow")
}

func TestInt32_Distance(t *testing.T) {
	d := domain.Int32{}

	v, err := d.Distance(math.MaxInt32, math.MaxInt32)
	NoError(t, err)
	Equal(t, int64(0), v)

	v, err = d.Distance(math.MaxInt32, math.MinInt32)
	NoError(t, err)
	Equal(t, int64(math.MinInt32)*2+1, v)

	v, err = d.Distance(math.MinInt32, math.MaxInt32)
	NoError(t, err)
	Equal(t, int64(math.MaxInt32)*2+1, v)
}

func TestInt32_String(t *testing.T) {
	Equal(t, "domain.Int32", domain.Int32{}.String())
}

func fst32(v int32, err error) int32 {
	if err != nil {
		panic(err)
	}
	return v
}
