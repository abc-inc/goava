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

package hostandport_test

import (
	. "github.com/abc-inc/goava/net/hostandport"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestFromStringWellFormed(t *testing.T) {
	checkFromStringCase(t, "google.com", 80, "google.com", 80, false)
	checkFromStringCase(t, "google.com", 80, "google.com", 80, false)
	checkFromStringCase(t, "192.0.2.1", 82, "192.0.2.1", 82, false)
	checkFromStringCase(t, "[2001::1]", 84, "2001::1", 84, false)
	checkFromStringCase(t, "2001::3", 86, "2001::3", 86, false)
	checkFromStringCase(t, "host:", 80, "host", 80, false)
}

func TestFromStringBadDefaultPort(t *testing.T) {
	// Well-formed strings with bad default ports.
	checkFromStringCase(t, "gmail.com:81", -1, "gmail.com", 81, true)
	checkFromStringCase(t, "192.0.2.2:83", -1, "192.0.2.2", 83, true)
	checkFromStringCase(t, "[2001::2]:85", -1, "2001::2", 85, true)
	checkFromStringCase(t, "goo.gl:65535", 65536, "goo.gl", 65535, true)
	// No port, bad default.
	checkFromStringCase(t, "google.com", -1, "google.com", -1, false)
	checkFromStringCase(t, "192.0.2.1", 65536, "192.0.2.1", -1, false)
	checkFromStringCase(t, "[2001::1]", -1, "2001::1", -1, false)
	checkFromStringCase(t, "2001::3", 65536, "2001::3", -1, false)
}

func TestFromStringUnusedDefaultPort(t *testing.T) {
	// Default port, but unused.
	checkFromStringCase(t, "gmail.com:81", 77, "gmail.com", 81, true)
	checkFromStringCase(t, "192.0.2.2:83", 77, "192.0.2.2", 83, true)
	checkFromStringCase(t, "[2001::2]:85", 77, "2001::2", 85, true)
}

func TestFromStringBadPort(t *testing.T) {
	// Out-of-range ports.
	checkFromStringCase(t, "google.com:65536", 1, "", 99, false)
	checkFromStringCase(t, "google.com:9999999999", 1, "", 99, false)
	// Invalid port parts.
	checkFromStringCase(t, "google.com:port", 1, "", 99, false)
	checkFromStringCase(t, "google.com:-25", 1, "", 99, false)
	checkFromStringCase(t, "google.com:+25", 1, "", 99, false)
	checkFromStringCase(t, "google.com:25  ", 1, "", 99, false)
	checkFromStringCase(t, "google.com:25\t", 1, "", 99, false)
	checkFromStringCase(t, "google.com:0x25 ", 1, "", 99, false)
}

func TestFromStringUnparseableNonsense(t *testing.T) {
	// Some nonsense that causes parse failures.
	checkFromStringCase(t, "[goo.gl]", 1, "", 99, false)
	checkFromStringCase(t, "[goo.gl]:80", 1, "", 99, false)
	checkFromStringCase(t, "[", 1, "", 99, false)
	checkFromStringCase(t, "[]:", 1, "", 99, false)
	checkFromStringCase(t, "[]:80", 1, "", 99, false)
	checkFromStringCase(t, "[]bad", 1, "", 99, false)

	checkFromStringCase(t, "[1::1]:bad", 1, "", 99, false)
}

func TestFromStringParseableNonsense(t *testing.T) {
	// Examples of nonsense that gets through.
	checkFromStringCase(t, "[[:]]", 86, "[:]", 86, false)
	checkFromStringCase(t, "x:y:z", 87, "x:y:z", 87, false)
	checkFromStringCase(t, "", 88, "", 88, false)
	checkFromStringCase(t, ":", 99, "", 99, false)
	checkFromStringCase(t, ":123", -1, "", 123, true)
	checkFromStringCase(t, "\nOMG\t", 89, "\nOMG\t", 89, false)
}

func checkFromStringCase(t *testing.T, hpString string, defaultPort int, expectHost string, expectPort int, expectHasExplicitPort bool) {
	hp, err := FromString(hpString)
	if err != nil {
		Empty(t, expectHost)
		return
	}

	// Apply withDefaultPort(), yielding hp2.
	badDefaultPort := defaultPort < 0 || defaultPort > 65535
	hp2, err := hp.WithDefaultPort(defaultPort)
	Equal(t, badDefaultPort, err != nil)

	// Check the pre-withDefaultPort() instance.
	if expectHasExplicitPort {
		True(t, hp.HasPort())
		port, err := hp.Port()
		NoError(t, err)
		Equal(t, expectPort, port)
	} else {
		False(t, hp.HasPort())
		port, err := hp.Port()
		Error(t, err)
		Equal(t, -1, port)
	}
	Equal(t, expectHost, hp.Host())

	// Check the post-withDefaultPort() instance (if any).
	if !badDefaultPort {
		port, err := hp2.Port()
		if err != nil {
			// Make sure we expected this to fail.
			Equal(t, -1, expectPort)
		} else {
			NotEqual(t, -1, expectPort)
			Equal(t, expectPort, port)
		}
	}

	Equal(t, expectHost, hp2.Host())
}

func TestFromParts(t *testing.T) {
	hp, err := FromParts("gmail.com", 81)
	NoError(t, err)
	Equal(t, "gmail.com", hp.Host())
	True(t, hp.HasPort())

	port, err := hp.Port()
	NoError(t, err)
	Equal(t, 81, port)

	hp, err = FromParts("gmail.com:80", 81)
	EqualError(t, err, "Host has a port: gmail.com:80")

	hp, err = FromParts("gmail.com", -1)
	EqualError(t, err, "Port out of range: -1")

	hp, err = FromParts("gmail.com:unknown", 80)
	EqualError(t, err, "Unparseable port number: gmail.com:unknown")
}

func TestFromHost(t *testing.T) {
	hp, err := FromHost("gmail.com")
	NoError(t, err)
	Equal(t, "gmail.com", hp.Host())
	False(t, hp.HasPort())

	hp, err = FromHost("[::1]")
	NoError(t, err)
	Equal(t, "::1", hp.Host())
	False(t, hp.HasPort())

	hp, err = FromHost("gmail.com:80")
	Error(t, err)

	hp, err = FromHost("[gmail.com]")
	Error(t, err)
}

func TestGetPortOrDefault(t *testing.T) {
	hp, _ := FromString("host:80")
	Equal(t, 80, hp.GetPortOrDefault(123))
	hp, _ = FromString("host")
	Equal(t, 123, hp.GetPortOrDefault(123))
}

func TestRequireBracketsForIPv6(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		// Bracketed IPv6 works fine.
		{"[::1]", "::1"},
		{"[::1]:80", "::1"},
		// Non-bracketed non-IPv6 works fine.
		{"x", "x"},
		{"x:80", "x"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			hp, err := FromString(test.str)
			NoError(t, err)

			hp, err = hp.RequireBracketsForIPv6()
			NoError(t, err)
			Equal(t, test.want, hp.Host())
		})
	}

	// Non-bracketed IPv6 fails.
	hp, err := FromString("::1")
	NoError(t, err)
	hp, err = hp.RequireBracketsForIPv6()
	Error(t, err)
}

func TestString(t *testing.T) {
	tests := []struct {
		str  string
		port int
		want string
	}{
		// With ports.
		{"foo:101", NoPort, "foo:101"},
		{":102", NoPort, ":102"},
		{"1::2", 103, "[1::2]:103"},
		{"[::1]:104", NoPort, "[::1]:104"},

		// Without ports.
		{"foo", NoPort, "foo"},
		{"", NoPort, ""},
		{"1::2", NoPort, "[1::2]"},
		{"[::1]", NoPort, "[::1]"},

		// Garbage in, garbage out.
		{"::]", 107, "[::]]:107"},
		{"[[:]]:108", NoPort, "[[:]]:108"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			var hp HostAndPort
			if test.port > NoPort {
				hp, _ = FromParts(test.str, test.port)
			} else {
				hp, _ = FromString(test.str)
			}
			Equal(t, test.want, hp.String())
		})
	}
}
