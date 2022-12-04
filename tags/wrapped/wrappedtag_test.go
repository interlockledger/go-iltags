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

package wrapped

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/tagtest"
	"github.com/stretchr/testify/assert"
)

func TestWrappedValueTagsSize(t *testing.T) {
	t1 := NewSampleWrappedTag()

	v1 := new(uint32)
	*v1 = 0x01234567
	v2 := new(uint32)
	*v2 = 0x089ABCDEF

	assert.Equal(t, uint64(8), WrappedValueTagsSize[uint32](t1, v1))
	assert.Nil(t, t1.Value)

	assert.Equal(t, uint64(16), WrappedValueTagsSize[uint32](t1, v1, v2))
	assert.Nil(t, t1.Value)
}

func TestWrappedValueTagsSizeOrNull(t *testing.T) {
	t1 := NewSampleWrappedTag()

	v1 := new(uint32)
	*v1 = 0x01234567
	v2 := new(uint32)
	*v2 = 0x089ABCDEF

	assert.Equal(t, uint64(8), WrappedValueTagsSizeOrNull[uint32](t1, v1))
	assert.Nil(t, t1.Value)

	assert.Equal(t, uint64(16), WrappedValueTagsSizeOrNull[uint32](t1, v1, v2))
	assert.Nil(t, t1.Value)

	assert.Equal(t, uint64(17), WrappedValueTagsSizeOrNull[uint32](t1, v1, nil, v2))
	assert.Nil(t, t1.Value)
}

func TestSerializeWrappedValueTags(t *testing.T) {
	t1 := NewSampleWrappedTag()

	v1 := new(uint32)
	*v1 = 0x01234567
	v2 := new(uint32)
	*v2 = 0x089ABCDEF

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeWrappedValueTags[uint32](t1, w, v1))
	assert.Nil(t, t1.Value)
	assert.Equal(t, []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeWrappedValueTags[uint32](t1, w, v1, v2))
	assert.Nil(t, t1.Value)
	assert.Equal(t, []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}, w.Bytes())

	wl := tagtest.NewLimitedWriter(15, false)
	assert.ErrorIs(t, SerializeWrappedValueTags[uint32](t1, wl, v1, v2), io.ErrShortWrite)
	assert.Nil(t, t1.Value)
}

func TestSerializeWrappedValueTagsOrNull(t *testing.T) {
	t1 := NewSampleWrappedTag()

	v1 := new(uint32)
	*v1 = 0x01234567
	v2 := new(uint32)
	*v2 = 0x089ABCDEF

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeWrappedValueTagsOrNull[uint32](t1, w, v1))
	assert.Nil(t, t1.Value)
	assert.Equal(t, []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeWrappedValueTagsOrNull[uint32](t1, w, v1, v2))
	assert.Nil(t, t1.Value)
	assert.Equal(t, []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeWrappedValueTagsOrNull[uint32](t1, w, v1, nil, v2, nil))
	assert.Nil(t, t1.Value)
	assert.Equal(t, []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0x0,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF,
		0x0}, w.Bytes())

	wl := tagtest.NewLimitedWriter(15, false)
	assert.ErrorIs(t, SerializeWrappedValueTagsOrNull[uint32](t1, wl, v1, v2), io.ErrShortWrite)
	assert.Nil(t, t1.Value)

	wl = tagtest.NewLimitedWriter(16, false)
	assert.ErrorIs(t, SerializeWrappedValueTagsOrNull[uint32](t1, wl, v1, v2, nil), io.ErrShortWrite)
	assert.Nil(t, t1.Value)
}

func TestDeserializeWrappedValueTags(t *testing.T) {
	t1 := NewSampleWrappedTag()

	bin := []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}
	r := bytes.NewReader(bin)
	l, err := DeserializeWrappedValueTags[uint32](nil, t1, len(bin), r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x01234567), *l[0])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(bin)
	l, err = DeserializeWrappedValueTags[uint32](nil, t1, len(bin), r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 2)
	assert.Equal(t, uint32(0x01234567), *l[0])
	assert.Equal(t, uint32(0x089ABCDEF), *l[1])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(bin)
	l, err = DeserializeWrappedValueTags[uint32](nil, t1, 8, r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x01234567), *l[0])

	l, err = DeserializeWrappedValueTags[uint32](nil, t1, 8, r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x089ABCDEF), *l[0])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}
	r = bytes.NewReader(bin)
	_, err = DeserializeWrappedValueTags[uint32](nil, t1, 7, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
	assert.Nil(t, t1.Value)
}

func TestDeserializeWrappedValueTagsOrNull(t *testing.T) {
	t1 := NewSampleWrappedTag()

	bin := []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}
	r := bytes.NewReader(bin)
	l, err := DeserializeWrappedValueTagsOrNull[uint32](nil, t1, len(bin), r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x01234567), *l[0])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(bin)
	l, err = DeserializeWrappedValueTagsOrNull[uint32](nil, t1, len(bin), r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 2)
	assert.Equal(t, uint32(0x01234567), *l[0])
	assert.Equal(t, uint32(0x089ABCDEF), *l[1])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0x0,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF,
		0x0}
	r = bytes.NewReader(bin)
	l, err = DeserializeWrappedValueTagsOrNull[uint32](nil, t1, len(bin), r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 4)
	assert.Equal(t, uint32(0x01234567), *l[0])
	assert.Nil(t, l[1])
	assert.Equal(t, uint32(0x089ABCDEF), *l[2])
	assert.Nil(t, l[3])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67,
		0xf9, 0x3, 0xda, 0x4, 0x89, 0xAB, 0xCD, 0xEF}
	r = bytes.NewReader(bin)
	l, err = DeserializeWrappedValueTagsOrNull[uint32](nil, t1, 8, r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x01234567), *l[0])

	l, err = DeserializeWrappedValueTagsOrNull[uint32](nil, t1, 8, r)
	assert.Nil(t, err)
	assert.Nil(t, t1.Value)
	assert.Len(t, l, 1)
	assert.Equal(t, uint32(0x089ABCDEF), *l[0])

	bin = []byte{
		0xf9, 0x3, 0xda, 0x4, 0x1, 0x23, 0x45, 0x67}
	r = bytes.NewReader(bin)
	_, err = DeserializeWrappedValueTagsOrNull[uint32](nil, t1, 7, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
	assert.Nil(t, t1.Value)
}
