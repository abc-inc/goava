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

package opt_test

import (
	"fmt"
	"github.com/abc-inc/goava/base/opt"
	"github.com/abc-inc/goava/collect/set"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestAbsent(t *testing.T) {
	o := opt.Absent()
	False(t, o.IsPresent())
}

func TestOf(t *testing.T) {
	o, err := opt.Of("training")
	NoError(t, err)

	v, err := o.Get()
	NoError(t, err)

	Equal(t, "training", v)
}

func TestOf_nil(t *testing.T) {
	_, err := opt.Of(nil)
	EqualError(t, err, "use Optional.FromNillable() instead of Optional.Of(nil)")
}

func TestFromNillable(t *testing.T) {
	v, err := opt.FromNillable("bob").Get()
	NoError(t, err)
	Equal(t, "bob", v)
}

func TestFromNillable_nil(t *testing.T) {
	Equal(t, opt.Absent(), opt.FromNillable(nil))
}

func TestAbsent_IsPresent(t *testing.T) {
	False(t, opt.Absent().IsPresent())
}

func TestPresent_IsPresent(t *testing.T) {
	o, _ := opt.Of("training")
	True(t, o.IsPresent())
}

func TestAbsent_Get(t *testing.T) {
	_, err := opt.Absent().Get()
	EqualError(t, err, "Optional.Get() cannot be called on an absent value")
}

func TestPresent_Get(t *testing.T) {
	o, _ := opt.Of("training")
	v, err := o.Get()
	NoError(t, err)
	Equal(t, "training", v)
}

func TestAbsent_Or(t *testing.T) {
	v, err := opt.Absent().Or("default")
	NoError(t, err)
	Equal(t, "default", v)
}

func TestPresent_Or(t *testing.T) {
	o, _ := opt.Of("a")
	v, err := o.Or("default")
	NoError(t, err)
	Equal(t, "a", v)
}

func TestAbsent_Or_Nil(t *testing.T) {
	_, err := opt.Absent().Or(nil)
	EqualError(t, err, "use Optional.OrNil() instead of Optional.Or(nil)")
}

func TestPresent_Or_Nil(t *testing.T) {
	o, _ := opt.Of("a")
	_, err := o.Or(nil)
	EqualError(t, err, "use Optional.OrNil() instead of Optional.Or(nil)")
}

func TestAbsent_OrGet(t *testing.T) {
	v, err := opt.Absent().OrGet(func() interface{} { return "fallback" })
	NoError(t, err)
	Equal(t, "fallback", v)
}

func TestPresent_OrGet(t *testing.T) {
	o, _ := opt.Of("a")
	v, err := o.Or(func() interface{} { return "fallback" })
	NoError(t, err)
	Equal(t, "a", v)
}

func TestAbsent_OrGet_FunctionReturnsNil(t *testing.T) {
	f := func() interface{} { return nil }
	_, err := opt.Absent().OrGet(f)
	EqualError(t, err, "use Optional.OrNil() instead of a function that returns nil")
}

func TestPresent_OrGet_FunctionReturnsNil(t *testing.T) {
	f := func() interface{} { return nil }
	o, _ := opt.Of("a")
	v, err := o.OrGet(f)
	NoError(t, err)
	Equal(t, "a", v)
}

func TestAbsent_OrGet_Nil(t *testing.T) {
	_, err := opt.Absent().OrGet(nil)
	EqualError(t, err, "the function passed to Optional.OrGet() must not be nil")
}

func TestPresent_OrGet_Nil(t *testing.T) {
	o, _ := opt.Of("a")
	_, err := o.OrGet(nil)
	EqualError(t, err, "the function passed to Optional.OrGet() must not be nil")
}

func TestAbsent_OrOpt(t *testing.T) {
	want, _ := opt.Of("fallback")
	a := opt.Absent()
	f, _ := opt.Of("fallback")
	o, _ := a.OrOpt(f)
	Equal(t, want, o)
}

func TestPresent_OrOpt(t *testing.T) {
	want, _ := opt.Of("a")
	a, _ := opt.Of("a")
	f, _ := opt.Of("fallback")
	o, _ := a.OrOpt(f)
	Equal(t, want, o)
}

func TestAbsent_OrOpt_Nil(t *testing.T) {
	_, err := opt.Absent().OrOpt(nil)
	EqualError(t, err, "use Optional.OrNil() instead of a nil Optional")
}

func TestPresent_OrOpt_Nil(t *testing.T) {
	a, _ := opt.Of("a")
	_, err := a.OrOpt(nil)
	EqualError(t, err, "use Optional.OrNil() instead of a nil Optional")
}

func TestAbsent_OrNil(t *testing.T) {
	Nil(t, opt.Absent().OrNil())
}

func TestPresent_OrNil(t *testing.T) {
	o, _ := opt.Of("a")
	Equal(t, "a", o.OrNil())
}

func TestAbsent_AsSet(t *testing.T) {
	Equal(t, set.Empty(), opt.Absent().AsSet())
}

func TestPresent_AsSet(t *testing.T) {
	o, _ := opt.Of("a")
	Equal(t, set.Singleton("a"), o.AsSet())
}

func TestAbsent_Transform(t *testing.T) {
	o, err := opt.Absent().Transform(func(v interface{}) interface{} { return v })
	NoError(t, err)
	Equal(t, opt.Absent(), o)
}

func TestPresent_Transform_Identity(t *testing.T) {
	want, _ := opt.Of("a")
	a, _ := opt.Of("a")
	o, err := a.Transform(func(v interface{}) interface{} { return v })
	NoError(t, err)
	Equal(t, want, o)
}

func TestPresent_Transform_String(t *testing.T) {
	want, _ := opt.Of("a")
	a, _ := opt.Of("a")
	o, err := a.Transform(func(v interface{}) interface{} { return fmt.Sprint(v) })
	NoError(t, err)
	Equal(t, want, o)
}

func TestAbsent_Transform_FunctionReturnsNil(t *testing.T) {
	o, err := opt.Absent().Transform(func(v interface{}) interface{} { return nil })
	NoError(t, err)
	Equal(t, opt.Absent(), o)
}

func TestPresent_Transform_FunctionReturnsNil(t *testing.T) {
	a, _ := opt.Of("a")
	_, err := a.Transform(func(v interface{}) interface{} { return nil })
	EqualError(t, err, "the function passed to Optional.Transform() must not return nil")
}

func TestAbsent_Transform_Nil(t *testing.T) {
	_, err := opt.Absent().Transform(nil)
	EqualError(t, err, "the function passed to Optional.Transform() must not be nil")
}

func TestPresent_Transform_Nil(t *testing.T) {
	a, _ := opt.Of("a")
	_, err := a.Transform(nil)
	EqualError(t, err, "the function passed to Optional.Transform() must not be nil")
}

func TestAbsent_String(t *testing.T) {
	Equal(t, "Optional.Absent()", opt.Absent().String())
}

func TestPresent_String(t *testing.T) {
	o, _ := opt.Of("training")
	Equal(t, "Optional.of(training)", o.String())
}
