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

package precond

import (
	"errors"
	"fmt"
	"strconv"
)

// IllegalArgumentError indicates that a function has been passed an illegal or inappropriate argument.
type IllegalArgumentError struct {
	msg string
}

func (e *IllegalArgumentError) Error() string {
	return e.msg
}

// IllegalStateError signals that a function has been invoked at an illegal or inappropriate time.
type IllegalStateError struct {
	msg string
}

func (e *IllegalStateError) Error() string {
	return e.msg
}

// IndexOutOfBoundsError indicates that an index of some sort (such as to an array or to a string) is out of range.
type IndexOutOfBoundsError struct {
	index  int
	size   int
	desc   string
	args   []interface{}
	msgFun func(int, int, string, ...interface{}) string
}

func (e *IndexOutOfBoundsError) Error() string {
	return e.msgFun(e.index, e.size, e.desc, e.args...)
}

// NilError indicates when an application attempts to use nil in a case where an object is required.
type NilError struct {
	msg string
}

func (e *NilError) Error() string {
	return e.msg
}

// CheckArgument ensures the truth of an expression involving one or more parameters to the calling method.
func CheckArgument(expr bool) error {
	return CheckArgumentf(expr, "invalid argument")
}

// CheckArgumentf ensures the truth of an expression involving one or more parameters to the calling method.
func CheckArgumentf(expr bool, desc string, args ...interface{}) error {
	if !expr {
		return &IllegalArgumentError{fmt.Sprintf(desc, args...)}
	}
	return nil
}

// CheckState ensures the truth of an expression involving the state of the calling instance, but not involving any
// parameters to the calling method.
func CheckState(expr bool) error {
	return CheckStatef(expr, "illegal state")
}

// CheckStatef ensures the truth of an expression involving the state of the calling instance, but not involving any
// parameters to the calling method.
func CheckStatef(expr bool, desc string, args ...interface{}) error {
	if !expr {
		return &IllegalStateError{fmt.Sprintf(desc, args...)}
	}
	return nil
}

// CheckNotNil ensures that an object passed as a parameter to the calling method is not nil.
func CheckNotNil(obj interface{}) (interface{}, error) {
	return CheckNotNilf(obj, "")
}

// CheckNotNilf ensures that an object passed as a parameter to the calling method is not nil.
func CheckNotNilf(obj interface{}, desc string, args ...interface{}) (interface{}, error) {
	if obj == nil {
		return nil, &NilError{fmt.Sprintf(desc, args...)}
	}
	return obj, nil
}

// CheckElementIndex ensures that index specifies a valid element in an array or string of the given size.
// An element index may range from zero, inclusive, to size, exclusive.
func CheckElementIndex(index, size int) (int, error) {
	return CheckElementIndexf(index, size, "index")
}

// CheckElementIndexf ensures that index specifies a valid element in an array or string of the given size.
// An element index may range from zero, inclusive, to size, exclusive.
func CheckElementIndexf(index, size int, desc string, args ...interface{}) (int, error) {
	if index < 0 || index >= size {
		return -1, &IndexOutOfBoundsError{index, size, desc, args, badElementIndex}
	}
	return index, nil
}

// CheckPositionIndex ensures that index specifies a valid position in an array, list or string of the given size.
// A position index may range from zero to size, inclusive.
func CheckPositionIndex(index, size int) (int, error) {
	return CheckPositionIndexf(index, size, "index")
}

// CheckPositionIndexf ensures that index specifies a valid position in an array, list or string of the given size.
// A position index may range from zero to size, inclusive.
func CheckPositionIndexf(index, size int, desc string, args ...interface{}) (int, error) {
	if index < 0 || index > size {
		return -1, &IndexOutOfBoundsError{index, size, desc, args, badPositionIndex}
	}
	return index, nil
}

func badElementIndex(index, size int, desc string, args ...interface{}) string {
	if args != nil {
		return fmt.Sprintf(desc, args...)
	} else if index < 0 {
		return fmt.Sprintf("%s (%d) must not be negative", desc, index)
	} else if size < 0 {
		return "negative size: " + strconv.Itoa(size)
	} else { // index >= size
		return fmt.Sprintf("%s (%d) must be less than size (%d)", desc, index, size)
	}
}

func badPositionIndex(index, size int, desc string, args ...interface{}) string {
	if args != nil {
		return fmt.Sprintf(desc, args...)
	} else if index < 0 {
		return fmt.Sprintf("%s (%d) must not be negative", desc, index)
	} else if size < 0 {
		return fmt.Sprintf("negative size: " + strconv.Itoa(size))
	} else { // index > size
		return fmt.Sprintf("%s (%d) must not be greater than size (%d)", desc, index, size)
	}
}

func CheckNonnegative(value int, name string) (int, error) {
	if value < 0 {
		return 0, errors.New(name + " cannot be negative but was: " + strconv.Itoa(value))
	}
	return value, nil
}

func CheckNonnegative64(value int64, name string) (int64, error) {
	if value < 0 {
		return 0, errors.New(name + " cannot be negative but was: " + strconv.FormatInt(value, 10))
	}
	return value, nil
}
