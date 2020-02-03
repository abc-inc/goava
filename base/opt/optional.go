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

package opt

import (
	"github.com/abc-inc/goava/base/precond"
	"github.com/abc-inc/goava/collect/set"
)

// Optional is an immutable object that may contain a non-nil reference to another object.
// Each instance of this type either contains a non-nil reference, or contains nothing (in which case we say that the
// reference is "absent"); it is never said to "contain nil".
//
// A non-nil Optional reference can be used as a replacement for a nillable reference.
// It allows you to represent "an object that must be present" and "an object that might be absent" as two distinct
// types in your program, which can aid clarity.
//
// # Some uses of this type include
//
// • As a method return type, as an alternative to returning nil to indicate that no value was available
//
// • To distinguish between "unknown" (for example, not present in a map) and "known to have no value" (present in the
// map, with value Absent())
//
// A common alternative to using this type is to find or create a suitable zero value for the type in question.
//
// This type is not intended as a direct analogue of any existing "option" or "maybe" construct from other programming
// environments, though it may bear some similarities.
type Optional interface {
	// IsPresent returns true if this holder contains a (non-nil) instance.
	IsPresent() bool

	// Get returns the contained instance, which must be present.
	//
	// If the instance might be absent, use Or(defValue) or OrNil() instead.
	Get() (interface{}, error)

	// Or returns the contained instance if it is present; defValue otherwise.
	//
	// If no default value should be required because the instance is known to be present, use Get() instead.
	// For a default value of nil, use OrNil().
	Or(defValue interface{}) (interface{}, error)

	// OrOpt returns this Optional if it has a value present; other otherwise.
	OrOpt(other Optional) (Optional, error)

	// OrGet returns the contained instance if it is present; the result of the provided function otherwise.
	OrGet(func() interface{}) (interface{}, error)

	// OrNil returns the contained instance if it is present; nil otherwise.
	//
	// If the instance is known to be present, use Get() instead.
	OrNil() interface{}

	// AsSet returns a Set whose only element is the contained instance if it is present; an empty Set otherwise.
	AsSet() set.Set

	// Transform applies the given function, if the instance is present; otherwise, absent is returned.
	Transform(func(interface{}) interface{}) (Optional, error)

	// String returns a string representation for this instance.
	String() string
}

var absentInstance = absent{}

// Absent returns an Optional instance with no contained reference.
func Absent() Optional {
	return absentInstance
}

// Of returns an Optional instance containing the given non-nil reference.
//
// To have nil treated as absent, use FromNillable() instead.
func Of(value interface{}) (Optional, error) {
	if _, err := precond.CheckNotNilf(value, "use Optional.FromNillable() instead of Optional.Of(nil)"); err != nil {
		return nil, err
	}
	return present{value}, nil
}

// FromNillable returns an Optional instance containing that reference, if it is non-nil.
// Otherwise, it returns an absent instance.
func FromNillable(value interface{}) Optional {
	if value == nil {
		return Absent()
	}
	return present{value}
}
