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

package tags

import (
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
)

// Implementation of the null payload.
type NullPayload struct {
}

func (p *NullPayload) ValueSize() uint64 {
	return 0
}

func (p *NullPayload) SerializeValue(writer io.Writer) error {
	return nil
}

func (p *NullPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	return nil
}

//------------------------------------------------------------------------------

// Implementation a basic boolean payload
type BoolPayload struct {
	Payload bool
}

func (p *BoolPayload) ValueSize() uint64 {
	return 1
}

func (p *BoolPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteBool(writer, p.Payload)
}

func (p *BoolPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 1 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadBool(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic uint8 payload
type UInt8Payload struct {
	Payload uint8
}

func (p *UInt8Payload) ValueSize() uint64 {
	return 1
}

func (p *UInt8Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt8(writer, p.Payload)
}

func (p *UInt8Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 1 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadUInt8(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic int8 payload
type Int8Payload struct {
	Payload int8
}

func (p *Int8Payload) ValueSize() uint64 {
	return 1
}

func (p *Int8Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt8(writer, p.Payload)
}

func (p *Int8Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 1 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadInt8(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic uint16 payload
type UInt16Payload struct {
	Payload uint16
}

func (p *UInt16Payload) ValueSize() uint64 {
	return 2
}

func (p *UInt16Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt16(writer, p.Payload)
}

func (p *UInt16Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 2 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadUInt16(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic int16 payload
type Int16Payload struct {
	Payload int16
}

func (p *Int16Payload) ValueSize() uint64 {
	return 2
}

func (p *Int16Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt16(writer, p.Payload)
}

func (p *Int16Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 2 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadInt16(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic uint32 payload
type UInt32Payload struct {
	Payload uint32
}

func (p *UInt32Payload) ValueSize() uint64 {
	return 4
}

func (p *UInt32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt32(writer, p.Payload)
}

func (p *UInt32Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 4 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadUInt32(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic int32 payload
type Int32Payload struct {
	Payload int32
}

func (p *Int32Payload) ValueSize() uint64 {
	return 4
}

func (p *Int32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt32(writer, p.Payload)
}

func (p *Int32Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 4 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadInt32(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic uint64 payload
type UInt64Payload struct {
	Payload uint64
}

func (p *UInt64Payload) ValueSize() uint64 {
	return 8
}

func (p *UInt64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt64(writer, p.Payload)
}

func (p *UInt64Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 8 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadUInt64(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic int64 payload
type Int64Payload struct {
	Payload int64
}

func (p *Int64Payload) ValueSize() uint64 {
	return 8
}

func (p *Int64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt64(writer, p.Payload)
}

func (p *Int64Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 8 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadInt64(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic float32 payload
type Float32Payload struct {
	Payload float32
}

func (p *Float32Payload) ValueSize() uint64 {
	return 4
}

func (p *Float32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteFloat32(writer, p.Payload)
}

func (p *Float32Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 4 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadFloat32(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic float64 payload
type Float64Payload struct {
	Payload float64
}

func (p *Float64Payload) ValueSize() uint64 {
	return 8
}

func (p *Float64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteFloat64(writer, p.Payload)
}

func (p *Float64Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 8 {
		return ErrBadTagFormat
	}
	if s, err := serialization.ReadFloat64(reader); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}

//------------------------------------------------------------------------------

// Implementation a basic ILInt payload
type ILIntPayload struct {
	Payload uint64
}

func (p *ILIntPayload) ValueSize() uint64 {
	return uint64(ilint.EncodedSize(p.Payload))
}

func (p *ILIntPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteILInt(writer, p.Payload)
}

func (p *ILIntPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	v, err := serialization.ReadILInt(reader)
	if err != nil {
		return err
	}
	p.Payload = v
	return nil
}

//------------------------------------------------------------------------------

// Implementation a basic ILInt payload
type SignedILIntPayload struct {
	Payload int64
}

func (p *SignedILIntPayload) ValueSize() uint64 {
	return uint64(ilint.SignedEncodedSize(p.Payload))
}

func (p *SignedILIntPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteSignedILInt(writer, p.Payload)
}

func (p *SignedILIntPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	v, err := serialization.ReadSignedILInt(reader)
	if err != nil {
		return err
	}
	p.Payload = v
	return nil
}

//------------------------------------------------------------------------------

// Implementation a basic string payload
type StringPayload struct {
	Payload string
}

func (p *StringPayload) ValueSize() uint64 {
	return uint64(len(p.Payload))
}

func (p *StringPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteString(writer, p.Payload)
}

func (p *StringPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if s, err := serialization.ReadString(reader, valueSize); err == nil {
		p.Payload = s
		return nil
	} else {
		return err
	}
}
