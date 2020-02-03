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

package stopwatch

import (
	"fmt"
	"github.com/abc-inc/goava/base/stopwatch"
	"testing"
	"time"
)

func BenchmarkStopwatch(b *testing.B) {
	total := int64(0)
	s := stopwatch.CreateStarted()
	for i := 0; i < b.N; i++ {
		_, _ = s.Reset().Start()
		// here is where you would do something
		total += s.ElapsedTime(time.Nanosecond)
	}
	fmt.Println("total:", total)
}

func BenchmarkManual(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		start := time.Now()
		// here is where you would do something
		total += time.Now().Nanosecond() - start.Nanosecond()
	}
	fmt.Println("total:", total)
}
