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
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
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

func serializeTagId(tagId tags.TagID, writer io.Writer) error {
	return serialization.WriteILInt(writer, tagId.UInt64())
}

func serializeSmallValueTagHeader(tagId tags.TagID, valueSize int, writer io.Writer) error {
	if err := serialization.WriteILInt(writer, tagId.UInt64()); err != nil {
		return err
	}
	return serialization.WriteUInt8(writer, uint8(valueSize))
}

func serializeStandardUInt8TagCore(tagId tags.TagID, v uint8, writer io.Writer) error {
	buff := make([]byte, 2)
	buff[0] = byte(tagId & 0xFF)
	buff[1] = v
	_, err := writer.Write(buff)
	return err
}

/*
Serializes a standard BoolTag directly into a writer.
*/
func SerializeStandardBoolTag(v bool, writer io.Writer) error {
	b := byte(0)
	if v {
		b = 1
	}
	return serializeStandardUInt8TagCore(tags.IL_BOOL_TAG_ID, b, writer)
}

/*
Serializes a standard UInt8Tag directly into a writer.
*/
func SerializeStandardUInt8Tag(v uint8, writer io.Writer) error {
	return serializeStandardUInt8TagCore(tags.IL_UINT8_TAG_ID, v, writer)
}

/*
Serializes a standard Int8Tag directly into a writer.
*/
func SerializeStandardInt8Tag(v int8, writer io.Writer) error {
	return serializeStandardUInt8TagCore(tags.IL_INT8_TAG_ID, uint8(v), writer)
}

/*
Serializes a explicit UInt8Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeUInt8Tag(tagId tags.TagID, v uint8, writer io.Writer) error {
	if err := serializeSmallValueTagHeader(tagId, 1, writer); err != nil {
		return err
	}
	return serialization.WriteUInt8(writer, v)
}

/*
Serializes a explicit Int8Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeInt8Tag(tagId tags.TagID, v int8, writer io.Writer) error {
	return SerializeUInt8Tag(tagId, uint8(v), writer)
}

func serializeStandardUInt16TagCore(tagId tags.TagID, v uint16, writer io.Writer) error {
	buff := make([]byte, 1+2)
	buff[0] = byte(tagId & 0xFF)
	buff[1] = byte((v >> 8) & 0xFF)
	buff[2] = byte(v & 0xFF)
	_, err := writer.Write(buff)
	return err
}

/*
Serializes a standard UInt16Tag directly into a writer.
*/
func SerializeStandardUInt16Tag(v uint16, writer io.Writer) error {
	return serializeStandardUInt16TagCore(tags.IL_UINT16_TAG_ID, v, writer)
}

/*
Serializes a standard Int16Tag directly into a writer.
*/
func SerializeStandardInt16Tag(v int16, writer io.Writer) error {
	return serializeStandardUInt16TagCore(tags.IL_INT16_TAG_ID, uint16(v), writer)
}

/*
Serializes a explicit UInt16Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeUInt16Tag(tagId tags.TagID, v uint16, writer io.Writer) error {
	if err := serializeSmallValueTagHeader(tagId, 2, writer); err != nil {
		return err
	}
	return serialization.WriteUInt16(writer, v)
}

/*
Serializes a explicit Int16Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeInt16Tag(tagId tags.TagID, v int16, writer io.Writer) error {
	return SerializeUInt16Tag(tagId, uint16(v), writer)
}
