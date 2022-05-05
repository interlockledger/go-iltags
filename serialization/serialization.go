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

package serialization

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"unicode/utf8"

	"github.com/interlockledger/go-iltags/ilint"
)

var (
	// This error happens when the serialized data format is invalid.
	ErrSerializationFormat = fmt.Errorf("Invalid UTF-8 string.")
	// This error happens when an UTF-8 string is invalid.
	ErrBadUTF8String = fmt.Errorf("Invalid UTF-8 string.")
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
	b := uint8(0)
	if v {
		b = 1
	}
	return WriteUInt8(writer, b)
}

/*
Writes an unsigned 8-bit value.
*/
func WriteUInt8(writer io.Writer, v uint8) error {
	var buff [1]byte
	buff[0] = v
	return WriteBytes(writer, buff[:])
}

/*
Writes a signed 8-bit value.
*/
func WriteInt8(writer io.Writer, v int8) error {
	var buff [1]byte
	buff[0] = uint8(v)
	return WriteBytes(writer, buff[:])
}

/*
Writes an unsigned 16-bit value.
*/
func WriteUInt16(writer io.Writer, v uint16) error {
	var buff [2]byte
	binary.BigEndian.PutUint16(buff[:], v)
	return WriteBytes(writer, buff[:])
}

/*
Writes an signed 16-bit value.
*/
func WriteInt16(writer io.Writer, v int16) error {
	return WriteUInt16(writer, uint16(v))
}

/*
Writes an unsigned 32-bit value.
*/
func WriteUInt32(writer io.Writer, v uint32) error {
	var buff [4]byte
	binary.BigEndian.PutUint32(buff[:], v)
	return WriteBytes(writer, buff[:])
}

/*
Writes an signed 32-bit value.
*/
func WriteInt32(writer io.Writer, v int32) error {
	return WriteUInt32(writer, uint32(v))
}

/*
Writes an unsigned 64-bit value.
*/
func WriteUInt64(writer io.Writer, v uint64) error {
	var buff [8]byte
	binary.BigEndian.PutUint64(buff[:], v)
	return WriteBytes(writer, buff[:])
}

/*
Writes an signed 64-bit value.
*/
func WriteInt64(writer io.Writer, v int64) error {
	return WriteUInt64(writer, uint64(v))
}

/*
Writes a 32-bit floating point.
*/
func WriteFloat32(writer io.Writer, v float32) error {
	i := math.Float32bits(v)
	return WriteUInt32(writer, i)
}

/*
Writes a 64-bit floating point.
*/
func WriteFloat64(writer io.Writer, v float64) error {
	i := math.Float64bits(v)
	return WriteUInt64(writer, i)
}

/*
Writes a string encoded in UTF-8.
*/
func WriteString(writer io.Writer, v string) error {
	if utf8.ValidString(v) {
		return WriteBytes(writer, []byte(v))
	} else {
		return ErrBadUTF8String
	}
}

/*
Writes an ILInt value.
*/
func WriteILInt(writer io.Writer, v uint64) error {
	_, err := ilint.EncodeToWriter(v, writer)
	return err
}

/*
Writes an ILInt value.
*/
func WriteSignedILInt(writer io.Writer, v int64) error {
	_, err := ilint.EncodeSignedToWriter(v, writer)
	return err
}

/*
Reads all bytes into b.
*/
func ReadBytes(reader io.Reader, b []byte) error {
	if n, err := reader.Read(b); err != nil {
		return err
	} else if n != len(b) {
		return io.ErrUnexpectedEOF
	} else {
		return nil
	}
}

/*
Reads a boolean value.
*/
func ReadBool(reader io.Reader) (bool, error) {
	if v, err := ReadUInt8(reader); err != nil {
		return false, err
	} else {
		switch v {
		case 0:
			return false, nil
		case 1:
			return true, nil
		default:
			return false, ErrSerializationFormat
		}
	}
}

/*
Reads an unsigned 8-bit value.
*/
func ReadUInt8(reader io.Reader) (uint8, error) {
	var buff [1]byte
	if err := ReadBytes(reader, buff[:]); err != nil {
		return 0, err
	} else {
		return buff[0], nil
	}
}

/*
Reads a signed 8-bit value.
*/
func ReadInt8(reader io.Reader) (int8, error) {
	if v, err := ReadUInt8(reader); err != nil {
		return 0, err
	} else {
		return int8(v), nil
	}
}

/*
Reads an unsigned 16-bit value.
*/
func ReadUInt16(reader io.Reader) (uint16, error) {
	var buff [2]byte
	if err := ReadBytes(reader, buff[:]); err != nil {
		return 0, err
	} else {
		return binary.BigEndian.Uint16(buff[:]), nil
	}
}

/*
Reads a signed 16-bit value.
*/
func ReadInt16(reader io.Reader) (int16, error) {
	if v, err := ReadUInt16(reader); err != nil {
		return 0, err
	} else {
		return int16(v), nil
	}
}

/*
Reads an unsigned 32-bit value.
*/
func ReadUInt32(reader io.Reader) (uint32, error) {
	var buff [4]byte
	if err := ReadBytes(reader, buff[:]); err != nil {
		return 0, err
	} else {
		return binary.BigEndian.Uint32(buff[:]), nil
	}
}

/*
Reads a signed 32-bit value.
*/
func ReadInt32(reader io.Reader) (int32, error) {
	if v, err := ReadUInt32(reader); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

/*
Reads an unsigned 64-bit value.
*/
func ReadUInt64(reader io.Reader) (uint64, error) {
	var buff [8]byte
	if err := ReadBytes(reader, buff[:]); err != nil {
		return 0, err
	} else {
		return binary.BigEndian.Uint64(buff[:]), nil
	}
}

/*
Reads a signed 64-bit value.
*/
func ReadInt64(reader io.Reader) (int64, error) {
	if v, err := ReadUInt64(reader); err != nil {
		return 0, err
	} else {
		return int64(v), nil
	}
}

/*
Reads a 32-bit floating point value.
*/
func ReadFloat32(reader io.Reader) (float32, error) {
	if v, err := ReadUInt32(reader); err != nil {
		return 0, err
	} else {
		return math.Float32frombits(v), nil
	}
}

/*
Reads a 64-bit floating point value.
*/
func ReadFloat64(reader io.Reader) (float64, error) {
	if v, err := ReadUInt64(reader); err != nil {
		return 0, err
	} else {
		return math.Float64frombits(v), nil
	}
}

/*
Reads an UTF-8 string. If fails if it is not possible to read all bytes or the
data read is not a valid UTF-8 string.
*/
func ReadString(reader io.Reader, size int) (string, error) {
	buff := make([]byte, size)
	if err := ReadBytes(reader, buff); err != nil {
		return "", err
	}
	if utf8.Valid(buff) {
		return string(buff), nil
	} else {
		return "", ErrBadUTF8String
	}
}

/*
Reads an ILInt.
*/
func ReadILInt(reader io.Reader) (uint64, error) {
	if v, _, err := ilint.DecodeFromReader(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Reads a signed ILInt.
*/
func ReadSignedILInt(reader io.Reader) (int64, error) {
	if v, _, err := ilint.DecodeSignedFromReader(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}
