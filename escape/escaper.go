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

package escape

// Escaper converts literal text into a format safe for inclusion in a particular context (such as an XML document).
// Typically (but not always), the inverse process of "unescaping" the text is performed automatically by the relevant parser.
//
// For example, an XML escaper would convert the literal string "Foo<Bar>" into "Foo&lt;Bar&gt;" to prevent "<Bar>" from being confused with an XML tag.
// When the resulting XML document is parsed, the parser API will return this text as the original literal string "Foo<Bar>".
//
// An escaper instance is required to be stateless, and safe when used concurrently by multiple go routines.
//
// Because, in general, escaping operates on the runes of a string and not on its individual byte values,
// it is not safe to assume that Escape(s) is equivalent to Escape(s[0:n]) + Escape(s[n:]) for arbitrary n.
// This is because of the possibility of splitting a surrogate pair.
// The only case in which it is safe to escape strings and concatenate the results is if you can rule out this possibility,
// either by splitting an existing long string into short strings adaptively around surrogate pairs,
// or by starting with short strings already known to be free of unpaired surrogates.
//
// The two primary implementations of this interface are bytes.Escaper and runes.Escaper.
// They are heavily optimized for performance and greatly simplify the task of implementing new escapers.
// It is strongly recommended that when implementing a new escaper you extend one of these implementations.
//
// Popular escapers are defined as constants:
//
// • html.Escaper
//
// • xml.Escaper
//
// To create your own escapers, use bytes.Builder, runes.Builder, or extend bytes.Escaper or runes.Escaper.
type Escaper interface {
	// Escape returns the escaped form of a given literal string.
	//
	// Note that this method may treat input characters differently depending on the specific escaper implementation.
	//
	// • runes.Escaper handles unicode correctly.
	//
	// • bytes.Escaper handles bytes independently and does not verify the input for well formed characters.
	Escape(str string) string
}

// NilEscaper is an escaper that efficiently performs no escaping.
type NilEscaper struct{}

// Escape returns the given literal string.
func (e NilEscaper) Escape(str string) string {
	return str
}
