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

import "time"

// Ticker is a time source; returns a time value representing the number of nanoseconds elapsed since some fixed but
// arbitrary point in time.
// Note that most users should use Stopwatch instead of interacting with this class directly.
//
// Warning: this interface can only be used to measure elapsed time, not wall time.
type Ticker interface {
	// Read returns the number of nanoseconds elapsed since this ticker's fixed point of reference.
	Read() int64
}

func System() Ticker {
	return system{}
}

type system struct {
}

func (t system) Read() int64 {
	return time.Now().UnixNano()
}
