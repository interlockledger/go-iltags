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
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------------------
func TestNullTag(t *testing.T) {
	var _ tags.ILTag = (*NullTag)(nil)

	var tag NullTag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, NullPayload{}))
}

func TestNewNullTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *NullTag = NewNullTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestBoolTag(t *testing.T) {
	var _ tags.ILTag = (*BoolTag)(nil)

	var tag BoolTag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, BoolPayload{}))
}

func TestNewBoolTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *BoolTag = NewBoolTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestUInt8Tag(t *testing.T) {
	var _ tags.ILTag = (*UInt8Tag)(nil)

	var tag UInt8Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, UInt8Payload{}))
}

func TestNewUInt8Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *UInt8Tag = NewUInt8Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestInt8Tag(t *testing.T) {
	var _ tags.ILTag = (*Int8Tag)(nil)

	var tag Int8Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Int8Payload{}))
}

func TestNewInt8Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Int8Tag = NewInt8Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestUInt16Tag(t *testing.T) {
	var _ tags.ILTag = (*UInt16Tag)(nil)

	var tag UInt16Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, UInt16Payload{}))
}

func TestNewUInt16Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *UInt16Tag = NewUInt16Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestInt16Tag(t *testing.T) {
	var _ tags.ILTag = (*Int16Tag)(nil)

	var tag Int16Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Int16Payload{}))
}

func TestNewInt16Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Int16Tag = NewInt16Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestUInt32Tag(t *testing.T) {
	var _ tags.ILTag = (*UInt32Tag)(nil)

	var tag UInt32Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, UInt32Payload{}))
}

func TestNewUInt32Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *UInt32Tag = NewUInt32Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestInt32Tag(t *testing.T) {
	var _ tags.ILTag = (*Int32Tag)(nil)

	var tag Int32Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Int32Payload{}))
}

func TestNewInt32Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Int32Tag = NewInt32Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestUInt64Tag(t *testing.T) {
	var _ tags.ILTag = (*UInt64Tag)(nil)

	var tag UInt64Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, UInt64Payload{}))
}

func TestNewUInt64Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *UInt64Tag = NewUInt64Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestInt64Tag(t *testing.T) {
	var _ tags.ILTag = (*Int64Tag)(nil)

	var tag Int64Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Int64Payload{}))
}

func TestNewInt64Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Int64Tag = NewInt64Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestFloat32Tag(t *testing.T) {
	var _ tags.ILTag = (*Float32Tag)(nil)

	var tag Float32Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Float32Payload{}))
}

func TestNewFloat32Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Float32Tag = NewFloat32Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestFloat64Tag(t *testing.T) {
	var _ tags.ILTag = (*Float64Tag)(nil)

	var tag Float64Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Float64Payload{}))
}

func TestNewFloat64Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Float64Tag = NewFloat64Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestFloat128Tag(t *testing.T) {
	var _ tags.ILTag = (*Float128Tag)(nil)

	var tag Float128Tag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, Float128Payload{}))
}

func TestNewFloat128Tag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *Float128Tag = NewFloat128Tag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestILIntTag(t *testing.T) {
	var _ tags.ILTag = (*ILIntTag)(nil)

	var tag ILIntTag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, ILIntPayload{}))
}

func TestNewILIntTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *ILIntTag = NewILIntTag(id)
	assert.Equal(t, id, tag.Id())
}

// ------------------------------------------------------------------------------
func TestSignedILIntTag(t *testing.T) {
	var _ tags.ILTag = (*SignedILIntTag)(nil)

	var tag SignedILIntTag
	assert.True(t, AssertStructEmbeds(tag, tags.ILTagHeaderImpl{}))
	assert.True(t, AssertStructEmbeds(tag, SignedILIntPayload{}))
}

func TestNewSignedILIntTag(t *testing.T) {
	id := tags.TagID(1234567)
	var tag *SignedILIntTag = NewSignedILIntTag(id)
	assert.Equal(t, id, tag.Id())
}
