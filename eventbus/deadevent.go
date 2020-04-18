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

package eventbus

import "fmt"

// DeadEvent wraps an event that was posted, but which had no subscribers and thus could not be delivered.
//
// Registering a DeadEvent subscriber is useful for debugging or logging, as it can detect misconfigurations in a
// system's event distribution.
type DeadEvent struct {
	event  interface{}
	source interface{}
}

// Event returns the wrapped, 'dead' event, which the system was unable to deliver to any registered subscriber.
func (e DeadEvent) Event() interface{} {
	return e.event
}

// Source returns the object that originated this event (not the object that originated the wrapped event).
// This is generally an EventBus.
func (e DeadEvent) Source() interface{} {
	return e.source
}

func (e DeadEvent) String() string {
	return fmt.Sprintf("DeadEvent{source=%s, event=%s)", e.source, e.event)
}
