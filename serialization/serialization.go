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

/*
This package implements the serialization primitives used to serialize the tags.

Although the standard library package binary can perform most of the tasks
inside this package we decided to go for a more explicit approach when dealing
with this API.
*/
package serialization

import (
	"encoding/binary"
	"io"
)

/*
Writes all bytes of b.
*/
func WriteBytes(writer io.Writer, b []byte) error {
	if n, err := writer.Write(b); err != nil || n != len(b) {
		return io.ErrShortWrite
	} else {
		return nil
	}
}

/*
Writes a boolean value.
*/
func WriteBool(writer io.Writer, v bool) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an unsigned 8-bit value.
*/
func WriteUInt8(writer io.Writer, v uint8) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes a signed 8-bit value.
*/
func WriteInt8(writer io.Writer, v int8) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an unsigned 16-bit value.
*/
func WriteUInt16(writer io.Writer, v uint16) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an signed 16-bit value.
*/
func WriteInt16(writer io.Writer, v int16) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an unsigned 32-bit value.
*/
func WriteUInt32(writer io.Writer, v uint32) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an signed 32-bit value.
*/
func WriteInt32(writer io.Writer, v int32) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an unsigned 64-bit value.
*/
func WriteUInt64(writer io.Writer, v uint64) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes an signed 64-bit value.
*/
func WriteInt64(writer io.Writer, v int64) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes a 32-bit floating point.
*/
func WriteFloat32(writer io.Writer, v float32) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Writes a 64-bit floating point.
*/
func WriteFloat64(writer io.Writer, v float64) error {
	return binary.Write(writer, binary.BigEndian, v)
}

/*
Reads all bytes into b.
*/
func ReadBytes(reader io.Reader, b []byte) error {
	if n, err := reader.Read(b); err != nil || n != len(b) {
		return io.ErrUnexpectedEOF
	} else {
		return nil
	}
}

/*
Reads a boolean value.
*/
func ReadBool(reader io.Reader) (bool, error) {
	var v bool
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return false, err
	} else {
		return v, nil
	}
}

/*
Reads an unsigned 8-bit value.
*/
func ReadUInt8(reader io.Reader) (uint8, error) {
	var v uint8
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a signed 8-bit value.
*/
func ReadInt8(reader io.Reader) (int8, error) {
	var v int8
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads an unsigned 16-bit value.
*/
func ReadUInt16(reader io.Reader) (uint16, error) {
	var v uint16
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a signed 16-bit value.
*/
func ReadInt16(reader io.Reader) (int16, error) {
	var v int16
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads an unsigned 32-bit value.
*/
func ReadUInt32(reader io.Reader) (uint32, error) {
	var v uint32
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a signed 32-bit value.
*/
func ReadInt32(reader io.Reader) (int32, error) {
	var v int32
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads an unsigned 64-bit value.
*/
func ReadUInt64(reader io.Reader) (uint64, error) {
	var v uint64
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a signed 64-bit value.
*/
func ReadInt64(reader io.Reader) (int64, error) {
	var v int64
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a 32-bit floating point value.
*/
func ReadFloat32(reader io.Reader) (float32, error) {
	var v float32
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a 32-bit floating point value.
*/
func ReadFloat64(reader io.Reader) (float64, error) {
	var v float64
	if err := binary.Read(reader, binary.BigEndian, &v); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}
