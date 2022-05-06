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
	. "github.com/interlockledger/go-iltags/tags"
)

// Create a new standard NullTag.
func NewStdNullTag() *NullTag {
	return NewNullTag(IL_NULL_TAG_ID)
}

// Create a new standard BoolTag.
func NewStdBoolTag() *BoolTag {
	return NewBoolTag(IL_BOOL_TAG_ID)
}

// Create a new standard UInt8Tag.
func NewStdInt8Tag() *UInt8Tag {
	return NewUInt8Tag(IL_UINT8_TAG_ID)
}

// Create a new standard UInt8Tag.
func NewStdUInt8Tag() *Int8Tag {
	return NewInt8Tag(IL_INT8_TAG_ID)
}

// Create a new standard UInt16Tag.
func NewStdInt16Tag() *UInt16Tag {
	return NewUInt16Tag(IL_UINT16_TAG_ID)
}

// Create a new standard UInt16Tag.
func NewStdUInt16Tag() *Int16Tag {
	return NewInt16Tag(IL_INT16_TAG_ID)
}

// Create a new standard UInt32Tag.
func NewStdInt32Tag() *UInt32Tag {
	return NewUInt32Tag(IL_UINT32_TAG_ID)
}

// Create a new standard UInt32Tag.
func NewStdUInt32Tag() *Int32Tag {
	return NewInt32Tag(IL_INT32_TAG_ID)
}

// Create a new standard UInt64Tag.
func NewStdInt64Tag() *UInt64Tag {
	return NewUInt64Tag(IL_UINT64_TAG_ID)
}

// Create a new standard UInt64Tag.
func NewStdUInt64Tag() *Int64Tag {
	return NewInt64Tag(IL_INT64_TAG_ID)
}

// Create a new standard ILIntTag.
func NewStdILIntTag() *ILIntTag {
	return NewILIntTag(IL_UINT64_TAG_ID)
}

// Create a new standard Float32Tag.
func NewStdFloat32Tag() *Float32Tag {
	return NewFloat32Tag(IL_BIN32_TAG_ID)
}

// Create a new standard Float32Tag.
func NewStdFloat64Tag() *Float64Tag {
	return NewFloat64Tag(IL_BIN64_TAG_ID)
}

// Create a new standard Float64Tag.
func NewStdFloat128Tag() *Float128Tag {
	return NewFloat128Tag(IL_BIN128_TAG_ID)
}

// Create a new standard SignedILIntTag.
func NewStdSignedILIntTag() *SignedILIntTag {
	return NewSignedILIntTag(IL_SIGNED_ILINT_TAG_ID)
}

// Create a new StringTag.
func NewStdStringTag() *StringTag {
	return NewStringTag(IL_STRING_TAG_ID)
}
