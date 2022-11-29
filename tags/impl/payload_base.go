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
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
)

//------------------------------------------------------------------------------

// Implementation a basic string payload
type StringPayload struct {
	Payload string
}

// Implementation of ILTagPayload.ValueSize().
func (p *StringPayload) ValueSize() uint64 {
	return uint64(len(p.Payload))
}

// Implementation of ILTagPayload.SerializeValue()
func (p *StringPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteString(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *StringPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 0 {
		return tags.ErrBadTagFormat
	} else if valueSize == 0 {
		p.Payload = ""
		return nil
	} else {
		if s, err := serialization.ReadString(reader, valueSize); err == nil {
			p.Payload = s
			return nil
		} else {
			return err
		}
	}
}

//------------------------------------------------------------------------------

// Implementation of the big int payload.
type BigIntPayload struct {
	Payload []byte
}

// Implementation of ILTagPayload.ValueSize().
func (p *BigIntPayload) ValueSize() uint64 {
	if l := len(p.Payload); l > 0 {
		return uint64(l)
	} else {
		return 1
	}
}

// Implementation of ILTagPayload.SerializeValue()
func (p *BigIntPayload) SerializeValue(writer io.Writer) error {
	if len(p.Payload) != 0 {
		return serialization.WriteBytes(writer, p.Payload)
	} else {
		return serialization.WriteBytes(writer, []byte{0})
	}
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *BigIntPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return tags.ErrBadTagFormat
	}
	p.Payload = make([]byte, valueSize)
	return serialization.ReadBytes(reader, p.Payload)
}

//------------------------------------------------------------------------------

// Implementation of the big decimal payload.
type BigDecPayload struct {
	BigIntPayload
	Scale int32
}

// Implementation of ILTagPayload.ValueSize().
func (p *BigDecPayload) ValueSize() uint64 {
	return uint64(p.BigIntPayload.ValueSize() + 4)
}

// Implementation of ILTagPayload.SerializeValue()
func (p *BigDecPayload) SerializeValue(writer io.Writer) error {
	if err := p.BigIntPayload.SerializeValue(writer); err != nil {
		return err
	}
	return serialization.WriteInt32(writer, p.Scale)
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *BigDecPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 5 {
		return tags.ErrBadTagFormat
	}
	if err := p.BigIntPayload.DeserializeValue(factory, valueSize-4, reader); err != nil {
		return err
	}
	if v, err := serialization.ReadInt32(reader); err != nil {
		return err
	} else {
		p.Scale = v
		return nil
	}
}

//------------------------------------------------------------------------------

// Implementation of the ILInt array payload. It can also be used to implement
// the ITU Object Identifier.
type ILIntArrayPayload struct {
	Payload []uint64
}

// Implementation of ILTagPayload.ValueSize().
func (p *ILIntArrayPayload) ValueSize() uint64 {
	size := ilint.EncodedSize(uint64(len(p.Payload)))
	for _, v := range p.Payload {
		size += ilint.EncodedSize(v)
	}
	return uint64(size)
}

// Implementation of ILTagPayload.SerializeValue()
func (p *ILIntArrayPayload) SerializeValue(writer io.Writer) error {
	if len(p.Payload) == 0 {
		return serialization.WriteUInt8(writer, 0)
	}
	if err := serialization.WriteILInt(writer, uint64(len(p.Payload))); err != nil {
		return err
	}
	for _, v := range p.Payload {
		if err := serialization.WriteILInt(writer, v); err != nil {
			return err
		}
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *ILIntArrayPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return tags.ErrBadTagFormat
	}
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	a, err := p.deserializeValueCore(r)
	if err != nil {
		return err
	}
	if r.N == 0 {
		p.Payload = a
		return nil
	} else {
		return tags.ErrBadTagFormat
	}
}

func (p *ILIntArrayPayload) deserializeValueCore(reader *io.LimitedReader) ([]uint64, error) {
	// Read the size first
	size, err := serialization.ReadILInt(reader)
	if err != nil {
		return nil, err
	}
	// Check the format. All 1 byte ILInt
	if size > uint64(reader.N) {
		return nil, tags.ErrBadTagFormat
	}
	a := make([]uint64, int(size))
	for i := 0; i < len(a); i++ {
		if v, err := serialization.ReadILInt(reader); err != nil {
			return nil, err
		} else {
			a[i] = v
		}
	}
	return a, nil
}

//------------------------------------------------------------------------------

// Implementation of the ILTag array payload.
type ILTagArrayPayload struct {
	Payload []tags.ILTag
}

// Implementation of ILTagPayload.ValueSize().
func (p *ILTagArrayPayload) ValueSize() uint64 {
	if len(p.Payload) == 0 {
		return 1
	}
	size := ilint.EncodedSize(uint64(len(p.Payload)))
	for _, v := range p.Payload {
		size += int(tags.ILTagSize(v))
	}
	return uint64(size)
}

// Implementation of ILTagPayload.SerializeValue()
func (p *ILTagArrayPayload) SerializeValue(writer io.Writer) error {
	if p.Payload == nil {
		return serialization.WriteUInt8(writer, 0)
	}
	if err := serialization.WriteILInt(writer, uint64(len(p.Payload))); err != nil {
		return err
	}
	for _, v := range p.Payload {
		if err := tags.ILTagSeralize(v, writer); err != nil {
			return err
		}
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *ILTagArrayPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return tags.ErrBadTagFormat
	}
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	a, err := p.deserializeValueCore(factory, r)
	if err != nil {
		return err
	}
	if r.N == 0 {
		p.Payload = a
		return nil
	} else {
		return tags.ErrBadTagFormat
	}
}

func (p *ILTagArrayPayload) deserializeValueCore(factory tags.ILTagFactory,
	reader *io.LimitedReader) ([]tags.ILTag, error) {
	// Read the size first
	size, err := serialization.ReadILInt(reader)
	if err != nil {
		return nil, err
	}
	// Basic format check. All null tags
	if size > uint64(reader.N) {
		return nil, tags.ErrBadTagFormat
	}
	a := make([]tags.ILTag, int(size))
	for i := 0; i < len(a); i++ {
		if v, err := tags.ILTagDeserialize(factory, reader); err != nil {
			return nil, err
		} else {
			a[i] = v
		}
	}
	return a, nil
}

//------------------------------------------------------------------------------

// Implementation of the ILTag sequence payload.
type ILTagSequencePayload struct {
	Payload []tags.ILTag
}

// Implementation of ILTagPayload.ValueSize().
func (p *ILTagSequencePayload) ValueSize() uint64 {
	if p.Payload == nil {
		return 0
	}
	size := 0
	for _, v := range p.Payload {
		size += int(tags.ILTagSize(v))
	}
	return uint64(size)
}

// Implementation of ILTagPayload.SerializeValue()
func (p *ILTagSequencePayload) SerializeValue(writer io.Writer) error {
	if p.Payload == nil {
		return nil
	}
	for _, v := range p.Payload {
		if err := tags.ILTagSeralize(v, writer); err != nil {
			return err
		}
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *ILTagSequencePayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 0 {
		return tags.ErrBadTagFormat
	} else if valueSize == 0 {
		p.Payload = make([]tags.ILTag, 0)
		return nil
	} else {
		r := io.LimitedReader{R: reader, N: int64(valueSize)}
		a := make([]tags.ILTag, 0, 16)
		for {
			if v, err := tags.ILTagDeserialize(factory, &r); err != nil {
				return err
			} else {
				a = append(a, v)
			}
			if r.N == 0 {
				p.Payload = a
				return nil
			}
		}
	}
}

//------------------------------------------------------------------------------

// Implementation of the range payload.
type RangePayload struct {
	Start uint64
	Count uint16
}

// Implementation of ILTagPayload.ValueSize().
func (p *RangePayload) ValueSize() uint64 {
	return uint64(ilint.EncodedSize(p.Start)) + 2
}

// Implementation of ILTagPayload.SerializeValue()
func (p *RangePayload) SerializeValue(writer io.Writer) error {
	if err := serialization.WriteILInt(writer, p.Start); err != nil {
		return err
	}
	if err := serialization.WriteUInt16(writer, p.Count); err != nil {
		return err
	} else {
		return nil
	}
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *RangePayload) DeserializeValue(factory tags.ILTagFactory,
	valueSize int, reader io.Reader) error {
	if valueSize < 3 {
		return tags.ErrBadTagFormat
	}
	r := io.LimitedReader{R: reader, N: int64(valueSize)}
	// Read start
	s, err := serialization.ReadILInt(&r)
	if err != nil {
		return err
	}
	// Read count
	c, err := serialization.ReadUInt16(&r)
	if err != nil {
		return err
	}
	if r.N == 0 {
		p.Start = s
		p.Count = c
		return nil
	} else {
		return tags.ErrBadTagFormat
	}
}

//------------------------------------------------------------------------------

// Implementation of the version payload.
type VersionPayload struct {
	Major    int32
	Minor    int32
	Revision int32
	Build    int32
}

// Implementation of ILTagPayload.ValueSize().
func (p *VersionPayload) ValueSize() uint64 {
	return uint64(4 + 4 + 4 + 4)
}

// Implementation of ILTagPayload.SerializeValue()
func (p *VersionPayload) SerializeValue(writer io.Writer) error {
	if err := serialization.WriteInt32(writer, p.Major); err != nil {
		return err
	}
	if err := serialization.WriteInt32(writer, p.Minor); err != nil {
		return err
	}
	if err := serialization.WriteInt32(writer, p.Revision); err != nil {
		return err
	}
	if err := serialization.WriteInt32(writer, p.Build); err != nil {
		return err
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *VersionPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 16 {
		return tags.ErrBadTagFormat
	}
	major, err := serialization.ReadInt32(reader)
	if err != nil {
		return err
	}
	minor, err := serialization.ReadInt32(reader)
	if err != nil {
		return err
	}
	revision, err := serialization.ReadInt32(reader)
	if err != nil {
		return err
	}
	build, err := serialization.ReadInt32(reader)
	if err != nil {
		return err
	}
	p.Major = major
	p.Minor = minor
	p.Revision = revision
	p.Build = build
	return nil
}

//------------------------------------------------------------------------------

// Implementation of the version payload.
type StringDictionaryPayload struct {
	Map StableStringMap
}

// Implementation of ILTagPayload.ValueSize().
func (p *StringDictionaryPayload) ValueSize() uint64 {
	size := uint64(ilint.EncodedSize(uint64(p.Map.Size())))
	for _, e := range p.Map.Entries() {
		size += StdStringTagSize(e.Key)
		size += StdStringTagSize(e.Value)
	}
	return size
}

// Implementation of ILTagPayload.SerializeValue()
func (p *StringDictionaryPayload) SerializeValue(writer io.Writer) error {

	if err := serialization.WriteILInt(writer, uint64(p.Map.Size())); err != nil {
		return err
	}
	for _, e := range p.Map.Entries() {
		if err := SerializeStdStringTag(e.Key, writer); err != nil {
			return err
		}
		if err := SerializeStdStringTag(e.Value, writer); err != nil {
			return err
		}
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *StringDictionaryPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return tags.ErrBadTagFormat
	}
	p.Map.Clear()
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	// Deserialize...
	if err := p.deserializeValueCore(r); err != nil {
		return err
	}
	// Check if something is left in the payload
	if r.N != 0 {
		return tags.ErrBadTagFormat
	}
	return nil
}

func (p *StringDictionaryPayload) deserializeValueCore(reader *io.LimitedReader) error {
	// Read size and see if it can be used
	size, err := serialization.ReadILInt(reader)
	if err != nil {
		return tags.ErrBadTagFormat
	}
	if size > uint64(reader.N/(2+2)) {
		// The maximum number of entries that can fit in N bytes is
		// remaining / (minStringTagSize + minStringTagSize) where
		// minStringTagSize is 2 (standard string tag size storing "")
		return tags.ErrBadTagFormat
	}
	for i := 0; i < int(size); i++ {
		k, err := DeserializeStdStringTag(reader)
		if err != nil {
			return err
		}
		v, err := DeserializeStdStringTag(reader)
		if err != nil {
			return err
		}
		p.Map.Put(k, v)
	}
	return nil
}

//------------------------------------------------------------------------------

// Implementation of the version payload.
type DictionaryPayload struct {
	Map StableILTagMap
}

// Implementation of ILTagPayload.ValueSize().
func (p *DictionaryPayload) ValueSize() uint64 {
	size := uint64(ilint.EncodedSize(uint64(p.Map.Size())))
	for _, k := range p.Map.Keys() {
		size += StdStringTagSize(k)
		t, _ := p.Map.Get(k)
		size += tags.ILTagSize(t)
	}
	return size
}

// Implementation of ILTagPayload.SerializeValue()
func (p *DictionaryPayload) SerializeValue(writer io.Writer) error {
	if err := serialization.WriteILInt(writer, uint64(p.Map.Size())); err != nil {
		return err
	}
	for _, k := range p.Map.Keys() {
		if err := SerializeStdStringTag(k, writer); err != nil {
			return err
		}
		t, _ := p.Map.Get(k)
		if err := tags.ILTagSeralize(t, writer); err != nil {
			return err
		}
	}
	return nil

}

// Implementation of ILTagPayload.DeserializeValue()
func (p *DictionaryPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return tags.ErrBadTagFormat
	}
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	// Deserialize...
	if err := p.deserializeValueCore(factory, r); err != nil {
		return err
	}
	// Check if something is left in the payload
	if r.N != 0 {
		return tags.ErrBadTagFormat
	}
	return nil
}

func (p *DictionaryPayload) deserializeValueCore(factory tags.ILTagFactory, reader *io.LimitedReader) error {
	// Read size and see if it can be used
	size, err := serialization.ReadILInt(reader)
	if err != nil {
		return tags.ErrBadTagFormat
	}
	if size > uint64(reader.N/(2+1)) {
		// The maximum number of entries that can fit in N bytes is
		// remaining / (minStringTagSize + minTagSize) where
		// minStringTagSize is 2 and minTagSize is 1 (standard string tag size storing "")
		return tags.ErrBadTagFormat
	}
	p.Map.Clear()
	for i := 0; i < int(size); i++ {
		k, err := DeserializeStdStringTag(reader)
		if err != nil {
			return err
		}
		t, err := tags.ILTagDeserialize(factory, reader)
		if err != nil {
			return err
		}
		p.Map.Put(k, t)
	}
	return nil
}
