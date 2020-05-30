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

package ticker_test

import (
	"fmt"
	"github.com/abc-inc/goava/base/ticker"
	"time"
)

func Example() {
	// tick := ticker.System()
	tick := ticker.NewFake()
	tick.AdvanceNanos(1577836801000000000) // 2020-01-01 00:00:00 GMT

	tick.SetAutoInc(100 * time.Nanosecond) // simulate reading time takes 100 ns

	// do something
	tick.Advance(1 * time.Second)

	fmt.Println(tick.Read())
	fmt.Println(tick.Read())
	// Output:
	// 1577836802000000000
	// 1577836802000000100
}
