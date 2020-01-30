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

package runematcher_test

import (
	"fmt"

	"github.com/abc-inc/goava/base/runematcher"
)

func Example_anyOfCollapseFrom() {
	str := runematcher.AnyOf("eko").CollapseFrom("bookkeeper", '-')
	fmt.Println(str)
	// Output: b-p-r
}

func Example_anyOfTrimFrom() {
	str := runematcher.AnyOf("ab").TrimFrom("abacatbab")
	fmt.Println(str)
	// Output: cat
}

func Example_anyOfTrimLeadingFrom() {
	str := runematcher.AnyOf("ab").TrimLeadingFrom("abacatbab")
	fmt.Println(str)
	// Output: catbab
}

func Example_anyOfTrimTrailingFrom() {
	str := runematcher.AnyOf("ab").TrimTrailingFrom("abacatbab")
	fmt.Println(str)
	// Output: abacat
}

func Example_asciiMatchesAllOf() {
	valid := runematcher.ASCII().MatchesAllOf("abc")
	fmt.Println(valid)
	// Output: true
}

func Example_isRemoveFrom() {
	str := runematcher.Is('a').RemoveFrom("bazaar")
	fmt.Println(str)
	// Output: bzr
}

func Example_isReplaceFromRune() {
	str := runematcher.Is('a').ReplaceFromRune("radar", 'o')
	fmt.Println(str)
	// Output: rodor
}

func Example_isReplaceFrom() {
	str := runematcher.Is('a').ReplaceFrom("yaha", "oo")
	fmt.Println(str)
	// Output: yoohoo
}

func Example_isRetainFrom() {
	str := runematcher.Is('a').RetainFrom("bazaar")
	fmt.Println(str)
	// Output: aaa
}

func Example_whitespaceTrimFrom() {
	trimmed := runematcher.Whitespace().TrimFrom("    charming    ")
	fmt.Println(trimmed)
	// Output: charming
}
