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
	"bytes"
	"io"
	"math/rand"
	"testing"
	"unicode/utf8"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringTag(t *testing.T) {
	var _ tags.ILTag = (*StringTag)(nil)

	var tag StringTag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, StringPayload{}))
}

func TestNewStringTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *StringTag = NewStringTag(id)
	assert.Equal(t, id, tag.Id())
}

func TestStringTagSize(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := GenerateRandomString()
		require.True(t, utf8.ValidString(s))
		size := ilint.EncodedSize(id.UInt64()) +
			ilint.EncodedSize(uint64(len(s))) +
			len(s)
		assert.Equal(t, uint64(size), StringTagSize(id, s))
	}
}

func TestSerializeStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := GenerateRandomString()
		b := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, b))

		exp := bytes.NewBuffer(nil)
		assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
		assert.Nil(t, serialization.WriteILInt(exp, uint64(len(s))))
		assert.Nil(t, serialization.WriteBytes(exp, []byte(s)))

		assert.Equal(t, exp.Bytes(), b.Bytes())
	}

	id := tags.TagID(256)
	s := "123456"
	w := &limitedDummyWriter{1}
	assert.NotNil(t, SerializeStringTag(id, s, w))

	w = &limitedDummyWriter{2}
	assert.NotNil(t, SerializeStringTag(id, s, w))

	w = &limitedDummyWriter{3}
	assert.NotNil(t, SerializeStringTag(id, s, w))
}

func TestDeserializeStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := GenerateRandomString()

		exp := bytes.NewBuffer(nil)
		assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
		assert.Nil(t, serialization.WriteILInt(exp, uint64(len(s))))
		assert.Nil(t, serialization.WriteBytes(exp, []byte(s)))
		assert.Nil(t, serialization.WriteUInt8(exp, 0))
		serialized := exp.Bytes()

		r := &io.LimitedReader{R: bytes.NewReader(serialized), N: int64(len(serialized) - 1)}
		a, err := DeserializeStringTag(id, r)
		assert.Nil(t, err)
		assert.Equal(t, s, a)
		assert.Equal(t, int64(0), r.N)

		r = &io.LimitedReader{R: bytes.NewReader(serialized), N: int64(len(serialized))}
		a, err = DeserializeStringTag(id, r)
		assert.Nil(t, err)
		assert.Equal(t, s, a)
		assert.Equal(t, int64(1), r.N)

		// Errors
		idSize := ilint.EncodedSize(id.UInt64())
		sizeSize := ilint.EncodedSize(uint64(len(s)))

		re := bytes.NewReader(serialized[:idSize-1])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		// Truncate the last byte
		re = bytes.NewReader(serialized[:idSize+sizeSize-2])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		re = bytes.NewReader(serialized[:idSize+sizeSize])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		re = bytes.NewReader(serialized)
		a, err = DeserializeStringTag(tags.TagID(id+1), re)
		assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)
		assert.Equal(t, "", a)
	}

	exp := bytes.NewBuffer(nil)
	id := tags.TagID(1234)
	assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
	assert.Nil(t, serialization.WriteILInt(exp, uint64(tags.MAX_TAG_SIZE+1)))
	re := bytes.NewReader(exp.Bytes())
	a, err := DeserializeStringTag(id, re)
	assert.ErrorIs(t, err, tags.ErrTagTooLarge)
	assert.Equal(t, "", a)
}

func TestStdStringTagSize(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := GenerateRandomString()
		require.True(t, utf8.ValidString(s))
		assert.Equal(t, StdStringTagSize(s), StringTagSize(id, s))
	}
}

func TestSerializeStdStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := GenerateRandomString()
		exp := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, exp))

		b := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStdStringTag(s, b))
		assert.Equal(t, exp.Bytes(), b.Bytes())
	}

	b := &limitedDummyWriter{N: 2}
	assert.Error(t, SerializeStdStringTag("1234", b))
}

func TestDeserializeStdStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := GenerateRandomString()
		serialized := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, serialized))

		a, err := DeserializeStdStringTag(bytes.NewReader(serialized.Bytes()))
		assert.Nil(t, err)
		assert.Equal(t, s, a)
	}

	// Bad String tag or anything else.
	id := tags.IL_STRING_TAG_ID + 1
	s := GenerateRandomString()
	serialized := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStringTag(id, s, serialized))
	a, err := DeserializeStdStringTag(bytes.NewReader(serialized.Bytes()))
	assert.Error(t, err)
	assert.Equal(t, "", a)
}
