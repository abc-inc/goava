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
	"github.com/abc-inc/goava/base/stopwatch"
	"github.com/jonboulle/clockwork"
	. "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateStarted(t *testing.T) {
	s := stopwatch.CreateStarted()
	True(t, s.IsRunning())
}

func TestCreateUnstarted(t *testing.T) {
	s := stopwatch.CreateUnstarted()
	False(t, s.IsRunning())
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestInitialState(t *testing.T) {
	s := stopwatch.CreateUnstarted()
	False(t, s.IsRunning())
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestStart(t *testing.T) {
	s := stopwatch.CreateUnstarted()
	s2, err := s.Start()
	NoError(t, err)
	Same(t, s, s2)
	True(t, s.IsRunning())
}

func TestStartWhileRunning(t *testing.T) {
	s := stopwatch.CreateStarted()
	_, err := s.Start()
	EqualError(t, err, "this stopwatch is already running")
	True(t, s.IsRunning())
}

func TestStop(t *testing.T) {
	s := stopwatch.CreateStarted()
	s2, err := s.Stop()
	NoError(t, err)
	Same(t, s, s2)
	False(t, s.IsRunning())
}

func TestStopNew(t *testing.T) {
	s := stopwatch.CreateUnstarted()
	_, err := s.Stop()
	EqualError(t, err, "this stopwatch is already stopped")
	False(t, s.IsRunning())
}

func TestStopAlreadyStopped(t *testing.T) {
	s := stopwatch.CreateStarted()
	_, _ = s.Stop()

	_, err := s.Stop()
	EqualError(t, err, "this stopwatch is already stopped")
	False(t, s.IsRunning())
}

func TestResetNew(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateUnstartedClock(c)
	c.Advance(1)

	s.Reset()
	False(t, s.IsRunning())
	c.Advance(2)
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))

	_, _ = s.Start()
	c.Advance(3)
	Equal(t, 3*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestResetWhileRunning(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateUnstartedClock(c)
	c.Advance(1)

	_, _ = s.Start()
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))
	c.Advance(2)
	Equal(t, 2*time.Nanosecond, s.Elapsed(time.Nanosecond))

	s.Reset()
	False(t, s.IsRunning())
	c.Advance(3)
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))

}

func TestElapsedWhileRunning(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateUnstartedClock(c)
	c.Advance(78)

	_, _ = s.Start()
	Equal(t, 0*time.Nanosecond, s.Elapsed(time.Nanosecond))
	c.Advance(345)
	Equal(t, 345*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestElapsedNotRunning(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateUnstartedClock(c)
	c.Advance(1)

	_, _ = s.Start()
	c.Advance(4)

	_, _ = s.Stop()
	c.Advance(9)
	Equal(t, 4*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestElapsedMultipleSegments(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(c)
	c.Advance(9)
	_, _ = s.Stop()
	c.Advance(16)
	_, _ = s.Start()
	Equal(t, 9*time.Nanosecond, s.Elapsed(time.Nanosecond))
	c.Advance(25)
	Equal(t, 34*time.Nanosecond, s.Elapsed(time.Nanosecond))

	_, _ = s.Stop()
	c.Advance(36)
	Equal(t, 34*time.Nanosecond, s.Elapsed(time.Nanosecond))
}

func TestElapsedMicros(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(c)
	c.Advance(999)
	Equal(t, 0*time.Microsecond, s.Elapsed(time.Microsecond))
	c.Advance(1)
	Equal(t, 1*time.Microsecond, s.Elapsed(time.Microsecond))
}

func TestElapsedMillis(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(c)
	c.Advance(999999)
	Equal(t, 0*time.Millisecond, s.Elapsed(time.Millisecond))
	c.Advance(1)
	Equal(t, 1*time.Millisecond, s.Elapsed(time.Millisecond))
}

func TestElapsedTime(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(c)
	c.Advance(999999)
	Equal(t, int64(999999), s.ElapsedTime(time.Nanosecond))
	c.Advance(1)
	Equal(t, int64(1), s.ElapsedTime(time.Millisecond))
}

func TestString(t *testing.T) {
	c := clockwork.NewFakeClock()
	s := stopwatch.CreateStartedClock(c)

	_, _ = s.Start()
	Equal(t, "0.000 ns", s.String())
	c.Advance(1)
	Equal(t, "1.000 ns", s.String())
	c.Advance(998)
	Equal(t, "999.0 ns", s.String())
	c.Advance(1)
	Equal(t, "1.000 \u03bcs", s.String())
	c.Advance(1)
	Equal(t, "1.001 \u03bcs", s.String())
	c.Advance(8998)
	Equal(t, "9.999 \u03bcs", s.String())

	s.Reset()
	_, _ = s.Start()
	c.Advance(1234567)
	Equal(t, "1.235 ms", s.String())

	s.Reset()
	_, _ = s.Start()
	c.Advance(5000000000)
	Equal(t, "5.000 s", s.String())

	s.Reset()
	_, _ = s.Start()
	c.Advance(1.5 * 60 * 1000000000)
	Equal(t, "1.500 min", s.String())

	s.Reset()
	_, _ = s.Start()
	c.Advance(2.5 * 60 * 60 * 1000000000)
	Equal(t, "2.500 h", s.String())

	s.Reset()
	_, _ = s.Start()
	c.Advance(7.25 * 24 * 60 * 60 * 1000000000)
	Equal(t, "7.250 d", s.String())
}
