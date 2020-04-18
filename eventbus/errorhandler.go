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
	"fmt"
	"log"
	"reflect"
)

// ErrorHandler TODO
type ErrorHandler interface {
	Handle(err error, sub Subscriber)
}

// LoggingHandler TODO
type LoggingHandler struct {
}

// Handle TODO
func (h LoggingHandler) Handle(event Event, sub Subscriber, err error) {
	log.Println(message(event, sub), err)
}

func message(event Event, sub Subscriber) string {
	return "Exception thrown by subscriber method " + sub.method.Name +
		"(" + sub.method.Type.In(1).Name() + ")" +
		" on subscriber " + reflect.TypeOf(sub.listener).Name() +
		" when dispatching event: " + fmt.Sprintln(event.Event())
}
