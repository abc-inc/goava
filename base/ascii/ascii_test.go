// Copyright 2021 The Goava authors.
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

package ascii_test

import (
	"strconv"
	"testing"

	. "github.com/abc-inc/goava/base/ascii"
	. "github.com/stretchr/testify/require"
)

// The Unicode points 00c1 and 00e1 are the upper- and lowercase forms of A-with-acute-accent,
// Á and á.
const ignored = "`10-=~!@#$%^&*()_+[]\\{}|;':\",./<>?'\u00c1\u00e1\n"

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestCharsIgnored(t *testing.T) {
	for _, c := range []byte(ignored) {
		Equal(t, c, ToLowerCase(c))
		Equal(t, c, ToUpperCase(c))
		False(t, IsLowerCase(c))
		False(t, IsUpperCase(c))
	}
}

func TestCharsLower(t *testing.T) {
	for _, c := range []byte(lower) {
		Equal(t, c, ToLowerCase(c))
		NotEqual(t, c, ToUpperCase(c))
		True(t, IsLowerCase(c))
		False(t, IsUpperCase(c))
	}
}

func TestCharsUpper(t *testing.T) {
	for _, c := range []byte(upper) {
		NotEqual(t, c, ToLowerCase(c))
		Equal(t, c, ToUpperCase(c))
		False(t, IsLowerCase(c))
		True(t, IsUpperCase(c))
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		seq                 string
		maxLen              int
		truncationIndicator string
	}
	tests := []struct {
		want string
		args args
	}{
		{"foobar", args{"foobar", 10, "..."}},
		{"fo...", args{"foobar", 5, "..."}},
		{"foobar", args{"foobar", 6, "..."}},
		{"...", args{"foobar", 3, "..."}},
		{"foobar", args{"foobar", 10, "…"}},
		{"foo…", args{"foobar", 4, "…"}},
		{"fo--", args{"foobar", 4, "--"}},
		{"foobar", args{"foobar", 6, "…"}},
		{"foob…", args{"foobar", 5, "…"}},
		{"foo", args{"foobar", 3, ""}},
		{"", args{"", 5, ""}},
		{"", args{"", 5, "..."}},
		{"", args{"", 0, ""}},
	}
	for _, tt := range tests {
		t.Run(tt.args.seq+strconv.Itoa(tt.args.maxLen)+tt.args.truncationIndicator, func(t *testing.T) {
			got, err := Truncate(tt.args.seq, tt.args.maxLen, tt.args.truncationIndicator)
			NoError(t, err)
			Equal(t, tt.want, got)
		})
	}
}

func TestTruncateIllegalArguments(t *testing.T) {
	_, err := Truncate("foobar", 2, "...")
	EqualError(t, err, "maxLen (2) must be >= length of the truncation indicator (3)")

	_, err = Truncate("foobar", 8, "1234567890")
	EqualError(t, err, "maxLen (8) must be >= length of the truncation indicator (10)")

	_, err = Truncate("foobar", -1, "...")
	EqualError(t, err, "maxLen (-1) must be >= length of the truncation indicator (3)")

	_, err = Truncate("foobar", -1, "")
	EqualError(t, err, "maxLen (-1) must be >= length of the truncation indicator (0)")
}
