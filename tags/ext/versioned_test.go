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
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

type DummyVersionedPayloadData struct {
	Value           uint64
	ReceivedVersion uint16
}

func (d *DummyVersionedPayloadData) Version() uint16 {
	return 1
}

func (d *DummyVersionedPayloadData) SupportedVersion(version uint16) bool {
	return version == 0 || version == 1
}

func (d *DummyVersionedPayloadData) Size() uint64 {
	return 8
}

func (d *DummyVersionedPayloadData) Serialize(writer io.Writer) error {
	return serialization.WriteUInt64(writer, d.Value)
}

func (d *DummyVersionedPayloadData) Deserialize(version uint16, factory tags.ILTagFactory, valueSize int,
	reader *io.LimitedReader) error {
	d.ReceivedVersion = version
	if v, err := serialization.ReadUInt64(reader); err != nil {
		return err
	} else {
		d.Value = v
	}
	return nil
}

func TestDummyVersionedPayloadData(t *testing.T) {
	var _ VersionedPayloadData = (*DummyVersionedPayloadData)(nil)
}

func TestDummyVersionedPayloadData_Version(t *testing.T) {
	var d DummyVersionedPayloadData
	assert.Equal(t, uint16(1), d.Version())
}

func TestDummyVersionedPayloadData_SupportedVersion(t *testing.T) {
	var d DummyVersionedPayloadData
	assert.True(t, d.SupportedVersion(0))
	assert.True(t, d.SupportedVersion(1))
	assert.False(t, d.SupportedVersion(2))
}

func TestDummyVersionedPayloadData_Size(t *testing.T) {
	var d DummyVersionedPayloadData
	assert.Equal(t, uint64(8), d.Size())
}

func TestDummyVersionedPayloadData_Serialize(t *testing.T) {
	var d DummyVersionedPayloadData

	d.Value = 0x0123456789ABCDEF
	w := bytes.NewBuffer(nil)
	assert.Nil(t, d.Serialize(w))
	assert.Equal(t, []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}, w.Bytes())
}

func TestDummyVersionedPayloadData_Deserialize(t *testing.T) {
	serialized := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	var d DummyVersionedPayloadData

	d.Value = 0x0
	r := &io.LimitedReader{R: bytes.NewReader(serialized), N: 8}
	assert.Nil(t, d.Deserialize(0, nil, 8, r))
	assert.Equal(t, uint64(0x0123456789ABCDEF), d.Value)
	assert.Equal(t, int64(0), r.N)

	d.Value = 0x0
	r = &io.LimitedReader{R: bytes.NewReader(serialized), N: 8}
	assert.Nil(t, d.Deserialize(1, nil, 8, r))
	assert.Equal(t, uint64(0x0123456789ABCDEF), d.Value)
	assert.Equal(t, int64(0), r.N)

	d.Value = 0x0
	r = &io.LimitedReader{R: bytes.NewReader(serialized), N: 7}
	assert.Error(t, d.Deserialize(1, nil, 8, r))
}

// -----------------------------------------------------------------------------
type BuggyDummyVersionedPayloadData struct {
	DummyVersionedPayloadData
}

func (d *BuggyDummyVersionedPayloadData) Deserialize(version uint16, factory tags.ILTagFactory, valueSize int,
	reader *io.LimitedReader) error {
	if !d.SupportedVersion(version) {
		return tags.ErrBadTagFormat
	}
	return nil
}

func TestBuggyDummyVersionedPayloadData(t *testing.T) {
	var _ VersionedPayloadData = (*BuggyDummyVersionedPayloadData)(nil)
}

func TestBuggyDummyVersionedPayloadData_Deserialize(t *testing.T) {
	serialized := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	var d BuggyDummyVersionedPayloadData

	d.Value = 0x0
	r := &io.LimitedReader{R: bytes.NewReader(serialized), N: 8}
	assert.Nil(t, d.Deserialize(0, nil, 8, r))
	assert.Equal(t, uint64(0), d.Value)
	assert.Equal(t, int64(8), r.N)
}

//------------------------------------------------------------------------------

func TestVersionedPayload(t *testing.T) {
	var _ tags.ILTagPayload = (*VersionedPayload[*DummyVersionedPayloadData])(nil)
}

func TestVersionedPayload_ValueSize(t *testing.T) {
	p := &VersionedPayload[*DummyVersionedPayloadData]{}

	assert.Equal(t, uint64(2+8), p.ValueSize())
}

func TestVersionedPayload_SerializeValue(t *testing.T) {
	p := &VersionedPayload[*DummyVersionedPayloadData]{}

	p.Data = &DummyVersionedPayloadData{}
	p.Data.Value = 0x0123456789ABCDEF
	w := bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
		w.Bytes())

	dw := &DummyWriter{Limit: 1}
	assert.Error(t, p.SerializeValue(dw))

	dw = &DummyWriter{Limit: 9}
	assert.Error(t, p.SerializeValue(dw))
}

func TestVersionedPayload_DeserializeValue(t *testing.T) {
	p := &VersionedPayload[*DummyVersionedPayloadData]{Data: &DummyVersionedPayloadData{}}
	p.Data.Value = 0

	r := bytes.NewReader([]byte{0x00, 0x00, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	assert.Nil(t, p.DeserializeValue(nil, 10, r))
	assert.Equal(t, uint64(0x0123456789ABCDEF), p.Data.Value)
	assert.Equal(t, uint16(0), p.Data.ReceivedVersion)

	r = bytes.NewReader([]byte{0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	assert.Nil(t, p.DeserializeValue(nil, 10, r))
	assert.Equal(t, uint64(0x0123456789ABCDEF), p.Data.Value)
	assert.Equal(t, uint16(1), p.Data.ReceivedVersion)

	r = bytes.NewReader([]byte{0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	assert.ErrorIs(t, p.DeserializeValue(nil, 1, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0x00})
	assert.ErrorIs(t, p.DeserializeValue(nil, 10, r), io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	assert.ErrorIs(t, p.DeserializeValue(nil, 9, r), io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0x00, 0xff, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	assert.ErrorIs(t, p.DeserializeValue(nil, 10, r), tags.ErrBadTagFormat)

	r = bytes.NewReader([]byte{0x00, 0x01, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF})
	p1 := &VersionedPayload[*BuggyDummyVersionedPayloadData]{
		Data: &BuggyDummyVersionedPayloadData{}}
	p1.Data.Value = 0
	assert.ErrorIs(t, p1.DeserializeValue(nil, 10, r), tags.ErrBadTagFormat)
}

//------------------------------------------------------------------------------

type DummyVersionedTag = VersionedPayloadTag[*DummyVersionedPayloadData]

func TestNewVersionedPayloadTag(t *testing.T) {
	var tag *DummyVersionedTag = NewVersionedPayloadTag(16, &DummyVersionedPayloadData{})
	var _ tags.ILTag = tag
	assert.Equal(t, tags.TagID(16), tag.Id())

	assert.Panics(t, func() {
		NewVersionedPayloadTag(15, &DummyVersionedPayloadData{})
	})
}
