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

package runematcher

import (
	"fmt"
)

//           N777777777NO
//         N7777777777777N
//        M777777777777777N
//        $N877777777D77777M
//       N M77777777ONND777M
//       MN777777777NN  D777
//     N7ZN777777777NN ~M7778
//    N777777777777MMNN88777N
//    N777777777777MNZZZ7777O
//    DZN7777O77777777777777
//     N7OONND7777777D77777N
//      8$M++++?N???$77777$
//       M7++++N+M77777777N
//        N77O777777777777$                              M
//          DNNM$$$$777777N                              D
//         N$N:=N$777N7777M                             NZ
//        77Z::::N777777777                          ODZZZ
//       77N::::::N77777777M                         NNZZZ$
//     $777:::::::77777777MN                        ZM8ZZZZZ
//     777M::::::Z7777777Z77                        N++ZZZZNN
//    7777M:::::M7777777$777M                       $++IZZZZM
//   M777$:::::N777777$M7777M                       +++++ZZZDN
//     NN$::::::7777$$M777777N                      N+++ZZZZNZ
//       N::::::N:7$O:77777777                      N++++ZZZZN
//       M::::::::::::N77777777+                   +?+++++ZZZM
//       8::::::::::::D77777777M                    O+++++ZZ
//        ::::::::::::M777777777N                      O+?D
//        M:::::::::::M77777777778                     77=
//        D=::::::::::N7777777777N                    777
//       INN===::::::=77777777777N                  I777N
//      ?777N========N7777777777787M               N7777
//      77777$D======N77777777777N777N?         N777777
//     I77777$$$N7===M$$77777777$77777777$MMZ77777777N
//      $$$$$$$$$$$NIZN$$$$$$$$$M$$7777777777777777ON
//       M$$$$$$$$M    M$$$$$$$$N=N$$$$7777777$$$ND
//      O77Z$$$$$$$     M$$$$$$$$MNI==$DNNNNM=~N
//   7 :N MNN$$$$M$      $$$777$8      8D8I
//     NMM.:7O           777777778
//                       7777777MN
//                       M NO .7:
//                       M   :   M
//                            8

// Matcher determines a true or false value for any rune.
// Also offers basic text processing methods based on this function.
// Implementations are strongly encouraged to be side-effect-free and immutable.
//
// Throughout the documentation of this class, the phrase "matching character" is used to mean
// "any rune value r for which Matches(r) returns true".
type Matcher interface {
	fmt.Stringer

	// Matches determines a true or false value for the given character.
	Matches(r rune) bool

	// Negate returns a matcher that matches any character not matched by this matcher.
	Negate() Matcher

	// And returns a matcher that matches any character matched by both this matcher and other.
	And(other Matcher) Matcher

	// Or returns a matcher that matches any character matched by either this matcher or other.
	Or(other Matcher) Matcher

	// MatchesAnyOf returns true if a character sequence contains at least one matching character.
	//
	// Equivalent to !MatchesNoneOf(sequence)
	MatchesAnyOf(str string) bool

	// MatchesAllOf returns true if a character sequence contains only matching characters.
	MatchesAllOf(str string) bool

	// MatchesNoneOf returns true if a character sequence contains no matching characters.
	//
	// Equivalent to !MatchesAnyOf(sequence).
	MatchesNoneOf(str string) bool

	// IndexIn returns the index of the first matching character in a character sequence,
	// starting from a given position, or -1 if no character matches after that position.
	IndexIn(str string, start int) int

	// IndexInRunes returns the index of the first matching character in a character sequence,
	// starting from a given position, or -1 if no character matches after that position.
	IndexInRunes(runes []rune, start int) int

	// LastIndexIn returns the index of the last matching character in a character sequence,
	// or -1 if no matching character is present.
	LastIndexIn(str string) int

	// CountIn returns the number of matching characters found in a character sequence.
	CountIn(str string) int

	// RemoveFrom returns a string containing all non-matching characters of a character sequence, in order.
	RemoveFrom(str string) string

	// RetainFrom returns a string containing all matching characters of a character sequence, in order.
	RetainFrom(str string) string

	// ReplaceFromRune returns a string copy of the input character sequence, with each matching character
	// replaced by a given replacement character.
	ReplaceFromRune(str string, replacement rune) string

	// ReplaceFrom returns a string copy of the input character sequence, with each matching character
	// replaced by a given replacement sequence.
	ReplaceFrom(str string, replacement string) string

	// TrimFrom returns a substring of the input character sequence that omits all matching characters
	// from the beginning and from the end of the string.
	TrimFrom(str string) string

	// TrimLeadingFrom returns a substring of the input character sequence that omits all matching characters
	// from the beginning of the string.
	TrimLeadingFrom(str string) string

	// TrimTrailingFrom returns a substring of the input character sequence that omits all matching characters
	// from the end of the string.
	TrimTrailingFrom(str string) string

	// CollapseFrom returns a string copy of the input character sequence, with each group of consecutive matching
	// characters replaced by a single replacement character.
	CollapseFrom(str string, replacement rune) string

	// TrimAndCollapseFrom collapses groups of matching characters exactly as CollapseFrom(str, replacement) does,
	// except that groups of matching characters at the start or end of the sequence are removed without replacement.
	TrimAndCollapseFrom(str string, replacement rune) string
}
