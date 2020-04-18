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

package opt

import (
	"errors"
	"fmt"
	"github.com/abc-inc/goava/base/precond"
	"github.com/abc-inc/goava/collect/set"
)

// present is an implementation of an Optional containing a value.
type present struct {
	value interface{}
}

// IsPresent returns true if this holder contains a (non-nil) instance.
func (o present) IsPresent() bool {
	return true
}

// Get returns the contained instance, which must be present.
//
// If the instance might be absent, use Or(defValue) or OrNil() instead.
func (o present) Get() (interface{}, error) {
	return o.value, nil
}

// Or returns the contained instance if it is present; defValue otherwise.
//
// If no default value should be required because the instance is known to be present, use Get() instead.
// For a default value of nil, use OrNil().
func (o present) Or(defValue interface{}) (interface{}, error) {
	if _, err := precond.CheckNotNilf(defValue, "use Optional.OrNil() instead of Optional.Or(nil)"); err != nil {
		return nil, err
	}
	return o.value, nil
}

// OrOpt returns this Optional if it has a value present; other otherwise.
func (o present) OrOpt(other Optional) (Optional, error) {
	if _, err := precond.CheckNotNilf(other, "use Optional.OrNil() instead of a nil Optional"); err != nil {
		return nil, err
	}
	return o, nil
}

// OrGet returns the contained instance if it is present; the result of the provided function otherwise.
func (o present) OrGet(supplier func() interface{}) (interface{}, error) {
	if supplier == nil {
		return nil, errors.New("the function passed to Optional.OrGet() must not be nil")
	}

	return o.value, nil
}

// OrNil returns the contained instance if it is present; nil otherwise.
//
// If the instance is known to be present, use Get() instead.
func (o present) OrNil() interface{} {
	return o.value
}

// AsSet returns a Set whose only element is the contained instance if it is present; an empty Set otherwise.
func (o present) AsSet() set.Set {
	return set.Singleton(o.value)
}

// Transform applies the given function, if the instance is present; otherwise, absent is returned.
func (o present) Transform(f func(interface{}) interface{}) (Optional, error) {
	if f == nil {
		return nil, errors.New("the function passed to Optional.Transform() must not be nil")
	}

	newVal, err := precond.CheckNotNilf(f(o.value), "the function passed to Optional.Transform() must not return nil")
	if err != nil {
		return nil, err
	}
	return present{newVal}, nil
}

// String returns a string representation for this instance.
func (o present) String() string {
	return fmt.Sprintf("Optional.of(%v)", o.value)
}
