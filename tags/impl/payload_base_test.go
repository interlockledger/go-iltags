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
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	tag.Payload = []byte{0xFF} // Just add garbage to ensure overwrite
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
	tag.Payload = []byte{0xFF}
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
	tag.Payload = []byte{0xFF}
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
	tag.Payload = []uint64{0xFF}
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
	tag.Payload = []ILTag{NewStdNullTag()}
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
	tag.Payload = []ILTag{NewStdNullTag()}
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

func TestVersionPayload(t *testing.T) {
	var _ ILTagPayload = (*VersionPayload)(nil)
	var tag VersionPayload
	enc := []byte{
		0x00, 0x00, 0x01, 0x23,
		0x00, 0x00, 0x45, 0x67,
		0x00, 0x00, 0x89, 0xAB,
		0x00, 0x00, 0xCD, 0xEF}

	// Size
	assert.Equal(t, uint64(16), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Major = int32(0x0123)
	tag.Minor = int32(0x4567)
	tag.Revision = int32(0x89AB)
	tag.Build = int32(0xCDEF)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, enc, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Major = 0
	tag.Minor = 0
	tag.Revision = 0
	tag.Build = 0
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t,
		make([]byte, 16), w.Bytes())

	for i := 1; i < 16; i += 4 {
		lw := &limitedDummyWriter{int64(i)}
		assert.Error(t, tag.SerializeValue(lw))
	}

	// Deserialize
	r := bytes.NewReader(enc)
	f := &mockFactory{}

	assert.Nil(t, tag.DeserializeValue(f, 16, r))
	assert.Equal(t, int32(0x0123), tag.Major)
	assert.Equal(t, int32(0x4567), tag.Minor)
	assert.Equal(t, int32(0x89AB), tag.Revision)
	assert.Equal(t, int32(0xCDEF), tag.Build)

	for i := 1; i < 16; i += 4 {
		r = bytes.NewReader(enc[:i])
		assert.Error(t, tag.DeserializeValue(f, 16, r))
	}
	assert.ErrorIs(t, tag.DeserializeValue(f, 15, r), ErrBadTagFormat)
	assert.ErrorIs(t, tag.DeserializeValue(f, 17, r), ErrBadTagFormat)
}

func TestStringDictionaryPayload(t *testing.T) {
	var _ ILTagPayload = (*StringDictionaryPayload)(nil)
	sample, binSample := CreateSampleStringArray(512)
	encoded := append(ilint.Encode(uint64(len(sample)/2), nil), binSample...)

	var tag StringDictionaryPayload

	// Size
	assert.Equal(t, uint64(1), tag.ValueSize())
	for i := 0; i < len(sample)/2; i++ {
		tag.Map.Put(sample[i*2], sample[i*2+1])
	}
	assert.Equal(t, uint64(len(encoded)), tag.ValueSize())

	// Serialize
	tag.Map.Clear()
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	for i := 0; i < len(sample)/2; i++ {
		tag.Map.Put(sample[i*2], sample[i*2+1])
	}
	w = bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, encoded, w.Bytes())

	tag.Map.Clear()
	tag.Map.Put("a1", "a2")

	we := &limitedDummyWriter{N: 0}
	assert.Error(t, tag.SerializeValue(we))
	we = &limitedDummyWriter{N: 1}
	assert.Error(t, tag.SerializeValue(we))
	we = &limitedDummyWriter{N: 5}
	assert.Error(t, tag.SerializeValue(we))

	// Deserialize
	f := NewStandardTagFactory(false)

	tag.Map.Clear()
	tag.Map.Put("a", "b")
	r := bytes.NewReader([]byte{0x00})
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, 0, tag.Map.Size())

	tag.Map.Clear()
	r = bytes.NewReader(encoded)
	assert.Nil(t, tag.DeserializeValue(f, len(encoded), r))
	assert.Equal(t, 256, tag.Map.Size())
	for i := 0; i < 256; i++ {
		k := sample[i*2]
		v := sample[i*2+1]
		assert.Equal(t, k, tag.Map.Keys()[i])
		vc, ok := tag.Map.Get(k)
		assert.True(t, ok)
		assert.Equal(t, v, vc)
	}

	// Deserialization failures
	encoded = []byte{0x01, 0x11, 0x00, 0x11, 0x00} // 1 entry with 2 empty strings

	r = bytes.NewReader(encoded[:0])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))
	r = bytes.NewReader(encoded[:1])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))
	r = bytes.NewReader(encoded[:4])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))

	r = bytes.NewReader(encoded[:1])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))

	// Too large
	r = bytes.NewReader(encoded)
	assert.Error(t, tag.DeserializeValue(f, len(encoded)+1, r))

	// More entries than bytes
	r = bytes.NewReader(encoded)
	assert.ErrorIs(t, tag.DeserializeValue(f, len(encoded)-1, r), ErrBadTagFormat)

	// Bad size
	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}

func TestDictionaryPayload(t *testing.T) {
	var _ ILTagPayload = (*DictionaryPayload)(nil)
	sampleKeys, _ := CreateSampleStringArray(256)
	sampleObjects, _ := CreateSampleILTagArray(256)

	b := bytes.NewBuffer(nil)
	require.Nil(t, serialization.WriteILInt(b, uint64(len(sampleKeys))))
	for i := 0; i < len(sampleKeys); i++ {
		k := sampleKeys[i]
		v := sampleObjects[i]
		require.Nil(t, SerializeStdStringTag(k, b))
		require.Nil(t, ILTagSeralize(v, b))
	}
	encoded := b.Bytes()

	// Size
	var tag DictionaryPayload
	assert.Equal(t, uint64(1), tag.ValueSize())
	for i := 0; i < len(sampleKeys); i++ {
		k := sampleKeys[i]
		v := sampleObjects[i]
		tag.Map.Put(k, v)
	}
	assert.Equal(t, uint64(len(encoded)), tag.ValueSize())

	// Serialize
	tag.Map.Clear()
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, []byte{0x00}, w.Bytes())

	for i := 0; i < len(sampleKeys); i++ {
		k := sampleKeys[i]
		v := sampleObjects[i]
		tag.Map.Put(k, v)
	}
	w = bytes.NewBuffer(nil)
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, encoded, w.Bytes())

	tag.Map.Clear()
	tag.Map.Put("a1", NewStdNullTag()) // 1 + (1 + 1 + 2) + (1)
	we := &limitedDummyWriter{N: 0}
	assert.Error(t, tag.SerializeValue(we))
	we = &limitedDummyWriter{N: 1}
	assert.Error(t, tag.SerializeValue(we))
	we = &limitedDummyWriter{N: 5}
	assert.Error(t, tag.SerializeValue(we))

	// Deserialize
	f := NewStandardTagFactory(false)
	tag.Map.Clear()
	tag.Map.Put("a1", NewStdNullTag())
	r := bytes.NewReader([]byte{0x00})
	assert.Nil(t, tag.DeserializeValue(f, 1, r))
	assert.Equal(t, 0, tag.Map.Size())

	tag.Map.Clear()
	r = bytes.NewReader(encoded)
	assert.Nil(t, tag.DeserializeValue(f, len(encoded), r))
	assert.Equal(t, 256, tag.Map.Size())
	for i := 0; i < len(sampleKeys); i++ {
		k := sampleKeys[i]
		v := sampleObjects[i]
		assert.Equal(t, k, tag.Map.Keys()[i])
		vc, ok := tag.Map.Get(k)
		assert.True(t, ok)
		assert.Equal(t, v, vc)
	}

	// Deserialization failures
	encoded = []byte{0x01, 0x11, 0x00, 0x11, 0x00} // 1 entry with 2 empty strings

	r = bytes.NewReader(encoded[:0])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))
	r = bytes.NewReader(encoded[:1])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))
	r = bytes.NewReader(encoded[:4])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))

	r = bytes.NewReader(encoded[:1])
	assert.Error(t, tag.DeserializeValue(f, len(encoded), r))

	// Too large
	r = bytes.NewReader(encoded)
	assert.Error(t, tag.DeserializeValue(f, len(encoded)+1, r))

	// More entries than bytes
	encoded = []byte{0x01, 0x11, 0x00, 0x00} // 1 entry with 1 empty and a null tag
	r = bytes.NewReader(encoded)
	assert.ErrorIs(t, tag.DeserializeValue(f, len(encoded)-1, r), ErrBadTagFormat)

	// Bad size
	assert.ErrorIs(t, tag.DeserializeValue(f, 0, r), ErrBadTagFormat)
}
