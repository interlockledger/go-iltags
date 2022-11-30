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
	"testing"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tagtest"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------------------
func TestBytesTag(t *testing.T) {
	var _ tags.ILTag = (*BytesTag)(nil)
	var _ tags.RawTag = BytesTag{}
}

func TestNewBytesTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *BytesTag = NewBytesTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestBigIntTag(t *testing.T) {
	var _ tags.ILTag = (*BigIntTag)(nil)

	var tag BigIntTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, BigIntPayload{}))
}

func TestNewBigIntTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *BigIntTag = NewBigIntTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestBigDecTag(t *testing.T) {
	var _ tags.ILTag = (*BigDecTag)(nil)

	var tag BigDecTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, BigDecPayload{}))
}

func TestNewBigDecTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *BigDecTag = NewBigDecTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestILIntArrayTag(t *testing.T) {
	var _ tags.ILTag = (*ILIntArrayTag)(nil)

	var tag ILIntArrayTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, ILIntArrayPayload{}))
}

func TestNewILIntArrayTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *ILIntArrayTag = NewILIntArrayTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestILTagArrayTag(t *testing.T) {
	var _ tags.ILTag = (*ILTagArrayTag)(nil)

	var tag ILTagArrayTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, ILTagArrayPayload{}))
}

func TestNewILTagArrayTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *ILTagArrayTag = NewILTagArrayTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestILTagSequenceTag(t *testing.T) {
	var _ tags.ILTag = (*ILTagSequenceTag)(nil)

	var tag ILTagSequenceTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, ILTagSequencePayload{}))
}

func TestNewILTagSequenceTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *ILTagSequenceTag = NewILTagSequenceTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestRangeTag(t *testing.T) {
	var _ tags.ILTag = (*RangeTag)(nil)

	var tag RangeTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, RangePayload{}))
}

func TestNewRangeTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *RangeTag = NewRangeTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestVersionTag(t *testing.T) {
	var _ tags.ILTag = (*VersionTag)(nil)

	var tag VersionTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, VersionPayload{}))
}

func TestNewVersionTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *VersionTag = NewVersionTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestOIDTag(t *testing.T) {
	var _ tags.ILTag = (*OIDTag)(nil)
	var _ ILIntArrayTag = OIDTag{}
}

func TestNewOIDTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *OIDTag = NewOIDTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestStringDictionaryTag(t *testing.T) {
	var _ tags.ILTag = (*StringDictionaryTag)(nil)

	var tag StringDictionaryTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, StringDictionaryPayload{}))
}

func TestNewStringDictionaryTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *StringDictionaryTag = NewStringDictionaryTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestDictionaryTag(t *testing.T) {
	var _ tags.ILTag = (*DictionaryTag)(nil)

	var tag DictionaryTag
	assert.True(t, tagtest.StructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, tagtest.StructEmbeds(tag, DictionaryPayload{}))
}

func TestNewDictionaryTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *DictionaryTag = NewDictionaryTag(id)
	assert.Equal(t, id, tag.Id())
}
