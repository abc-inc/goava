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

package io

import "log"

// Flusher is a destination of data that can be flushed.
//
// The Flush() method is invoked to write any buffers.
type Flusher interface {
	Flush() error
}

// Flush a Flusher, with control over whether an error may be thrown.
//
// If swallow is true, then we don't rethrow an error, but merely log it.
func Flush(flusher Flusher, swallow bool) (err error) {
	if err = flusher.Flush(); err == nil || !swallow {
		return err
	}
	log.Println("error thrown while flushing Flusher.", err)
	return nil
}

// FlushQuietly is equivalent to calling Flush(flusher, true), but with no error in the signature.
func FlushQuietly(flusher Flusher) {
	if err := Flush(flusher, true); err != nil {
		log.Fatalln("error should not have been thrown.", err)
	}
}
