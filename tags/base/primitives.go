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

package base

import (
	. "github.com/interlockledger/go-iltags/tags"
	. "github.com/interlockledger/go-iltags/tags/payloads"
)

//------------------------------------------------------------------------------

// Implementation of the null tag.
type NullTag struct {
	ILTagHeaderImpl
	NullPayload
}

// Create a new NullTag.
func NewNullTag(id TagID) *NullTag {
	var t NullTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the bool tag.
type BoolTag struct {
	ILTagHeaderImpl
	BoolPayload
}

// Create a new BoolTag.
func NewBoolTag(id TagID) *BoolTag {
	var t BoolTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the uint8 tag.
type UInt8Tag struct {
	ILTagHeaderImpl
	UInt8Payload
}

// Create a new UInt8Tag.
func NewUInt8Tag(id TagID) *UInt8Tag {
	var t UInt8Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the int8 tag.
type Int8Tag struct {
	ILTagHeaderImpl
	Int8Payload
}

// Create a new Int8Tag.
func NewInt8Tag(id TagID) *Int8Tag {
	var t Int8Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the uint16 tag.
type UInt16Tag struct {
	ILTagHeaderImpl
	UInt16Payload
}

// Create a new UInt16Tag.
func NewUInt16Tag(id TagID) *UInt16Tag {
	var t UInt16Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the int16 tag.
type Int16Tag struct {
	ILTagHeaderImpl
	Int16Payload
}

// Create a new Int16Tag.
func NewInt16Tag(id TagID) *Int16Tag {
	var t Int16Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the uint32 tag.
type UInt32Tag struct {
	ILTagHeaderImpl
	UInt32Payload
}

// Create a new UInt32Tag.
func NewUInt32Tag(id TagID) *UInt32Tag {
	var t UInt32Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the int32 tag.
type Int32Tag struct {
	ILTagHeaderImpl
	Int32Payload
}

// Create a new Int32Tag.
func NewInt32Tag(id TagID) *Int32Tag {
	var t Int32Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the uint64 tag.
type UInt64Tag struct {
	ILTagHeaderImpl
	UInt64Payload
}

// Create a new UInt64Tag.
func NewUInt64Tag(id TagID) *UInt64Tag {
	var t UInt64Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the int64 tag.
type Int64Tag struct {
	ILTagHeaderImpl
	Int64Payload
}

// Create a new Int64Tag.
func NewInt64Tag(id TagID) *Int64Tag {
	var t Int64Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the float32 tag.
type Float32Tag struct {
	ILTagHeaderImpl
	Float32Payload
}

// Create a new Float32Tag.
func NewFloat32Tag(id TagID) *Float32Tag {
	var t Float32Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the float64 tag.
type Float64Tag struct {
	ILTagHeaderImpl
	Float64Payload
}

// Create a new Float64Tag.
func NewFloat64Tag(id TagID) *Float64Tag {
	var t Float64Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the float64 tag.
type Float128Tag struct {
	ILTagHeaderImpl
	Float128Payload
}

// Create a new Float64Tag.
func NewFloat128Tag(id TagID) *Float128Tag {
	var t Float128Tag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the ILInt tag.
type ILIntTag struct {
	ILTagHeaderImpl
	ILIntPayload
}

// Create a new ILIntTag.
func NewILIntTag(id TagID) *ILIntTag {
	var t ILIntTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the signed ILInt tag.
type SignedILIntTag struct {
	ILTagHeaderImpl
	SignedILIntPayload
}

// Create a new SignedILIntTag.
func NewSignedILIntTag(id TagID) *SignedILIntTag {
	var t SignedILIntTag
	t.SetId(id)
	return &t
}

//------------------------------------------------------------------------------

// Implementation of the StringTag tag.
type StringTag struct {
	ILTagHeaderImpl
	StringPayload
}

// Create a new StringTag.
func NewStringTag(id TagID) *StringTag {
	var t StringTag
	t.SetId(id)
	return &t
}
