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

	"github.com/interlockledger/go-iltags/ilint"
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

//------------------------------------------------------------------------------

func TestNullPayload(t *testing.T) {
	var _ ILTagPayload = (*NullPayload)(nil)

	var tag NullPayload

	// Size
	assert.Equal(t, uint64(0), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 0, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 1, r), ErrBadTagFormat)
}

func TestBoolPayload(t *testing.T) {
	var _ ILTagPayload = (*BoolPayload)(nil)

	var tag BoolPayload
	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = true
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x01}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x00})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.False(t, tag.Payload)

	r = bytes.NewReader([]byte{0x01})
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.True(t, tag.Payload)

	r = bytes.NewReader([]byte{0x02})
	assert.Error(t, tag.DeserializeValue(f, 1, r))
	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 2, r), ErrBadTagFormat)
}

func TestUInt8Payload(t *testing.T) {
	var _ ILTagPayload = (*UInt8Payload)(nil)

	var tag UInt8Payload
	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0x33
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x33}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x44})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, uint8(0x44), tag.Payload)

	r = bytes.NewReader([]byte{})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 2, r), ErrBadTagFormat)
}

func TestInt8Payload(t *testing.T) {
	var _ ILTagPayload = (*Int8Payload)(nil)

	var tag Int8Payload
	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = -2
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0xFE}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0xFA})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, int8(-6), tag.Payload)

	r = bytes.NewReader([]byte{})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 2, r), ErrBadTagFormat)
}

func TestUInt16Payload(t *testing.T) {
	var _ ILTagPayload = (*UInt16Payload)(nil)

	var tag UInt16Payload
	// Size
	assert.Equal(t, uint64(2), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0x3344
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x33, 0x44}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x44, 0x33})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 2, r))
	assert.Equal(t, uint16(0x4433), tag.Payload)

	r = bytes.NewReader([]byte{0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 1, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 3, r), ErrBadTagFormat)
}

func TestInt16Payload(t *testing.T) {
	var _ ILTagPayload = (*Int16Payload)(nil)

	var tag Int16Payload
	// Size
	assert.Equal(t, uint64(2), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = -2
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0xFF, 0xFE}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0xFF, 0xFA})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 2, r))
	assert.Equal(t, int16(-6), tag.Payload)

	r = bytes.NewReader([]byte{0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 1, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 3, r), ErrBadTagFormat)
}

func TestUInt32Payload(t *testing.T) {
	var _ ILTagPayload = (*UInt32Payload)(nil)

	var tag UInt32Payload
	// Size
	assert.Equal(t, uint64(4), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0x3344
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x33, 0x44}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x00, 0x00, 0x44, 0x33})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 4, r))
	assert.Equal(t, uint32(0x4433), tag.Payload)

	r = bytes.NewReader([]byte{0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 4, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 3, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 5, r), ErrBadTagFormat)
}

func TestInt32Payload(t *testing.T) {
	var _ ILTagPayload = (*Int32Payload)(nil)

	var tag Int32Payload
	// Size
	assert.Equal(t, uint64(4), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = -2
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0xFF, 0xFF, 0xFF, 0xFE}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFA})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 4, r))
	assert.Equal(t, int32(-6), tag.Payload)

	r = bytes.NewReader([]byte{0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 4, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 1, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 3, r), ErrBadTagFormat)
}

func TestUInt64Payload(t *testing.T) {
	var _ ILTagPayload = (*UInt64Payload)(nil)

	var tag UInt64Payload
	// Size
	assert.Equal(t, uint64(8), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0x3344
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x44}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44, 0x33})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 8, r))
	assert.Equal(t, uint64(0x4433), tag.Payload)

	r = bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 8, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 7, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 9, r), ErrBadTagFormat)
}

func TestInt64Payload(t *testing.T) {
	var _ ILTagPayload = (*Int64Payload)(nil)

	var tag Int64Payload
	// Size
	assert.Equal(t, uint64(8), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = -2
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFA})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 8, r))
	assert.Equal(t, int64(-6), tag.Payload)

	r = bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 8, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 7, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 9, r), ErrBadTagFormat)
}

func TestFloat32Payload(t *testing.T) {
	var _ ILTagPayload = (*Float32Payload)(nil)

	var tag Float32Payload
	// Size
	assert.Equal(t, uint64(4), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0.1234
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x3d, 0xfc, 0xb9, 0x24}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x3d, 0xfc, 0xb9, 0x24})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 4, r))
	assert.Equal(t, float32(0.1234), tag.Payload)

	r = bytes.NewReader([]byte{0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 4, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 3, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 5, r), ErrBadTagFormat)
}

func TestFloat64Payload(t *testing.T) {
	var _ ILTagPayload = (*Float64Payload)(nil)

	var tag Float64Payload
	// Size
	assert.Equal(t, uint64(8), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = 0.1234
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3}, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 8, r))
	assert.Equal(t, float64(0.1234), tag.Payload)

	r = bytes.NewReader([]byte{0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 8, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 7, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 9, r), ErrBadTagFormat)
}

func TestFloat128Payload(t *testing.T) {
	var _ ILTagPayload = (*Float128Payload)(nil)
	sample := []byte{0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3, 0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3}

	var tag Float128Payload
	// Size
	assert.Equal(t, uint64(16), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, make([]byte, 16), w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.SetPayload(sample)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader(sample)
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 16, r))
	assert.Equal(t, sample, tag.Payload[:])

	r = bytes.NewReader([]byte{0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x44})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 8, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 15, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 17, r), ErrBadTagFormat)
}

func TestFloat128PayloadSetPayload(t *testing.T) {
	sample := []byte{0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3, 0x3f, 0xbf, 0x97, 0x24, 0x74, 0x53, 0x8e, 0xf3}

	var tag Float128Payload
	assert.Equal(t, make([]byte, 16), tag.Payload[:])

	tag.SetPayload(sample)
	assert.Equal(t, sample, tag.Payload[:])

	assert.Panics(t, func() {
		tag.SetPayload(nil)
	})

	assert.Panics(t, func() {
		tag.SetPayload(sample[1:])
	})

	assert.Panics(t, func() {
		tag.SetPayload(make([]byte, 17))
	})
}

func TestILIntPayload(t *testing.T) {
	var _ ILTagPayload = (*ILIntPayload)(nil)
	val := uint64(0x1231231313132)
	sample := ilint.Encode(val, nil)

	var tag ILIntPayload

	// Size
	tag.Payload = val
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())
	tag.Payload = 0
	assert.Equal(t, uint64(1), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = val
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader(sample)
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, -1, r))
	assert.Equal(t, val, tag.Payload)

	r = bytes.NewReader(sample[:len(sample)-1])
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, -1, r))
}

func TestSignedILIntPayload(t *testing.T) {
	var _ ILTagPayload = (*SignedILIntPayload)(nil)
	val := int64(-231231313132)
	sample := ilint.EncodeSigned(val, nil)

	var tag SignedILIntPayload

	// Size
	tag.Payload = val
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())
	tag.Payload = 0
	assert.Equal(t, uint64(1), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = val
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader(sample)
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, -1, r))
	assert.Equal(t, val, tag.Payload)

	r = bytes.NewReader(sample[:len(sample)-1])
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, -1, r))
}
