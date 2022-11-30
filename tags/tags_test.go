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

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ------------------------------------------------------------------------------
type mockTag struct {
	mock.Mock
}

func (t *mockTag) Id() TagID {
	ret := t.Called()
	return ret.Get(0).(TagID)
}

func (t *mockTag) Implicit() bool {
	t.Called()
	return t.Id().Implicit()
}

func (t *mockTag) Reserved() bool {
	t.Called()
	return t.Id().Implicit()
}

func (t *mockTag) ValueSize() uint64 {
	ret := t.Called()
	return ret.Get(0).(uint64)
}

func (t *mockTag) SerializeValue(writer io.Writer) error {
	ret := t.Called(writer)
	return ret.Error(0)
}

func (t *mockTag) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	ret := t.Called(factory, valueSize, reader)
	return ret.Error(0)
}

type mockFactory struct {
	mock.Mock
}

func (t *mockFactory) CreateTag(tagId TagID) (ILTag, error) {
	ret := t.Called(tagId)
	err := ret.Error(1)
	if err != nil {
		return nil, err
	} else {
		return ret.Get(0).(ILTag), nil
	}
}

type mockWriter struct {
	mock.Mock
}

func (w *mockWriter) Write(b []byte) (int, error) {
	ret := w.Called(len(b))
	return ret.Int(0), ret.Error(1)
}

//------------------------------------------------------------------------------

func TestTagHeaderSize(t *testing.T) {

	// Implicit
	tag := mockTag{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	assert.Equal(t, uint64(1), tagHeaderSize(&tag))

	// Explicit
	tag = mockTag{}
	tag.On("Id").Return(TagID(123456))
	tag.On("ValueSize").Return(uint64(12312312312313))
	exp := ilint.EncodedSize(123456) + ilint.EncodedSize(12312312312313)
	assert.Equal(t, uint64(exp), tagHeaderSize(&tag))
}

func TestSeralizeTagHeader(t *testing.T) {

	// Implicit
	tag := mockTag{}
	w := mockWriter{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	w.On("Write", 1).Return(1, nil)
	assert.Nil(t, seralizeTagHeader(&tag, &w))

	// Implicit error
	tag = mockTag{}
	w = mockWriter{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	w.On("Write", 1).Return(0, io.ErrUnexpectedEOF)
	assert.Equal(t, uint64(1), tagHeaderSize(&tag))
	assert.ErrorIs(t, seralizeTagHeader(&tag, &w), io.ErrUnexpectedEOF)

	// Explicit
	tag = mockTag{}
	w = mockWriter{}
	tag.On("Id").Return(TagID(123456))
	tag.On("ValueSize").Return(uint64(12312312312313))
	w.On("Write", ilint.EncodedSize(123456)).Return(ilint.EncodedSize(123456), nil)
	w.On("Write", ilint.EncodedSize(12312312312313)).Return(ilint.EncodedSize(12312312312313), nil)
	assert.Nil(t, seralizeTagHeader(&tag, &w))

	tag = mockTag{}
	w = mockWriter{}
	tag.On("Id").Return(TagID(123456))
	tag.On("ValueSize").Return(uint64(12312312312313))
	w.On("Write", ilint.EncodedSize(123456)).Return(0, io.ErrUnexpectedEOF)
	w.On("Write", ilint.EncodedSize(12312312312313)).Return(ilint.EncodedSize(12312312312313), nil)
	assert.ErrorIs(t, seralizeTagHeader(&tag, &w), io.ErrUnexpectedEOF)

	tag = mockTag{}
	w = mockWriter{}
	tag.On("Id").Return(TagID(123456))
	tag.On("ValueSize").Return(uint64(12312312312313))
	w.On("Write", ilint.EncodedSize(123456)).Return(ilint.EncodedSize(123456), nil)
	w.On("Write", ilint.EncodedSize(12312312312313)).Return(0, io.ErrUnexpectedEOF)
	assert.ErrorIs(t, seralizeTagHeader(&tag, &w), io.ErrUnexpectedEOF)
}

func TestILTagSize(t *testing.T) {

	// Implicit
	tag := mockTag{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(10))
	assert.Equal(t, uint64(11), ILTagSize(&tag))

	// Explicit
	tag = mockTag{}
	tag.On("Id").Return(TagID(123456))
	tag.On("ValueSize").Return(uint64(12312312312313))
	exp := uint64(ilint.EncodedSize(123456)+ilint.EncodedSize(12312312312313)) + uint64(12312312312313)
	assert.Equal(t, exp, ILTagSize(&tag))
}

func TestILTagSeralize(t *testing.T) {

	// Implicit
	tag := mockTag{}
	w := bytes.NewBuffer(nil)
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", w).Run(func(args mock.Arguments) {
		args.Get(0).(io.Writer).Write([]byte{1, 2, 3, 4, 5})
	}).Return(nil)
	assert.Nil(t, ILTagSeralize(&tag, w))
	assert.Equal(t, []byte{1, 1, 2, 3, 4, 5}, w.Bytes())

	// Explicit
	tag = mockTag{}
	w = bytes.NewBuffer(nil)
	tag.On("Id").Return(TagID(16))
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", w).Run(func(args mock.Arguments) {
		args.Get(0).(io.Writer).Write([]byte{1, 2, 3, 4, 5})
	}).Return(nil)
	assert.Nil(t, ILTagSeralize(&tag, w))
	assert.Equal(t, []byte{16, 5, 1, 2, 3, 4, 5}, w.Bytes())

	// Failure on header
	tag = mockTag{}
	we := mockWriter{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", &we).Return(nil)
	we.On("Write", 1).Return(0, io.ErrShortWrite)
	assert.ErrorIs(t, ILTagSeralize(&tag, &we), io.ErrShortWrite)

	// Failure on payload serialization
	tag = mockTag{}
	we = mockWriter{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", &we).Return(io.ErrShortWrite)
	we.On("Write", 1).Return(1, nil)
	assert.ErrorIs(t, ILTagSeralize(&tag, &we), io.ErrShortWrite)
}

func TestILTagToBytes(t *testing.T) {

	tag := mockTag{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", mock.Anything).Run(func(args mock.Arguments) {
		args.Get(0).(io.Writer).Write([]byte{1, 2, 3, 4, 5})
	}).Return(nil)
	b, err := ILTagToBytes(&tag)
	assert.Nil(t, err)
	assert.Equal(t, []byte{1, 1, 2, 3, 4, 5}, b)

	// Fail
	tag = mockTag{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("ValueSize").Return(uint64(5))
	tag.On("SerializeValue", mock.Anything).Return(io.ErrShortWrite)
	b, err = ILTagToBytes(&tag)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Nil(t, b)
}

func TestImplicitPayloadSize(t *testing.T) {
	assert.Equal(t, int(0), implicitPayloadSize(IL_NULL_TAG_ID))
	assert.Equal(t, int(1), implicitPayloadSize(IL_BOOL_TAG_ID))
	assert.Equal(t, int(1), implicitPayloadSize(IL_INT8_TAG_ID))
	assert.Equal(t, int(1), implicitPayloadSize(IL_UINT8_TAG_ID))
	assert.Equal(t, int(2), implicitPayloadSize(IL_INT16_TAG_ID))
	assert.Equal(t, int(2), implicitPayloadSize(IL_UINT16_TAG_ID))
	assert.Equal(t, int(4), implicitPayloadSize(IL_INT32_TAG_ID))
	assert.Equal(t, int(4), implicitPayloadSize(IL_UINT32_TAG_ID))
	assert.Equal(t, int(8), implicitPayloadSize(IL_INT64_TAG_ID))
	assert.Equal(t, int(8), implicitPayloadSize(IL_UINT64_TAG_ID))
	assert.Equal(t, int(-1), implicitPayloadSize(IL_ILINT_TAG_ID))
	assert.Equal(t, int(4), implicitPayloadSize(IL_BIN32_TAG_ID))
	assert.Equal(t, int(8), implicitPayloadSize(IL_BIN64_TAG_ID))
	assert.Equal(t, int(16), implicitPayloadSize(IL_BIN128_TAG_ID))
	assert.Equal(t, int(-1), implicitPayloadSize(IL_SIGNED_ILINT_TAG_ID))
	assert.Equal(t, int(-1), implicitPayloadSize(TagID(15)))
	assert.Equal(t, int(-1), implicitPayloadSize(TagID(16)))
	assert.Equal(t, int(-1), implicitPayloadSize(TagID(16123123)))
}

func TestReadTagID(t *testing.T) {

	r := bytes.NewReader([]byte{20})
	v, err := readTagID(r)
	assert.Nil(t, err)
	assert.Equal(t, TagID(20), v)

	r = bytes.NewReader([]byte{0xFF})
	v, err = readTagID(r)
	assert.NotNil(t, err)
	assert.Equal(t, TagID(0), v)
}

func TestReadTagHeader(t *testing.T) {

	// Implicit
	for i := 0; i < 16; i++ {
		expId := TagID(i)
		r := bytes.NewReader([]byte{byte(i)})
		v, s, err := readTagHeader(r)
		assert.Nil(t, err)
		assert.Equal(t, expId, v)
		expSize := int64(implicitPayloadSize(expId))
		assert.Equal(t, uint64(expSize), s)
	}

	// Explicit
	r := bytes.NewReader([]byte{0x10, 0x5})
	v, s, err := readTagHeader(r)
	assert.Nil(t, err)
	assert.Equal(t, TagID(0x10), v)
	assert.Equal(t, uint64(0x5), s)

	// Fail to get the tag
	r = bytes.NewReader([]byte{})
	v, s, err = readTagHeader(r)
	assert.NotNil(t, err)
	assert.Equal(t, TagID(0), v)
	assert.Equal(t, uint64(0), s)

	// Fail to reade the size
	r = bytes.NewReader([]byte{0x10, 0xFF})
	v, s, err = readTagHeader(r)
	assert.NotNil(t, err)
	assert.Equal(t, TagID(0), v)
	assert.Equal(t, uint64(0), s)
}

func TestReadTagPayload(t *testing.T) {

	// Read ILInt
	f := &mockFactory{}
	tag := &mockTag{}
	r := bytes.NewBuffer([]byte{})
	s := uint64(0xFFFF_FFFF_FFFF_FFFF)
	tag.On("Id").Return(IL_ILINT_TAG_ID)
	tag.On("DeserializeValue", f, -1, r).Return(nil)
	assert.Nil(t, readTagPayload(f, r, s, tag))

	// Read Signed ILInt
	f = &mockFactory{}
	tag = &mockTag{}
	s = uint64(0xFFFF_FFFF_FFFF_FFFF)
	tag.On("Id").Return(IL_SIGNED_ILINT_TAG_ID)
	tag.On("DeserializeValue", f, -1, r).Return(nil)
	assert.Nil(t, readTagPayload(f, r, s, tag))

	// Read Tag with size > MAX_TAG_SIZE
	f = &mockFactory{}
	tag = &mockTag{}
	s = uint64(MAX_TAG_SIZE + 1)
	tag.On("Id").Return(IL_BYTES_TAG_ID)
	tag.On("DeserializeValue", f, int(s), r).Return(nil)
	assert.ErrorIs(t, readTagPayload(f, r, s, tag), ErrTagTooLarge)

	// Read Tag with size == 0
	f = &mockFactory{}
	tag = &mockTag{}
	s = uint64(0)
	tag.On("Id").Return(IL_BYTES_TAG_ID)
	tag.On("DeserializeValue", f, int(s), mock.Anything).Return(nil)
	assert.Nil(t, readTagPayload(f, r, s, tag))

	// Read a normal tag
	f = &mockFactory{}
	tag = &mockTag{}
	r = bytes.NewBuffer([]byte{1, 2, 3, 4, 5})
	s = uint64(5)
	tag.On("Id").Return(IL_BYTES_TAG_ID)
	tag.On("DeserializeValue", f, int(s), mock.Anything).Run(func(args mock.Arguments) {
		var tmp [5]byte
		args.Get(2).(io.Reader).Read(tmp[:])
	}).Return(nil)
	assert.Nil(t, readTagPayload(f, r, s, tag))

	// Fail to read all bytes
	f = &mockFactory{}
	tag = &mockTag{}
	r = bytes.NewBuffer([]byte{1, 2, 3, 4, 5})
	s = uint64(5)
	tag.On("Id").Return(IL_BYTES_TAG_ID)
	tag.On("DeserializeValue", f, int(s), mock.Anything).Run(func(args mock.Arguments) {
		var tmp [4]byte
		args.Get(2).(io.Reader).Read(tmp[:])
	}).Return(nil)
	assert.ErrorIs(t, readTagPayload(f, r, s, tag), ErrBadTagFormat)

	// Fail with error
	f = &mockFactory{}
	tag = &mockTag{}
	r = bytes.NewBuffer([]byte{})
	s = uint64(5)
	tag.On("Id").Return(IL_BYTES_TAG_ID)
	tag.On("DeserializeValue", f, int(s), mock.Anything).Return(io.ErrUnexpectedEOF)
	assert.ErrorIs(t, readTagPayload(f, r, s, tag), io.ErrUnexpectedEOF)
}

func TestILTagDeserialize(t *testing.T) {

	// Read Null Tag
	r := bytes.NewBuffer([]byte{0x00})
	tag := &mockTag{}
	f := &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	tag.On("DeserializeValue", f, mock.Anything, mock.Anything).Return(nil)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	nt, err := ILTagDeserialize(f, r)
	assert.Nil(t, err)
	assert.Same(t, nt, tag)

	// Fail on the ID read
	r = bytes.NewBuffer([]byte{})
	tag = &mockTag{}
	f = &mockFactory{}
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	nt, err = ILTagDeserialize(f, r)
	assert.ErrorIs(t, err, io.EOF)
	assert.Nil(t, nt)

	// Bad tag creation
	r = bytes.NewBuffer([]byte{0x00})
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(nil, ErrUnsupportedTagId)
	nt, err = ILTagDeserialize(f, r)
	assert.ErrorIs(t, err, ErrUnsupportedTagId)
	assert.Nil(t, nt)

	// Bad payload
	r = bytes.NewBuffer([]byte{0x01})
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("DeserializeValue", f, mock.Anything, mock.Anything).Return(io.ErrUnexpectedEOF)
	f.On("CreateTag", IL_BOOL_TAG_ID).Return(tag, nil)
	nt, err = ILTagDeserialize(f, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
	assert.Nil(t, nt)
}

func TestILTagDeserializeInto(t *testing.T) {

	// Read Null Tag
	r := bytes.NewBuffer([]byte{0x00})
	tag := &mockTag{}
	f := &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	tag.On("DeserializeValue", f, mock.Anything, mock.Anything).Return(nil)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	err := ILTagDeserializeInto(f, r, tag)
	assert.Nil(t, err)

	// No match
	r = bytes.NewBuffer([]byte{0x01})
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	err = ILTagDeserializeInto(f, r, tag)
	assert.ErrorIs(t, err, ErrUnexpectedTagId)
	assert.ErrorContains(t, err, "expecting tag with id 1 but got the id 0")

	// Bad header
	r = bytes.NewBuffer([]byte{})
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	err = ILTagDeserializeInto(f, r, tag)
	assert.ErrorIs(t, err, io.EOF)

	// Error on unserialize
	r = bytes.NewBuffer([]byte{0x01})
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("DeserializeValue", f, 1, mock.Anything).Return(nil, ErrBadTagFormat)
	f.On("CreateTag", IL_BOOL_TAG_ID).Return(tag, nil)
	err = ILTagDeserializeInto(f, r, tag)
	assert.ErrorIs(t, err, ErrBadTagFormat)
}

func TestILTagFromBytes(t *testing.T) {

	// Read Null Tag
	bin := []byte{0x00}
	tag := &mockTag{}
	f := &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	tag.On("DeserializeValue", f, mock.Anything, mock.Anything).Return(nil)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	nt, err := ILTagFromBytes(f, bin)
	assert.Nil(t, err)
	assert.Same(t, tag, nt)

	// Too much bytes
	bin = []byte{0x00, 0x01}
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_NULL_TAG_ID)
	tag.On("DeserializeValue", f, mock.Anything, mock.Anything).Return(nil)
	f.On("CreateTag", IL_NULL_TAG_ID).Return(tag, nil)
	nt, err = ILTagFromBytes(f, bin)
	assert.ErrorIs(t, err, ErrBadTagFormat)
	assert.Nil(t, nt)

	// Too much bytes
	bin = []byte{0x01}
	tag = &mockTag{}
	f = &mockFactory{}
	tag.On("Id").Return(IL_BOOL_TAG_ID)
	tag.On("DeserializeValue", f, 1, mock.Anything).Return(nil, ErrBadTagFormat)
	f.On("CreateTag", IL_BOOL_TAG_ID).Return(tag, nil)
	nt, err = ILTagFromBytes(f, bin)
	assert.ErrorIs(t, err, ErrBadTagFormat)
	assert.Nil(t, nt)

	// Empty
	bin = []byte{}
	nt, err = ILTagFromBytes(f, bin)
	assert.ErrorIs(t, err, ErrBadTagFormat)
	assert.Nil(t, nt)

	// Nil
	nt, err = ILTagFromBytes(f, nil)
	assert.ErrorIs(t, err, ErrBadTagFormat)
	assert.Nil(t, nt)
}

func TestILTagSeralizeWithNull(t *testing.T) {

	// Serialize nil
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralizeWithNull(nil, w))
	assert.Equal(t, []byte{0}, w.Bytes())

	// Serialize something
	tag := NewRawTag(64)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralizeWithNull(tag, w))
	assert.Equal(t, []byte{0x40, 0x0}, w.Bytes())
}

func TestILTagDeserializeIntoOrNull(t *testing.T) {
	tag := NewRawTag(64)

	// Deserialize nil
	r := bytes.NewReader([]byte{0})
	nullTag, err := ILTagDeserializeIntoOrNull(nil, r, tag)
	assert.Nil(t, err)
	assert.True(t, nullTag)

	// Deserialize something else
	r = bytes.NewReader([]byte{0x40, 0x1, 0x23})
	nullTag, err = ILTagDeserializeIntoOrNull(nil, r, tag)
	assert.Nil(t, err)
	assert.False(t, nullTag)
	assert.Equal(t, []byte{0x23}, tag.Payload)

	// Errors
	r = bytes.NewReader([]byte{})
	nullTag, err = ILTagDeserializeIntoOrNull(nil, r, tag)
	assert.ErrorIs(t, err, io.EOF)
	assert.False(t, nullTag)

	r = bytes.NewReader([]byte{0x41, 0x1, 0x23})
	nullTag, err = ILTagDeserializeIntoOrNull(nil, r, tag)
	assert.ErrorIs(t, err, ErrUnexpectedTagId)
	assert.False(t, nullTag)

	r = bytes.NewReader([]byte{0x40, 0x1})
	nullTag, err = ILTagDeserializeIntoOrNull(nil, r, tag)
	assert.ErrorIs(t, err, io.EOF)
	assert.False(t, nullTag)
}

func TestGetExplicitTagSize(t *testing.T) {

	// The smallest tag possible
	assert.Equal(t, uint64(1+1), GetExplicitTagSize(16, 0))

	// A typical tag
	assert.Equal(t, uint64(1+1+10), GetExplicitTagSize(16, 10))

	// ILInt limits - From ilint_test.go
	assert.Equal(t, uint64(1+9+0x123456789ABCEE7), GetExplicitTagSize(0xF7, 0x123456789ABCEE7))
	assert.Equal(t, uint64(2+8+0x123456789ACC5), GetExplicitTagSize(0xF8, 0x123456789ACC5))
	assert.Equal(t, uint64(3+7+0x012345678AA3), GetExplicitTagSize(0x021B, 0x012345678AA3))
	assert.Equal(t, uint64(4+6+0x0123456881), GetExplicitTagSize(0x01243D, 0x0123456881))
	assert.Equal(t, uint64(5+5+0x0123465F), GetExplicitTagSize(0x0123465F, 0x0123465F))
	assert.Equal(t, uint64(6+4+0x01243D), GetExplicitTagSize(0x0123456881, 0x01243D))
	assert.Equal(t, uint64(7+3+0x021B), GetExplicitTagSize(0x012345678AA3, 0x021B))
	assert.Equal(t, uint64(8+2+0xF8), GetExplicitTagSize(0x123456789ACC5, 0xF8))
	assert.Equal(t, uint64(9+1+0xF7), GetExplicitTagSize(0x123456789ABCEE7, 0xF7))

	assert.Equal(t, uint64(1+1+10),
		GetExplicitTagSize(16, 10))

	// Maximum possible tag size
	assert.Equal(t, uint64(1+9+0xFFFF_FFFF_FFFF_FFF5),
		GetExplicitTagSize(16, 0xFFFF_FFFF_FFFF_FFF5))
	assert.Equal(t, uint64(9+9+0xFFFF_FFFF_FFFF_FFED),
		GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF,
			0xFFFF_FFFF_FFFF_FFED))
}
