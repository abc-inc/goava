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
	"golang.org/x/tools/container/intsets"
	"math"
	"testing"
)

func TestInt_Offset(t *testing.T) {
	d := domain.Int{}

	offset, err := d.Offset(0, 1)
	NoError(t, err)
	Equal(t, int64(1), offset)

	offset, err = d.Offset(math.MinInt32, math.MaxInt64)
	NoError(t, err)
	Equal(t, int64(math.MinInt32+math.MaxInt64), offset)
}

func TestInt_OffsetErrors(t *testing.T) {
	d := domain.Int{}

	v, err := d.Offset(intsets.MaxInt, 1)
	if intsets.MaxInt == math.MaxInt32 {
		Equal(t, int64(math.MaxInt32)+1, v)
	} else {
		EqualError(t, err, "int overflow")
	}

	_, err = d.Offset(intsets.MinInt, math.MaxInt64+1)
	Error(t, err)
}

func TestInt_Next(t *testing.T) {
	d := domain.Int{}

	Equal(t, 3, fst(d.Next(fst(d.Next(fst(d.Next(0)))))))

	_, err := d.Next(intsets.MaxInt)
	EqualError(t, err, "int overflow")
}

func TestInt_Previous(t *testing.T) {
	d := domain.Int{}

	Equal(t, -3, fst(d.Previous(fst(d.Previous(fst(d.Previous(0)))))))

	_, err := d.Previous(intsets.MinInt)
	EqualError(t, err, "int underflow")
}

func TestInt_Distance(t *testing.T) {
	d := domain.Int{}

	v, err := d.Distance(intsets.MaxInt, intsets.MaxInt)
	NoError(t, err)
	Equal(t, int64(0), v)

	v, err = d.Distance(intsets.MaxInt, intsets.MinInt)
	if intsets.MaxInt == math.MaxInt32 {
		Equal(t, int64(math.MinInt32)-math.MaxInt32, v)
	} else {
		EqualError(t, err, "int underflow")
	}

	_, err = d.Distance(intsets.MinInt, intsets.MaxInt)
	if intsets.MaxInt == math.MaxInt32 {
		Equal(t, int64(math.MinInt32)-math.MaxInt32, v)
	} else {
		EqualError(t, err, "int overflow")

	}
}

func TestInt_String(t *testing.T) {
	Equal(t, "domain.Int", domain.Int{}.String())
}

func fst(v int, err error) int {
	if err != nil {
		panic(err)
	}
	return v
}
