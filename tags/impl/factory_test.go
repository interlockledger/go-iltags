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

	. "github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//------------------------------------------------------------------------------
type mockFactory struct {
	mock.Mock
}

func (f *mockFactory) CreateTag(tagId TagID) (ILTag, error) {
	arg := f.Called(tagId)
	return arg.Get(0).(ILTag), arg.Error(1)
}

func TestNewStandardTagFactory(t *testing.T) {

	f := NewStandardTagFactory(false)
	assert.False(t, f.Strict)
	assert.Nil(t, f.tagCreators)

	f = NewStandardTagFactory(true)
	assert.True(t, f.Strict)
	assert.Nil(t, f.tagCreators)
}

func TestStandardTagFactoryRegisterTag(t *testing.T) {
	for _, strict := range []bool{false, true} {
		var cf TagCreatorFunc

		f := StandardTagFactory{Strict: strict}

		cf = func(id TagID) ILTag {
			return NewInt8Tag(id)
		}
		id := TagID(1231245)
		assert.Nil(t, f.tagCreators[id])
		f.RegisterTag(id, cf)
		assert.NotNil(t, f.tagCreators[id])
		f.RegisterTag(id, nil)
		assert.Nil(t, f.tagCreators[id])

		assert.Panics(t, func() {
			f.RegisterTag(IL_STRING_DICTIONARY_TAG_ID, cf)
		})
	}

}

func TestStandardTagFactoryCreateTag(t *testing.T) {

	for _, strict := range []bool{false, true} {
		var tag ILTag
		var err error

		f := StandardTagFactory{Strict: strict}

		tag, err = f.CreateTag(IL_NULL_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &NullTag{}, tag)
		assert.Equal(t, IL_NULL_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BOOL_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &BoolTag{}, tag)
		assert.Equal(t, IL_BOOL_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_UINT8_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &UInt8Tag{}, tag)
		assert.Equal(t, IL_UINT8_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_INT8_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Int8Tag{}, tag)
		assert.Equal(t, IL_INT8_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_UINT16_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &UInt16Tag{}, tag)
		assert.Equal(t, IL_UINT16_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_INT16_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Int16Tag{}, tag)
		assert.Equal(t, IL_INT16_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_UINT32_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &UInt32Tag{}, tag)
		assert.Equal(t, IL_UINT32_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_INT32_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Int32Tag{}, tag)
		assert.Equal(t, IL_INT32_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_UINT64_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &UInt64Tag{}, tag)
		assert.Equal(t, IL_UINT64_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_INT64_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Int64Tag{}, tag)
		assert.Equal(t, IL_INT64_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BIN32_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Float32Tag{}, tag)
		assert.Equal(t, IL_BIN32_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BIN64_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Float64Tag{}, tag)
		assert.Equal(t, IL_BIN64_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BIN128_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &Float128Tag{}, tag)
		assert.Equal(t, IL_BIN128_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_ILINT_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &ILIntTag{}, tag)
		assert.Equal(t, IL_ILINT_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_SIGNED_ILINT_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &SignedILIntTag{}, tag)
		assert.Equal(t, IL_SIGNED_ILINT_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BYTES_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &BytesTag{}, tag)
		assert.Equal(t, IL_BYTES_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_STRING_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &StringTag{}, tag)
		assert.Equal(t, IL_STRING_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BINT_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &BigIntTag{}, tag)
		assert.Equal(t, IL_BINT_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_BDEC_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &BigDecTag{}, tag)
		assert.Equal(t, IL_BDEC_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_ILINTARRAY_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &ILIntArrayTag{}, tag)
		assert.Equal(t, IL_ILINTARRAY_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_ILTAGARRAY_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &ILTagArrayTag{}, tag)
		assert.Equal(t, IL_ILTAGARRAY_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_ILTAGSEQ_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &ILTagSequenceTag{}, tag)
		assert.Equal(t, IL_ILTAGSEQ_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_RANGE_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &RangeTag{}, tag)
		assert.Equal(t, IL_RANGE_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_VERSION_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &VersionTag{}, tag)
		assert.Equal(t, IL_VERSION_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_OID_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &OIDTag{}, tag)
		assert.Equal(t, IL_OID_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_STRING_DICTIONARY_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &StringDictionaryTag{}, tag)
		assert.Equal(t, IL_STRING_DICTIONARY_TAG_ID, tag.Id())

		tag, err = f.CreateTag(IL_DICTIONARY_TAG_ID)
		assert.Nil(t, err)
		assert.IsType(t, &DictionaryTag{}, tag)
		assert.Equal(t, IL_DICTIONARY_TAG_ID, tag.Id())

		for _, v := range []TagID{15, 26, 27, 28, 29} {
			tag, err = f.CreateTag(v)
			assert.ErrorIs(t, err, ErrUnsupportedTagId)
			assert.Nil(t, tag)
		}

		// Custom tag
		id := TagID(1231245)
		f.RegisterTag(id, func(id TagID) ILTag {
			return NewNullTag(id)
		})

		if f.Strict {
			tag, err = f.CreateTag(id)
			assert.Nil(t, err)
			assert.IsType(t, &NullTag{}, tag)
			assert.Equal(t, id, tag.Id())

			id = 1231313
			tag, err = f.CreateTag(id)
			assert.ErrorIs(t, err, ErrUnsupportedTagId)
			assert.Nil(t, tag)
		} else {
			tag, err = f.CreateTag(id)
			assert.Nil(t, err)
			assert.IsType(t, &NullTag{}, tag)
			assert.Equal(t, id, tag.Id())

			id = 1231313
			tag, err = f.CreateTag(id)
			assert.Nil(t, err)
			assert.IsType(t, &RawTag{}, tag)
			assert.Equal(t, id, tag.Id())
		}
	}
}
