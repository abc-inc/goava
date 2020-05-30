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

package ticker

import (
	"github.com/abc-inc/goava/base/precond"
	"sync/atomic"
	"time"
)

// Fake is a Ticker whose value can be advanced programmatically in test.
//
// The ticker can be configured so that the time is incremented whenever Read() is called (see SetAutoInc()).
type Fake struct {
	nanos        *int64
	autoIncNanos *int64
}

func NewFake() *Fake {
	n, i := int64(0), int64(0)
	return &Fake{&n, &i}
}

// AdvanceNanos advances the ticker value by nanoseconds.
func (t *Fake) AdvanceNanos(nanoseconds int64) *Fake {
	atomic.AddInt64(t.nanos, nanoseconds)
	return t
}

// Advance advances the ticker value by duration.
func (t *Fake) Advance(duration time.Duration) *Fake {
	return t.AdvanceNanos(duration.Nanoseconds())
}

// Sets the increment applied to the ticker whenever it is queried.
//
// The default behavior is to auto increment by zero. i.e: The ticker is left unchanged when queried.
func (t *Fake) SetAutoInc(autoInc time.Duration) *Fake {
	if err := precond.CheckArgumentf(autoInc >= 0, "May not auto-increment by a negative amount"); err != nil {
		panic(err)
	}
	atomic.CompareAndSwapInt64(t.autoIncNanos, *t.autoIncNanos, autoInc.Nanoseconds())
	return t
}

func (t *Fake) Read() int64 {
	n := *t.nanos
	atomic.AddInt64(t.nanos, *t.autoIncNanos)
	return n
}
