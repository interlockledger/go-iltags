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
	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/tags"
)

const (
	// Standard null tag size.
	IL_NULL_TAG_ID_SIZE uint64 = 1

	// Standard bool tag size.
	IL_BOOL_TAG_ID_SIZE uint64 = 1 + 1

	// Standard signed 8-bit integer tag size.
	IL_INT8_TAG_ID_SIZE uint64 = 1 + 1

	// Standard unsigned 8-bit integer tag size.
	IL_UINT8_TAG_ID_SIZE uint64 = 1 + 1

	// Standard signed 16-bit integer tag size.
	IL_INT16_TAG_ID_SIZE uint64 = 1 + 2

	// Standard unsigned 16-bit integer tag size.
	IL_UINT16_TAG_ID_SIZE uint64 = 1 + 2

	// Standard signed 32-bit integer tag size.
	IL_INT32_TAG_ID_SIZE uint64 = 1 + 4

	// Standard unsigned 32-bit integer tag size.
	IL_UINT32_TAG_ID_SIZE uint64 = 1 + 4

	// Standard signed 64-bit integer tag size.
	IL_INT64_TAG_ID_SIZE uint64 = 1 + 8

	// Standard unsigned 64-bit integer tag size.
	IL_UINT64_TAG_ID_SIZE uint64 = 1 + 8

	// Standard 32-bit floating point tag size.
	IL_BIN32_TAG_ID_SIZE uint64 = 1 + 4

	// Standard 64-bit floating point tag size.
	IL_BIN64_TAG_ID_SIZE uint64 = 1 + 8

	// Standard 128-bit floating point tag size.
	IL_BIN128_TAG_ID_SIZE uint64 = 1 + 16
)

/*
Computes the size of an explicit tag with a small payload size.

The result of this function is meaningless if the value size exceeds 0xF7.
*/
func getSmallValueExplicitTagSize(tagId tags.TagID, valueSize int) uint64 {
	return uint64(ilint.EncodedSize(tagId.UInt64()) + 1 + valueSize)
}

/*
Returns the size of an explicit Int8Tag.
*/
func ExplicitInt8TagSize(tagId tags.TagID) uint64 {
	return getSmallValueExplicitTagSize(tagId, 1)
}

/*
Returns the size of an explicit Int16Tag.
*/
func ExplicitInt16TagSize(tagId tags.TagID) uint64 {
	return getSmallValueExplicitTagSize(tagId, 2)
}

/*
Returns the size of an explicit Int32Tag.
*/
func ExplicitInt32TagSize(tagId tags.TagID) uint64 {
	return getSmallValueExplicitTagSize(tagId, 4)
}

/*
Returns the size of an explicit Int64Tag.
*/
func ExplicitInt64TagSize(tagId tags.TagID) uint64 {
	return getSmallValueExplicitTagSize(tagId, 8)
}

/*
Returns the size of an explicit Float128Tag.
*/
func ExplicitFloat128TagSize(tagId tags.TagID) uint64 {
	return getSmallValueExplicitTagSize(tagId, 16)
}

/*
Returns the size of a standard ILIntTag .
*/
func StdILIntTagSize(value uint64) uint64 {
	return 1 + uint64(ilint.EncodedSize(value))
}

/*
Returns the size of an explicit ILIntTag .
*/
func ExplicitILIntTagSize(tagId tags.TagID, value uint64) uint64 {
	return getSmallValueExplicitTagSize(tagId, ilint.EncodedSize(value))
}

/*
Returns the size of a standard SignedILIntTag .
*/
func StdSignedILIntTagSize(value int64) uint64 {
	return 1 + uint64(ilint.SignedEncodedSize(value))
}

/*
Returns the size of an explicit SignedILIntTag .
*/
func ExplicitSignedILIntTagSize(tagId tags.TagID, value int64) uint64 {
	return getSmallValueExplicitTagSize(tagId, ilint.SignedEncodedSize(value))
}

var (
	/*
		Returns the size of an explicit Bool8Tag.
	*/
	ExplicitBoolTagSize = ExplicitInt8TagSize
	/*
		Returns the size of an explicit UInt8Tag.
	*/
	ExplicitUInt8TagSize = ExplicitInt8TagSize
	/*
		Returns the size of an explicit UInt16Tag.
	*/
	ExplicitUInt16TagSize = ExplicitInt16TagSize
	/*
		Returns the size of an explicit UInt32Tag.
	*/
	ExplicitUInt32TagSize = ExplicitInt32TagSize
	/*
		Returns the size of an explicit UInt64Tag.
	*/
	ExplicitUInt64TagSize = ExplicitInt64TagSize
	/*
		Returns the size of an explicit Float32Tag.
	*/
	ExplicitFloat32TagSize = ExplicitInt32TagSize
	/*
		Returns the size of an explicit Float64Tag.
	*/
	ExplicitFloat64TagSize = ExplicitInt64TagSize
)

/*
Returns the size of a RawTag. The provided tag ID must be an explicit tag id.
*/
func RawTagSize(tagId tags.TagID, value []byte) uint64 {
	return tags.GetExplicitTagSize(tagId, uint64(len(value)))
}
