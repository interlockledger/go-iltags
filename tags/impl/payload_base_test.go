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
	"fmt"
	"io"
	"math/rand"
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

type limitedDummyWriter struct {
	N int64
}

func (w *limitedDummyWriter) Write(b []byte) (int, error) {
	if w.N == 0 {
		return 0, io.EOF
	}
	n := int64(len(b))
	if n > w.N {
		n = w.N
	}
	w.N -= n
	return int(n), nil
}

func TestRawPayload(t *testing.T) {
	var _ ILTagPayload = (*RawPayload)(nil)
	sample := []byte("And so it begins. You have forgotten something, Commander.")

	var tag RawPayload

	// Size
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = []byte{}
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = sample
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []byte{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = sample
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, []byte{}, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, []byte{}, tag.Payload)

	assert.Nil(t, tag.DeserializeValue(f, len(sample), r))
	assert.Equal(t, sample, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Error(t, tag.DeserializeValue(f, len(sample)+1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
}

func TestStringPayload(t *testing.T) {
	var _ ILTagPayload = (*StringPayload)(nil)
	sample := "You have always been here."
	encSample := []byte(sample)

	var tag StringPayload

	// Size
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = ""
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = sample
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = ""
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = sample
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, encSample, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, "", tag.Payload)

	r = bytes.NewReader(encSample)
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, "", tag.Payload)

	assert.Nil(t, tag.DeserializeValue(f, len(encSample), r))
	assert.Equal(t, sample, tag.Payload)

	r = bytes.NewReader(encSample)
	assert.Error(t, tag.DeserializeValue(f, len(encSample)+1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
}

func TestBigIntPayload(t *testing.T) {
	var _ ILTagPayload = (*BigIntPayload)(nil)
	sample := []byte("If you go to Z'ha'dum, you will die.")

	var tag BigIntPayload

	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())
	tag.Payload = []byte{}
	assert.Equal(t, uint64(1), tag.ValueSize())

	tag.Payload = sample
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []byte{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = sample
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{0x00})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, []byte{0x00}, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Nil(t, tag.DeserializeValue(f, len(sample), r))
	assert.Equal(t, sample, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Error(t, tag.DeserializeValue(f, len(sample)+1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}

func TestBigDecPayload(t *testing.T) {
	var _ ILTagPayload = (*BigDecPayload)(nil)
	sample := []byte("If you go to Z'ha'dum, you will die.")
	encoded := append(sample, 0x01, 0x23, 0x45, 0x67)

	var tag BigDecPayload

	// Size
	assert.Equal(t, uint64(5), tag.ValueSize())
	tag.Payload = []byte{}
	assert.Equal(t, uint64(5), tag.ValueSize())

	tag.Payload = sample
	assert.Equal(t, uint64(len(sample)+4), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	tag.Scale = 0x01234567
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []byte{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = sample
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, encoded, w.Bytes())

	lw := &limitedDummyWriter{1}
	tag.Payload = sample
	assert.Error(t, tag.SerializeValue(lw))

	// Deserialize
	r := bytes.NewReader([]byte{0x00, 0x01, 0x23, 0x45, 0x67})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 5, r))
	assert.Equal(t, []byte{0x00}, tag.Payload)
	assert.Equal(t, int32(0x01234567), tag.Scale)

	r = bytes.NewReader(encoded)
	assert.Nil(t, tag.DeserializeValue(f, len(encoded), r))
	assert.Equal(t, sample, tag.Payload)
	assert.Equal(t, int32(0x01234567), tag.Scale)

	r = bytes.NewReader([]byte{})
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))

	r = bytes.NewReader(encoded)
	assert.Error(t, tag.DeserializeValue(f, len(encoded)+1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}

func CreateSampleILInt64Array(n int) ([]uint64, []byte) {
	l := make([]uint64, n)
	b := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		l[i] = rand.Uint64()
		if err := serialization.WriteILInt(b, l[i]); err != nil {
			panic("Unable to serialize the ILInt")
		}
	}
	return l, b.Bytes()
}

func TestILIntArrayPayload(t *testing.T) {
	var _ ILTagPayload = (*ILIntArrayPayload)(nil)

	var tag ILIntArrayPayload

	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())
	tag.Payload = []uint64{}
	assert.Equal(t, uint64(1), tag.ValueSize())
	for i := range []int{1, 256} {
		sample, bin := CreateSampleILInt64Array(i)
		tag.Payload = sample
		assert.Equal(t, uint64(ilint.EncodedSize(uint64(i))+len(bin)), tag.ValueSize())
	}

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []uint64{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILInt64Array(i)
		tag.Payload = sample
		w = bytes.NewBuffer(nil)
		assert.Nil(t, tag.SerializeValue(w))
		enc := append(ilint.Encode(uint64(i), nil), bin...)
		assert.Equal(t, enc, w.Bytes())
	}

	lw := &limitedDummyWriter{0}
	tag.Payload = []uint64{1}
	assert.Error(t, tag.SerializeValue(lw))

	lw = &limitedDummyWriter{1}
	assert.Error(t, tag.SerializeValue(lw))

	// Deserialize
	r := bytes.NewReader([]byte{0x00})
	f := &mockFactory{}
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, []uint64{}, tag.Payload)

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILInt64Array(i)
		enc := append(ilint.Encode(uint64(i), nil), bin...)
		r := bytes.NewReader(enc)
		f := &mockFactory{}
		assert.Nil(t, tag.DeserializeValue(f, len(enc), r))
		assert.Equal(t, sample, tag.Payload)
	}

	r = bytes.NewReader([]byte{})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 1, r))

	r = bytes.NewReader([]byte{0x01})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	r = bytes.NewReader([]byte{0x02, 0x00})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	r = bytes.NewReader([]byte{0x00, 0x00})
	f = &mockFactory{}
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}

func CreateSampleILTagArray(n int) ([]ILTag, []byte) {
	l := make([]ILTag, n)
	b := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		var t ILTag
		switch i % 3 {
		case 0:
			r := NewStdBoolTag()
			r.Payload = rand.Int()&0x1 == 0
			t = r
		case 1:
			r := NewStdFloat32Tag()
			r.Payload = rand.Float32()
			t = r
		case 2:
			r := NewStdStringTag()
			r.Payload = fmt.Sprintf("%d", rand.Uint64())
			t = r
		}
		l[i] = t
		if err := ILTagSeralize(t, b); err != nil {
			panic("Unable to serialize the ILTag")
		}
	}
	return l, b.Bytes()
}

func TestILTagArrayPayload(t *testing.T) {
	var _ ILTagPayload = (*ILTagArrayPayload)(nil)

	var tag ILTagArrayPayload

	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())
	tag.Payload = []ILTag{}
	assert.Equal(t, uint64(1), tag.ValueSize())
	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		tag.Payload = sample
		assert.Equal(t, uint64(ilint.EncodedSize(uint64(i))+len(bin)), tag.ValueSize())
	}

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []ILTag{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		tag.Payload = sample
		w = bytes.NewBuffer(nil)
		assert.Nil(t, tag.SerializeValue(w))
		enc := append(ilint.Encode(uint64(i), nil), bin...)
		assert.Equal(t, enc, w.Bytes())
	}

	lw := &limitedDummyWriter{0}
	tag.Payload = []ILTag{NewStdNullTag()}
	assert.Error(t, tag.SerializeValue(lw))

	lw = &limitedDummyWriter{1}
	assert.Error(t, tag.SerializeValue(lw))

	// Deserialize
	r := bytes.NewReader([]byte{0x00})
	f := &StandardTagFactory{}

	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, []ILTag{}, tag.Payload)

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		enc := append(ilint.Encode(uint64(i), nil), bin...)
		r := bytes.NewReader(enc)
		assert.Nil(t, tag.DeserializeValue(f, len(enc), r))
		assert.Equal(t, sample, tag.Payload)
	}

	r = bytes.NewReader([]byte{})
	assert.Error(t, tag.DeserializeValue(f, 1, r))

	r = bytes.NewReader([]byte{0x01})
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	r = bytes.NewReader([]byte{0x02, 0x00})
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	r = bytes.NewReader([]byte{0x00, 0x00})
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}

func TestILTagSequencePayload(t *testing.T) {
	var _ ILTagPayload = (*ILTagSequencePayload)(nil)

	var tag ILTagSequencePayload

	// Size
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = []ILTag{}
	assert.Equal(t, uint64(0), tag.ValueSize())
	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		tag.Payload = sample
		assert.Equal(t, uint64(len(bin)), tag.ValueSize())
	}

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []ILTag{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		tag.Payload = sample
		w = bytes.NewBuffer(nil)
		assert.Nil(t, tag.SerializeValue(w))
		assert.Equal(t, bin, w.Bytes())
	}

	lw := &limitedDummyWriter{0}
	tag.Payload = []ILTag{NewStdNullTag()}
	assert.Error(t, tag.SerializeValue(lw))

	// Deserialize
	r := bytes.NewReader([]byte{})
	f := &StandardTagFactory{}

	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, []ILTag{}, tag.Payload)

	for i := range []int{1, 256} {
		sample, bin := CreateSampleILTagArray(i)
		r := bytes.NewReader(bin)
		assert.Nil(t, tag.DeserializeValue(f, len(bin), r))
		assert.Equal(t, sample, tag.Payload)
	}

	r = bytes.NewReader([]byte{})
	assert.Error(t, tag.DeserializeValue(f, 1, r))

	r = bytes.NewReader([]byte{0x01})
	assert.Error(t, tag.DeserializeValue(f, 2, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
}

func TestRangePayload(t *testing.T) {
	var _ ILTagPayload = (*RangePayload)(nil)

	var tag RangePayload

	// Size
	assert.Equal(t, uint64(3), tag.ValueSize())
	for i := range []int{1, 256} {
		tag.Start = uint64(i)
		assert.Equal(t, uint64(ilint.EncodedSize(tag.Start))+2, tag.ValueSize())
	}

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Start = 0
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0, 0, 0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Start = 0xDAD0
	tag.Count = 0xFACA
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0xf9, 0xd9, 0xd8, 0xfa, 0xca}, w.Bytes())

	lw := &limitedDummyWriter{2}
	assert.Error(t, tag.SerializeValue(lw))
	lw = &limitedDummyWriter{4}
	assert.Error(t, tag.SerializeValue(lw))

	// Deserialize
	r := bytes.NewReader([]byte{0, 0, 0})
	f := &mockFactory{}

	assert.Nil(t, tag.DeserializeValue(f, 3, r))
	assert.Equal(t, uint64(0), tag.Start)
	assert.Equal(t, uint16(0), tag.Count)

	r = bytes.NewReader([]byte{0xf9, 0xd9, 0xd8, 0xfa, 0xca})
	assert.Nil(t, tag.DeserializeValue(f, 5, r))
	assert.Equal(t, uint64(0xDAD0), tag.Start)
	assert.Equal(t, uint16(0xFACA), tag.Count)

	r = bytes.NewReader([]byte{0, 0, 0, 0})
	assert.Error(t, tag.DeserializeValue(f, 4, r))

	r = bytes.NewReader([]byte{0xf9, 0xd9})
	assert.Error(t, tag.DeserializeValue(f, 5, r))

	r = bytes.NewReader([]byte{0xf9, 0xd9, 0xd8, 0xfa, 0xca})
	assert.Error(t, tag.DeserializeValue(f, 3, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, 2, r), ErrBadTagFormat)
}
