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

package hostandport

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/abc-inc/goava/base/precond"
)

// NoPort indicates the absence of a port number.
const NoPort = -1

// HostAndPort is an immutable representation of a host and port.
type HostAndPort struct {
	// host represents a Hostname, IPv4/IPv6 literal, or unvalidated nonsense.
	host string

	// port is a validated port number in the range [0..65535], or NoPort
	port int

	// hasBracketlessColons is true if the parsed host has colons, but no surrounding brackets.
	hasBracketlessColons bool
}

// Host returns the portion of this HostAndPort instance that should represent the hostname or IPv4/IPv6 literal.
//
// A successful parse does not imply any degree of sanity in this field.
func (hp HostAndPort) Host() string {
	return hp.host
}

// HasPort returns true if this instance has a defined port.
func (hp HostAndPort) HasPort() bool {
	return hp.port > NoPort
}

// Port gets the current port number, failing if no port is defined.
func (hp HostAndPort) Port() (int, error) {
	if err := precond.CheckState(hp.HasPort()); err != nil {
		return NoPort, err
	}
	return hp.port, nil
}

// GetPortOrDefault returns the current port number, with a default if no port is defined.
func (hp HostAndPort) GetPortOrDefault(defaultPort int) int {
	if hp.HasPort() {
		return hp.port
	}
	return defaultPort
}

// String rebuilds the host:port string, including brackets if necessary.
func (hp HostAndPort) String() string {
	builder := strings.Builder{}
	builder.Grow(len(hp.host) + 8)

	if strings.ContainsRune(hp.host, ':') {
		builder.WriteByte('[')
		builder.WriteString(hp.host)
		builder.WriteByte(']')
	} else {
		builder.WriteString(hp.host)
	}
	if hp.HasPort() {
		builder.WriteByte(':')
		builder.WriteString(strconv.Itoa(hp.port))
	}
	return builder.String()
}

// FromParts builds a HostAndPort instance from separate host and port values.
//
// Note: Non-bracketed IPv6 literals are allowed. Use requireBracketsForIPv6() to prohibit these.
func FromParts(host string, port int) (HostAndPort, error) {
	if err := precond.CheckArgumentf(isValidPort(port), "Port out of range: %d", port); err != nil {
		return HostAndPort{}, err
	}

	hp, err := FromString(host)
	if err != nil {
		return HostAndPort{}, err
	}

	if err := precond.CheckArgumentf(!hp.HasPort(), "Host has a port: %s", host); err != nil {
		return HostAndPort{}, err
	}

	return HostAndPort{hp.Host(), port, hp.hasBracketlessColons}, nil
}

// FromHost builds a HostAndPort instance from a host only.
//
// Note: Non-bracketed IPv6 literals are allowed. Use requireBracketsForIPv6() to prohibit these.
func FromHost(host string) (HostAndPort, error) {
	parsedHost, err := FromString(host)
	if err != nil {
		return parsedHost, err
	}

	if err := precond.CheckArgumentf(!parsedHost.HasPort(), "Host has a port: %s", host); err != nil {
		return parsedHost, err
	}

	return parsedHost, nil
}

// FromString splits a freeform string into a host and port, without strict validation.
//
// Note that the host-only formats will leave the port field undefined.
// You can use withDefaultPort(int) to patch in a default value.
func FromString(hostPort string) (HostAndPort, error) {
	var hp HostAndPort
	var host string
	var portString string
	hasBracketlessColons := false

	if strings.HasPrefix(hostPort, "[") {
		parts, err := getHostAndPortFromBracketedHost(hostPort)
		if err != nil {
			return hp, err
		}
		host = parts[0]
		portString = parts[1]
	} else {
		colonPos := strings.IndexByte(hostPort, ':')
		if colonPos >= 0 && strings.IndexByte(hostPort[colonPos+1:], ':') == -1 {
			// Exactly 1 colon. Split into host:port.
			host = hostPort[0:colonPos]
			portString = hostPort[colonPos+1:]
		} else {
			// 0 or 2+ colons. Bare hostname or IPv6 literal.
			host = hostPort
			hasBracketlessColons = colonPos >= 0
		}
	}

	port := NoPort
	if len(portString) != 0 {
		var err error
		// Try to parse the whole port string as a number.
		if err = precond.CheckArgumentf(!strings.HasPrefix(portString, "+"), "Unparseable port number: %s", hostPort); err != nil {
			return hp, err
		}

		if port, err = strconv.Atoi(portString); err != nil {
			return hp, precond.CheckArgumentf(false, "Unparseable port number: "+hostPort)
		}

		if err = precond.CheckArgumentf(isValidPort(port), "Port number out of range: %s", hostPort); err != nil {
			return hp, err
		}
	}

	return HostAndPort{host, port, hasBracketlessColons}, nil
}

// getHostAndPortFromBracketedHost parses a bracketed host-port string, returning an error if parsing fails.
func getHostAndPortFromBracketedHost(hostPort string) ([]string, error) {
	closeBracketIndex := 0
	if err := precond.CheckArgumentf(hostPort[0] == '[', "Bracketed host-port string must start with a bracket: %s", hostPort); err != nil {
		return nil, err
	}

	colonIndex := strings.IndexByte(hostPort, ':')
	closeBracketIndex = strings.LastIndexByte(hostPort, ']')
	if err := precond.CheckArgumentf(colonIndex > -1 && closeBracketIndex > colonIndex, "Invalid bracketed host/port: %s", hostPort); err != nil {
		return nil, err
	}

	host := hostPort[1:closeBracketIndex]
	if closeBracketIndex+1 == len(hostPort) {
		return []string{host, ""}, nil
	}

	if err := precond.CheckArgumentf(hostPort[closeBracketIndex+1] == ':', "Only a colon may follow a close bracket: %s", hostPort); err != nil {
		return nil, err
	}

	for i := closeBracketIndex + 2; i < len(hostPort); i++ {
		if err := precond.CheckArgumentf(unicode.IsDigit(rune(hostPort[i])), "Port must be numeric: %s", hostPort); err != nil {
			return nil, err
		}
	}
	return []string{host, hostPort[closeBracketIndex+2:]}, nil
}

// WithDefaultPort provides a default port if the parsed string contained only a host.
//
// You can use this after fromString(String) to include a port in case the port was omitted from the input string.
// If a port was already provided, then this method is a no-op.
func (hp HostAndPort) WithDefaultPort(defaultPort int) (HostAndPort, error) {
	if err := precond.CheckArgumentf(isValidPort(defaultPort), "Port out of range: %d", defaultPort); err != nil {
		return hp, err
	}

	if hp.HasPort() {
		return hp, nil
	}

	return HostAndPort{hp.host, defaultPort, hp.hasBracketlessColons}, nil
}

// RequireBracketsForIPv6 generate an error if the host might be a non-bracketed IPv6 literal.
//
// Use this call after fromString(String) to increase the strictness of the parser, and disallow IPv6 literals that
// don't contain these brackets.
//
// Note that this parser identifies IPv6 literals solely based on the presence of a colon.
func (hp HostAndPort) RequireBracketsForIPv6() (HostAndPort, error) {
	if err := precond.CheckArgumentf(!hp.hasBracketlessColons, "Possible bracketless IPv6 literal: %s", hp.Host()); err != nil {
		return hp, err
	}

	return hp, nil
}

// isValidPort returns true for valid port numbers.
func isValidPort(port int) bool {
	return port >= 0 && port <= 65535
}
