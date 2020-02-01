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

package compchain_test

import (
	"fmt"
	"sort"

	"github.com/abc-inc/goava/collect/compchain"
)

type user struct {
	ID   int
	Name string
	Mail string
}

func Example() {
	users := []user{
		{2, "John Doe", "john.william.doe@mail.com"},
		{1, "John Doe", "john.doe@mail.com"},
	}

	less := func(i, j int) bool {
		u1 := users[i]
		u2 := users[j]
		return compchain.Start().
			CompareString(u1.Name, u2.Name).
			CompareString(u1.Mail, u2.Mail).
			CompareInt(u1.ID, u2.ID).
			Result() < 0
	}

	fmt.Println(sort.SliceIsSorted(users, less))
	sort.Slice(users, less)
	fmt.Println(sort.SliceIsSorted(users, less))
	// Output:
	// false
	// true
}
