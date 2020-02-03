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

package stopwatch_test

import (
	"fmt"
	"time"

	"github.com/abc-inc/goava/base/stopwatch"
	"github.com/jonboulle/clockwork"
)

func Example() {
	clock := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(clock)

	// doSomething()
	clock.Advance(12300 * time.Microsecond) // for the sake of the example, we turn the watch hand

	s.Stop() // optional
	duration := s.ElapsedTime(time.Microsecond)
	fmt.Println(duration)
	fmt.Println("time:", s.String())
	// Output:
	// 12300
	// time: 12.30 ms
}
