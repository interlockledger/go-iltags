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

package direct

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

func TestDeserializeHeader(t *testing.T) {

	r := bytes.NewReader([]byte{0})
	assert.Nil(t, deserializeHeader(0, r))

	r = bytes.NewReader([]byte{0xf7})
	assert.Nil(t, deserializeHeader(0xf7, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	assert.Nil(t, deserializeHeader(0x1234567890ABCDEF, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc})
	assert.ErrorIs(t, deserializeHeader(0x1234567890ABCDEF, r), io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	assert.ErrorIs(t, deserializeHeader(0x1234567890ABCDEE, r), tags.ErrUnexpectedTagId)
}

func TestDeserializeSmallValueHeader(t *testing.T) {
	r := bytes.NewReader([]byte{0, 0})
	s, err := deserializeSmallValueHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), s)

	r = bytes.NewReader([]byte{0, 0xF7})
	s, err = deserializeSmallValueHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	s, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0x1, 0xF7})
	_, err = deserializeSmallValueHeader(0, r)
	assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)

	r = bytes.NewReader([]byte{0xff})
	_, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.ErrorIs(t, err, io.EOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	_, err = deserializeSmallValueHeader(0x1234567890ABCDEF, r)
	assert.ErrorIs(t, err, io.EOF)

	r = bytes.NewReader([]byte{0x1, 0xF8})
	_, err = deserializeSmallValueHeader(1, r)
	assert.ErrorIs(t, err, tags.ErrBadTagFormat)
}

func TestDeserializeSmallValueHeaderWithSize(t *testing.T) {

	r := bytes.NewReader([]byte{0, 0})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0, 0, r))

	r = bytes.NewReader([]byte{0, 0xF7})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0, 0xF7, r))

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	assert.Nil(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF7, r))

	r = bytes.NewReader([]byte{0xff})
	assert.ErrorIs(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF7, r), io.EOF)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	assert.ErrorIs(t, deserializeSmallValueHeaderWithSize(0x1234567890ABCDEF, 0xF6, r), tags.ErrBadTagFormat)
}

func TestDeserializeExplicitHeader(t *testing.T) {

	r := bytes.NewReader([]byte{0, 0})
	s, err := deserializeExplicitHeader(0, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), s)

	r = bytes.NewReader([]byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7, 0xF7})
	s, err = deserializeExplicitHeader(0x1234567890ABCDEF, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0xF7), s)

	r = bytes.NewReader([]byte{0xF7, 0xFB, 0x1f, 0xff, 0xff, 0x08})
	s, err = deserializeExplicitHeader(0xF7, r)
	assert.Nil(t, err)
	assert.Equal(t, uint64(536870912), s)

	r = bytes.NewReader([]byte{0xF7, 0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc})
	_, err = deserializeExplicitHeader(0xF7, r)
	assert.ErrorIs(t, err, io.ErrUnexpectedEOF)

	r = bytes.NewReader([]byte{0xF7, 0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7})
	_, err = deserializeExplicitHeader(0xF7, r)
	assert.ErrorIs(t, err, tags.ErrTagTooLarge)
}
