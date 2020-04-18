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

// Package url provides escaper instances suitable for strings to be included in particular sections of URLs.
//
// If the resulting URLs are inserted into an HTML or XML document, they will require additional escaping with HTML escapers or XML escapers.
package url

import (
	"github.com/abc-inc/goava/escape"
)

const urlFormParameterOtherSafeChars = "-._*~!'()"

const urlPathOtherSafeCharsLackingPlus = "" +
	"-._~" + // Unreserved characters.
	"!$'()*,;&=" + // The subdelim characters (excluding '+').
	"@:" // The gendelim characters permitted in paths.

// FormParameterEscaper escapes strings so they can be safely included in URL form parameter names and values.
// Escaping is performed with the UTF-8 character encoding.
// The caller is responsible for replacing any unpaired carriage return or line feed characters with a CR+LF pair
// on any non-file inputs before escaping them with this escaper.
//
// When escaping a string, the following rules apply:
//
// • The alphanumeric characters "a" through "z", "A" through "Z" and "0" through "9" remain the same.
//
// • The special characters ".", "-", "*", and "_" remain the same.
//
// • The space character " " is converted into a plus sign "+".
//
// • All other characters are converted into one or more bytes using UTF-8 encoding and each byte is then represented
// by the 3-character string "%XY", where "XY" is the two-digit, uppercase, hexadecimal representation of the byte value.
//
// This escaper is suitable for escaping parameter names and values even when using the non-standard semicolon,
// rather than the ampersand, as a parameter delimiter.
// Nevertheless, we recommend using the ampersand unless you must interoperate with systems that require semicolons.
//
// Note: Unlike other escapers, URL escapers produce uppercase hexadecimal sequences.
var FormParameterEscaper escape.Escaper

// PathSegmentEscaper escapes strings so they can be safely included in URL path segments.
// The escaper escapes all non-ASCII characters, even though many of these are accepted in modern URLs.
// (If the escaper were to leave these characters unescaped, they would be escaped by the consumer at parse time, anyway.)
// Additionally, the escaper escapes the slash character ("/").
// While slashes are acceptable in URL paths, they are considered by the specification to be separators between "path segments."
// This implies that, if you wish for your path to contain slashes, you must escape each segment separately and then join them.
//
// When escaping a string, the following rules apply:
//
// • The alphanumeric characters "a" through "z", "A" through "Z" and "0" through "9" remain the same.
//
// • The special characters ".", "-", "*", and "_" remain the same.
//
// • The general delimiters "@" and ":" remain the same.
//
// • The subdelimiters "!", "$", "&amp;", "'", "(", ")", "*", "+", ",", ";", and "=" remain the same.
//
// • The space character " " is converted into %20.
//
// • All other characters are converted into one or more bytes using UTF-8 encoding and each byte is then represented
// by the 3-character string "%XY", where "XY" is the two-digit, uppercase, hexadecimal representation of the byte value.
//
// Note: Unlike other escapers, URL escapers produce uppercase hexadecimal sequences.
var PathSegmentEscaper escape.Escaper

// FragmentEscaper escapes strings so they can be safely included in a URL fragment.
// The escaper escapes all non-ASCII characters, even though many of these are accepted in modern URLs.
//
// When escaping a String, the following rules apply:
//
// • The alphanumeric characters "a" through "z", "A" through "Z" and "0" through "9" remain the same.
//
// • The unreserved characters ".", "-", "~", and "_" remain the same.
//
// • The general delimiters "@" and ":" remain the same.
//
// • The subdelimiters "!", "$", "&amp;", "'", "(", ")", "*", "+", ",", ";", and "=" remain the same.
//
// • The space character " " is converted into %20.
//
// • Fragments allow unescaped "/" and "?", so they remain the same.
//
// • All other characters are converted into one or more bytes using UTF-8 encoding and each byte is then represented
// by the 3-character string "%XY", where "XY" is the two-digit, uppercase, hexadecimal representation of the byte value.
//
// Note: Unlike other escapers, URL escapers produce uppercase hexadecimal sequences.
var FragmentEscaper escape.Escaper

func init() {
	FormParameterEscaper = NewPercentEscaper(urlFormParameterOtherSafeChars, true)
	PathSegmentEscaper = NewPercentEscaper(urlPathOtherSafeCharsLackingPlus+"+", false)
	FragmentEscaper = NewPercentEscaper(urlPathOtherSafeCharsLackingPlus+"+/?", false)
}
