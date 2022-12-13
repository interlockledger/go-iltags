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
	"testing"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

func TestChainNameBlockRefPayload_ValueSize(t *testing.T) {
	// Standard IDs
	p := ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(tags.IL_STRING_TAG_ID)
	p.BlockIdTag.SetId(tags.IL_ILINT_TAG_ID)

	assert.Equal(t, uint64(2+2), p.ValueSize())
	p.ChainNameTag.Payload = "12345"
	p.BlockIdTag.Payload = 0x0123456

	assert.Equal(t, uint64(2+5+1+4), p.ValueSize())

	// Custom IDs
	p = ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(255)
	p.BlockIdTag.SetId(0x123456)

	assert.Equal(t, uint64(2+1+0+4+1+1), p.ValueSize())

	p.ChainNameTag.Payload = "12345"
	p.BlockIdTag.Payload = 0x0123456
	assert.Equal(t, uint64(2+1+5+4+1+4), p.ValueSize())
}

func TestChainNameBlockRefPayload_SerializeValue(t *testing.T) {
	// Standard IDs
	p := ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(tags.IL_STRING_TAG_ID)
	p.BlockIdTag.SetId(tags.IL_ILINT_TAG_ID)

	w := bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{
		0x11, 0x0,
		0xa, 0x0}, w.Bytes())

	p.SetChainName("12345")
	p.SetBlockId(0x0123456)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{
		0x11, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xa, 0xfa, 0x12, 0x33, 0x5e}, w.Bytes())

	// Custom IDs
	p = ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(255)
	p.BlockIdTag.SetId(0x123456)

	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{
		0xf8, 0x7, 0x0,
		0xfa, 0x12, 0x33, 0x5e, 0x1, 0x0}, w.Bytes())

	p.SetChainName("12345")
	p.SetBlockId(0x0123456)
	w = bytes.NewBuffer(nil)
	assert.Nil(t, p.SerializeValue(w))
	assert.Equal(t, []byte{
		0xf8, 0x7, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xfa, 0x12, 0x33, 0x5e, 0x4, 0xfa, 0x12, 0x33, 0x5e}, w.Bytes())
}

func TestChainNameBlockRefPayload_DeserializeValue(t *testing.T) {
	// Standard IDs
	p := ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(tags.IL_STRING_TAG_ID)
	p.BlockIdTag.SetId(tags.IL_ILINT_TAG_ID)

	bin := []byte{
		0x11, 0x0,
		0xa, 0x0}
	r := bytes.NewReader(bin)
	assert.Nil(t, p.DeserializeValue(nil, len(bin), r))
	assert.Equal(t, "", p.ChainName())
	assert.Equal(t, uint64(0), p.BlockId())

	bin = []byte{
		0x11, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xa, 0xfa, 0x12, 0x33, 0x5e}
	r = bytes.NewReader(bin)
	assert.Nil(t, p.DeserializeValue(nil, len(bin), r))
	assert.Equal(t, "12345", p.ChainName())
	assert.Equal(t, uint64(0x0123456), p.BlockId())

	bin = []byte{
		0x12, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xa, 0xfa, 0x12, 0x33, 0x5e}
	r = bytes.NewReader(bin)
	assert.ErrorIs(t, p.DeserializeValue(nil, len(bin), r), tags.ErrUnexpectedTagId)

	bin = []byte{
		0x11, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xb, 0xfa, 0x12, 0x33, 0x5e}
	r = bytes.NewReader(bin)
	assert.ErrorIs(t, p.DeserializeValue(nil, len(bin), r), tags.ErrUnexpectedTagId)

	bin = []byte{
		0x11, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xa, 0xfa, 0x12, 0x33, 0x5e}
	r = bytes.NewReader(bin)
	assert.ErrorIs(t, p.DeserializeValue(nil, len(bin)+1, r), tags.ErrBadTagFormat)

	// Custom ID
	p = ChainNameBlockRefPayload{}
	p.ChainNameTag.SetId(255)
	p.BlockIdTag.SetId(0x123456)

	bin = []byte{
		0xf8, 0x7, 0x0,
		0xfa, 0x12, 0x33, 0x5e, 0x1, 0x0}
	r = bytes.NewReader(bin)
	assert.Nil(t, p.DeserializeValue(nil, len(bin), r))
	assert.Equal(t, "", p.ChainName())
	assert.Equal(t, uint64(0), p.BlockId())

	bin = []byte{
		0xf8, 0x7, 0x5, 0x31, 0x32, 0x33, 0x34, 0x35,
		0xfa, 0x12, 0x33, 0x5e, 0x4, 0xfa, 0x12, 0x33, 0x5e}
	r = bytes.NewReader(bin)
	assert.Nil(t, p.DeserializeValue(nil, len(bin), r))
	assert.Equal(t, "12345", p.ChainName())
	assert.Equal(t, uint64(0x0123456), p.BlockId())
}

//------------------------------------------------------------------------------

func TestNewChainNameBlockRefTag(t *testing.T) {
	var _ tags.ILTag = (*ChainNameBlockRefTag)(nil)

	tag := NewChainNameBlockRefTag(1234)
	assert.Equal(t, tags.TagID(1234), tag.Id())
	assert.Equal(t, tags.IL_STRING_TAG_ID, tag.ChainNameTag.Id())
	assert.Equal(t, tags.IL_ILINT_TAG_ID, tag.BlockIdTag.Id())

	tag = NewChainNameBlockRefTag(1234, 5678)
	assert.Equal(t, tags.TagID(1234), tag.Id())
	assert.Equal(t, tags.TagID(5678), tag.ChainNameTag.Id())
	assert.Equal(t, tags.IL_ILINT_TAG_ID, tag.BlockIdTag.Id())

	tag = NewChainNameBlockRefTag(1234, 5678, 9012)
	assert.Equal(t, tags.TagID(1234), tag.Id())
	assert.Equal(t, tags.TagID(5678), tag.ChainNameTag.Id())
	assert.Equal(t, tags.TagID(9012), tag.BlockIdTag.Id())

	assert.Panics(t, func() {
		NewChainNameBlockRefTag(1234, 5678, 5678)
	})
}
