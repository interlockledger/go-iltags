/*
 * BSD 3-Clause License
 *
 * Copyright (c) 2022, InterlockLedger Network
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 * * Neither the name of the copyright holder nor the names of its
 *   contributors may be used to endorse or promote products derived from
 *   this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
 * CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
 * OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package tags

// Type of the TagID.
type TagID uint64

const (
	// Maximum tag id value for implicit tags.
	IMPLICIT_ID_MAX TagID = 0xF

	// Maximum tag id value for reserved tags.
	RESERVED_ID_MAX TagID = 0x1F

	// Standard null tag ID.
	IL_NULL_TAG_ID TagID = 0

	// Standard bool tag ID.
	IL_BOOL_TAG_ID TagID = 1

	// Standard signed 8-bit integer tag ID.
	IL_INT8_TAG_ID TagID = 2

	// Standard unsigned 8-bit integer tag ID.
	IL_UINT8_TAG_ID TagID = 3

	// Standard signed 16-bit integer tag ID.
	IL_INT16_TAG_ID TagID = 4

	// Standard unsigned 16-bit integer tag ID.
	IL_UINT16_TAG_ID TagID = 5

	// Standard signed 32-bit integer tag ID.
	IL_INT32_TAG_ID TagID = 6

	// Standard unsigned 32-bit integer tag ID.
	IL_UINT32_TAG_ID TagID = 7

	// Standard signed 64-bit integer tag ID.
	IL_INT64_TAG_ID TagID = 8

	// Standard unsigned 64-bit integer tag ID.
	IL_UINT64_TAG_ID TagID = 9

	// Standard ILInt tag ID.
	IL_ILINT_TAG_ID TagID = 10

	// Standard 32-bit floating point tag ID.
	IL_BIN32_TAG_ID TagID = 11

	// Standard 64-bit floating point tag ID.
	IL_BIN64_TAG_ID TagID = 12

	// Standard 128-bit floating point tag ID.
	IL_BIN128_TAG_ID TagID = 13

	// Standard Signed ILInt tag ID.
	IL_SIGNED_ILINT_TAG_ID TagID = 14

	// Standard byte array tag ID.
	IL_BYTES_TAG_ID TagID = 16

	// Standard string tag ID.
	IL_STRING_TAG_ID TagID = 17

	// Standard big integer tag ID.
	IL_BINT_TAG_ID TagID = 18

	// Standard big decimal tag ID.
	IL_BDEC_TAG_ID TagID = 19

	// Standard ILInt array tag ID.
	IL_ILINTARRAY_TAG_ID TagID = 20

	// Standard ILTag array tag ID.
	IL_ILTAGARRAY_TAG_ID TagID = 21

	// Standard ILTag sequence tag ID.
	IL_ILTAGSEQ_TAG_ID TagID = 22

	// Standard range tag ID.
	IL_RANGE_TAG_ID TagID = 23

	// Standard version tag ID.
	IL_VERSION_TAG_ID TagID = 24

	// Standard OID tag ID.
	IL_OID_TAG_ID TagID = 25

	// Standard dictionary tag ID.
	IL_DICTIONARY_TAG_ID TagID = 30

	// Standard string-only dictionary tag ID.
	IL_STRING_DICTIONARY_TAG_ID TagID = 31
)

// Returns true if this TagID is implicit.
func (i TagID) Implicit() bool {
	return i <= IMPLICIT_ID_MAX
}

// Returns true if this TagID is reserved.
func (i TagID) Reserved() bool {
	return i <= RESERVED_ID_MAX
}

// Returns the Id as an uint64.
func (i TagID) UInt64() uint64 {
	return uint64(i)
}
