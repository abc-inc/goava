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

// Package html provides escaper instances suitable for strings to be included in HTML attribute values and most elements' text contents.
// When possible, avoid manual escaping by using templating systems and high-level APIs that provide autoescaping.
//
// HTML escaping is particularly tricky: For example, some elements' text contents must not be HTML escaped (http://goo.gl/5TgZb).
// As a result, it is impossible to escape an HTML document correctly without domain-specific knowledge beyond what html.Escaper provides.
// We strongly encourage the use of HTML templating systems.
package html

import (
	"github.com/abc-inc/goava/escape"
	"github.com/abc-inc/goava/escape/runes"
)

// Escaper escapes HTML metacharacters as specified by HTML 4.01 (http://www.w3.org/TR/html4/).
// The resulting strings can be used both in attribute values and in most elements' text contents,
// provided that the HTML document's character encoding can encode any non-ASCII code points in the input
// (as UTF-8 and other Unicode encodings can).
//
// Note: This escaper only performs minimal escaping to make content structurally compatible with HTML.
// Specifically, it does not perform entity replacement (symbolic or numeric), so it does not replace non-ASCII code points with character references.
// This escaper escapes only the following five ASCII characters: '"&<>.
var Escaper escape.Escaper

func init() {
	Escaper = runes.NewBuilder().
		AddEscape('"', "&quot;").
		// Note: "&apos;" is not defined in HTML 4.01.
		AddEscape('\'', "&#39;").
		AddEscape('&', "&amp;").
		AddEscape('<', "&lt;").
		AddEscape('>', "&gt;").
		Build()
}
