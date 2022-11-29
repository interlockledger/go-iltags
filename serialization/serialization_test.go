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

package serialization

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var SAMPLE_BIN = []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

type mockWriter struct {
	mock.Mock
}

func (w *mockWriter) Write(b []byte) (int, error) {
	args := w.Mock.Called(b)
	return args.Int(0), args.Error(1)
}

type mockReader struct {
	mock.Mock
}

func (r *mockReader) Read(b []byte) (int, error) {
	args := r.Mock.Called(b)
	return args.Int(0), args.Error(1)
}

func TestWriteBytes(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteBytes(w, SAMPLE_BIN))
	assert.Equal(t, SAMPLE_BIN, w.Bytes())

	we := mockWriter{}
	we.On("Write", mock.Anything).Return(len(SAMPLE_BIN)-1, nil)
	assert.NotNil(t, WriteBytes(&we, SAMPLE_BIN))

	we.On("Write", mock.Anything).Return(0, fmt.Errorf("dummy"))
	assert.NotNil(t, WriteBytes(&we, SAMPLE_BIN))
}

func TestWriteBool(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteBool(w, false))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, WriteBool(w, true))
	assert.Equal(t, []byte{0x01}, w.Bytes())
}

func TestWriteUint8(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteUInt8(w, uint8(0xFA)))
	assert.Equal(t, []byte{0xFA}, w.Bytes())
}

func TestWriteInt8(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteInt8(w, int8(-2)))
	assert.Equal(t, []byte{0xFE}, w.Bytes())
}

func TestWriteUint16(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteUInt16(w, uint16(0xFACA)))
	assert.Equal(t, []byte{0xFA, 0xCA}, w.Bytes())
}

func TestWriteInt16(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteInt16(w, int16(-2)))
	assert.Equal(t, []byte{0xFF, 0xFE}, w.Bytes())
}

func TestWriteUint32(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteUInt32(w, uint32(0xFACADA01)))
	assert.Equal(t, []byte{0xFA, 0xCA, 0xDA, 0x01}, w.Bytes())
}

func TestWriteInt32(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteInt32(w, int32(-2)))
	assert.Equal(t, []byte{0xFF, 0xFF, 0xFF, 0xFE}, w.Bytes())
}

func TestWriteUint64(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteUInt64(w, uint64(0xFACADA01FACADA01)))
	assert.Equal(t, []byte{0xFA, 0xCA, 0xDA, 0x01, 0xFA, 0xCA, 0xDA, 0x01}, w.Bytes())
}

func TestWriteInt64(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteInt64(w, int64(-2)))
	assert.Equal(t, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE}, w.Bytes())
}

func TestWriteFlaot32(t *testing.T) {
	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteFloat32(w, float32(3.14159274101257324)))
	assert.Equal(t, []byte{0x40, 0x49, 0x0f, 0xdb}, w.Bytes())
}

func TestWriteFlaot64(t *testing.T) {
	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteFloat64(w, float64(0.333333333333333314829616256247390992939472198486328125)))
	assert.Equal(t, []byte{0x3F, 0xD5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55}, w.Bytes())
}

func TestWriteString(t *testing.T) {
	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteString(w, "コーヒー"))
	assert.Equal(t, []byte{0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x92, 0xe3, 0x83, 0xbc}, w.Bytes())

	w = bytes.NewBuffer(nil)
	s := string([]byte{0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x92, 0xe3, 0x83})
	assert.ErrorIs(t, WriteString(w, s), ErrBadUTF8String)
}

func TestWriteILInt(t *testing.T) {
	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteILInt(w, uint64(1234567890)))
	assert.Equal(t, []byte{0xfb, 0x49, 0x96, 0x1, 0xda}, w.Bytes())

	we := mockWriter{}
	we.On("Write", mock.Anything).Return(4, nil)
	assert.NotNil(t, WriteILInt(&we, uint64(1234567890)))

	we = mockWriter{}
	we.On("Write", mock.Anything).Return(0, fmt.Errorf("Dummy"))
	assert.NotNil(t, WriteILInt(&we, uint64(1234567890)))
}

func TestWriteSignedILInt(t *testing.T) {
	w := bytes.NewBuffer(nil)
	assert.Nil(t, WriteSignedILInt(w, int64(-1234567890)))
	assert.Equal(t, []byte{0xfb, 0x93, 0x2c, 0x4, 0xab}, w.Bytes())

	we := mockWriter{}
	we.On("Write", mock.Anything).Return(4, nil)
	assert.Nil(t, WriteSignedILInt(w, int64(-1234567890)))

	we = mockWriter{}
	we.On("Write", mock.Anything).Return(0, fmt.Errorf("Dummy"))
	assert.Nil(t, WriteSignedILInt(w, int64(-1234567890)))
}

func TestReadBytes(t *testing.T) {
	buff := make([]byte, len(SAMPLE_BIN))

	r := bytes.NewReader(SAMPLE_BIN)
	assert.Nil(t, ReadBytes(r, buff))
	assert.Equal(t, SAMPLE_BIN, buff)

	mr := mockReader{}
	mr.On("Read", mock.Anything).Return(0, io.EOF)
	assert.ErrorIs(t, ReadBytes(&mr, buff), io.EOF)

	mr = mockReader{}
	mr.On("Read", mock.Anything).Return(len(buff)-1, nil)
	assert.ErrorIs(t, ReadBytes(&mr, buff), io.ErrUnexpectedEOF)
}

func TestReadBool(t *testing.T) {

	v, err := ReadBool(bytes.NewReader([]byte{0x00}))
	assert.Nil(t, err)
	assert.False(t, v)

	v, err = ReadBool(bytes.NewReader([]byte{0x01}))
	assert.Nil(t, err)
	assert.True(t, v)

	_, err = ReadBool(bytes.NewReader([]byte{0x02}))
	assert.ErrorIs(t, err, ErrSerializationFormat)

	_, err = ReadBool(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadUInt8(t *testing.T) {

	v, err := ReadUInt8(bytes.NewReader([]byte{0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, uint8(0xFA), v)

	_, err = ReadUInt8(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadInt8(t *testing.T) {

	v, err := ReadInt8(bytes.NewReader([]byte{0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, int8(-6), v)

	_, err = ReadInt8(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)
}

func TestReadUInt16(t *testing.T) {

	v, err := ReadUInt16(bytes.NewReader([]byte{0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, uint16(0xFFFA), v)

	_, err = ReadUInt16(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadUInt16(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadInt16(t *testing.T) {

	v, err := ReadInt16(bytes.NewReader([]byte{0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, int16(-6), v)

	_, err = ReadInt16(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadInt16(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadUInt32(t *testing.T) {

	v, err := ReadUInt32(bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, uint32(0xFFFFFFFA), v)

	_, err = ReadUInt32(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadUInt32(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadInt32(t *testing.T) {

	v, err := ReadInt32(bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, int32(-6), v)

	_, err = ReadInt32(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadInt32(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadUInt64(t *testing.T) {

	v, err := ReadUInt64(bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFA), v)

	_, err = ReadUInt64(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadUInt64(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadInt64(t *testing.T) {

	v, err := ReadInt64(bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFA}))
	assert.Nil(t, err)
	assert.Equal(t, int64(-6), v)

	_, err = ReadInt64(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadInt64(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadFloat32(t *testing.T) {

	v, err := ReadFloat32(bytes.NewReader([]byte{0x40, 0x49, 0x0f, 0xdb}))
	assert.Nil(t, err)
	assert.Equal(t, float32(3.14159274101257324), v)

	_, err = ReadFloat32(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadFloat32(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadFloat64(t *testing.T) {

	v, err := ReadFloat64(bytes.NewReader([]byte{0x3F, 0xD5, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55}))
	assert.Nil(t, err)
	assert.Equal(t, float64(0.333333333333333314829616256247390992939472198486328125), v)

	_, err = ReadFloat64(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadFloat64(bytes.NewReader([]byte{0x00}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadString(t *testing.T) {

	v, err := ReadString(bytes.NewReader([]byte{0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x92, 0xe3, 0x83, 0xbc}), 12)
	assert.Nil(t, err)
	assert.Equal(t, "コーヒー", v)

	_, err = ReadString(bytes.NewReader([]byte{}), 12)
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadString(bytes.NewReader([]byte{0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x92, 0xe3, 0x83}), 12)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)

	_, err = ReadString(bytes.NewReader([]byte{0xe3, 0x82, 0xb3, 0xe3, 0x83, 0xbc, 0xe3, 0x83, 0x92, 0xe3, 0x83}), 11)
	assert.ErrorIs(t, err, ErrBadUTF8String)
}

func TestReadILInt(t *testing.T) {

	v, err := ReadILInt(bytes.NewReader([]byte{0xfb, 0x49, 0x96, 0x1, 0xda}))
	assert.Nil(t, err)
	assert.Equal(t, uint64(1234567890), v)

	_, err = ReadILInt(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadILInt(bytes.NewReader([]byte{0xfb, 0x49, 0x96, 0x1}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}

func TestReadSignedILInt(t *testing.T) {

	v, err := ReadSignedILInt(bytes.NewReader([]byte{0xfb, 0x93, 0x2c, 0x4, 0xab}))
	assert.Nil(t, err)
	assert.Equal(t, int64(-1234567890), v)

	_, err = ReadSignedILInt(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, err, io.EOF)

	_, err = ReadSignedILInt(bytes.NewReader([]byte{0xfb, 0x93, 0x2c, 0x4}))
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
}
