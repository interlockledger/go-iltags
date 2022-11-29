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
	"testing"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------------------
func TestNewStdNullTag(t *testing.T) {
	var tag *NullTag = NewStdNullTag()
	assert.Equal(t, tags.IL_NULL_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdBoolTag(t *testing.T) {
	var tag *BoolTag = NewStdBoolTag()
	assert.Equal(t, tags.IL_BOOL_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdUInt8Tag(t *testing.T) {
	var tag *UInt8Tag = NewStdUInt8Tag()
	assert.Equal(t, tags.IL_UINT8_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdInt8Tag(t *testing.T) {
	var tag *Int8Tag = NewStdInt8Tag()
	assert.Equal(t, tags.IL_INT8_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdUInt16Tag(t *testing.T) {
	var tag *UInt16Tag = NewStdUInt16Tag()
	assert.Equal(t, tags.IL_UINT16_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdInt16Tag(t *testing.T) {
	var tag *Int16Tag = NewStdInt16Tag()
	assert.Equal(t, tags.IL_INT16_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdUInt32Tag(t *testing.T) {
	var tag *UInt32Tag = NewStdUInt32Tag()
	assert.Equal(t, tags.IL_UINT32_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdInt32Tag(t *testing.T) {
	var tag *Int32Tag = NewStdInt32Tag()
	assert.Equal(t, tags.IL_INT32_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdUInt64Tag(t *testing.T) {
	var tag *UInt64Tag = NewStdUInt64Tag()
	assert.Equal(t, tags.IL_UINT64_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdInt64Tag(t *testing.T) {
	var tag *Int64Tag = NewStdInt64Tag()
	assert.Equal(t, tags.IL_INT64_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdFloat32Tag(t *testing.T) {
	var tag *Float32Tag = NewStdFloat32Tag()
	assert.Equal(t, tags.IL_BIN32_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdFloat64Tag(t *testing.T) {
	var tag *Float64Tag = NewStdFloat64Tag()
	assert.Equal(t, tags.IL_BIN64_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdFloat128Tag(t *testing.T) {
	var tag *Float128Tag = NewStdFloat128Tag()
	assert.Equal(t, tags.IL_BIN128_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdILIntTag(t *testing.T) {
	var tag *ILIntTag = NewStdILIntTag()
	assert.Equal(t, tags.IL_ILINT_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdSignedILIntTag(t *testing.T) {
	var tag *SignedILIntTag = NewStdSignedILIntTag()
	assert.Equal(t, tags.IL_SIGNED_ILINT_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdBytesTag(t *testing.T) {
	var tag *BytesTag = NewStdBytesTag()
	assert.Equal(t, tags.IL_BYTES_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdStringTag(t *testing.T) {
	var tag *StringTag = NewStdStringTag()
	assert.Equal(t, tags.IL_STRING_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdBigIntTag(t *testing.T) {
	var tag *BigIntTag = NewStdBigIntTag()
	assert.Equal(t, tags.IL_BINT_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdBigDecTag(t *testing.T) {
	var tag *BigDecTag = NewStdBigDecTag()
	assert.Equal(t, tags.IL_BDEC_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdILIntArrayTag(t *testing.T) {
	var tag *ILIntArrayTag = NewStdILIntArrayTag()
	assert.Equal(t, tags.IL_ILINTARRAY_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdILTagArrayTag(t *testing.T) {
	var tag *ILTagArrayTag = NewStdILTagArrayTag()
	assert.Equal(t, tags.IL_ILTAGARRAY_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdILTagSequenceTag(t *testing.T) {
	var tag *ILTagSequenceTag = NewStdILTagSequenceTag()
	assert.Equal(t, tags.IL_ILTAGSEQ_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdRangeTag(t *testing.T) {
	var tag *RangeTag = NewStdRangeTag()
	assert.Equal(t, tags.IL_RANGE_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdVersionTag(t *testing.T) {
	var tag *VersionTag = NewStdVersionTag()
	assert.Equal(t, tags.IL_VERSION_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdOIDTag(t *testing.T) {
	var tag *OIDTag = NewStdOIDTag()
	assert.Equal(t, tags.IL_OID_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdStringDictionaryTag(t *testing.T) {
	var tag *StringDictionaryTag = NewStdStringDictionaryTag()
	assert.Equal(t, tags.IL_STRING_DICTIONARY_TAG_ID, tag.Id())
}

// ------------------------------------------------------------------------------
func TestNewStdDictionaryTag(t *testing.T) {
	var tag *DictionaryTag = NewStdDictionaryTag()
	assert.Equal(t, tags.IL_DICTIONARY_TAG_ID, tag.Id())
}

func TestNewStandardTag(t *testing.T) {

	var tag tags.ILTag
	var err error

	tag, err = NewStandardTag(tags.IL_NULL_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &NullTag{}, tag)
	assert.Equal(t, tags.IL_NULL_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BOOL_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &BoolTag{}, tag)
	assert.Equal(t, tags.IL_BOOL_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_UINT8_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &UInt8Tag{}, tag)
	assert.Equal(t, tags.IL_UINT8_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_INT8_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Int8Tag{}, tag)
	assert.Equal(t, tags.IL_INT8_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_UINT16_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &UInt16Tag{}, tag)
	assert.Equal(t, tags.IL_UINT16_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_INT16_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Int16Tag{}, tag)
	assert.Equal(t, tags.IL_INT16_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_UINT32_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &UInt32Tag{}, tag)
	assert.Equal(t, tags.IL_UINT32_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_INT32_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Int32Tag{}, tag)
	assert.Equal(t, tags.IL_INT32_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_UINT64_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &UInt64Tag{}, tag)
	assert.Equal(t, tags.IL_UINT64_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_INT64_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Int64Tag{}, tag)
	assert.Equal(t, tags.IL_INT64_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BIN32_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Float32Tag{}, tag)
	assert.Equal(t, tags.IL_BIN32_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BIN64_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Float64Tag{}, tag)
	assert.Equal(t, tags.IL_BIN64_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BIN128_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &Float128Tag{}, tag)
	assert.Equal(t, tags.IL_BIN128_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_ILINT_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &ILIntTag{}, tag)
	assert.Equal(t, tags.IL_ILINT_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_SIGNED_ILINT_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &SignedILIntTag{}, tag)
	assert.Equal(t, tags.IL_SIGNED_ILINT_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BYTES_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &BytesTag{}, tag)
	assert.Equal(t, tags.IL_BYTES_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_STRING_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &StringTag{}, tag)
	assert.Equal(t, tags.IL_STRING_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BINT_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &BigIntTag{}, tag)
	assert.Equal(t, tags.IL_BINT_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_BDEC_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &BigDecTag{}, tag)
	assert.Equal(t, tags.IL_BDEC_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_ILINTARRAY_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &ILIntArrayTag{}, tag)
	assert.Equal(t, tags.IL_ILINTARRAY_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_ILTAGARRAY_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &ILTagArrayTag{}, tag)
	assert.Equal(t, tags.IL_ILTAGARRAY_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_ILTAGSEQ_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &ILTagSequenceTag{}, tag)
	assert.Equal(t, tags.IL_ILTAGSEQ_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_RANGE_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &RangeTag{}, tag)
	assert.Equal(t, tags.IL_RANGE_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_VERSION_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &VersionTag{}, tag)
	assert.Equal(t, tags.IL_VERSION_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_OID_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &OIDTag{}, tag)
	assert.Equal(t, tags.IL_OID_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_STRING_DICTIONARY_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &StringDictionaryTag{}, tag)
	assert.Equal(t, tags.IL_STRING_DICTIONARY_TAG_ID, tag.Id())

	tag, err = NewStandardTag(tags.IL_DICTIONARY_TAG_ID)
	assert.Nil(t, err)
	assert.IsType(t, &DictionaryTag{}, tag)
	assert.Equal(t, tags.IL_DICTIONARY_TAG_ID, tag.Id())

	for _, v := range []tags.TagID{15, 26, 27, 28, 29, 32} {
		tag, err = NewStandardTag(v)
		assert.ErrorIs(t, err, tags.ErrUnsupportedTagId)
		assert.Nil(t, tag)
	}
}
