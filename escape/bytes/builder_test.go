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

package bytes_test

import (
	. "github.com/abc-inc/goava/escape/bytes"
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestBuilderInitialStateNoReplacement(t *testing.T) {
	// Unsafe bytes aren't modified by default (unsafeRepl == null).
	e := NewBuilder().SetSafeRange('a', 'z').Build()
	Equal(t, "The Quick Brown Fox", e.Escape("The Quick Brown Fox"))
}

func TestBuilderInitialStateNoneUnsafe(t *testing.T) {
	// No bytes are unsafe by default (safeMin == 0, safeMax == 0xFFFF).
	e := NewBuilder().SetUnsafeReplacement("X").Build()
	Equal(t, "\u0000\uFFFF", e.Escape("\u0000\uFFFF"))
}

func TestBuilderRetainsState(t *testing.T) {
	// Setting a safe range and unsafe replacement works as expected.
	b := NewBuilder()
	b.SetSafeRange('a', 'z')
	b.SetUnsafeReplacement("X")
	Equal(t, "XheXXuickXXrownXXoxX", b.Build().Escape("The Quick Brown Fox!"))
	// Explicit replacements take priority over unsafe bytes.
	b.AddEscapes([]byte{' ', '!'}, "_")
	Equal(t, "Xhe_Xuick_Xrown_Xox_", b.Build().Escape("The Quick Brown Fox!"))
	// Explicit replacements take priority over safe bytes.
	b.SetSafeRange(' ', '~')
	Equal(t, "The_Quick_Brown_Fox_", b.Build().Escape("The Quick Brown Fox!"))
}

func TestBuilderCreatesIndependentEscapers(t *testing.T) {
	// Setup a simple builder and create the first escaper.
	b := NewBuilder()
	b.SetSafeRange('a', 'z')
	b.SetUnsafeReplacement("X")
	b.AddEscape(' ', "_")
	first := b.Build()
	// Modify one of the existing mappings before creating a new escaper.
	b.AddEscape(' ', "-")
	b.AddEscape('!', "$")
	second := b.Build()
	// This should have no effect on existing escapers.
	b.AddEscape(' ', "*")
	// Test both escapers after modifying the builder.
	Equal(t, "Xhe_Xuick_Xrown_XoxX", first.Escape("The Quick Brown Fox!"))
	Equal(t, "Xhe-Xuick-Xrown-Xox$", second.Escape("The Quick Brown Fox!"))
}
