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

package ext

import (
	"bytes"
	"testing"
	"time"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/impl"
	"github.com/stretchr/testify/assert"
)

func TestNewTimestapTag(t *testing.T) {
	id := tags.TagID(1234)
	tag := NewTimestapTag(id)

	assert.Equal(t, id, tag.Id())
	assert.Equal(t, int64(0), tag.Payload)

	assert.Panics(t, func() {
		NewTimestapTag(tags.TagID(15))
	})
}

func TestTimestapTag(t *testing.T) {
	id := tags.TagID(1234)
	tag := NewTimestapTag(id)

	// Ensure it has the correct aggreations.
	_ = tag.ILTagHeaderImpl.Id()
	_ = tag.SignedILIntPayload.Payload
}

func TestTimestapTag_SetTimestamp(t *testing.T) {
	id := tags.TagID(1234)

	tag := NewTimestapTag(id)
	tag.SetTimestamp(time.UnixMilli(0))
	assert.Equal(t, int64(0), tag.Payload)

	tag.SetTimestamp(time.UnixMicro(-1234))
	assert.Equal(t, int64(-1234), tag.Payload)

	tag.SetTimestamp(time.UnixMicro(1234))
	assert.Equal(t, int64(1234), tag.Payload)

	now := time.Now()
	tag.SetTimestamp(now)
	assert.Equal(t, now.UnixMicro(), tag.Payload)
}

func TestTimestapTag_GetTimestamp(t *testing.T) {
	now := time.Now()

	id := tags.TagID(1234)
	tag := NewTimestapTag(id)

	tag.Payload = 0
	tse := time.UnixMicro(0)
	ts := tag.GetTimestamp()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.Payload = -123
	tse = time.UnixMicro(-123)
	ts = tag.GetTimestamp()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.Payload = 123
	tse = time.UnixMicro(123)
	ts = tag.GetTimestamp()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.SetTimestamp(now)
	ts = tag.GetTimestamp()
	assert.Equal(t, now.UnixMicro(), ts.UnixMicro())
	assert.Equal(t, now.Location(), ts.Location())
}

func TestTimestapTag_GetTimestampUTC(t *testing.T) {
	now := time.Now().UTC()

	id := tags.TagID(1234)
	tag := NewTimestapTag(id)

	tag.Payload = 0
	tse := time.UnixMicro(0).UTC()
	ts := tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.Payload = -123
	tse = time.UnixMicro(-123).UTC()
	ts = tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.Payload = 123
	tse = time.UnixMicro(123).UTC()
	ts = tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)
	assert.Equal(t, now.Location(), ts.Location())

	tag.SetTimestamp(now)
	ts = tag.GetTimestampUTC()
	assert.Equal(t, now.UnixMicro(), ts.UnixMicro())
	assert.Equal(t, now.Location(), ts.Location())
}

func TestSerializeTimestapTag(t *testing.T) {
	now := time.Now()

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeTimestapTag(12345, now, w))

	tag := NewTimestapTag(12345)
	r := bytes.NewReader(w.Bytes())
	assert.Nil(t, tags.ILTagDeserializeInto(nil, r, tag))
	assert.Equal(t, time.UnixMicro(now.UnixMicro()), tag.GetTimestamp())
	assert.Equal(t, now.UnixMicro(), tag.GetTimestamp().UnixMicro())
}

func TestDeerializeTimestapTag(t *testing.T) {
	now := time.Now()

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeTimestapTag(12345, now, w))

	ts, err := DeserializeTimestapTag(12345, bytes.NewReader(w.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, now.UnixMicro(), ts.UnixMicro())

	_, err = DeserializeTimestapTag(12347, bytes.NewReader(w.Bytes()))
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)
}

//------------------------------------------------------------------------------

func TestTimestampTZPayload(t *testing.T) {
	p := TimestampTZPayload{}

	var _ tags.ILTagPayload = &p
	_ = p.SignedILIntPayload.Payload
	_ = p.Offset
}

func TestTimestampTZPayload_ValueSize(t *testing.T) {
	p := TimestampTZPayload{}

	assert.Equal(t, uint64(1+2), p.ValueSize())

	p.Payload = 137438953472
	p.Offset = int16(1)
	assert.Equal(t, uint64(6+2), p.ValueSize())

	p.Payload = 9223372036854775807
	p.Offset = int16(32767)
	assert.Equal(t, uint64(9+2), p.ValueSize())

	p.Payload = -9223372036854775808
	p.Offset = int16(-32768)
	assert.Equal(t, uint64(9+2), p.ValueSize())
}

func TestTimestampTZPayload_SerializeValue(t *testing.T) {
	p := TimestampTZPayload{}

	w := bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{0, 0, 0}, w.Bytes())

	p.Payload = 137438953472
	p.Offset = int16(1)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{0xfc, 0x3f, 0xff, 0xff, 0xff, 0x8, 0x0, 0x1}, w.Bytes())

	p.Payload = 9223372036854775807
	p.Offset = int16(32767)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6, 0x7f, 0xff}, w.Bytes())

	p.Payload = -9223372036854775808
	p.Offset = int16(-32768)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0}, w.Bytes())

	// Simulate errors.
	p.Payload = 0
	p.Offset = 0
	w2 := DummyWriter{0}
	assert.Error(t, p.SerializeValue(&w2))

	p.Payload = 0
	p.Offset = 0
	w2 = DummyWriter{2}
	assert.Error(t, p.SerializeValue(&w2))

}

func TestTimestampTZPayload_DeserializeValue(t *testing.T) {
	f := impl.NewStandardTagFactory(true)
	p := TimestampTZPayload{}

	r := bytes.NewReader([]byte{0, 0, 0})
	assert.Nil(t, p.DeserializeValue(f, 3, r))
	assert.Equal(t, int64(0), p.Payload)
	assert.Equal(t, int16(0), p.Offset)

	r = bytes.NewReader([]byte{0xfc, 0x3f, 0xff, 0xff, 0xff, 0x8, 0x0, 0x1})
	assert.Nil(t, p.DeserializeValue(f, 8, r))
	assert.Equal(t, int64(137438953472), p.Payload)
	assert.Equal(t, int16(1), p.Offset)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x6, 0x7f, 0xff})
	assert.Nil(t, p.DeserializeValue(f, 11, r))
	assert.Equal(t, int64(9223372036854775807), p.Payload)
	assert.Equal(t, int16(32767), p.Offset)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0})
	assert.Nil(t, p.DeserializeValue(f, 11, r))
	assert.Equal(t, int64(-9223372036854775808), p.Payload)
	assert.Equal(t, int16(-32768), p.Offset)

	// Errors
	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0})
	assert.ErrorIs(t, p.DeserializeValue(f, 2, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0})
	assert.ErrorIs(t, p.DeserializeValue(f, 8, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0})
	assert.ErrorIs(t, p.DeserializeValue(f, 9, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0})
	assert.ErrorIs(t, p.DeserializeValue(f, 10, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80, 0x0, 0x00})
	assert.ErrorIs(t, p.DeserializeValue(f, 12, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7, 0x80})
	assert.ErrorIs(t, p.DeserializeValue(f, 11, r), tags.ErrBadTagFormat)
}

//------------------------------------------------------------------------------

func TestNewTimestapTZTag(t *testing.T) {
	id := tags.TagID(1234)
	tag := NewTimestapTZTag(id)

	assert.Equal(t, id, tag.Id())
	assert.Equal(t, int64(0), tag.Payload)
	assert.Equal(t, int16(0), tag.Offset)

	assert.Panics(t, func() {
		NewTimestapTZTag(tags.TagID(15))
	})
}

func TestTimestapTZTag(t *testing.T) {
	var tag TimestapTZTag

	_ = tag.ILTagHeaderImpl.Id()
	_ = tag.TimestampTZPayload.Payload
}

func TestTimestapTZTag_SetTimestamp(t *testing.T) {
	id := tags.TagID(1234)

	tag := NewTimestapTZTag(id)
	ts := time.UnixMicro(0).UTC()
	tag.SetTimestamp(ts)
	assert.Equal(t, int64(0), tag.Payload)
	assert.Equal(t, int16(0), tag.Offset)

	ts = time.UnixMicro(-1234).In(time.FixedZone("", 12344))
	tag.SetTimestamp(ts)
	assert.Equal(t, int64(-1234), tag.Payload)
	assert.Equal(t, int16(205), tag.Offset)

	ts = time.UnixMicro(1234).In(time.FixedZone("", -12345))
	tag.SetTimestamp(ts)
	assert.Equal(t, int64(1234), tag.Payload)
	assert.Equal(t, int16(-205), tag.Offset)

	now := time.Now()
	tag.SetTimestamp(now)
	assert.Equal(t, now.UnixMicro(), tag.Payload)
	_, offset := now.Zone()
	assert.Equal(t, int16(offset/60), tag.Offset)
}

func TestTimestapTZTag_GetTimestamp(t *testing.T) {
	now := time.Now()

	id := tags.TagID(1234)
	tag := NewTimestapTZTag(id)

	tag.Payload = 0
	tag.Offset = 0
	tse := time.UnixMicro(0).In(time.FixedZone("", 0))
	ts := tag.GetTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = -123
	tag.Offset = -1234
	tse = time.UnixMicro(-123).In(time.FixedZone("", -1234*60))
	ts = tag.GetTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = 123
	tag.Offset = 1234
	tse = time.UnixMicro(123).In(time.FixedZone("", 1234*60))
	ts = tag.GetTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = now.UnixMicro()
	_, offs := now.Zone()
	tag.Offset = int16(offs / 60)
	tse = time.UnixMicro(now.UnixMicro()).In(time.FixedZone("", offs))
	ts = tag.GetTimestamp()
	assert.Equal(t, tse, ts)
}

func TestTimestapTZTag_GetLocalTimestamp(t *testing.T) {
	now := time.Now()

	id := tags.TagID(1234)
	tag := NewTimestapTZTag(id)

	tag.Payload = 0
	tag.Offset = 0
	tse := time.UnixMicro(0)
	ts := tag.GetLocalTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = -123
	tag.Offset = -1234
	tse = time.UnixMicro(-123)
	ts = tag.GetLocalTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = 123
	tag.Offset = 1234
	tse = time.UnixMicro(123)
	ts = tag.GetLocalTimestamp()
	assert.Equal(t, tse, ts)

	tag.Payload = now.UnixMicro()
	_, offs := now.Zone()
	tag.Offset = int16(offs / 60)
	tse = time.UnixMicro(now.UnixMicro())
	ts = tag.GetLocalTimestamp()
	assert.Equal(t, tse, ts)
}

func TestTimestapTZTag_GetTimestampUTC(t *testing.T) {
	now := time.Now().UTC()

	id := tags.TagID(1234)
	tag := NewTimestapTZTag(id)

	tag.Payload = 0
	tag.Offset = 0
	tse := time.UnixMicro(0).UTC()
	ts := tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)

	tag.Payload = -123
	tag.Offset = -1234
	tse = time.UnixMicro(-123).UTC()
	ts = tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)

	tag.Payload = 123
	tag.Offset = 1234
	tse = time.UnixMicro(123).UTC()
	ts = tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)

	tag.Payload = now.UnixMicro()
	_, offs := now.Zone()
	tag.Offset = int16(offs / 60)
	tse = time.UnixMicro(now.UnixMicro()).UTC()
	ts = tag.GetTimestampUTC()
	assert.Equal(t, tse, ts)
}
