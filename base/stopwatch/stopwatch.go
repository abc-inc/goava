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
	"github.com/abc-inc/goava/base/precond"
	"github.com/jonboulle/clockwork"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
	"time"
)

// Stopwatch measures elapsed time in nanoseconds.
//
// It is useful to measure elapsed time using this type instead of direct calls to time.Now() for a few reasons:
//
// • An alternate time source can be substituted, for testing or performance reasons.
//
// • As documented by the time package, the value returned by time.Now() has no absolute meaning, and can only be
// interpreted as relative to another timestamp returned by time.Now() at a different time.
// Stopwatch is a more effective abstraction because it exposes only these relative values, not the absolute ones.
//
// Stopwatch methods are not idempotent; it is an error to start or stop a stopwatch that is already in the desired
// state.
//
// When testing code that uses Stopwatch, use CreateUnstarted(Clock) or CreateStarted(Clock) to supply a fake or mock
// clock. This allows you to simulate any valid behavior of the Stopwatch.
//
// Note: This struct is not thread-safe.
type Stopwatch struct {
	clock        clockwork.Clock
	isRunning    bool
	elapsedNanos time.Duration
	startTick    time.Time
}

// CreateUnstarted creates (but does not start) a new stopwatch using the wall clock as its time source.
func CreateUnstarted() *Stopwatch {
	return CreateUnstartedClock(clockwork.NewRealClock())
}

// CreateUnstartedClock creates (but does not start) a new stopwatch, using the specified time source.
func CreateUnstartedClock(clock clockwork.Clock) *Stopwatch {
	return &Stopwatch{clock, false, 0, time.Now()}
}

// CreateStarted creates (and starts) a new stopwatch using the wall clock as its time source.
func CreateStarted() *Stopwatch {
	s, _ := CreateUnstarted().Start()
	return s
}

// CreateStartedClock creates (and starts) a new stopwatch, using the specified time source.
func CreateStartedClock(clock clockwork.Clock) *Stopwatch {
	s, _ := CreateUnstartedClock(clock).Start()
	return s
}

// IsRunning returns true if Start has been called on this stopwatch, and Stop has not been called since the last call
// to Start.
func (s Stopwatch) IsRunning() bool {
	return s.isRunning
}

// Start starts the stopwatch.
func (s *Stopwatch) Start() (*Stopwatch, error) {
	if err := precond.CheckStatef(!s.isRunning, "this stopwatch is already running"); err != nil {
		return nil, err
	}

	s.isRunning = true
	s.startTick = s.clock.Now()
	return s, nil
}

// Stop stops the stopwatch.
// Future reads will return the fixed duration that had elapsed up to this point.
func (s *Stopwatch) Stop() (*Stopwatch, error) {
	if err := precond.CheckStatef(s.isRunning, "this stopwatch is already stopped"); err != nil {
		return nil, err
	}

	tick := s.clock.Now()
	s.isRunning = false
	s.elapsedNanos += tick.Sub(s.startTick)
	return s, nil
}

// Reset sets the elapsed time for this stopwatch to zero, and places it in a stopped state.
func (s *Stopwatch) Reset() *Stopwatch {
	s.elapsedNanos = 0
	s.isRunning = false
	return s
}

// calcElapsedNanos calculates the number of nanoseconds this stopwatch was in started state without being reset.
func (s Stopwatch) calcElapsedNanos() time.Duration {
	if s.isRunning {
		t := s.clock.Now()
		return t.Sub(s.startTick) + s.elapsedNanos
	}
	return s.elapsedNanos
}

// Elapsed returns the current elapsed time shown on this stopwatch as a Duration.
func (s Stopwatch) Elapsed(unit time.Duration) time.Duration {
	return s.calcElapsedNanos().Truncate(unit)
}

// ElapsedTime returns the current elapsed time shown on this stopwatch, expressed in the desired time unit, with any
// fraction rounded down.
//
// Note: the overhead of measurement can be more than a microsecond, so it is generally not useful to specify
// time.Nanosecond precision here.
//
// It is generally not a good idea to use an ambiguous, unitless int64 to represent elapsed time.
// Therefore, we recommend using Elapsed(Duration) instead, which returns a strongly-typed Duration instance.
func (s Stopwatch) ElapsedTime(unit time.Duration) int64 {
	return int64(s.calcElapsedNanos() / unit)
}

// String returns a string representation of the current elapsed time.
func (s Stopwatch) String() string {
	nanos := s.calcElapsedNanos()
	unit := chooseUnit(nanos)
	p := message.NewPrinter(language.English)
	v := p.Sprintf("%4.3f", float64(nanos)/float64(unit))
	return strings.TrimSuffix(v[0:5], ".") + " " + abbreviate(unit)
}

// chooseUnit returns a human-readable duration of the Duration.
func chooseUnit(nanos time.Duration) time.Duration {
	if nanos.Truncate(24*time.Hour) > 0 {
		return 24 * time.Hour
	}
	if nanos.Truncate(time.Hour) > 0 {
		return time.Hour
	}
	if nanos.Truncate(time.Minute) > 0 {
		return time.Minute
	}
	if nanos.Truncate(time.Second) > 0 {
		return time.Second
	}
	if nanos.Truncate(time.Millisecond) > 0 {
		return time.Millisecond
	}
	if nanos.Truncate(time.Microsecond) > 0 {
		return time.Microsecond
	}
	return time.Nanosecond
}

// abbreviate returns a string representation of the time unit.
func abbreviate(unit time.Duration) string {
	switch unit {
	case time.Nanosecond:
		return "ns"
	case time.Microsecond:
		return "\u03bcs" // μs
	case time.Millisecond:
		return "ms"
	case time.Second:
		return "s"
	case time.Minute:
		return "min"
	case time.Hour:
		return "h"
	default:
		return "d"
	}
}
