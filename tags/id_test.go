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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagIDConst(t *testing.T) {
	assert.Equal(t, TagID(0xF), IMPLICIT_ID_MAX)
	assert.Equal(t, TagID(0x1F), RESERVED_ID_MAX)
	assert.Equal(t, TagID(0), IL_NULL_TAG_ID)
	assert.Equal(t, TagID(1), IL_BOOL_TAG_ID)
	assert.Equal(t, TagID(2), IL_INT8_TAG_ID)
	assert.Equal(t, TagID(3), IL_UINT8_TAG_ID)
	assert.Equal(t, TagID(4), IL_INT16_TAG_ID)
	assert.Equal(t, TagID(5), IL_UINT16_TAG_ID)
	assert.Equal(t, TagID(6), IL_INT32_TAG_ID)
	assert.Equal(t, TagID(7), IL_UINT32_TAG_ID)
	assert.Equal(t, TagID(8), IL_INT64_TAG_ID)
	assert.Equal(t, TagID(9), IL_UINT64_TAG_ID)
	assert.Equal(t, TagID(10), IL_ILINT_TAG_ID)
	assert.Equal(t, TagID(11), IL_BIN32_TAG_ID)
	assert.Equal(t, TagID(12), IL_BIN64_TAG_ID)
	assert.Equal(t, TagID(13), IL_BIN128_TAG_ID)
	assert.Equal(t, TagID(14), IL_SIGNED_ILINT_TAG_ID)
	assert.Equal(t, TagID(16), IL_BYTES_TAG_ID)
	assert.Equal(t, TagID(17), IL_STRING_TAG_ID)
	assert.Equal(t, TagID(18), IL_BINT_TAG_ID)
	assert.Equal(t, TagID(19), IL_BDEC_TAG_ID)
	assert.Equal(t, TagID(20), IL_ILINTARRAY_TAG_ID)
	assert.Equal(t, TagID(21), IL_ILTAGARRAY_TAG_ID)
	assert.Equal(t, TagID(22), IL_ILTAGSEQ_TAG_ID)
	assert.Equal(t, TagID(23), IL_RANGE_TAG_ID)
	assert.Equal(t, TagID(24), IL_VERSION_TAG_ID)
	assert.Equal(t, TagID(25), IL_OID_TAG_ID)
	assert.Equal(t, TagID(30), IL_DICTIONARY_TAG_ID)
	assert.Equal(t, TagID(31), IL_STRING_DICTIONARY_TAG_ID)
}

func TestTagID(t *testing.T) {

	// Reserved
	for i := 0; i < 16; i++ {
		id := TagID(i)
		assert.True(t, id.Implicit())
		assert.True(t, id.Reserved())
		assert.Equal(t, uint64(i), id.UInt64())
	}
	for i := 16; i < 32; i++ {
		id := TagID(i)
		assert.False(t, id.Implicit())
		assert.True(t, id.Reserved())
		assert.Equal(t, uint64(i), id.UInt64())
	}
	id := TagID(32)
	assert.False(t, id.Implicit())
	assert.False(t, id.Reserved())
	assert.Equal(t, uint64(32), id.UInt64())
}
