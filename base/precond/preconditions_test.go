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

package precond_test

import (
	. "github.com/abc-inc/goava/base/precond"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestCheckArgument(t *testing.T) {
	NoError(t, CheckArgument(true))
	EqualError(t, CheckArgument(false), "invalid argument")
}

func TestCheckState(t *testing.T) {
	NoError(t, CheckState(true))
	EqualError(t, CheckState(false), "illegal state")
}

func TestCheckNotNil(t *testing.T) {
	v, err := CheckNotNil("")
	NotNil(t, v)
	NoError(t, err)

	v, err = CheckNotNilf(nil, "%v", nil)
	Nil(t, v)
	EqualError(t, err, "<nil>")
}

func TestCheckElementIndex(t *testing.T) {
	tests := []struct {
		index int
		size  int
		msg   string
	}{
		{0, 1, ""},
		{0, -1, "negative size: -1"},
		{-1, 1, "index (-1) must not be negative"},
		{2, 1, "index (2) must be less than size (1)"},
	}

	for _, tc := range tests {
		i, err := CheckElementIndex(tc.index, tc.size)
		if tc.msg == "" {
			Equal(t, tc.index, i)
			NoError(t, err)
		} else {
			Equal(t, -1, i)
			EqualError(t, err, tc.msg)
		}
	}

	_, err := CheckElementIndexf(1, 0, "custom %s", "message")
	EqualError(t, err, "custom message")
}

func TestCheckPositionIndex(t *testing.T) {
	tests := []struct {
		index int
		size  int
		msg   string
	}{
		{1, 1, ""},
		{0, -1, "negative size: -1"},
		{-1, 1, "index (-1) must not be negative"},
		{2, 1, "index (2) must not be greater than size (1)"},
	}

	for _, tc := range tests {
		i, err := CheckPositionIndex(tc.index, tc.size)
		if tc.msg == "" {
			Equal(t, tc.index, i)
			NoError(t, err)
		} else {
			Equal(t, -1, i)
			EqualError(t, err, tc.msg)
		}
	}

	_, err := CheckPositionIndexf(1, 0, "custom %s", "message")
	EqualError(t, err, "custom message")
}

func TestCheckNonnegative(t *testing.T) {
	v, err := CheckNonnegative(0, "a")
	NoError(t, err)
	Equal(t, 0, v)

	v, err = CheckNonnegative(-1, "a")
	EqualError(t, err, "a cannot be negative but was: -1")
}

func TestCheckNonnegative64(t *testing.T) {
	v, err := CheckNonnegative64(0, "a")
	NoError(t, err)
	Equal(t, int64(0), v)

	v, err = CheckNonnegative64(-1, "a")
	EqualError(t, err, "a cannot be negative but was: -1")
}
