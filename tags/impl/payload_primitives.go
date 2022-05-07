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
	. "github.com/interlockledger/go-iltags/tags"
)

// Implementation of the null payload.
type NullPayload struct {
}

// Implementation of ILTagPayload.ValueSize().
func (p *NullPayload) ValueSize() uint64 {
	return 0
}

// Implementation of ILTagPayload.SerializeValue()
func (p *NullPayload) SerializeValue(writer io.Writer) error {
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *NullPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 0 {
		return ErrBadTagFormat
	}
	return nil
}

//------------------------------------------------------------------------------

// Implementation a basic boolean payload
type BoolPayload struct {
	Payload bool
}

// Implementation of ILTagPayload.ValueSize().
func (p *BoolPayload) ValueSize() uint64 {
	return 1
}

// Implementation of ILTagPayload.SerializeValue()
func (p *BoolPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteBool(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *UInt8Payload) ValueSize() uint64 {
	return 1
}

// Implementation of ILTagPayload.SerializeValue()
func (p *UInt8Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt8(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Int8Payload) ValueSize() uint64 {
	return 1
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Int8Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt8(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *UInt16Payload) ValueSize() uint64 {
	return 2
}

// Implementation of ILTagPayload.SerializeValue()
func (p *UInt16Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt16(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Int16Payload) ValueSize() uint64 {
	return 2
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Int16Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt16(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *UInt32Payload) ValueSize() uint64 {
	return 4
}

// Implementation of ILTagPayload.SerializeValue()
func (p *UInt32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt32(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Int32Payload) ValueSize() uint64 {
	return 4
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Int32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt32(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *UInt64Payload) ValueSize() uint64 {
	return 8
}

// Implementation of ILTagPayload.SerializeValue()
func (p *UInt64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteUInt64(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Int64Payload) ValueSize() uint64 {
	return 8
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Int64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteInt64(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Float32Payload) ValueSize() uint64 {
	return 4
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Float32Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteFloat32(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation of ILTagPayload.ValueSize().
func (p *Float64Payload) ValueSize() uint64 {
	return 8
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Float64Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteFloat64(writer, p.Payload)
}

// Implementation of ILTagPayload.DeserializeValue()
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

// Implementation a basic float128 payload. This version handle this type of
// value as a binary object.
type Float128Payload struct {
	Payload [16]byte
}

// Sets the payload. The array v must have exact 16 bytes or this method panics.
func (p *Float128Payload) SetPayload(v []byte) {
	if len(v) != 16 {
		panic("v must have exact 16 bytes.")
	}
	copy(p.Payload[:], v)
}

// Implementation of ILTagPayload.ValueSize().
func (p *Float128Payload) ValueSize() uint64 {
	return 16
}

// Implementation of ILTagPayload.SerializeValue()
func (p *Float128Payload) SerializeValue(writer io.Writer) error {
	return serialization.WriteBytes(writer, p.Payload[:])
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *Float128Payload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize != 16 {
		return ErrBadTagFormat
	}
	return serialization.ReadBytes(reader, p.Payload[:])
}

//------------------------------------------------------------------------------

// Implementation a basic ILInt payload
type ILIntPayload struct {
	Payload uint64
}

// Implementation of ILTagPayload.ValueSize().
func (p *ILIntPayload) ValueSize() uint64 {
	return uint64(ilint.EncodedSize(p.Payload))
}

// Implementation of ILTagPayload.SerializeValue()
func (p *ILIntPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteILInt(writer, p.Payload)
}

/*
Implementation of ILTagPayload.DeserializeValue(). Since the payload does not
have a fixed size, valueSize is ignored.
*/
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

// Implementation of ILTagPayload.ValueSize().
func (p *SignedILIntPayload) ValueSize() uint64 {
	return uint64(ilint.SignedEncodedSize(p.Payload))
}

// Implementation of ILTagPayload.SerializeValue()
func (p *SignedILIntPayload) SerializeValue(writer io.Writer) error {
	return serialization.WriteSignedILInt(writer, p.Payload)
}

/*
Implementation of ILTagPayload.DeserializeValue(). Since the payload does not
have a fixed size, valueSize is ignored.
*/
func (p *SignedILIntPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	v, err := serialization.ReadSignedILInt(reader)
	if err != nil {
		return err
	}
	p.Payload = v
	return nil
}
