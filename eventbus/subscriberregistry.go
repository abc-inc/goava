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
	"sync"

	"github.com/abc-inc/goava/collect/set"
)

var methodsByType sync.Map

// SubscriberRegistry TODO
type SubscriberRegistry struct {
	subscribers sync.Map
}

// Register registers all subscriber methods on the given listener.
func (r *SubscriberRegistry) register(listener interface{}) {
	ls := findAllSubscribers(listener)
	for _, l := range ls {
		r.subscribersForType(l.kind()).Add(l)
	}
}

// unregister unregisters all subscribers on the given listener object.
func (r *SubscriberRegistry) unregister(listener interface{}) {
	ls := findAllSubscribers(listener)
	for _, l := range ls {
		r.subscribersForType(l.kind()).Remove(l)
	}
}

func findAllSubscribers(listener interface{}) []Subscriber {
	t := reflect.TypeOf(listener)
	subs, ok := methodsByType.Load(t)
	if ok {
		return subs.([]Subscriber)
	}

	ms := []Subscriber{}
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		modelType := reflect.TypeOf((*Event)(nil)).Elem()
		if method.Type.NumIn() == 2 && method.Type.In(1).Implements(modelType) {
			ms = append(ms, create(listener, method))
		}
	}
	methodsByType.Store(t, ms)
	return ms
}

func (r *SubscriberRegistry) subscribersForType(t reflect.Type) set.Set {
	var subs, _ = r.subscribers.Load(t)
	if subs == nil {
		subs = set.Empty()
		r.subscribers.Store(t, subs)
	}
	return subs.(set.Set)
}
