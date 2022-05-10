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
	"testing"

	. "github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

// This file contains a few serialization/deserialization tests used to ensure
// compatibility with the defined standard

var testTagFactory = NewStandardTagFactory(false)

func TestNullTagReference(t *testing.T) {

	serialized := []byte{0x00}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_NULL_TAG_ID, tag.Id())

	assert.IsType(t, &NullTag{}, tag)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestBoolTagReference(t *testing.T) {

	serialized := []byte{0x01,
		0x01}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BOOL_TAG_ID, tag.Id())
	assert.IsType(t, &BoolTag{}, tag)
	assert.True(t, tag.(*BoolTag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestInt8TagReference(t *testing.T) {

	serialized := []byte{0x02,
		0xFE}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_INT8_TAG_ID, tag.Id())
	assert.IsType(t, &Int8Tag{}, tag)
	assert.Equal(t, int8(-2), tag.(*Int8Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestUInt8TagReference(t *testing.T) {

	serialized := []byte{0x03,
		0xFE}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_UINT8_TAG_ID, tag.Id())
	assert.IsType(t, &UInt8Tag{}, tag)
	assert.Equal(t, uint8(0xFE), tag.(*UInt8Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestInt16TagReference(t *testing.T) {

	serialized := []byte{0x04,
		0xFE, 0xDC}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_INT16_TAG_ID, tag.Id())
	assert.IsType(t, &Int16Tag{}, tag)
	assert.Equal(t, int16(-292), tag.(*Int16Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestUInt16TagReference(t *testing.T) {

	serialized := []byte{0x05,
		0xFE, 0xDC}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_UINT16_TAG_ID, tag.Id())
	assert.IsType(t, &UInt16Tag{}, tag)
	assert.Equal(t, uint16(0xFEDC), tag.(*UInt16Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestInt32TagReference(t *testing.T) {

	serialized := []byte{0x06,
		0xFE, 0xDC, 0xBA, 0x98}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_INT32_TAG_ID, tag.Id())
	assert.IsType(t, &Int32Tag{}, tag)
	assert.Equal(t, int32(-19088744), tag.(*Int32Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestUInt32TagReference(t *testing.T) {

	serialized := []byte{0x07,
		0xFE, 0xDC, 0xBA, 0x98}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_UINT32_TAG_ID, tag.Id())
	assert.IsType(t, &UInt32Tag{}, tag)
	assert.Equal(t, uint32(0xFEDCBA98), tag.(*UInt32Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestInt64TagReference(t *testing.T) {

	serialized := []byte{0x08,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_INT64_TAG_ID, tag.Id())
	assert.IsType(t, &Int64Tag{}, tag)
	assert.Equal(t, int64(-81985529216486896), tag.(*Int64Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestUInt64TagReference(t *testing.T) {

	serialized := []byte{0x09,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_UINT64_TAG_ID, tag.Id())
	assert.IsType(t, &UInt64Tag{}, tag)
	assert.Equal(t, uint64(0xFEDCBA9876543210), tag.(*UInt64Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestILIntTagReference(t *testing.T) {

	serialized := []byte{0x0A,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILINT_TAG_ID, tag.Id())
	assert.IsType(t, &ILIntTag{}, tag)
	assert.Equal(t, uint64(0xdcba9876543308), tag.(*ILIntTag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestFloat32TagReference(t *testing.T) {

	serialized := []byte{0x0B,
		0xFE, 0xDC, 0xBA, 0x98}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BIN32_TAG_ID, tag.Id())
	assert.IsType(t, &Float32Tag{}, tag)
	assert.Equal(t, float32(-1.466995e+38), tag.(*Float32Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestFloat64TagReference(t *testing.T) {

	serialized := []byte{0x0C,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BIN64_TAG_ID, tag.Id())
	assert.IsType(t, &Float64Tag{}, tag)
	assert.Equal(t, float64(-1.2313300687736946e+303), tag.(*Float64Tag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestFloat128TagReference(t *testing.T) {

	serialized := []byte{0x0D,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BIN128_TAG_ID, tag.Id())
	assert.IsType(t, &Float128Tag{}, tag)
	assert.Equal(t, serialized[1:], tag.(*Float128Tag).Payload[:])
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestSignedILIntTagReference(t *testing.T) {

	serialized := []byte{0x0E, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x11}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_SIGNED_ILINT_TAG_ID, tag.Id())
	assert.IsType(t, &SignedILIntTag{}, tag)
	assert.Equal(t, int64(-31064829429684613), tag.(*SignedILIntTag).Payload)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestBytesTagReference(t *testing.T) {

	serialized := []byte{
		0x10, // ID
		0x00, // Size
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BYTES_TAG_ID, tag.Id())
	assert.IsType(t, &BytesTag{}, tag)
	assert.Equal(t, []byte{}, tag.(*BytesTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x10,       // ID
		0xF8, 0x08, // Size
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BYTES_TAG_ID, tag.Id())
	assert.IsType(t, &BytesTag{}, tag)
	assert.Equal(t, serialized[3:], tag.(*BytesTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestStringTagReference(t *testing.T) {

	serialized := []byte{
		0x11, // ID
		0x00, // Size
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_STRING_TAG_ID, tag.Id())
	assert.IsType(t, &StringTag{}, tag)
	assert.Equal(t, "", tag.(*StringTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x11, // ID
		0x25, // Size
		0x41, 0x20, 0x70, 0x72, 0x65, 0x73, 0x73, 0x61,
		0x20, 0xc3, 0xa9, 0x20, 0x61, 0x20, 0x69, 0x6e,
		0x69, 0x6d, 0x69, 0x67, 0x61, 0x20, 0x64, 0x61,
		0x20, 0x70, 0x65, 0x72, 0x66, 0x65, 0x69, 0xc3,
		0xa7, 0xc3, 0xa3, 0x6f, 0x2e}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_STRING_TAG_ID, tag.Id())
	assert.IsType(t, &StringTag{}, tag)
	assert.Equal(t, "A pressa é a inimiga da perfeição.", tag.(*StringTag).Payload)
	w = bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestBigIntTagReference(t *testing.T) {

	serialized := []byte{
		0x12, // ID
		0x01, // Size
		0x23, // Payload
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BINT_TAG_ID, tag.Id())
	assert.IsType(t, &BigIntTag{}, tag)
	assert.Equal(t, []byte{0x23}, tag.(*BigIntTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x12,       // ID
		0xF8, 0x08, // Size
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BINT_TAG_ID, tag.Id())
	assert.IsType(t, &BigIntTag{}, tag)
	assert.Equal(t, serialized[3:], tag.(*BigIntTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x12, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestBigDecTagReference(t *testing.T) {

	serialized := []byte{
		0x13, // ID
		0x05, // Size
		0x23, 0x45, 0x67, 0x89, 0xAB}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BDEC_TAG_ID, tag.Id())
	assert.IsType(t, &BigDecTag{}, tag)
	assert.Equal(t, []byte{0x23}, tag.(*BigDecTag).Payload)
	assert.Equal(t, int32(0x456789AB), tag.(*BigDecTag).Scale)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x13,       // ID
		0xF8, 0x08, // Size
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_BDEC_TAG_ID, tag.Id())
	assert.IsType(t, &BigDecTag{}, tag)
	assert.Equal(t, serialized[3:len(serialized)-4], tag.(*BigDecTag).Payload)
	assert.Equal(t, int32(-1985229329), tag.(*BigDecTag).Scale)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x13, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestILIntArrayTagReference(t *testing.T) {

	serialized := []byte{
		0x14, // ID
		0x01, // Size
		0x00, // Payload
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILINTARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILIntArrayTag{}, tag)
	assert.Equal(t, []uint64{}, tag.(*ILIntArrayTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x14, // ID
		0x37, // Size
		0x0A,
		0xF7,
		0xF8, 0x00,
		0xF9, 0x01, 0x23,
		0xFA, 0x01, 0x23, 0x45,
		0xFB, 0x01, 0x23, 0x45, 0x67,
		0xFC, 0x01, 0x23, 0x45, 0x67, 0x89,
		0xFD, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
		0xFE, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
		0xFF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x07}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILINTARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILIntArrayTag{}, tag)
	assert.Equal(t, []uint64{
		0xf7,
		0xf8,
		0x21b,
		0x1243d,
		0x123465f,
		0x123456881,
		0x12345678aa3,
		0x123456789acc5,
		0x123456789abcee7,
		0xffffffffffffffff}, tag.(*ILIntArrayTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	// Long zeroes
	serialized = append(
		[]byte{
			0x14,       // ID
			0xF8, 0x0A, // Size
			0xF8, 0x08, // Count
		},
		make([]byte, 256)...)
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILINTARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILIntArrayTag{}, tag)
	assert.Equal(t, make([]uint64, 256), tag.(*ILIntArrayTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x14, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestILTagArrayTagReference(t *testing.T) {

	serialized := []byte{
		0x15, // ID
		0x01, // Size
		0x00, // Payload
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagArrayTag{}, tag)
	assert.Equal(t, []ILTag{}, tag.(*ILTagArrayTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x15,       // ID
		0x07,       // Size
		0x03,       // Count
		0x00,       // NullTag
		0x01, 0x00, // Bool tag
		0x04, 0x00, 0x00, // Int16 tag
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagArrayTag{}, tag)
	exp := []ILTag{
		NewStdNullTag(),
		NewStdBoolTag(),
		NewStdInt16Tag()}
	assert.Equal(t, exp, tag.(*ILTagArrayTag).Payload)

	serialized = append(
		[]byte{
			0x15,       // ID
			0xF8, 0x0A, // Size
			0xF8, 0x08, // Count
		},
		make([]byte, 256)...)
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGARRAY_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagArrayTag{}, tag)
	exp = make([]ILTag, 256)
	for i := 0; i < 256; i++ {
		exp[i] = NewStdNullTag()
	}
	assert.Equal(t, exp, tag.(*ILTagArrayTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x15, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestILTagSequenceTagReference(t *testing.T) {

	serialized := []byte{
		0x16, // ID
		0x00, // Size
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGSEQ_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagSequenceTag{}, tag)
	assert.Equal(t, []ILTag{}, tag.(*ILTagSequenceTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x16,       // ID
		0x06,       // Size
		0x00,       // NullTag
		0x01, 0x00, // Bool tag
		0x04, 0x00, 0x00, // Int16 tag
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGSEQ_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagSequenceTag{}, tag)
	exp := []ILTag{
		NewStdNullTag(),
		NewStdBoolTag(),
		NewStdInt16Tag()}
	assert.Equal(t, exp, tag.(*ILTagSequenceTag).Payload)

	serialized = append(
		[]byte{
			0x16,       // ID
			0xF8, 0x08, // Size
		},
		make([]byte, 256)...)
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_ILTAGSEQ_TAG_ID, tag.Id())
	assert.IsType(t, &ILTagSequenceTag{}, tag)
	exp = make([]ILTag, 256)
	for i := 0; i < 256; i++ {
		exp[i] = NewStdNullTag()
	}
	assert.Equal(t, exp, tag.(*ILTagSequenceTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())
}

func TestRangeTagReference(t *testing.T) {

	serialized := []byte{
		0x17, // ID
		0x03, // Size
		0x00,
		0x00, 0x00}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_RANGE_TAG_ID, tag.Id())
	assert.IsType(t, &RangeTag{}, tag)
	assert.Equal(t, uint64(0), tag.(*RangeTag).Start)
	assert.Equal(t, uint16(0), tag.(*RangeTag).Count)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x17, // ID
		0x04, // Size
		0xF8, 0x08,
		0xFE, 0xDC}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_RANGE_TAG_ID, tag.Id())
	assert.IsType(t, &RangeTag{}, tag)
	assert.Equal(t, uint64(256), tag.(*RangeTag).Start)
	assert.Equal(t, uint16(0xFEDC), tag.(*RangeTag).Count)
	w = bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x17, // ID
		0x0A, // Size
		0xFE, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
		0xFE, 0xDC}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_RANGE_TAG_ID, tag.Id())
	assert.IsType(t, &RangeTag{}, tag)
	assert.Equal(t, uint64(0x123456789ACC5), tag.(*RangeTag).Start)
	assert.Equal(t, uint16(0xFEDC), tag.(*RangeTag).Count)
	w = bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x17, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestVersionTagReference(t *testing.T) {

	serialized := []byte{
		0x18, // ID
		0x10, // Size
		0xFE, 0xDC, 0xBA, 0x98,
		0x76, 0x54, 0x32, 0x10,
		0x01, 0x23, 0x45, 0x67,
		0x89, 0xAB, 0xCD, 0xEF}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_VERSION_TAG_ID, tag.Id())
	assert.IsType(t, &VersionTag{}, tag)
	assert.Equal(t, int32(-19088744), tag.(*VersionTag).Major)
	assert.Equal(t, int32(1985229328), tag.(*VersionTag).Minor)
	assert.Equal(t, int32(19088743), tag.(*VersionTag).Revision)
	assert.Equal(t, int32(-1985229329), tag.(*VersionTag).Build)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x18, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}

func TestOIDTagReference(t *testing.T) {

	serialized := []byte{
		0x19, // ID
		0x01, // Size
		0x00, // Payload
	}
	r := bytes.NewReader(serialized)
	tag, err := ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_OID_TAG_ID, tag.Id())
	assert.IsType(t, &OIDTag{}, tag)
	assert.Equal(t, []uint64{}, tag.(*OIDTag).Payload)
	w := bytes.NewBuffer([]byte{})
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x19, // ID
		0x37, // Size
		0x0A,
		0xF7,
		0xF8, 0x00,
		0xF9, 0x01, 0x23,
		0xFA, 0x01, 0x23, 0x45,
		0xFB, 0x01, 0x23, 0x45, 0x67,
		0xFC, 0x01, 0x23, 0x45, 0x67, 0x89,
		0xFD, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
		0xFE, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
		0xFF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x07}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_OID_TAG_ID, tag.Id())
	assert.IsType(t, &OIDTag{}, tag)
	assert.Equal(t, []uint64{
		0xf7,
		0xf8,
		0x21b,
		0x1243d,
		0x123465f,
		0x123456881,
		0x12345678aa3,
		0x123456789acc5,
		0x123456789abcee7,
		0xffffffffffffffff}, tag.(*OIDTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	// Long zeroes
	serialized = append(
		[]byte{
			0x19,       // ID
			0xF8, 0x0A, // Size
			0xF8, 0x08, // Count
		},
		make([]byte, 256)...)
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Nil(t, err)
	assert.Equal(t, IL_OID_TAG_ID, tag.Id())
	assert.IsType(t, &OIDTag{}, tag)
	assert.Equal(t, make([]uint64, 256), tag.(*OIDTag).Payload)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, ILTagSeralize(tag, w))
	assert.Equal(t, serialized, w.Bytes())

	serialized = []byte{
		0x19, // ID
		0x00, // Size
	}
	r = bytes.NewReader(serialized)
	tag, err = ILTagDeserialize(testTagFactory, r)
	assert.Error(t, err)
	assert.Nil(t, tag)
}
