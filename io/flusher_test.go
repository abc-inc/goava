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

package io

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var flusher funcFlusher

type funcFlusher struct {
	f func() error
}

func (m funcFlusher) Flush() error {
	return m.f()
}

func TestFlush_clean(t *testing.T) {
	// make sure that no exception is thrown regardless of value of 'swallow' when the mock does not throw an exception.
	setupFlusher(false)
	doFlush(t, flusher, false, false)

	setupFlusher(false)
	doFlush(t, flusher, true, false)
}

func TestFlush_flusherWithEatenException(t *testing.T) {
	// make sure that no exception is thrown if 'swallow' is true when the mock does throw an exception.
	setupFlusher(true)
	doFlush(t, flusher, true, false)
}

func TestFlush_flusherWithThrownException(t *testing.T) {
	// make sure that the exception is thrown if 'swallow' is false when the mock does throw an exception.
	setupFlusher(true)
	doFlush(t, flusher, false, true)
}

func TestFlushQuietly_flusherWithEatenException(t *testing.T) {
	// make sure that no exception is thrown by FlushQuietly() when the mock does throw an exception.
	setupFlusher(true)
	FlushQuietly(flusher)
}

// setupFlusher sets up a Flusher to expect to be flushed, and optionally to throw an exception.
func setupFlusher(shouldThrowOnFlush bool) {
	flusher = funcFlusher{func() error {
		if shouldThrowOnFlush {
			return errors.New("this should only appear in the logs - it should not be rethrown")
		}
		return nil
	}}
}

// doFlush flushes the Flusher, passing in the swallow parameter.
// expectThrown determines whether we expect an exception to be thrown by Flush().
func doFlush(t *testing.T, flusher Flusher, swallow, expectThrown bool) {
	err := Flush(flusher, swallow)
	if err == nil && expectThrown {
		require.Fail(t, "Didn't throw exception.")
	}
	if err != nil && !expectThrown {
		require.Fail(t, "Threw exception")
	}
}
