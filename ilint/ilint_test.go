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
[ILInt Specification](https://github.com/interlockledger/specification/blob/master/ILInt/README.md).
*/
package ilint

import (
	"bytes"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockWriter struct {
	mock.Mock
}

func (w *mockWriter) Write(b []byte) (int, error) {
	args := w.Mock.Called(b)
	return args.Int(0), args.Error(1)
}

type mockReader struct {
	mock.Mock
	reader io.Reader
}

func (r *mockReader) Read(b []byte) (int, error) {
	args := r.Mock.Called(len(b))
	// Priorize error first
	err := args.Error(1)
	if err != nil {
		return 0, err
	}
	// Limit the number of bytes read.
	expN := args.Int(0)
	n, err := r.reader.Read(b[:expN])
	if err != nil {
		return n, nil
	} else {
		return n, err
	}
}

var sample_values = []struct {
	Value       uint64
	EncodedSize int
	Encoded     []byte
}{
	{
		0xF7,
		1,
		[]byte{0xF7},
	},
	{
		0xF8,
		2,
		[]byte{
			0xF8, 0x00,
		},
	},

	{
		0x021B,
		3,
		[]byte{
			0xF9, 0x01, 0x23,
		},
	},
	{
		0x01243D,
		4,
		[]byte{
			0xFA, 0x01, 0x23, 0x45,
		},
	},
	{
		0x0123465F,
		5,
		[]byte{
			0xFB, 0x01, 0x23, 0x45, 0x67,
		},
	},
	{
		0x0123456881,
		6,
		[]byte{
			0xFC, 0x01, 0x23, 0x45, 0x67, 0x89,
		},
	},
	{
		0x012345678AA3,
		7,
		[]byte{
			0xFD, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB,
		},
	},
	{
		0x123456789ACC5,
		8,
		[]byte{
			0xFE, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD,
		},
	},
	{
		0x123456789ABCEE7,
		9,
		[]byte{0xFF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
	},
	{
		0xFFFFFFFFFFFFFFFF,
		9,
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x07},
	},
}

func TestConstants(t *testing.T) {

	assert.Equal(t, uint8(0xF8), ILINT_BASE)
	assert.Equal(t, uint64(0xF8), ILINT_BASE64)
}

func TestILIntEncodedSize(t *testing.T) {

	assert.Equal(t, 1, EncodedSize(0))
	assert.Equal(t, 1, EncodedSize(ILINT_BASE64-1))
	assert.Equal(t, 2, EncodedSize(ILINT_BASE64))
	assert.Equal(t, 2, EncodedSize(ILINT_BASE64+0x0000_0000_0000_00FF))
	assert.Equal(t, 3, EncodedSize(ILINT_BASE64+0x0000_0000_0000_0100))
	assert.Equal(t, 3, EncodedSize(ILINT_BASE64+0x0000_0000_0000_FFFF))
	assert.Equal(t, 4, EncodedSize(ILINT_BASE64+0x0000_0000_0001_0000))
	assert.Equal(t, 4, EncodedSize(ILINT_BASE64+0x0000_0000_00FF_FFFF))
	assert.Equal(t, 5, EncodedSize(ILINT_BASE64+0x0000_0000_0100_0000))
	assert.Equal(t, 5, EncodedSize(ILINT_BASE64+0x0000_0000_FFFF_FFFF))
	assert.Equal(t, 6, EncodedSize(ILINT_BASE64+0x0000_0001_0000_0000))
	assert.Equal(t, 6, EncodedSize(ILINT_BASE64+0x0000_00FF_FFFF_FFFF))
	assert.Equal(t, 7, EncodedSize(ILINT_BASE64+0x0000_0100_0000_0000))
	assert.Equal(t, 7, EncodedSize(ILINT_BASE64+0x0000_FFFF_FFFF_FFFF))
	assert.Equal(t, 8, EncodedSize(ILINT_BASE64+0x0001_0000_0000_0000))
	assert.Equal(t, 8, EncodedSize(ILINT_BASE64+0x00FF_FFFF_FFFF_FFFF))
	assert.Equal(t, 9, EncodedSize(ILINT_BASE64+0x0100_0000_0000_0000))
	assert.Equal(t, 9, EncodedSize(0xFFFF_FFFF_FFFF_FFFF))
}

func TestEncode(t *testing.T) {

	for i := 0; i < int(ILINT_BASE); i++ {
		b := make([]byte, 1)
		c := Encode(uint64(i), b[0:0])
		assert.Equal(t, byte(i), b[0])
		assert.Same(t, &b[0], &c[0])

		c = Encode(uint64(i), b)
		assert.Equal(t, byte(i), c[1])

		c = Encode(uint64(i), nil)
		assert.Equal(t, byte(i), c[0])
	}

	for _, s := range sample_values {
		b := make([]byte, s.EncodedSize)
		c := Encode(s.Value, b[0:0])
		assert.Equal(t, s.Encoded, c)
		assert.Same(t, &b[0], &c[0])

		c = Encode(s.Value, b)
		assert.Equal(t, s.Encoded, c[s.EncodedSize:])

		c = Encode(s.Value, nil)
		assert.Equal(t, s.Encoded, c)
	}
}

func TestEncodeToWriter(t *testing.T) {

	for i := 0; i < int(ILINT_BASE); i++ {
		w := bytes.NewBuffer(nil)
		n, err := EncodeToWriter(uint64(i), w)
		assert.Nil(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, byte(i), w.Bytes()[0])
	}

	for _, s := range sample_values {
		w := bytes.NewBuffer(nil)
		n, err := EncodeToWriter(s.Value, w)
		assert.Nil(t, err)
		assert.Equal(t, s.EncodedSize, n)
		assert.Equal(t, s.Encoded, w.Bytes())
	}

	// Failed due to error
	for _, s := range sample_values {
		mw := mockWriter{}
		mw.On("Write", mock.Anything).Return(s.EncodedSize-1, nil)
		_, err := EncodeToWriter(s.Value, &mw)
		assert.ErrorIs(t, err, io.ErrShortWrite)

		mw = mockWriter{}
		mw.On("Write", mock.Anything).Return(s.EncodedSize, io.ErrShortWrite)
		_, err = EncodeToWriter(s.Value, &mw)
		assert.ErrorIs(t, err, io.ErrShortWrite)
	}
}

func TestEncodedSizeFromHeader(t *testing.T) {

	for i := 0; i < int(ILINT_BASE); i++ {
		assert.Equal(t, 1, EncodedSizeFromHeader(byte(i)))
	}
	assert.Equal(t, 2, EncodedSizeFromHeader(ILINT_BASE))
	assert.Equal(t, 3, EncodedSizeFromHeader(ILINT_BASE+1))
	assert.Equal(t, 4, EncodedSizeFromHeader(ILINT_BASE+2))
	assert.Equal(t, 5, EncodedSizeFromHeader(ILINT_BASE+3))
	assert.Equal(t, 6, EncodedSizeFromHeader(ILINT_BASE+4))
	assert.Equal(t, 7, EncodedSizeFromHeader(ILINT_BASE+5))
	assert.Equal(t, 8, EncodedSizeFromHeader(ILINT_BASE+6))
	assert.Equal(t, 9, EncodedSizeFromHeader(ILINT_BASE+7))
}

func TestDecodeBody(t *testing.T) {

	for _, s := range sample_values[2:] {
		v, err := DecodeBody(s.Encoded[1:])
		assert.Nil(t, err)
		assert.Equal(t, s.Value, v)
	}

	v, err := DecodeBody([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x07})
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xFFFF_FFFF_FFFF_FFFF), v)

	_, err = DecodeBody([]byte{})
	assert.ErrorIs(t, ErrInvalidILInt, err)
	_, err = DecodeBody(make([]byte, 9))
	assert.ErrorIs(t, ErrInvalidILInt, err)

	_, err = DecodeBody([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x08})
	assert.ErrorIs(t, ErrOverflow, err)
}

func TestDecode(t *testing.T) {

	for _, s := range sample_values {
		v, n, err := Decode(s.Encoded)
		assert.Nil(t, err)
		assert.Equal(t, s.EncodedSize, n)
		assert.Equal(t, s.Value, v)

		v, n, err = Decode(append(s.Encoded, 1))
		assert.Nil(t, err)
		assert.Equal(t, s.EncodedSize, n)
		assert.Equal(t, s.Value, v)

		_, _, err = Decode(s.Encoded[:len(s.Encoded)-1])
		assert.ErrorIs(t, ErrInvalidILInt, err)
	}

	_, _, err := Decode([]byte{})
	assert.ErrorIs(t, ErrInvalidILInt, err)

	_, _, err = Decode([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x08})
	assert.ErrorIs(t, ErrOverflow, err)

}

func TestDecodeFromReader(t *testing.T) {

	for _, s := range sample_values {
		v, n, err := DecodeFromReader(bytes.NewReader(s.Encoded))
		assert.Nil(t, err)
		assert.Equal(t, s.EncodedSize, n)
		assert.Equal(t, s.Value, v)
	}

	_, _, err := DecodeFromReader(bytes.NewReader([]byte{}))
	assert.ErrorIs(t, io.EOF, err)

	_, _, err = DecodeFromReader(bytes.NewReader([]byte{0xFF}))
	assert.ErrorIs(t, io.EOF, err)

	_, _, err = DecodeFromReader(bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}))
	assert.ErrorIs(t, io.ErrUnexpectedEOF, err)

	_, _, err = DecodeFromReader(
		bytes.NewReader([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x08}))
	assert.ErrorIs(t, ErrOverflow, err)

	sample := []byte{0xFF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}
	// Fail with an error on the first call
	mr := mockReader{reader: bytes.NewReader(sample)}
	mr.On("Read", int(1)).Return(0, io.EOF)
	v, n, err := DecodeFromReader(&mr)
	assert.ErrorIs(t, err, io.EOF)
	assert.Equal(t, uint64(0), v)
	assert.Equal(t, 0, n)

	mr = mockReader{reader: bytes.NewReader(sample)}
	mr.On("Read", int(1)).Return(0, nil)
	v, n, err = DecodeFromReader(&mr)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
	assert.Equal(t, uint64(0), v)
	assert.Equal(t, 0, n)

	mr = mockReader{reader: bytes.NewReader(sample)}
	mr.On("Read", int(1)).Return(1, nil)
	mr.On("Read", int(len(sample)-1)).Return(0, io.EOF)
	v, n, err = DecodeFromReader(&mr)
	assert.ErrorIs(t, err, io.EOF)
	assert.Equal(t, uint64(0), v)
	assert.Equal(t, 0, n)

	mr = mockReader{reader: bytes.NewReader(sample)}
	mr.On("Read", int(1)).Return(1, nil)
	mr.On("Read", int(len(sample)-1)).Return(len(sample)-2, nil)
	v, n, err = DecodeFromReader(&mr)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)
	assert.Equal(t, uint64(0), v)
	assert.Equal(t, len(sample)-1, n)
}

func TestEncodeDecode(t *testing.T) {

	for i := 0; i < 128; i++ {
		v := rand.Uint64()
		enc := Encode(v, nil)
		dec, n, err := Decode(enc)
		assert.Nil(t, err)
		assert.Equal(t, v, dec)
		assert.Equal(t, len(enc), n)

		dec, n, err = DecodeFromReader(bytes.NewReader(enc))
		assert.Nil(t, err)
		assert.Equal(t, v, dec)
		assert.Equal(t, len(enc), n)
	}
}

func TestSignedEncode(t *testing.T) {

	assert.Equal(t, uint64(0x0000_0000_0000_0000), SignedEncode(0))
	assert.Equal(t, uint64(0x0000_0000_0000_0001), SignedEncode(-1))
	assert.Equal(t, uint64(0x0000_0000_0000_0002), SignedEncode(1))
	assert.Equal(t, uint64(0xFFFF_FFFF_FFFF_FFFE), SignedEncode(int64(9223372036854775807)))
	assert.Equal(t, uint64(0xFFFF_FFFF_FFFF_FFFF), SignedEncode(int64(-9223372036854775808)))
}

func TestSignedDecode(t *testing.T) {

	assert.Equal(t, int64(0), SignedDecode(0x0000_0000_0000_0000))
	assert.Equal(t, int64(-1), SignedDecode(0x0000_0000_0000_0001))
	assert.Equal(t, int64(1), SignedDecode(0x0000_0000_0000_0002))
	assert.Equal(t, int64(9223372036854775807), SignedDecode(0xFFFF_FFFF_FFFF_FFFE))
	assert.Equal(t, int64(-9223372036854775808), SignedDecode(0xFFFF_FFFF_FFFF_FFFF))
}

func TestSignedEncodeDecode(t *testing.T) {

	for i := 0; i < 128; i++ {
		v := int64(rand.Uint64() & 0x7FFF_FFFF_FFFF_FFFF)
		enc := SignedEncode(v)
		dec := SignedDecode(enc)
		assert.Equal(t, v, dec)

		v = -v
		enc = SignedEncode(v)
		dec = SignedDecode(enc)
		assert.Equal(t, v, dec)
	}
}

func TestSignedEncodedSize(t *testing.T) {

	mask := uint64(0x7FFF_FFFF_FFFF_FFFF)
	for j := 0; j < 8; j++ {
		for i := 0; i < 16; i++ {
			v := int64(rand.Uint64() & mask)
			enc := SignedEncode(v)
			assert.Equal(t, EncodedSize(enc), SignedEncodedSize(v))

			v = -v
			enc = SignedEncode(v)
			assert.Equal(t, EncodedSize(enc), SignedEncodedSize(v))
		}
		mask = mask >> 8
	}
}

func TestEncodeDecodeSigned(t *testing.T) {

	mask := uint64(0x7FFF_FFFF_FFFF_FFFF)
	for j := 0; j < 8; j++ {
		for i := 0; i < 16; i++ {
			v := int64(rand.Uint64() & mask)
			enc := EncodeSigned(v, nil)

			b := bytes.NewBuffer(nil)
			n, err := EncodeSignedToWriter(v, b)
			assert.Nil(t, err)
			assert.Equal(t, len(enc), n)
			assert.Equal(t, enc, b.Bytes())

			dec, n, err := DecodeSigned(enc)
			assert.Nil(t, err)
			assert.Equal(t, v, dec)
			assert.Equal(t, len(enc), n)

			dec, n, err = DecodeSignedFromReader(bytes.NewReader(enc))
			assert.Nil(t, err)
			assert.Equal(t, v, dec)
			assert.Equal(t, len(enc), n)

			v = -v
			enc = EncodeSigned(v, nil)

			b = bytes.NewBuffer(nil)
			n, err = EncodeSignedToWriter(v, b)
			assert.Nil(t, err)
			assert.Equal(t, len(enc), n)
			assert.Equal(t, enc, b.Bytes())

			dec, n, err = DecodeSigned(enc)
			assert.Nil(t, err)
			assert.Equal(t, v, dec)
			assert.Equal(t, len(enc), n)

			dec, n, err = DecodeSignedFromReader(bytes.NewReader(enc))
			assert.Nil(t, err)
			assert.Equal(t, v, dec)
			assert.Equal(t, len(enc), n)
		}
		mask = mask >> 8
	}
}
