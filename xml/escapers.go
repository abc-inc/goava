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

// Package xml contains escaper instances suitable for strings to be included in XML attribute values and elements' text contents.
// When possible, avoid manual escaping by using templating systems and high-level APIs that provide autoescaping.
//
// Note: Currently the escapers provided by this package do not escape any characters outside the ASCII character range.
// Unlike HTML escaping the XML escapers will not escape non-ASCII characters to their numeric entity replacements.
// These XML escapers provide the minimal level of escaping to ensure that the output can be safely included in a Unicode XML document.
//
// For details on the behavior of the escapers in this package, see sections
// 2.2 (http://www.w3.org/TR/2008/REC-xml-20081126/#charsets) and
// 2.4 (http://www.w3.org/TR/2008/REC-xml-20081126/#syntax) of the XML specification.
package xml

import (
	"github.com/abc-inc/goava/escape"
	"github.com/abc-inc/goava/escape/runes"
)

const (
	minASCIIControlChar = byte(0x00)
	maxASCIIControlChar = byte(0x1F)
)

// ContentEscaper escapes special characters in a string so it can safely be included in an XML document as element content.
// See section 2.4 (http://www.w3.org/TR/2008/REC-xml-20081126/#syntax) of the XML specification.
//
// Note double and single quotes are not escaped, so it is not safe to use this escaper to escape attribute values.
// Use ContentEscaper if the output can appear in element content or AttributeEscaper in attribute values.
//
// This escaper substitutes 0xFFFD for non-whitespace control characters and the character values 0xFFFE and 0xFFFF which are not permitted in XML.
// For more detail see section 2.2 (http://www.w3.org/TR/2008/REC-xml-20081126/#charsets) of the XML specification.
//
// This escaper does not escape non-ASCII characters to their numeric character references (NCR).
// Any non-ASCII characters appearing in the input will be preserved in the output.
// Specifically "\r" (carriage return) is preserved in the output, which may result in it being silently converted to "\n" when the XML is parsed.
//
// This escaper does not treat surrogate pairs specially and does not perform Unicode validation on its input.
var ContentEscaper escape.Escaper

// Escaper escapes special characters in a string so it can safely be included in an XML document as element content.
// See section 2.4 (http://www.w3.org/TR/2008/REC-xml-20081126/#syntax) of the XML specification.
//
// Note double and single quotes are escaped.
//
// This escaper substitutes 0xFFFD for non-whitespace control characters and the character values 0xFFFE and 0xFFFF which are not permitted in XML.
// For more detail see section 2.2 (http://www.w3.org/TR/2008/REC-xml-20081126/#charsets) of the XML specification.
//
// This escaper does not escape non-ASCII characters to their numeric character references (NCR).
// Any non-ASCII characters appearing in the input will be preserved in the output.
// Specifically "\r" (carriage return) is preserved in the output, which may result in it being silently converted to "\n" when the XML is parsed.
//
// This escaper does not treat surrogate pairs specially and does not perform Unicode validation on its input.
var Escaper escape.Escaper

// AttributeEscaper escapes special characters in a string so it can safely be included in XML document as an attribute value.
// See section 3.3.3 (http://www.w3.org/TR/2008/REC-xml-20081126/#AVNormalize) of the XML specification.
//
// This escaper substitutes 0xFFFD for non-whitespace control characters and the character values 0xFFFE and 0xFFFF which are not permitted in XML.
// For more detail see section 2.2 (http://www.w3.org/TR/2008/REC-xml-20081126/#charsets) of the XML specification.
//
// This escaper does not escape non-ASCII characters to their numeric character references (NCR).
// However, horizontal tab '\t', line feed '\n' and carriage return '\r' are escaped to a corresponding NCR "&#x9;", "&#xA;", and "&#xD;" respectively.
// Any other non-ASCII characters appearing in the input will be preserved in the output.
//
// This escaper does not treat surrogate pairs specially and does not perform Unicode validation on its input.
var AttributeEscaper escape.Escaper

func init() {
	b := runes.NewBuilder()
	// The char values \uFFFE and \uFFFF are explicitly not allowed in XML
	// (Unicode code points above \uFFFF are represented via surrogate pairs which means they are treated as pairs of safe characters).
	b.SetSafeRange(0, '\uFFFD')
	// Unsafe characters are replaced with the Unicode replacement character.
	b.SetUnsafeReplacement("\uFFFD")

	/*
	 * Except for \n, \t, and \r, all ASCII control characters are replaced with the Unicode replacement character.
	 *
	 * Implementation note: An alternative to the following would be to make a map that simply replaces the allowed
	 * ASCII whitespace characters with themselves and to set the minimum safe character to 0x20.
	 * However this would slow down the escaping of simple strings that contain \t, \n, or \r.
	 */
	for c := minASCIIControlChar; c <= maxASCIIControlChar; c++ {
		if c != '\t' && c != '\n' && c != '\r' {
			b.AddEscape(rune(c), "\uFFFD")
		}
	}

	// Build the content escaper first and then add quote escaping for the general escaper.
	b.AddEscape('&', "&amp;")
	b.AddEscape('<', "&lt;")
	b.AddEscape('>', "&gt;")
	ContentEscaper = b.Build()
	b.AddEscape('\'', "&apos;")
	b.AddEscape('"', "&quot;")
	Escaper = b.Build()
	b.AddEscape('\t', "&#x9;")
	b.AddEscape('\n', "&#xA;")
	b.AddEscape('\r', "&#xD;")
	AttributeEscaper = b.Build()
}
