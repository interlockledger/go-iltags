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

package direct

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeTagId(t *testing.T) {

	r := bytes.NewReader([]byte{0})
	assert.Nil(t, deserializeTagId(0, r))

	r = bytes.NewReader([]byte{0xf7})
	assert.Nil(t, deserializeTagId(0xf7, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	assert.Nil(t, deserializeTagId(0x1234567890ABCDEF, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc})
	assert.ErrorIs(t, deserializeTagId(0x1234567890ABCDEF, r), io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	assert.ErrorIs(t, deserializeTagId(0x1234567890ABCDEE, r), tags.ErrUnexpectedTagId)
}

func TestDeserializeSmallValueHeader(t *testing.T) {
	r := bytes.NewReader([]byte{0, 0})
	s, err := deserializeSmallValueHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), s)

	r = bytes.NewReader([]byte{0, 0xF7})
	s, err = deserializeSmallValueHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	s, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0x1, 0xF7})
	_, err = deserializeSmallValueHeader(0, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0xff})
	_, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.ErrorIs(t, err, io.EOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	_, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.ErrorIs(t, err, io.EOF)

	r = bytes.NewReader([]byte{0x1, 0xF8})
	_, err = deserializeSmallValueHeader(1, r)
	assert.ErrorIs(t, err, tags.ErrBadTagFormat)
}

func TestDeserializeSmallValueHeaderWithSize(t *testing.T) {

	r := bytes.NewReader([]byte{0, 0})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0, 0, r))

	r = bytes.NewReader([]byte{0, 0xF7})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0, 0xF7, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF7, r))

	r = bytes.NewReader([]byte{0xff})
	assert.ErrorIs(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF7, r), io.EOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	assert.ErrorIs(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF6, r), tags.ErrBadTagFormat)
}

func TestDeserializeExplicitHeader(t *testing.T) {

	r := bytes.NewReader([]byte{0, 0})
	s, err := deserializeExplicitHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), s)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	s, err = deserializeExplicitHeader(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0xF7, 0xFB, 0x1f, 0xff, 0xff, 0x08})
	s, err = deserializeExplicitHeader(0xF7, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(536870912), s)

	r = bytes.NewReader([]byte{0xF7, 0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc})
	_, err = deserializeExplicitHeader(0xF7, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0xF7, 0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	_, err = deserializeExplicitHeader(0xF7, r)
	assert.ErrorIs(t, err, tags.ErrTagTooLarge)
}

//------------------------------------------------------------------------------

func TestDeserializeStdNullTag(t *testing.T) {

	r := bytes.NewReader([]byte{0x0})
	assert.Nil(t, DeserializeStdNullTag(r))

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdNullTag(w))
	r = bytes.NewReader(w.Bytes())
	assert.Nil(t, DeserializeStdNullTag(r))

	r = bytes.NewReader([]byte{0x1})
	assert.ErrorIs(t, DeserializeStdNullTag(r), tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{})
	assert.ErrorIs(t, DeserializeStdNullTag(r), io.EOF)
}

func TestDeserializeNullTag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x0})
	assert.Nil(t, DeserializeNullTag(16, r))

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeNullTag(0x1234567890ABCDEF, w))
	r = bytes.NewReader(w.Bytes())
	assert.Nil(t, DeserializeNullTag(0x1234567890ABCDEF, r))

	r = bytes.NewReader([]byte{0x10, 0x1})
	assert.ErrorIs(t, DeserializeNullTag(16, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0x1})
	assert.ErrorIs(t, DeserializeNullTag(16, r), tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{})
	assert.ErrorIs(t, DeserializeNullTag(16, r), io.EOF)
}

//------------------------------------------------------------------------------

func TestDeserializeStdBoolTag(t *testing.T) {

	r := bytes.NewReader([]byte{0x1, 0x0})
	v, err := DeserializeStdBoolTag(r)
	assert.Nil(t, err)
	assert.False(t, v)

	r = bytes.NewReader([]byte{0x1, 0x1})
	v, err = DeserializeStdBoolTag(r)
	assert.Nil(t, err)
	assert.True(t, v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdBoolTag(false, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdBoolTag(r)
	assert.Nil(t, err)
	assert.False(t, v)

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdBoolTag(true, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdBoolTag(r)
	assert.Nil(t, err)
	assert.True(t, v)

	r = bytes.NewReader([]byte{0x1, 0x2})
	_, err = DeserializeStdBoolTag(r)
	assert.ErrorIs(t, err, tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0x0, 0x0})
	_, err = DeserializeStdBoolTag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x0})
	_, err = DeserializeStdBoolTag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)
}

func TestDeserializeBoolTag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x01, 0x0})
	v, err := DeserializeBoolTag(16, r)
	assert.Nil(t, err)
	assert.False(t, v)

	r = bytes.NewReader([]byte{0x10, 0x01, 0x1})
	v, err = DeserializeBoolTag(16, r)
	assert.Nil(t, err)
	assert.True(t, v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeBoolTag(0x1234567890ABCDEF, false, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeBoolTag(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.False(t, v)

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeBoolTag(0x1234567890ABCDEF, true, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeBoolTag(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.True(t, v)

	r = bytes.NewReader([]byte{0x10, 0x1, 0x2})
	_, err = DeserializeBoolTag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0x17, 0x1, 0x0})
	_, err = DeserializeBoolTag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x1})
	_, err = DeserializeBoolTag(0x10, r)
	assert.ErrorIs(t, err, io.EOF)
}

//------------------------------------------------------------------------------

func TestDeserializeStdUInt8TagCore(t *testing.T) {

	r := bytes.NewReader([]byte{0x1, 0x2})
	v, err := deserializeStdUInt8TagCore(0x1, r)
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), v)

	r = bytes.NewReader([]byte{0x1, 0x2})
	_, err = deserializeStdUInt8TagCore(0x2, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x1})
	_, err = deserializeStdUInt8TagCore(0x1, r)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDeserializeStdUInt8Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x3, 0x2})
	v, err := DeserializeStdUInt8Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt8Tag(0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdUInt8Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint8(0xFA), v)

	r = bytes.NewReader([]byte{0x1, 0x2})
	_, err = DeserializeStdUInt8Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x3})
	_, err = DeserializeStdUInt8Tag(r)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDeserializeUInt8Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x1, 0x2})
	v, err := DeserializeUInt8Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt8Tag(0x10, 0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeUInt8Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint8(0xFA), v)

	r = bytes.NewReader([]byte{0x11, 0x01, 0x2})
	_, err = DeserializeUInt8Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x1})
	_, err = DeserializeUInt8Tag(0x10, r)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDeserializeStdInt8Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x2, 0x2})
	v, err := DeserializeStdInt8Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int8(2), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt8Tag(-0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdInt8Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int8(-0x12), v)

	r = bytes.NewReader([]byte{0x1, 0x2})
	_, err = DeserializeStdInt8Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x2})
	_, err = DeserializeStdInt8Tag(r)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDeserializeInt8Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x1, 0x2})
	v, err := DeserializeInt8Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int8(2), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt8Tag(0x10, -0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeInt8Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int8(-0x12), v)

	r = bytes.NewReader([]byte{0x11, 0x01, 0x2})
	_, err = DeserializeInt8Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x1})
	_, err = DeserializeInt8Tag(0x10, r)
	assert.ErrorIs(t, err, io.EOF)
}

//------------------------------------------------------------------------------

func TestDeserializeStdUInt16TagCore(t *testing.T) {

	r := bytes.NewReader([]byte{0x1, 0x2, 0x3})
	v, err := deserializeStdUInt16TagCore(0x1, r)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0x0203), v)

	r = bytes.NewReader([]byte{0x1, 0x2, 0x3})
	_, err = deserializeStdUInt16TagCore(0x2, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x1, 0x2})
	_, err = deserializeStdUInt16TagCore(0x1, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdUInt16Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x5, 0x2, 0x3})
	v, err := DeserializeStdUInt16Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0x0203), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt16Tag(0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdUInt16Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0xFA), v)

	r = bytes.NewReader([]byte{0x6, 0x2, 0x3})
	_, err = DeserializeStdUInt16Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x5, 0x2})
	_, err = DeserializeStdUInt16Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeUInt16Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x2, 0x2, 0x3})
	v, err := DeserializeUInt16Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0x0203), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x10, 0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeUInt16Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint16(0xFA), v)

	r = bytes.NewReader([]byte{0x11, 0x2, 0x2, 0x3})
	_, err = DeserializeUInt16Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x2, 0x2})
	_, err = DeserializeUInt16Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdInt16Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x4, 0x2, 0x3})
	v, err := DeserializeStdInt16Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int16(0x0203), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt16Tag(-0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdInt16Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int16(-0x12), v)

	r = bytes.NewReader([]byte{0x5, 0x2, 0x3})
	_, err = DeserializeStdInt16Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x4, 0x2})
	_, err = DeserializeStdInt16Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeInt16Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x2, 0x2, 0x3})
	v, err := DeserializeInt16Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int16(0x0203), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x10, -0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeInt16Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int16(-0x12), v)

	r = bytes.NewReader([]byte{0x11, 0x2, 0x2, 0x3})
	_, err = DeserializeInt16Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x2, 0x2})
	_, err = DeserializeInt16Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

//------------------------------------------------------------------------------

func TestDeserializeStdUInt32TagCore(t *testing.T) {

	r := bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4, 0x5})
	v, err := deserializeStdUInt32TagCore(0x1, r)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0x02030405), v)

	r = bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4, 0x5})
	_, err = deserializeStdUInt32TagCore(0x2, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4})
	_, err = deserializeStdUInt32TagCore(0x1, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdUInt32Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x7, 0x2, 0x3, 0x4, 0x5})
	v, err := DeserializeStdUInt32Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0x02030405), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt32Tag(0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdUInt32Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0xFA), v)

	r = bytes.NewReader([]byte{0x8, 0x2, 0x3, 0x4, 0x5})
	_, err = DeserializeStdUInt32Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x7, 0x2, 0x3, 0x4})
	_, err = DeserializeStdUInt32Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeUInt32Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x4, 0x2, 0x3, 0x4, 0x5})
	v, err := DeserializeUInt32Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0x02030405), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt32Tag(0x10, 0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeUInt32Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint32(0xFA), v)

	r = bytes.NewReader([]byte{0x11, 0x4, 0x2, 0x3, 0x4, 0x5})
	_, err = DeserializeUInt32Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x4, 0x2, 0x3, 0x4})
	_, err = DeserializeUInt32Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdInt32Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x6, 0x2, 0x3, 0x4, 0x5})
	v, err := DeserializeStdInt32Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int32(0x02030405), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt32Tag(-0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdInt32Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int32(-0x12), v)

	r = bytes.NewReader([]byte{0x7, 0x2, 0x3, 0x4, 0x5})
	_, err = DeserializeStdInt32Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x6, 0x2, 0x3, 0x4})
	_, err = DeserializeStdInt32Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeInt32Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x4, 0x2, 0x3, 0x4, 0x5})
	v, err := DeserializeInt32Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int32(0x02030405), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt32Tag(0x10, -0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeInt32Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int32(-0x12), v)

	r = bytes.NewReader([]byte{0x11, 0x4, 0x2, 0x3, 0x4, 0x5})
	_, err = DeserializeInt32Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x4, 0x2, 0x3, 0x4})
	_, err = DeserializeInt32Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

//------------------------------------------------------------------------------

func TestDeserializeStdUInt64TagCore(t *testing.T) {

	r := bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	v, err := deserializeStdUInt64TagCore(0x1, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0x0203040506070809), v)

	r = bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	_, err = deserializeStdUInt64TagCore(0x2, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8})
	_, err = deserializeStdUInt64TagCore(0x1, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdUInt64Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x9, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	v, err := DeserializeStdUInt64Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0x0203040506070809), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt64Tag(0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdUInt64Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFA), v)

	r = bytes.NewReader([]byte{0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	_, err = DeserializeStdUInt64Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x9, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8})
	_, err = DeserializeStdUInt64Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeUInt64Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	v, err := DeserializeUInt64Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0x0203040506070809), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt64Tag(0x10, 0xFA, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeUInt64Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFA), v)

	r = bytes.NewReader([]byte{0x11, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	_, err = DeserializeUInt64Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8})
	_, err = DeserializeUInt64Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeStdInt64Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	v, err := DeserializeStdInt64Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int64(0x0203040506070809), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt64Tag(-0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeStdInt64Tag(r)
	assert.Nil(t, err)
	assert.Equal(t, int64(-0x12), v)

	r = bytes.NewReader([]byte{0x9, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	_, err = DeserializeStdInt64Tag(r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8})
	_, err = DeserializeStdInt64Tag(r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestDeserializeInt64Tag(t *testing.T) {

	r := bytes.NewReader([]byte{0x10, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	v, err := DeserializeInt64Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int64(0x0203040506070809), v)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt64Tag(0x10, -0x12, w))
	r = bytes.NewReader(w.Bytes())
	v, err = DeserializeInt64Tag(0x10, r)
	assert.Nil(t, err)
	assert.Equal(t, int64(-0x12), v)

	r = bytes.NewReader([]byte{0x11, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9})
	_, err = DeserializeInt64Tag(0x10, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0x10, 0x8, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8})
	_, err = DeserializeInt64Tag(0x10, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}
