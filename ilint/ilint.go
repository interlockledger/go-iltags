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
This package implements the ILInt standard as defined in
ILInt Specification (https://github.com/interlockledger/specification/blob/master/ILInt/README.md).

The code inside this package is partially based on the Rust version of this
library (https://github.com/interlockledger/rust-il2-iltags).
*/
package ilint

import (
	"fmt"
	"io"
)

/*
This error is returned when the encoded ILInt is invalid.
*/
var ErrInvalidILInt = fmt.Errorf("Invalid ILInt.")

/*
This error is returned when the encoded ILInt is invalid due to an overflow in
the 64 bit value.
*/
var ErrOverflow = fmt.Errorf("Overflow.")

/*
LInt base value. All values smaller than this value are encoded asa single byte.
*/
const ILINT_BASE uint8 = 0xF8

/*
The base ILInt value as a 64 bit integer.
*/
const ILINT_BASE64 = uint64(ILINT_BASE)

/*
Returns the encoded size of the given value in bytes.
*/
func EncodedSize(v uint64) int {
	if v < ILINT_BASE64 {
		return 1
	} else if v <= (0xFF + ILINT_BASE64) {
		return 2
	} else if v <= (0xFFFF + ILINT_BASE64) {
		return 3
	} else if v <= (0x00FF_FFFF + ILINT_BASE64) {
		return 4
	} else if v <= (0xFFFF_FFFF + ILINT_BASE64) {
		return 5
	} else if v <= (0x00FF_FFFF_FFFF + ILINT_BASE64) {
		return 6
	} else if v <= (0xFFFF_FFFF_FFFF + ILINT_BASE64) {
		return 7
	} else if v <= (0x00FF_FFFF_FFFF_FFFF + ILINT_BASE64) {
		return 8
	} else {
		return 9
	}
}

/*
Encodes the current value into a byte array using the ILInt standard. The
encoded value is appended to the end of the slice b.
*/
func Encode(v uint64, b []byte) []byte {
	var tmp [9]byte
	size := EncodedSize(v)

	if size == 1 {
		tmp[0] = byte(v)
	} else {
		tmp[0] = ILINT_BASE + byte(size-2)
		t := uint64(v) - ILINT_BASE64
		for i := size - 1; i > 0; i-- {
			tmp[i] = byte(t & 0xFF)
			t = t >> 8
		}
	}
	if b == nil {
		return tmp[0:size]
	} else {
		return append(b, tmp[0:size]...)
	}
}

/*
Encodes the current value and write the result into a Writer. It fails if the
value cannot be fully written to the writer.
*/
func EncodeToWriter(v uint64, writer io.Writer) (int, error) {
	var bytes [9]byte
	ret := Encode(v, bytes[0:0])
	if n, err := writer.Write(ret); err != nil {
		return n, err
	} else if n != len(ret) {
		return n, io.ErrShortWrite
	} else {
		return n, nil
	}
}

/*
Returns the size of the encoded ILInt based on its header.
*/
func EncodedSizeFromHeader(header byte) int {
	if header < ILINT_BASE {
		return 1
	} else {
		return int(header - ILINT_BASE + 2)
	}
}

/*
Decodes the body of an ILInt without the header.
*/
func DecodeBody(body []byte) (uint64, error) {

	if len(body) == 0 || len(body) > 8 {
		return 0, ErrInvalidILInt
	}
	v := uint64(0)
	for _, b := range body {
		v = (v << 8) + uint64(b)
	}
	if v > 0xFFFF_FFFF_FFFF_FF07 {
		return 0, ErrOverflow
	}
	return v + ILINT_BASE64, nil
}

/*
Decodes an ILInt from a byte array. It returns the value, the number of bytes
consumed or an error in case of a failure.
*/
func Decode(bytes []byte) (uint64, int, error) {
	if len(bytes) == 0 {
		return 0, 0, ErrInvalidILInt
	}
	size := EncodedSizeFromHeader(bytes[0])
	if size == 1 {
		return uint64(bytes[0]), 1, nil
	}
	if len(bytes) < size {
		return (0), 0, ErrInvalidILInt
	}
	v, err := DecodeBody(bytes[1:size])
	if err != nil {
		return (0), 0, err
	}
	return v, size, nil
}

/*
Decodes an ILInt from a Reader. It returns the value, the number of bytes
consumed or an error in case of a failure.
*/
func DecodeFromReader(reader io.Reader) (uint64, int, error) {
	var bytes [8]byte

	n, err := reader.Read(bytes[0:1])
	if err != nil {
		return 0, 0, err
	}
	if n != 1 {
		return 0, 0, io.ErrUnexpectedEOF
	}
	size := EncodedSizeFromHeader(bytes[0])
	if size == 1 {
		return uint64(bytes[0]), 1, nil
	}
	size--
	n, err = reader.Read(bytes[0:size])
	if err != nil {
		return 0, 0, err
	}
	if n != size {
		return 0, n + 1, io.ErrUnexpectedEOF
	}
	v, err := DecodeBody(bytes[:size])
	if err != nil {
		return (0), 0, err
	}
	return v, size + 1, nil
}

/*
Encodes a signed value as an ILInt. It uses an alternative signed integer
encoding that guarantees the smallest encoding size based on the absolute value
of the integer. See the ILInt specification for further details about how this
encoding works.
*/
func SignedEncode(v int64) uint64 {
	tmp := uint64(v)

	if tmp&0x8000_0000_0000_0000 == 0 {
		return (tmp << 1)
	} else {
		return (^(tmp << 1))
	}
}

/*
Decodes a signed value encoded with SignedEncode().
*/
func SignedDecode(v uint64) int64 {
	if v&0x1 == 0 {
		v = (v >> 1)
	} else {
		v = ^(v >> 1)
	}
	return int64(v)
}

/*
Returns the encoded size of a signed value.
*/
func SignedEncodedSize(v int64) int {
	return EncodedSize(SignedEncode(v))
}

/*
Encodes a signed value using the ILInt format. It is equivalent of
Encode(SignedEncode(v), b).
*/
func EncodeSigned(v int64, b []byte) []byte {
	return Encode(SignedEncode(v), b)
}

/*
Encodes a signed value using the ILInt format and write the result into a writer.
It is equivalent of EncodeToWriter(SignedEncode(v), b).
*/
func EncodeSignedToWriter(v int64, writer io.Writer) (int, error) {
	return EncodeToWriter(SignedEncode(v), writer)
}

/*
Decodes a signed value encoded using the ILInt format. It is equivalent of
calling Decode() followed by a call to SignedDecode().
*/
func DecodeSigned(bytes []byte) (int64, int, error) {
	v, int, err := Decode(bytes)
	return SignedDecode(v), int, err
}

/*
Decodes a signed value encoded using the ILInt format stored in a reader.
It is equivalent of calling DecodeFromReader() followed by a call to
SignedDecode().
*/
func DecodeSignedFromReader(reader io.Reader) (int64, int, error) {
	v, int, err := DecodeFromReader(reader)
	return SignedDecode(v), int, err
}
