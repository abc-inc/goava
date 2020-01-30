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

package casefmt_test

import (
	. "github.com/abc-inc/goava/base/casefmt"
	. "github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/huandu/xstrings"
	"github.com/iancoleman/strcase"
	"github.com/objectundefined/caseformat"
	"github.com/ozgio/strutil"
	gostrcase "github.com/stoewer/go-strcase"
)

// Note that some test cases fail because the frameworks do not implement all format conversions identically.
// Nevertheless, the benchmark gives a good indication of the performance to expect.
type testCase struct {
	str  string
	name string
	want string
}

// n is used for artificially concatenating the string n times to benchmark longer strings.
//
// It is recommended to increase n by one order of magnitude at one point in time because of the poor performance of
// some frameworks and the increasing memory consumption for strings.
// Note that values greater than 10000 could lead to long execution times and high memory consumption.
const n = 1_000

// fmtByName maps logical names to framework specific functions (if available)
var fmtByName = map[string]func(string) string{}

var tests = []testCase{
	{"lower-hyphen-str", "ToLowerHyphen", "lower-hyphen-str"},
	{"lower-hyphen-str", "ToLowerUnderscore", "lower_hyphen_str"},
	{"lower-hyphen-str", "ToLowerCamel", "lowerHyphenStr"},
	{"lower-hyphen-str", "ToUpperCamel", "LowerHyphenStr"},
	{"lower-hyphen-str", "ToUpperUnderscore", "LOWER_HYPHEN_STR"},

	{"lower_underscore_str", "ToLowerHyphen", "lower-underscore-str"},
	{"lower_underscore_str", "ToLowerUnderscore", "lower_underscore_str"},
	{"lower_underscore_str", "ToLowerCamel", "lowerUnderscoreStr"},
	{"lower_underscore_str", "ToUpperCamel", "LowerUnderscoreStr"},
	{"lower_underscore_str", "ToUpperUnderscore", "LOWER_UNDERSCORE_STR"},

	{"lowerCamelStr", "ToLowerHyphen", "lower-camel-str"},
	{"lowerCamelStr", "ToLowerUnderscore", "lower_camel_str"},
	{"lowerCamelStr", "ToLowerCamel", "lowerCamelStr"},
	{"lowerCamelStr", "ToUpperCamel", "LowerCamelStr"},
	{"lowerCamelStr", "ToUpperUnderscore", "LOWER_CAMEL_STR"},

	{"UpperCamelStr", "ToLowerHyphen", "upper-camel-str"},
	{"UpperCamelStr", "ToLowerUnderscore", "upper_camel_str"},
	{"UpperCamelStr", "ToLowerCamel", "upperCamelStr"},
	{"UpperCamelStr", "ToUpperCamel", "UpperCamelStr"},
	{"UpperCamelStr", "ToUpperUnderscore", "UPPER_CAMEL_STR"},

	{"UPPER_UNDERSCORE_STR", "ToLowerHyphen", "upper-underscore-str"},
	{"UPPER_UNDERSCORE_STR", "ToLowerUnderscore", "upper_underscore_str"},
	{"UPPER_UNDERSCORE_STR", "ToLowerCamel", "upperUnderscoreStr"},
	{"UPPER_UNDERSCORE_STR", "ToUpperCamel", "UpperUnderscoreStr"},
	{"UPPER_UNDERSCORE_STR", "ToUpperUnderscore", "UPPER_UNDERSCORE_STR"},
}

func runTestCase(b *testing.B, tc testCase) bool {
	tcName := tc.name + "(" + tc.str + ")=" + tc.want
	f := fmtByName[tc.name]

	if got := f(tc.str); got != "" {
		Equalf(b, tc.want, got, tcName)
	}

	str := strings.Repeat(tc.str, n)
	return b.Run(tcName, func(b *testing.B) {
		f(str)
	})
}

func BenchmarkCaseFormat(b *testing.B) {
	// ToLowerHyphen not implemented

	fmtByName["ToLowerHyphen"] = notImplemented
	fmtByName["ToLowerUnderscore"] = caseformat.ToLowerUnderscore
	fmtByName["ToLowerCamel"] = caseformat.ToLowerCamel
	fmtByName["ToUpperCamel"] = caseformat.ToUpperCamel
	fmtByName["ToUpperUnderscore"] = caseformat.ToUpperUnderscore

	for _, tc := range tests {
		runTestCase(b, tc)
	}
}

func BenchmarkGoStrCase(b *testing.B) {
	// ToLowerHyphen not implemented

	fmtByName["ToLowerHyphen"] = notImplemented
	fmtByName["ToLowerUnderscore"] = gostrcase.SnakeCase
	fmtByName["ToLowerCamel"] = gostrcase.LowerCamelCase
	fmtByName["ToUpperCamel"] = gostrcase.UpperCamelCase
	fmtByName["ToUpperUnderscore"] = func(str string) string { return strings.ToUpper(gostrcase.SnakeCase(str)) }

	for _, tc := range tests {
		runTestCase(b, tc)
	}
}

func BenchmarkStrCase(b *testing.B) {
	// ToLowerHyphen not implemented
	// ToLowerCamel(UPPER_UNDERSCORE_STR)=upperUnderscoreStr produces uPPERUNDERSCORESTR
	// ToCamel(UPPER_UNDERSCORE_STR)=UpperUnderscoreStr produces UPPERUNDERSCORESTR

	fmtByName["ToLowerHyphen"] = notImplemented
	fmtByName["ToLowerUnderscore"] = strcase.ToSnake
	fmtByName["ToLowerCamel"] = strcase.ToLowerCamel
	fmtByName["ToUpperCamel"] = strcase.ToCamel
	fmtByName["ToUpperUnderscore"] = strcase.ToScreamingSnake

	for _, tc := range tests {
		runTestCase(b, tc)
	}
}

func BenchmarkStrUtil(b *testing.B) {
	// ToLowerHyphen not implemented
	// ToSnakeCase partly implemented (uppercase after whitespace)
	// ToCamelCase partly implemented (uppercase after whitespace)

	fmtByName["ToLowerHyphen"] = notImplemented
	fmtByName["ToLowerUnderscore"] = strutil.ToSnakeCase
	fmtByName["ToLowerCamel"] = strutil.ToCamelCase
	fmtByName["ToUpperCamel"] = func(str string) string { return xstrings.FirstRuneToUpper(strutil.ToCamelCase(str)) }
	fmtByName["ToUpperUnderscore"] = func(str string) string { return strings.ToUpper(strutil.ToSnakeCase(str)) }

	for _, tc := range tests {
		runTestCase(b, tc)
	}
}

func BenchmarkXStrings(b *testing.B) {
	// ToLowerHyphen not implemented
	// ToLowerCamel partly implemented (uppercase after underscore)

	fmtByName["ToLowerHyphen"] = notImplemented
	fmtByName["ToLowerUnderscore"] = xstrings.ToSnakeCase
	fmtByName["ToLowerCamel"] = xstrings.ToCamelCase
	fmtByName["ToUpperCamel"] = func(str string) string { return xstrings.FirstRuneToUpper(xstrings.ToCamelCase(str)) }
	fmtByName["ToUpperUnderscore"] = func(str string) string { return strings.ToUpper(xstrings.ToSnakeCase(str)) }

	for _, tc := range tests {
		runTestCase(b, tc)
	}
}

func BenchmarkGoavaCaseFormat(b *testing.B) {
	lHyp := LowerHyphen{}
	lUnd := LowerUnderscore{}
	lCml := LowerCamel{}
	uCml := UpperCamel{}
	uUnd := UpperUnderscore{}
	caseFormats := []CaseFormat{lHyp, lUnd, lCml, uCml, uUnd}

	for i, srcFmt := range caseFormats {
		fmtByName["ToLowerHyphen"] = conv(srcFmt, lHyp)
		fmtByName["ToLowerUnderscore"] = conv(srcFmt, lUnd)
		fmtByName["ToLowerCamel"] = conv(srcFmt, lCml)
		fmtByName["ToUpperCamel"] = conv(srcFmt, uCml)
		fmtByName["ToUpperUnderscore"] = conv(srcFmt, uUnd)

		for _, tc := range tests[i*len(caseFormats) : (i+1)*len(caseFormats)] {
			runTestCase(b, tc)
		}
	}
}

func conv(srcFmt, tgtFmt CaseFormat) func(string) string {
	return func(in string) string {
		return srcFmt.To(tgtFmt, in)
	}
}

func notImplemented(str string) string {
	return ""
}
