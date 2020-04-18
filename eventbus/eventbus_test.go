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

package eventbus_test

import (
	"testing"

	"github.com/abc-inc/goava/eventbus"
	"github.com/stretchr/testify/require"
)

type stringCatcher struct {
	events []string
}

func (s *stringCatcher) HereHaveAString(e eventbus.SimpleEvent) {
	s.events = append(s.events, e.Event().(string))
}

func (s stringCatcher) MethodWithoutAnnotation(str string) {
	panic("Event bus must not call methods without @Subscribe!")
}

func TestEventBus_BasicCatcherDistribution(t *testing.T) {
	b := eventbus.EventBus{}

	catcher := &stringCatcher{}
	b.Register(catcher)
	b.Post("nothing")

	require.Equal(t, 1, len(catcher.events))
	require.Equal(t, "nothing", catcher.events[0])

	_ = b.Unregister(catcher)
	b.Post("again")
	require.Equal(t, 1, len(catcher.events))
}
