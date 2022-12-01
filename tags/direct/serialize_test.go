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
	"math"
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tagtest"
	"github.com/stretchr/testify/assert"
)

var SAMPLE_TAG_IDS = []tags.TagID{
	0x12,
	0x1234,
	0x123456,
	0x12345678,
	0x1234567890,
	0x1234567890AB,
	0x1234567890ABCD,
	0x1234567890ABCDEF,
	0xFFFFFFFFFFFFFF0A,
}

var SAMPLE_TAG_SIZES = []uint64{
	0x12,
	0x1234,
	0x123456,
	0x12345678,
	0x1234567890,
	0x1234567890AB,
	0x1234567890ABCD,
	0x1234567890ABCDEF,
	0xFFFFFFFFFFFFFF0A,
}

func TestSerializeTagId(t *testing.T) {

	for _, id := range SAMPLE_TAG_IDS {
		w := bytes.NewBuffer(nil)
		assert.Nil(t, serializeTagId(id, w))
		assert.Equal(t, ilint.Encode(id.UInt64(), nil), w.Bytes())
	}
}

func TestSerializeSmallValueTagHeader(t *testing.T) {

	for _, id := range SAMPLE_TAG_IDS {
		w := bytes.NewBuffer(nil)
		assert.Nil(t, serializeSmallValueTagHeader(id, 0xf7, w))
		exp := bytes.NewBuffer(nil)
		_, err := ilint.EncodeToWriter(id.UInt64(), exp)
		assert.Nil(t, err)
		assert.Nil(t, exp.WriteByte(0xf7))
		assert.Equal(t, exp.Bytes(), w.Bytes())
	}

	w := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, serializeSmallValueTagHeader(0xF8, 0xf7, w), io.ErrShortWrite)
}

// ------------------------------------------------------------------------------
func TestSerializeStdNullTag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdNullTag(w))
	assert.Equal(t, []byte{0x0}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(0, false)
	assert.ErrorIs(t, SerializeStdNullTag(w1),
		io.ErrShortWrite)
}

func TestSerializeNullTag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeNullTag(0x16, w))
	assert.Equal(t, []byte{
		0x16,
		0x0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeNullTag(0x1234567890ABCDEF, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x0}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeNullTag(0xFA, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdBoolTag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdBoolTag(false, w))
	assert.Equal(t, []byte{0x1, 0x0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdBoolTag(true, w))
	assert.Equal(t, []byte{0x1, 0x1}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdBoolTag(true, w1),
		io.ErrShortWrite)
}

func TestSerializeBoolTag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeBoolTag(0x16, false, w))
	assert.Equal(t, []byte{
		0x16,
		0x1,
		0x0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeBoolTag(0x1234567890ABCDEF, true, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x1,
		0x01}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeBoolTag(0xFA, true, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt8Tag(-4, w))
	assert.Equal(t, []byte{0x2, 0xfc}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt8Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStdUInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt8Tag(0xFA, w))
	assert.Equal(t, []byte{0x3, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt8Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt8Tag(0x16, 0xFA, w))
	assert.Equal(t, []byte{
		0x16,
		0x1,
		0xfa}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt8Tag(0x1234567890ABCDEF, 0xFA, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x1,
		0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt8Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt8Tag(0x16, 123, w))
	assert.Equal(t, []byte{
		0x16,
		0x1,
		0x7b}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt8Tag(0x1234567890ABCDEF, -123, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x1, 0x85}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt8Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt16Tag(0x0123, w))
	assert.Equal(t, []byte{
		0x4,
		0x1, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt16Tag(-0x0123, w))
	assert.Equal(t, []byte{
		0x4,
		0xfe, 0xdd}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStdUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt16Tag(0xFA, w))
	assert.Equal(t, []byte{
		0x5,
		0x0, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x16, 0x0123, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x01, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x1234567890ABCDEF, 0x0123, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0x01, 0x23}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt16Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x16, 0x0123, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x01, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x1234567890ABCDEF,
		-0x0123, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0xfe, 0xdd}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt16Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt32Tag(0x01234567, w))
	assert.Equal(t, []byte{0x6,
		0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt32Tag(-0x01234567, w))
	assert.Equal(t, []byte{0x6,
		0xfe, 0xdc, 0xba, 0x99}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt32Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStdUInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt32Tag(0x01234567, w))
	assert.Equal(t, []byte{
		0x7,
		0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt32Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt32Tag(0x16, 0x01234567, w))
	assert.Equal(t, []byte{
		0x16,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt32Tag(0x1234567890ABCDEF, 0x01234567, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt32Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt32Tag(0x16, 0x01234567, w))
	assert.Equal(t, []byte{
		0x16,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt32Tag(0x1234567890ABCDEF, -0x01234567, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x4,
		0xfe, 0xdc, 0xba, 0x99}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt32Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt64Tag(0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdInt64Tag(-0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{0x8,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt64Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStdUInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdUInt64Tag(0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x9,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStdUInt64Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt64Tag(0x16, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x16,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt64Tag(0x1234567890ABCDEF, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt64Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt64Tag(0x16, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x16,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt64Tag(0x1234567890ABCDEF, -0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x8,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt64Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdFloat32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdFloat32Tag(float32(3.14159274101257324), w))
	assert.Equal(t, []byte{
		0xb,
		0x40, 0x49, 0xf, 0xdb}, w.Bytes())
}

func TestSerializeFloat32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat32Tag(16, float32(3.14159274101257324), w))
	assert.Equal(t, []byte{
		0x10,
		0x4,
		0x40, 0x49, 0xf, 0xdb}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat32Tag(0x1234567890ABCDEF, float32(3.14159274101257324), w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x4,
		0x40, 0x49, 0xf, 0xdb}, w.Bytes())
}

//------------------------------------------------------------------------------

func TestSerializeStdFloat64Tag(t *testing.T) {
	v := math.Float64frombits(0x400921FB54442D18)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdFloat64Tag(v, w))
	assert.Equal(t, []byte{
		0xc,
		0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18}, w.Bytes())
}

func TestSerializeFloat64Tag(t *testing.T) {
	v := math.Float64frombits(0x400921FB54442D18)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat64Tag(16, v, w))
	assert.Equal(t, []byte{
		0x10,
		0x8,
		0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat64Tag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x8,
		0x40, 0x09, 0x21, 0xFB, 0x54, 0x44, 0x2D, 0x18}, w.Bytes())
}

//------------------------------------------------------------------------------

func TestSerializeStdFloat128Tag(t *testing.T) {
	v := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdFloat128Tag(v, w))
	assert.Equal(t, []byte{
		0xd,
		0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(0, false)
	assert.ErrorIs(t, SerializeStdFloat128Tag(v, w1), io.ErrShortWrite)
}

func TestSerializeFloat128Tag(t *testing.T) {
	v := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat128Tag(16, v, w))
	assert.Equal(t, []byte{
		0x10,
		0x10,
		0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeFloat128Tag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x10,
		0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(0, false)
	assert.ErrorIs(t, SerializeFloat128Tag(0x1234567890ABCDEF, v, w1), io.ErrShortWrite)

	w1 = tagtest.NewLimitedWriter(9, false)
	assert.ErrorIs(t, SerializeFloat128Tag(0x1234567890ABCDEF, v, w1), io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStdILIntTag(t *testing.T) {
	v := uint64(0x400921FB54442D18)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdILIntTag(v, w))
	assert.Equal(t, []byte{
		0xa,
		0xff, 0x40, 0x9, 0x21, 0xfb, 0x54, 0x44, 0x2c, 0x20}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(0, false)
	assert.ErrorIs(t, SerializeStdILIntTag(v, w1), io.ErrShortWrite)
}

func TestSerializeILIntTag(t *testing.T) {
	v := uint64(0x400921FB54442D18)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeILIntTag(16, v, w))
	assert.Equal(t, []byte{
		0x10,
		0x09,
		0xff, 0x40, 0x9, 0x21, 0xfb, 0x54, 0x44, 0x2c, 0x20}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeILIntTag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x09,
		0xff, 0x40, 0x9, 0x21, 0xfb, 0x54, 0x44, 0x2c, 0x20}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeILIntTag(0x1234567890ABCDEF, v, w1), io.ErrShortWrite)
}

func TestSerializeStdSignedILIntTag(t *testing.T) {

	v := int64(0)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdSignedILIntTag(v, w))
	assert.Equal(t, []byte{
		0xe,
		0x0}, w.Bytes())

	v = int64(1)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdSignedILIntTag(v, w))
	assert.Equal(t, []byte{
		0xe,
		0x02}, w.Bytes())

	v = int64(-1)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdSignedILIntTag(v, w))
	assert.Equal(t, []byte{
		0xe,
		0x01}, w.Bytes())

	v = int64(9223372036854775807)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdSignedILIntTag(v, w))
	assert.Equal(t, []byte{
		0xe,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6}, w.Bytes())

	v = int64(-9223372036854775808)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStdSignedILIntTag(v, w))
	assert.Equal(t, []byte{
		0xe,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(0, false)
	assert.ErrorIs(t, SerializeStdSignedILIntTag(v, w1), io.ErrShortWrite)
}

func TestSerializeSignedILIntTag(t *testing.T) {

	v := int64(0)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeSignedILIntTag(16, v, w))
	assert.Equal(t, []byte{
		0x10,
		0x01,
		0x0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	v = 1
	assert.Nil(t, SerializeSignedILIntTag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x01,
		0x02}, w.Bytes())

	w = bytes.NewBuffer(nil)
	v = -1
	assert.Nil(t, SerializeSignedILIntTag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x01,
		0x01}, w.Bytes())

	w = bytes.NewBuffer(nil)
	v = 9223372036854775807
	assert.Nil(t, SerializeSignedILIntTag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x09,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6}, w.Bytes())

	w = bytes.NewBuffer(nil)
	v = -9223372036854775808
	assert.Nil(t, SerializeSignedILIntTag(0x1234567890ABCDEF, v, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x09,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeSignedILIntTag(0x1234567890ABCDEF, v, w1), io.ErrShortWrite)
}
