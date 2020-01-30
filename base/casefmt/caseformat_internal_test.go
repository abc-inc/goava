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

package casefmt

import (
	. "github.com/stretchr/testify/require"
	"testing"
)

func TestFirstCharOnlyToUpper(t *testing.T) {
	Equal(t, "", firstCharOnlyToUpper(""))
	Equal(t, " ", firstCharOnlyToUpper(" "))
	Equal(t, "Foo", firstCharOnlyToUpper("foo"))
	Equal(t, "Foobar", firstCharOnlyToUpper("fooBar"))
	Equal(t, "Foo-bar", firstCharOnlyToUpper("foo-bar"))
	Equal(t, "Http", firstCharOnlyToUpper("HTTP"))
	Equal(t, "H_t_t_p", firstCharOnlyToUpper("H_T_T_P"))
}
