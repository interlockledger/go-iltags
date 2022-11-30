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
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

func TestDirectSizeConstants(t *testing.T) {
	assert.Equal(t, uint64(1), IL_NULL_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+1), IL_BOOL_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+1), IL_INT8_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+1), IL_UINT8_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+2), IL_INT16_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+2), IL_UINT16_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+4), IL_INT32_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+4), IL_UINT32_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+8), IL_INT64_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+8), IL_UINT64_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+4), IL_BIN32_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+8), IL_BIN64_TAG_ID_SIZE)
	assert.Equal(t, uint64(1+16), IL_BIN128_TAG_ID_SIZE)
}

func TestGetSmallValueExplicitTagSize(t *testing.T) {
	// Test borders
	assert.Equal(t, tags.GetExplicitTagSize(16, 0),
		getSmallValueExplicitTagSize(16, 0))
	assert.Equal(t, tags.GetExplicitTagSize(16, 0xF7),
		getSmallValueExplicitTagSize(16, 0xF7))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF7),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF7))

	// Test why it fails.
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF8),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF8)+1)
}

func TestExplicitXXTagSize(t *testing.T) {

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitBoolTagSize(16))

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitInt8TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 1),
		ExplicitInt8TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitUInt8TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 1),
		ExplicitUInt8TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 2),
		ExplicitInt16TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 2),
		ExplicitInt16TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 2),
		ExplicitUInt16TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 2),
		ExplicitUInt16TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitInt32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitInt32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitUInt32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitUInt32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitInt64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitInt64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitUInt64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitUInt64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitFloat32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitFloat32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitFloat64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitFloat64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 16),
		ExplicitFloat128TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 16),
		ExplicitFloat128TagSize(0xFFFF_FFFF_FFFF_FFFF))
}

func TestILIntTagSize(t *testing.T) {

	v := uint64(0)
	tagId := tags.TagID(16)
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.EncodedSize(v))),
		ExplicitILIntTagSize(16, v))
	assert.Equal(t, uint64(ilint.EncodedSize(v))+1,
		StdILIntTagSize(v))

	v = uint64(0xFFFF_FFFF_FFFF_FF00)
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId,
		uint64(ilint.EncodedSize(v))),
		ExplicitILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.EncodedSize(v))+1,
		StdILIntTagSize(v))
}

func TestSignedILIntTagSize(t *testing.T) {

	v := int64(0)
	tagId := tags.TagID(16)
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = 1
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = -1
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = 9223372036854775807
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = -9223372036854775808
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))
}

func TestRawTagSize(t *testing.T) {

	for i := 0; i < 128; i++ {
		b := make([]byte, i)
		assert.Equal(t,
			tags.GetExplicitTagSize(16, uint64(i)),
			RawTagSize(16, b))

		assert.Equal(t,
			tags.GetExplicitTagSize(0x0123456789ABCDEF, uint64(i)),
			RawTagSize(0x0123456789ABCDEF, b))
	}
}
