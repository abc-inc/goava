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

package ticker_test

import (
	"sync"
	"testing"
	"time"

	"github.com/abc-inc/goava/base/ticker"
	. "github.com/stretchr/testify/require"
)

func TestAdvance(t *testing.T) {
	tick := ticker.NewFake()
	Equal(t, int64(0), tick.Read())
	Same(t, tick, tick.AdvanceNanos(10))
	Equal(t, int64(10), tick.Read())
	Same(t, tick, tick.Advance(1*time.Millisecond))
	Equal(t, int64(1000010), tick.Read())
}

func TestAutoIncrementStepReturnsSameInstance(t *testing.T) {
	tick := ticker.NewFake()
	Same(t, tick, tick.SetAutoInc(10*time.Nanosecond))
}

func TestAutoIncrementStepNanos(t *testing.T) {
	tick := ticker.NewFake().SetAutoInc(10 * time.Nanosecond)
	Equal(t, int64(0), tick.Read())
	Equal(t, int64(10), tick.Read())
	Equal(t, int64(20), tick.Read())
}

func TestAutoIncrementStepMillis(t *testing.T) {
	tick := ticker.NewFake().SetAutoInc(1 * time.Millisecond)
	Equal(t, int64(0), tick.Read())
	Equal(t, int64(1000000), tick.Read())
	Equal(t, int64(2000000), tick.Read())
}

func TestAutoIncrementStep_seconds(t *testing.T) {
	tick := ticker.NewFake().SetAutoInc(3 * time.Second)
	Equal(t, int64(0), tick.Read())
	Equal(t, int64(3000000000), tick.Read())
	Equal(t, int64(6000000000), tick.Read())
}

func TestAutoIncrementStepResetToZero(t *testing.T) {
	tick := ticker.NewFake().SetAutoInc(10 * time.Nanosecond)
	Equal(t, int64(0), tick.Read())
	Equal(t, int64(10), tick.Read())
	Equal(t, int64(20), tick.Read())
	tick.SetAutoInc(0)
	Equal(t, int64(30), tick.Read(), "Expected no auto-increment when setting autoInc to 0")
}

func TestAutoIncrementNegative(t *testing.T) {
	tick := ticker.NewFake()
	Panics(t, func() { tick.SetAutoInc(-1 * time.Nanosecond) })
}

func TestConcurrentAdvance(t *testing.T) {
	tick := ticker.NewFake()

	numberOfThreads := 64
	runConcurrentTest(
		numberOfThreads,
		func() {
			// adds two nanoseconds to the ticker
			tick.AdvanceNanos(1)
			time.Sleep(10 * time.Millisecond)
			tick.AdvanceNanos(1)
		})

	Equal(t, int64(numberOfThreads*2), tick.Read())
}

func TestConcurrentAutoIncrementStep(t *testing.T) {
	incrNanos := 3
	tick := ticker.NewFake().SetAutoInc(time.Duration(incrNanos) * time.Nanosecond)

	n := 64
	runConcurrentTest(n, func() { tick.Read() })

	Equal(t, int64(incrNanos*n), tick.Read())
}

// runConcurrentTest runs callable concurrently n times.
func runConcurrentTest(n int, callable func()) {
	var startWg, doneWg sync.WaitGroup
	startWg.Add(n)
	doneWg.Add(n)

	for i := n; i > 0; i-- {
		go func() {
			startWg.Done()
			startWg.Wait()
			callable()
			doneWg.Done()
		}()
	}
	doneWg.Wait()
}
