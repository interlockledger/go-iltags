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

package impl

import (
	"github.com/interlockledger/go-iltags/tags"
)

// Create a new standard NullTag.
func NewStdNullTag() *NullTag {
	return NewNullTag(tags.IL_NULL_TAG_ID)
}

// Create a new standard BoolTag.
func NewStdBoolTag() *BoolTag {
	return NewBoolTag(tags.IL_BOOL_TAG_ID)
}

// Create a new standard UInt8Tag.
func NewStdUInt8Tag() *UInt8Tag {
	return NewUInt8Tag(tags.IL_UINT8_TAG_ID)
}

// Create a new standard UInt8Tag.
func NewStdInt8Tag() *Int8Tag {
	return NewInt8Tag(tags.IL_INT8_TAG_ID)
}

// Create a new standard UInt16Tag.
func NewStdUInt16Tag() *UInt16Tag {
	return NewUInt16Tag(tags.IL_UINT16_TAG_ID)
}

// Create a new standard UInt16Tag.
func NewStdInt16Tag() *Int16Tag {
	return NewInt16Tag(tags.IL_INT16_TAG_ID)
}

// Create a new standard UInt32Tag.
func NewStdUInt32Tag() *UInt32Tag {
	return NewUInt32Tag(tags.IL_UINT32_TAG_ID)
}

// Create a new standard Int32Tag.
func NewStdInt32Tag() *Int32Tag {
	return NewInt32Tag(tags.IL_INT32_TAG_ID)
}

// Create a new standard UInt64Tag.
func NewStdUInt64Tag() *UInt64Tag {
	return NewUInt64Tag(tags.IL_UINT64_TAG_ID)
}

// Create a new standard Int64Tag.
func NewStdInt64Tag() *Int64Tag {
	return NewInt64Tag(tags.IL_INT64_TAG_ID)
}

// Create a new standard ILIntTag.
func NewStdILIntTag() *ILIntTag {
	return NewILIntTag(tags.IL_ILINT_TAG_ID)
}

// Create a new standard Float32Tag.
func NewStdFloat32Tag() *Float32Tag {
	return NewFloat32Tag(tags.IL_BIN32_TAG_ID)
}

// Create a new standard Float32Tag.
func NewStdFloat64Tag() *Float64Tag {
	return NewFloat64Tag(tags.IL_BIN64_TAG_ID)
}

// Create a new standard Float64Tag.
func NewStdFloat128Tag() *Float128Tag {
	return NewFloat128Tag(tags.IL_BIN128_TAG_ID)
}

// Create a new standard SignedILIntTag.
func NewStdSignedILIntTag() *SignedILIntTag {
	return NewSignedILIntTag(tags.IL_SIGNED_ILINT_TAG_ID)
}

// Create a new standard BytesTag.
func NewStdBytesTag() *tags.RawTag {
	return tags.NewRawTag(tags.IL_BYTES_TAG_ID)
}

// Create a new standard StringTag.
func NewStdStringTag() *StringTag {
	return NewStringTag(tags.IL_STRING_TAG_ID)
}

// Create a new standard BigIntTag.
func NewStdBigIntTag() *BigIntTag {
	return NewBigIntTag(tags.IL_BINT_TAG_ID)
}

// Create a new standard BigDecTag.
func NewStdBigDecTag() *BigDecTag {
	return NewBigDecTag(tags.IL_BDEC_TAG_ID)
}

// Create a new standard ILIntArrayTag.
func NewStdILIntArrayTag() *ILIntArrayTag {
	return NewILIntArrayTag(tags.IL_ILINTARRAY_TAG_ID)
}

// Create a new standard ILTagArrayTag.
func NewStdILTagArrayTag() *ILTagArrayTag {
	return NewILTagArrayTag(tags.IL_ILTAGARRAY_TAG_ID)
}

// Create a new standard ILTagSequenceTag.
func NewStdILTagSequenceTag() *ILTagSequenceTag {
	return NewILTagSequenceTag(tags.IL_ILTAGSEQ_TAG_ID)
}

// Create a new standard RangeTag.
func NewStdRangeTag() *RangeTag {
	return NewRangeTag(tags.IL_RANGE_TAG_ID)
}

// Create a new standard VersionTag.
func NewStdVersionTag() *VersionTag {
	return NewVersionTag(tags.IL_VERSION_TAG_ID)
}

// Create a new standard OIDTag.
func NewStdOIDTag() *OIDTag {
	return NewOIDTag(tags.IL_OID_TAG_ID)
}

// Create a new standard DictionaryTag.
func NewStdDictionaryTag() *DictionaryTag {
	return NewDictionaryTag(tags.IL_DICTIONARY_TAG_ID)
}

// Create a new standard StringDictionaryTag.
func NewStdStringDictionaryTag() *StringDictionaryTag {
	return NewStringDictionaryTag(tags.IL_STRING_DICTIONARY_TAG_ID)
}

// Creates a new standard ILTag. It returns an error if the ID is not a standard
// tag or if the tag is not defined.
func NewStandardTag(id tags.TagID) (tags.ILTag, error) {
	var t tags.ILTag
	switch id {
	case tags.IL_NULL_TAG_ID:
		t = NewStdNullTag()
	case tags.IL_BOOL_TAG_ID:
		t = NewStdBoolTag()
	case tags.IL_INT8_TAG_ID:
		t = NewStdInt8Tag()
	case tags.IL_UINT8_TAG_ID:
		t = NewStdUInt8Tag()
	case tags.IL_INT16_TAG_ID:
		t = NewStdInt16Tag()
	case tags.IL_UINT16_TAG_ID:
		t = NewStdUInt16Tag()
	case tags.IL_INT32_TAG_ID:
		t = NewStdInt32Tag()
	case tags.IL_UINT32_TAG_ID:
		t = NewStdUInt32Tag()
	case tags.IL_INT64_TAG_ID:
		t = NewStdInt64Tag()
	case tags.IL_UINT64_TAG_ID:
		t = NewStdUInt64Tag()
	case tags.IL_ILINT_TAG_ID:
		t = NewStdILIntTag()
	case tags.IL_BIN32_TAG_ID:
		t = NewStdFloat32Tag()
	case tags.IL_BIN64_TAG_ID:
		t = NewStdFloat64Tag()
	case tags.IL_BIN128_TAG_ID:
		t = NewStdFloat128Tag()
	case tags.IL_SIGNED_ILINT_TAG_ID:
		t = NewStdSignedILIntTag()
	case tags.IL_BYTES_TAG_ID:
		t = NewStdBytesTag()
	case tags.IL_STRING_TAG_ID:
		t = NewStdStringTag()
	case tags.IL_BINT_TAG_ID:
		t = NewStdBigIntTag()
	case tags.IL_BDEC_TAG_ID:
		t = NewStdBigDecTag()
	case tags.IL_ILINTARRAY_TAG_ID:
		t = NewStdILIntArrayTag()
	case tags.IL_ILTAGARRAY_TAG_ID:
		t = NewStdILTagArrayTag()
	case tags.IL_ILTAGSEQ_TAG_ID:
		t = NewStdILTagSequenceTag()
	case tags.IL_RANGE_TAG_ID:
		t = NewStdRangeTag()
	case tags.IL_VERSION_TAG_ID:
		t = NewStdVersionTag()
	case tags.IL_OID_TAG_ID:
		t = NewStdOIDTag()
	case tags.IL_DICTIONARY_TAG_ID:
		t = NewStdDictionaryTag()
	case tags.IL_STRING_DICTIONARY_TAG_ID:
		t = NewStdStringDictionaryTag()
	default:
		return nil, tags.NewErrUnsupportedTagId(id)
	}
	return t, nil
}
