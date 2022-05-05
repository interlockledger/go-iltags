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
	"math/big"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/utils"
)

//------------------------------------------------------------------------------

// Implementation of the raw payload.
type RawPayload struct {
	Payload []byte
}

func (p *RawPayload) ValueSize() uint64 {
	if p.Payload != nil {
		return uint64(len(p.Payload))
	} else {
		return 0
	}
}

func (p *RawPayload) SerializeValue(writer io.Writer) error {
	if p.Payload != nil {
		return serialization.WriteBytes(writer, p.Payload)
	}
	return nil
}

func (p *RawPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	p.Payload = make([]byte, valueSize)
	if valueSize > 0 {
		return serialization.ReadBytes(reader, p.Payload)
	} else {
		return nil
	}
}

//------------------------------------------------------------------------------

// Implementation of the big int payload. The serialization and deserialization
// methods of this payload will properly cleanup the temporary buffers whenever
// possible.
type BigIntPayload struct {
	Payload big.Int
}

func (p *BigIntPayload) ValueSize() uint64 {
	return uint64((p.Payload.BitLen() + 7) / 8)
}

func (p *BigIntPayload) SerializeValue(writer io.Writer) error {
	tmp := p.Payload.Bytes()
	defer utils.ShredBytes(tmp)
	return serialization.WriteBytes(writer, tmp)
}

func (p *BigIntPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return ErrBadTagFormat
	}
	tmp := make([]byte, valueSize)
	defer utils.ShredBytes(tmp)
	err := serialization.ReadBytes(reader, tmp)
	if err != nil {
		return err
	}
	p.Payload.SetBytes(tmp)
	return nil
}

//------------------------------------------------------------------------------

// Implementation of the big decimal payload. The serialization and deserialization
// methods of this payload will properly cleanup the temporary buffers whenever
// possible.
type BigDecPayload struct {
	BigIntPayload
	Scale int32
}

func (p *BigDecPayload) ValueSize() uint64 {
	return uint64(p.BigIntPayload.ValueSize() + 4)
}

func (p *BigDecPayload) SerializeValue(writer io.Writer) error {
	if err := p.BigIntPayload.SerializeValue(writer); err != nil {
		return err
	}
	return serialization.WriteInt32(writer, p.Scale)
}

func (p *BigDecPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 5 {
		return ErrBadTagFormat
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

func (p *ILIntArrayPayload) ValueSize() uint64 {
	if p.Payload == nil {
		return 0
	}
	size := ilint.EncodedSize(uint64(len(p.Payload)))
	for _, v := range p.Payload {
		size += ilint.EncodedSize(v)
	}
	return uint64(size)
}

func (p *ILIntArrayPayload) SerializeValue(writer io.Writer) error {
	if p.Payload == nil {
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

func (p *ILIntArrayPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return ErrBadTagFormat
	}
	r := io.LimitedReader{R: reader, N: int64(valueSize)}
	// Read the size first
	size, err := serialization.ReadILInt(&r)
	if err != nil {
		return err
	}
	if size > uint64(valueSize) {
		return ErrBadTagFormat
	}
	a := make([]uint64, int(size))
	for i := 0; i < len(a); i++ {
		if v, err := serialization.ReadILInt(&r); err != nil {
			return err
		} else {
			a[i] = v
		}
	}
	if r.N == 0 {
		p.Payload = a
		return nil
	} else {
		return ErrBadTagFormat
	}
}

//------------------------------------------------------------------------------

// Implementation of the ILTag array payload.
type ILTagArrayPayload struct {
	Payload []ILTag
}

func (p *ILTagArrayPayload) ValueSize() uint64 {
	if p.Payload == nil {
		return 0
	}
	size := ilint.EncodedSize(uint64(len(p.Payload)))
	for _, v := range p.Payload {
		size += int(ILTagSize(v))
	}
	return uint64(size)
}

func (p *ILTagArrayPayload) SerializeValue(writer io.Writer) error {
	if p.Payload == nil {
		return serialization.WriteUInt8(writer, 0)
	}
	if err := serialization.WriteILInt(writer, uint64(len(p.Payload))); err != nil {
		return err
	}
	for _, v := range p.Payload {
		if err := ILTagSeralize(v, writer); err != nil {
			return err
		}
	}
	return nil
}

func (p *ILTagArrayPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 1 {
		return ErrBadTagFormat
	}
	r := io.LimitedReader{R: reader, N: int64(valueSize)}
	// Read the size first
	size, err := serialization.ReadILInt(&r)
	if err != nil {
		return err
	}
	if size > uint64(valueSize) {
		return ErrBadTagFormat
	}
	a := make([]ILTag, int(size))
	for i := 0; i < len(a); i++ {
		if v, err := factory.Deserialize(&r); err != nil {
			return err
		} else {
			a[i] = v
		}
	}
	if r.N == 0 {
		p.Payload = a
		return nil
	} else {
		return ErrBadTagFormat
	}
}

//------------------------------------------------------------------------------

// Implementation of the ILTag sequence payload.
type ILTagSequencePayload struct {
	Payload []ILTag
}

func (p *ILTagSequencePayload) ValueSize() uint64 {
	if p.Payload == nil {
		return 0
	}
	size := 0
	for _, v := range p.Payload {
		size += int(ILTagSize(v))
	}
	return uint64(size)
}

func (p *ILTagSequencePayload) SerializeValue(writer io.Writer) error {
	if p.Payload == nil {
		return nil
	}
	for _, v := range p.Payload {
		if err := ILTagSeralize(v, writer); err != nil {
			return err
		}
	}
	return nil
}

func (p *ILTagSequencePayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	r := io.LimitedReader{R: reader, N: int64(valueSize)}

	a := make([]ILTag, 0, 16)
	for {
		if v, err := factory.Deserialize(&r); err != nil {
			return err
		} else {
			a = append(a, v)
		}
		if r.N == 0 {
			break
		}
	}
	p.Payload = a
	return nil
}

//------------------------------------------------------------------------------

// Implementation of the range payload.
type RangePayload struct {
	Start uint64
	Count uint16
}

func (p *RangePayload) ValueSize() uint64 {
	return uint64(ilint.EncodedSize(p.Start)) + 4
}

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

func (p *RangePayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 5 {
		return ErrBadTagFormat
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
		return ErrBadTagFormat
	}
}
