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
	"github.com/interlockledger/go-iltags/tags"
)

//------------------------------------------------------------------------------

// Implementation of the BytesTag. It is actually an alias to RawTag as both
// share the same serialization.
type BytesTag = tags.RawTag

// Create a new BytesTag.
func NewBytesTag(id tags.TagID) *BytesTag {
	return tags.NewRawTag(id)
}

//------------------------------------------------------------------------------

// Implementation of the BigIntTag.
type BigIntTag struct {
	tags.ILTagHeaderImpl
	BigIntPayload
}

// Create a new RawTag.
func NewBigIntTag(id tags.TagID) *BigIntTag {
	var t BigIntTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the BigDecTag.
type BigDecTag struct {
	tags.ILTagHeaderImpl
	BigDecPayload
}

// Create a new BigDecTag.
func NewBigDecTag(id tags.TagID) *BigDecTag {
	var t BigDecTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the ILIntArrayTag.
type ILIntArrayTag struct {
	tags.ILTagHeaderImpl
	ILIntArrayPayload
}

// Create a new ILIntArrayTag.
func NewILIntArrayTag(id tags.TagID) *ILIntArrayTag {
	var t ILIntArrayTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the ILTagArrayTag.
type ILTagArrayTag struct {
	tags.ILTagHeaderImpl
	ILTagArrayPayload
}

// Create a new ILTagArrayTag.
func NewILTagArrayTag(id tags.TagID) *ILTagArrayTag {
	var t ILTagArrayTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the ILTagSequenceTag.
type ILTagSequenceTag struct {
	tags.ILTagHeaderImpl
	ILTagSequencePayload
}

// Create a new ILTagSequenceTag.
func NewILTagSequenceTag(id tags.TagID) *ILTagSequenceTag {
	var t ILTagSequenceTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the RangeTag.
type RangeTag struct {
	tags.ILTagHeaderImpl
	RangePayload
}

// Create a new RangeTag.
func NewRangeTag(id tags.TagID) *RangeTag {
	var t RangeTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the VersionTag.
type VersionTag struct {
	tags.ILTagHeaderImpl
	VersionPayload
}

// Create a new VersionTag.
func NewVersionTag(id tags.TagID) *VersionTag {
	var t VersionTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the OIDTag.
type OIDTag = ILIntArrayTag

// Create a new OIDTag.
func NewOIDTag(id tags.TagID) *OIDTag {
	return NewILIntArrayTag(id)
}

//------------------------------------------------------------------------------

// Implementation of the StringDictionaryTag.
type StringDictionaryTag struct {
	tags.ILTagHeaderImpl
	StringDictionaryPayload
}

// Create a new StringDictionaryTag.
func NewStringDictionaryTag(id tags.TagID) *StringDictionaryTag {
	var t StringDictionaryTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the DictionaryTag.
type DictionaryTag struct {
	tags.ILTagHeaderImpl
	DictionaryPayload
}

// Create a new ILTagDictionaryTag.
func NewDictionaryTag(id tags.TagID) *DictionaryTag {
	var t DictionaryTag
	t.SetId(id)
	return &t
}
