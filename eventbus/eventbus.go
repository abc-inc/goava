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

import (
	"reflect"

	"github.com/abc-inc/goava/collect/set"
)

// EventBus ...
type EventBus struct {
	SubscriberRegistry
}

// Register registers all subscriber methods on interface to receive events.
func (e *EventBus) Register(object interface{}) {
	register := e.register
	register(object)
}

// Unregister unregisters all subscriber methods on a registered interface.
func (e *EventBus) Unregister(object interface{}) error {
	e.unregister(object)
	return nil
}

// Post posts an event to all registered subscribers.
//
// This method will return successfully after the event has been posted to all subscribers, and regardless of any
// exceptions thrown by subscribers.
//
// If no subscribers have been subscribed for event's type, and event is not already a DeadEvent, it will be wrapped in
// a DeadEvent and reposted.
func (e *EventBus) Post(event interface{}) {
	if _, ok := event.(Event); !ok {
		event = SimpleEvent{event, nil}
	}
	eventSubscribers := e.getSubscribers(event).ToArray()
	for _, s := range eventSubscribers {
		s.(Subscriber).DispatchEvent(event.(Event))
	}
	/*
		if eventSubscribers.hasNext() {
			e.dispatcher.dispatch(event, eventSubscribers)
		} else if _, ok := event.(DeadEvent); !ok {
			// the event had no subscribers and was not itself a DeadEvent
			e.post(DeadEvent{e, event})
		}
	*/
}

func (e *EventBus) getSubscribers(event interface{}) set.Set {
	return e.SubscriberRegistry.subscribersForType(reflect.TypeOf(event))
}
