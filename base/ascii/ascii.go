/*
 * Copyright 2021 The Goava authors
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

// Package ascii provides static methods pertaining to ASCII characters
// (those in the range of values 0x00 through 0x7F), and to strings containing such characters.
package ascii

import (
	"github.com/abc-inc/goava/base/precond"
	"unicode/utf8"
)

// The ASCII control characters, per RFC 20.

// Null ('\0'): The all-zeros character which may serve to accomplish time fill and media fill.
// Normally used as a C string terminator.
// Although RFC 20 names this as "Null", note that it is distinct from the C/C++ "NULL" pointer.
const NUL byte = 0

// Start of Heading: A communication control character used at the beginning of a sequence of
// characters which constitute a machine-sensible address or routing information. Such a sequence
// is referred to as the "heading." An STX character has the effect of terminating a heading.
const SOH byte = 1

// Start of Text: A communication control character which precedes a sequence of characters that
// is to be treated as an entity and entirely transmitted through to the ultimate destination.
// Such a sequence is referred to as "text." STX may be used to terminate a sequence of characters
// started by SOH.
const STX byte = 2

// End of Text: A communication control character used to terminate a sequence of characters
// started with STX and transmitted as an entity.
const ETX byte = 3

// End of Transmission: A communication control character used to indicate the conclusion of a
// transmission, which may have contained one or more texts and any associated headings.
const EOT byte = 4

// Enquiry: A communication control character used in data communication systems as a request for
// a response from a remote station. It may be used as a "Who Are You" (WRU) to obtain
// identification, or may be used to obtain station status, or both.
const ENQ byte = 5

// Acknowledge: A communication control character transmitted by a receiver as an affirmative
// response to a sender.
const ACK byte = 6

// Bell ('\a'): A character for use when there is a need to call for human attention. It may
// control alarm or attention devices.
const BEL byte = 7

// Backspace ('\b'): A format effector which controls the movement of the printing position one
// printing space backward on the same printing line. (Applicable also to display devices.)
const BS byte = 8

// Horizontal Tabulation ('\t'): A format effector which controls the movement of the printing
// position to the next in a series of predetermined positions along the printing line.
// (Applicable also to display devices and the skip function on punched cards.)
const HT byte = 9

// Line Feed ('\n'): A format effector which controls the movement of the printing position to the
// next printing line. (Applicable also to display devices.) Where appropriate, this character may
// have the meaning "New Line" (NL), a format effector which controls the movement of the printing
// point to the first printing position on the next printing line. Use of this convention requires
// agreement between sender and recipient of data.
const LF byte = 10

// Alternate name for LF. (LF is preferred.)
const NL byte = 10

// Vertical Tabulation ('\v'): A format effector which controls the movement of the printing
// position to the next in a series of predetermined printing lines. (Applicable also to display
// devices.
const VT byte = 11

// Form Feed ('\f'): A format effector which controls the movement of the printing position to the
// first pre-determined printing line on the next form or page. (Applicable also to display
// devices.)
const FF byte = 12

// Carriage Return ('\r'): A format effector which controls the movement of the printing position
// to the first printing position on the same printing line. (Applicable also to display devices.)
const CR byte = 13

// Shift Out: A control character indicating that the code combinations which follow shall be
// interpreted as outside of the character set of the standard code table until a Shift In
// character is reached.
const SO byte = 14

// Shift In: A control character indicating that the code combinations which follow shall be
// interpreted according to the standard code table.
const SI byte = 15

// Data Link Escape: A communication control character which will change the meaning of a limited
// number of contiguously following characters. It is used exclusively to provide supplementary
// controls in data communication networks.
const DLE byte = 16

// Device Control 1. Characters for the control of ancillary devices associated with data
// processing or telecommunication systems, more especially switching devices "on" or "off." (If a
// single "stop" control is required to interrupt or turn off ancillary devices, DC4 is the
// preferred assignment.)
const DC1 byte = 17 // aka XON

// Transmission On: Although originally defined as DC1, this ASCII control character is now better
// known as the XON code used for software flow control in serial communications. The main use is
// restarting the transmission after the communication has been stopped by the XOFF control code.
const XON byte = 17 // aka DC1

// Device Control 2. Characters for the control of ancillary devices associated with data
// processing or telecommunication systems, more especially switching devices "on" or "off." (If a
// single "stop" control is required to interrupt or turn off ancillary devices, DC4 is the
// preferred assignment.)
const DC2 byte = 18

// Device Control 3. Characters for the control of ancillary devices associated with data
// processing or telecommunication systems, more especially switching devices "on" or "off." (If a
// single "stop" control is required to interrupt or turn off ancillary devices, DC4 is the
// preferred assignment.)
const DC3 byte = 19 // aka XOFF

// Transmission off. See XON for explanation.
const XOFF byte = 19 // aka DC3

// Device Control 4. Characters for the control of ancillary devices associated with data
// processing or telecommunication systems, more especially switching devices "on" or "off." (If a
// single "stop" control is required to interrupt or turn off ancillary devices, DC4 is the
// preferred assignment.)
const DC4 byte = 20

// Negative Acknowledge: A communication control character transmitted by a receiver as a negative
// response to the sender.
const NAK byte = 21

// Synchronous Idle: A communication control character used by a synchronous transmission system
// in the absence of any other character to provide a signal from which synchronism may be
// achieved or retained.
const SYN byte = 22

// End of Transmission Block: A communication control character used to indicate the end of a
// block of data for communication purposes. ETB is used for blocking data where the block
// structure is not necessarily related to the processing format.
const ETB byte = 23

// Cancel: A control character used to indicate that the data with which it is sent is in error or
// is to be disregarded.
const CAN byte = 24

// End of Medium: A control character associated with the sent data which may be used to identify
// the physical end of the medium, or the end of the used, or wanted, portion of information
// recorded on a medium. (The position of this character does not necessarily correspond to the
// physical end of the medium.)
const EM byte = 25

// Substitute: A character that may be substituted for a character which is determined to be
// invalid or in error.
const SUB byte = 26

// Escape: A control character intended to provide code extension (supplementary characters) in
// general information interchange. The Escape character itself is a prefix affecting the
// interpretation of a limited number of contiguously following characters.
const ESC byte = 27

// File Separator: These four information separators may be used within data in optional fashion,
// except that their hierarchical relationship shall be: FS is the most inclusive, then GS, then
// RS, and US is least inclusive. (The content and length of a File, Group, Record, or Unit are
// not specified.)
const FS byte = 28

// Group Separator: These four information separators may be used within data in optional fashion,
// except that their hierarchical relationship shall be: FS is the most inclusive, then GS, then
// RS, and US is least inclusive. (The content and length of a File, Group, Record, or Unit are
// not specified.)
const GS byte = 29

// Record Separator: These four information separators may be used within data in optional
// fashion, except that their hierarchical relationship shall be: FS is the most inclusive, then
// GS, then RS, and US is least inclusive. (The content and length of a File, Group, Record, or
// Unit are not specified.)
const RS byte = 30

// Unit Separator: These four information separators may be used within data in optional fashion,
// except that their hierarchical relationship shall be: FS is the most inclusive, then GS, then
// RS, and US is least inclusive. (The content and length of a File, Group, Record, or Unit are
// not specified.)
const US byte = 31

// Space: A normally non-printing graphic character used to separate words. It is also a format
// effector which controls the movement of the printing position, one printing position forward.
// (Applicable also to display devices.)
const SP byte = 32

// Alternate name for SP.
const SPACE byte = 32

// Delete: This character is used primarily to "erase" or "obliterate" erroneous or unwanted
// characters in perforated tape.
const DEL byte = 127

// The minimum value of an ASCII character.
const MIN byte = 0

// The maximum value of an ASCII character.
const MAX byte = 127

/** A bit mask which selects the bit encoding ASCII character case. */
const CASE_MASK byte = 0x20

// ToLowerCase returns the lowercase equivalent if the argument is an uppercase ASCII character.
// Otherwise returns the argument.
func ToLowerCase(c byte) byte {
	if IsUpperCase(c) {
		return (c ^ CASE_MASK)
	}
	return c
}

// ToUpperCase returns the uppercase equivalent if the argument is a lowercase ASCII character.
// Otherwise returns the argument.
func ToUpperCase(c byte) byte {
	if IsLowerCase(c) {
		return (c ^ CASE_MASK)
	}
	return c
}

// IsLowerCase indicates whether b is one of the twenty-six lowercase ASCII alphabetic characters
// between 'a' and 'z' inclusive. All others (including non-ASCII characters) return false.
func IsLowerCase(c byte) bool {
	return (c >= 'a') && (c <= 'z')
}

// IsUpperCase indicates whether b is one of the twenty-six uppercase ASCII alphabetic characters
// between 'A' and 'Z' inclusive. All others (including non-ASCII characters) return false.
func IsUpperCase(c byte) bool {
	return (c >= 'A') && (c <= 'Z')
}

// Truncates the given string to the given maximum length. If the length of the string is greater
// than maxLen, the returned string will be exactly maxLen chars in length and will end with the
// given truncInd. Otherwise, the string will be returned with no changes to the content.
//
// Examples:
//
//  ascii.Truncate("foobar", 7, "...") // returns "foobar"
//  ascii.Truncate("foobar", 5, "...") // returns "fo..."
//
// Note: This function may work with certain non-ASCII text but is not safe for usewith arbitrary
// Unicode text. It is mostly intended for use with text that is known to be safe for use with it
// (such as all-ASCII text) and for simple debugging text.
// When using this method, consider the following:
//
// • it may split surrogate pairs
//
// • it may split characters and combining characters
//
// • it does not consider word boundaries
//
// • if truncating for display to users, there are other considerations that
// must be taken into account
//
// • the appropriate truncation indicator may be locale-dependent
//
// • it is safe to use non-ASCII characters in the truncation indicator
func Truncate(str string, maxLen int, truncInd string) (string, error) {
	// Length to truncate the string to, not including the truncation indicator.
	nTruncInd := utf8.RuneCountInString(truncInd)
	n := maxLen - nTruncInd

	// In this worst case, this allows a maxLen equal to the length of the truncInd,
	// meaning that a string will be truncated to just the truncation indicator itself.
	err := precond.CheckArgumentf(
		n >= 0,
		"maxLen (%d) must be >= length of the truncation indicator (%d)",
		maxLen,
		nTruncInd)
	if err != nil {
		return "", err
	}

	if utf8.RuneCountInString(str) <= maxLen {
		return str, nil
	}

	return str[0:n] + truncInd, nil
}
