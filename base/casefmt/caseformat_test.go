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

package casefmt_test

import (
	"reflect"
	"testing"

	. "github.com/abc-inc/goava/base/casefmt"
	. "github.com/stretchr/testify/require"
)

func TestIdentity(t *testing.T) {
	tests := []CaseFormat{LowerCamel{}, LowerHyphen{}, LowerUnderscore{}, UpperCamel{}, UpperUnderscore{}}

	for _, cf := range tests {
		t.Run(reflect.TypeOf(cf).Name()+" to "+reflect.TypeOf(cf).Name(), func(t *testing.T) {
			Equal(t, "foo", cf.To(cf, "foo"))
			Equal(t, "", cf.To(cf, ""))
			Equal(t, " ", cf.To(cf, " "))
		})
	}
}

func TestNilArgs(t *testing.T) {
	tests := []CaseFormat{LowerCamel{}, LowerHyphen{}, LowerUnderscore{}, UpperCamel{}, UpperUnderscore{}}

	for _, cf := range tests {
		t.Run(reflect.TypeOf(cf).Name()+" to "+reflect.TypeOf(cf).Name(), func(t *testing.T) {
			Panics(t, func() { cf.To(nil, "foo") })
		})
	}
}

func Test(t *testing.T) {
	lCml := LowerCamel{}
	lHyp := LowerHyphen{}
	lUnd := LowerUnderscore{}
	uCml := UpperCamel{}
	uUnd := UpperUnderscore{}

	tests := []struct {
		srcFmt CaseFormat
		tgtFmt CaseFormat
		in     string
		want   string
	}{
		{srcFmt: lHyp, tgtFmt: lHyp, in: "foo", want: "foo"},
		{srcFmt: lHyp, tgtFmt: lHyp, in: "foo-bar", want: "foo-bar"},
		{srcFmt: lHyp, tgtFmt: lUnd, in: "foo", want: "foo"},
		{srcFmt: lHyp, tgtFmt: lUnd, in: "foo-bar", want: "foo_bar"},
		{srcFmt: lHyp, tgtFmt: lCml, in: "foo", want: "foo"},
		{srcFmt: lHyp, tgtFmt: lCml, in: "foo-bar", want: "fooBar"},
		{srcFmt: lHyp, tgtFmt: uCml, in: "foo", want: "Foo"},
		{srcFmt: lHyp, tgtFmt: uCml, in: "foo-bar", want: "FooBar"},
		{srcFmt: lHyp, tgtFmt: uUnd, in: "foo", want: "FOO"},
		{srcFmt: lHyp, tgtFmt: uUnd, in: "foo-bar", want: "FOO_BAR"},

		{srcFmt: lUnd, tgtFmt: lHyp, in: "foo", want: "foo"},
		{srcFmt: lUnd, tgtFmt: lHyp, in: "foo_bar", want: "foo-bar"},
		{srcFmt: lUnd, tgtFmt: lUnd, in: "foo", want: "foo"},
		{srcFmt: lUnd, tgtFmt: lUnd, in: "foo_bar", want: "foo_bar"},
		{srcFmt: lUnd, tgtFmt: lCml, in: "foo", want: "foo"},
		{srcFmt: lUnd, tgtFmt: lCml, in: "foo_bar", want: "fooBar"},
		{srcFmt: lUnd, tgtFmt: uCml, in: "foo", want: "Foo"},
		{srcFmt: lUnd, tgtFmt: uCml, in: "foo_bar", want: "FooBar"},
		{srcFmt: lUnd, tgtFmt: uUnd, in: "foo", want: "FOO"},
		{srcFmt: lUnd, tgtFmt: uUnd, in: "foo_bar", want: "FOO_BAR"},

		{srcFmt: lCml, tgtFmt: lHyp, in: "foo", want: "foo"},
		{srcFmt: lCml, tgtFmt: lHyp, in: "fooBar", want: "foo-bar"},
		{srcFmt: lCml, tgtFmt: lHyp, in: "HTTP", want: "h-t-t-p"},
		{srcFmt: lCml, tgtFmt: lUnd, in: "foo", want: "foo"},
		{srcFmt: lCml, tgtFmt: lUnd, in: "fooBar", want: "foo_bar"},
		{srcFmt: lCml, tgtFmt: lUnd, in: "hTTP", want: "h_t_t_p"},
		{srcFmt: lCml, tgtFmt: lCml, in: "foo", want: "foo"},
		{srcFmt: lCml, tgtFmt: lCml, in: "fooBar", want: "fooBar"},
		{srcFmt: lCml, tgtFmt: uCml, in: "foo", want: "Foo"},
		{srcFmt: lCml, tgtFmt: uCml, in: "fooBar", want: "FooBar"},
		{srcFmt: lCml, tgtFmt: uCml, in: "hTTP", want: "HTTP"},
		{srcFmt: lCml, tgtFmt: uUnd, in: "foo", want: "FOO"},
		{srcFmt: lCml, tgtFmt: uUnd, in: "fooBar", want: "FOO_BAR"},

		{srcFmt: uCml, tgtFmt: lHyp, in: "Foo", want: "foo"},
		{srcFmt: uCml, tgtFmt: lHyp, in: "FooBar", want: "foo-bar"},
		{srcFmt: uCml, tgtFmt: lUnd, in: "Foo", want: "foo"},
		{srcFmt: uCml, tgtFmt: lUnd, in: "FooBar", want: "foo_bar"},
		{srcFmt: uCml, tgtFmt: lCml, in: "Foo", want: "foo"},
		{srcFmt: uCml, tgtFmt: lCml, in: "FooBar", want: "fooBar"},
		{srcFmt: uCml, tgtFmt: lCml, in: "HTTP", want: "hTTP"},
		{srcFmt: uCml, tgtFmt: uCml, in: "Foo", want: "Foo"},
		{srcFmt: uCml, tgtFmt: uCml, in: "FooBar", want: "FooBar"},
		{srcFmt: uCml, tgtFmt: uUnd, in: "Foo", want: "FOO"},
		{srcFmt: uCml, tgtFmt: uUnd, in: "FooBar", want: "FOO_BAR"},
		{srcFmt: uCml, tgtFmt: uUnd, in: "HTTP", want: "H_T_T_P"},
		{srcFmt: uCml, tgtFmt: uUnd, in: "H_T_T_P", want: "H__T__T__P"},

		{srcFmt: uUnd, tgtFmt: lHyp, in: "FOO", want: "foo"},
		{srcFmt: uUnd, tgtFmt: lHyp, in: "FOO_BAR", want: "foo-bar"},
		{srcFmt: uUnd, tgtFmt: lUnd, in: "FOO", want: "foo"},
		{srcFmt: uUnd, tgtFmt: lUnd, in: "FOO_BAR", want: "foo_bar"},
		{srcFmt: uUnd, tgtFmt: lCml, in: "FOO", want: "foo"},
		{srcFmt: uUnd, tgtFmt: lCml, in: "FOO_BAR", want: "fooBar"},
		{srcFmt: uUnd, tgtFmt: uCml, in: "FOO", want: "Foo"},
		{srcFmt: uUnd, tgtFmt: uCml, in: "FOO_BAR", want: "FooBar"},
		{srcFmt: uUnd, tgtFmt: uCml, in: "H_T_T_P", want: "HTTP"},
		{srcFmt: uUnd, tgtFmt: uUnd, in: "FOO", want: "FOO"},
		{srcFmt: uUnd, tgtFmt: uUnd, in: "FOO_BAR", want: "FOO_BAR"},
	}

	for _, tt := range tests {
		t.Run(reflect.TypeOf(tt.srcFmt).Name()+"To"+reflect.TypeOf(tt.tgtFmt).Name()+"_"+tt.in, func(t *testing.T) {
			Equal(t, tt.want, tt.srcFmt.To(tt.tgtFmt, tt.in))
		})
	}
}
