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
	"encoding/binary"
	"io"
	"math"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
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

func serializeStdUInt8TagCore(tagId tags.TagID, v uint8, writer io.Writer) error {
	buff := make([]byte, 2)
	buff[0] = byte(tagId & 0xFF)
	buff[1] = v
	_, err := writer.Write(buff)
	return err
}

//------------------------------------------------------------------------------

/*
Serializes a standard NullTag directly into a writer.
*/
func SerializeStdNullTag(writer io.Writer) error {
	return serialization.WriteInt8(writer, 0)
}

/*
Serializes a NullTag directly into a writer.
*/
func SerializeNullTag(tagId tags.TagID, writer io.Writer) error {
	return serializeSmallValueTagHeader(tagId, 0, writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard BoolTag directly into a writer.
*/
func SerializeStdBoolTag(v bool, writer io.Writer) error {
	b := byte(0)
	if v {
		b = 1
	}
	return serializeStdUInt8TagCore(tags.IL_BOOL_TAG_ID, b, writer)
}

/*
Serializes a standard BoolTag directly into a writer.
*/
func SerializeBoolTag(tagId tags.TagID, v bool, writer io.Writer) error {
	b := byte(0)
	if v {
		b = 1
	}
	return SerializeUInt8Tag(tagId, b, writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard UInt8Tag directly into a writer.
*/
func SerializeStdUInt8Tag(v uint8, writer io.Writer) error {
	return serializeStdUInt8TagCore(tags.IL_UINT8_TAG_ID, v, writer)
}

/*
Serializes a standard Int8Tag directly into a writer.
*/
func SerializeStdInt8Tag(v int8, writer io.Writer) error {
	return serializeStdUInt8TagCore(tags.IL_INT8_TAG_ID, uint8(v), writer)
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

//------------------------------------------------------------------------------

func serializeStdUInt16TagCore(tagId tags.TagID, v uint16, writer io.Writer) error {
	buff := make([]byte, 1+2)
	buff[0] = byte(tagId & 0xFF)
	binary.BigEndian.PutUint16(buff[1:], v)
	_, err := writer.Write(buff)
	return err
}

/*
Serializes a standard UInt16Tag directly into a writer.
*/
func SerializeStdUInt16Tag(v uint16, writer io.Writer) error {
	return serializeStdUInt16TagCore(tags.IL_UINT16_TAG_ID, v, writer)
}

/*
Serializes a standard Int16Tag directly into a writer.
*/
func SerializeStdInt16Tag(v int16, writer io.Writer) error {
	return serializeStdUInt16TagCore(tags.IL_INT16_TAG_ID, uint16(v), writer)
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

//------------------------------------------------------------------------------

func serializeStdUInt32TagCore(tagId tags.TagID, v uint32, writer io.Writer) error {
	buff := make([]byte, 1+4)
	buff[0] = byte(tagId & 0xFF)
	binary.BigEndian.PutUint32(buff[1:], v)
	_, err := writer.Write(buff)
	return err
}

/*
Serializes a standard UInt32Tag directly into a writer.
*/
func SerializeStdUInt32Tag(v uint32, writer io.Writer) error {
	return serializeStdUInt32TagCore(tags.IL_UINT32_TAG_ID, v, writer)
}

/*
Serializes a standard Int32Tag directly into a writer.
*/
func SerializeStdInt32Tag(v int32, writer io.Writer) error {
	return serializeStdUInt32TagCore(tags.IL_INT32_TAG_ID, uint32(v), writer)
}

/*
Serializes a explicit UInt32Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeUInt32Tag(tagId tags.TagID, v uint32, writer io.Writer) error {
	if err := serializeSmallValueTagHeader(tagId, 4, writer); err != nil {
		return err
	}
	return serialization.WriteUInt32(writer, v)
}

/*
Serializes a explicit Int32Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeInt32Tag(tagId tags.TagID, v int32, writer io.Writer) error {
	return SerializeUInt32Tag(tagId, uint32(v), writer)
}

//------------------------------------------------------------------------------

func serializeStdUInt64TagCore(tagId tags.TagID, v uint64, writer io.Writer) error {
	buff := make([]byte, 1+8)
	buff[0] = byte(tagId & 0xFF)
	binary.BigEndian.PutUint64(buff[1:], v)
	_, err := writer.Write(buff)
	return err
}

/*
Serializes a standard UInt64Tag directly into a writer.
*/
func SerializeStdUInt64Tag(v uint64, writer io.Writer) error {
	return serializeStdUInt64TagCore(tags.IL_UINT64_TAG_ID, v, writer)
}

/*
Serializes a standard Int64Tag directly into a writer.
*/
func SerializeStdInt64Tag(v int64, writer io.Writer) error {
	return serializeStdUInt64TagCore(tags.IL_INT64_TAG_ID, uint64(v), writer)
}

/*
Serializes a explicit UInt64Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeUInt64Tag(tagId tags.TagID, v uint64, writer io.Writer) error {
	if err := serializeSmallValueTagHeader(tagId, 8, writer); err != nil {
		return err
	}
	return serialization.WriteUInt64(writer, v)
}

/*
Serializes a explicit Int64Tag directly into a writer. The provided tagId must
be an explicit tag id or this function will fail.
*/
func SerializeInt64Tag(tagId tags.TagID, v int64, writer io.Writer) error {
	return SerializeUInt64Tag(tagId, uint64(v), writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard Float32Tag directly into a writer.
*/
func SerializeStdFloat32Tag(v float32, writer io.Writer) error {
	i := math.Float32bits(v)
	return serializeStdUInt32TagCore(tags.IL_BIN32_TAG_ID, i, writer)
}

/*
Serializes a Float32Tag directly into a writer.
*/
func SerializeFloat32Tag(tagId tags.TagID, v float32, writer io.Writer) error {
	i := math.Float32bits(v)
	return SerializeUInt32Tag(tagId, i, writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard Float64Tag directly into a writer.
*/
func SerializeStdFloat64Tag(v float64, writer io.Writer) error {
	i := math.Float64bits(v)
	return serializeStdUInt64TagCore(tags.IL_BIN64_TAG_ID, i, writer)
}

/*
Serializes a Float32Tag directly into a writer.
*/
func SerializeFloat64Tag(tagId tags.TagID, v float64, writer io.Writer) error {
	i := math.Float64bits(v)
	return SerializeUInt64Tag(tagId, i, writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard Float128Tag directly into a writer. Tne value v must have
at least 16 bytes in length or this function will panic.
*/
func SerializeStdFloat128Tag(v []byte, writer io.Writer) error {
	if err := serializeTagId(tags.IL_BIN128_TAG_ID, writer); err != nil {
		return err
	}
	_, err := writer.Write(v[:16])
	return err
}

/*
Serializes a Float32Tag directly into a writer. Tne value v must have
at least 16 bytes in length or this function will panic.
*/
func SerializeFloat128Tag(tagId tags.TagID, v []byte, writer io.Writer) error {
	return SerializeRawTag(tagId, v[:16], writer)
}

//------------------------------------------------------------------------------

/*
Serializes a standard ILIntTag directly into a writer.
*/
func SerializeStdILIntTag(v uint64, writer io.Writer) error {
	if err := serializeTagId(tags.IL_ILINT_TAG_ID, writer); err != nil {
		return err
	}
	_, err := ilint.EncodeToWriter(v, writer)
	return err
}

/*
Serializes a standard ILIntTag directly into a writer.
*/
func SerializeILIntTag(tagId tags.TagID, v uint64, writer io.Writer) error {
	tmp := ilint.Encode(v, nil)
	if err := serializeSmallValueTagHeader(tagId, len(tmp), writer); err != nil {
		return err
	}
	_, err := writer.Write(tmp)
	return err
}

/*
Serializes a standard ILIntTag directly into a writer.
*/
func SerializeStdSignedILIntTag(v int64, writer io.Writer) error {
	if err := serializeTagId(tags.IL_SIGNED_ILINT_TAG_ID, writer); err != nil {
		return err
	}
	_, err := ilint.EncodeSignedToWriter(v, writer)
	return err
}

/*
Serializes a standard ILIntTag directly into a writer.
*/
func SerializeSignedILIntTag(tagId tags.TagID, v int64, writer io.Writer) error {
	tmp := ilint.EncodeSigned(v, nil)
	if err := serializeSmallValueTagHeader(tagId, len(tmp), writer); err != nil {
		return err
	}
	_, err := writer.Write(tmp)
	return err
}

//------------------------------------------------------------------------------

/*
Serializes a RawTag directly into a writer. The tagId must be an explict tag
or the behavior of ths function is undefined.
*/
func SerializeRawTag(tagId tags.TagID, v []byte, writer io.Writer) error {
	if err := serializeTagId(tagId, writer); err != nil {
		return err
	}
	if _, err := ilint.EncodeToWriter(uint64(len(v)), writer); err != nil {
		return err
	}
	_, err := writer.Write(v)
	return err
}

/*
Serializes a BytesTag directly into a writer. The tagId must be an explict tag
or the behavior of ths function is undefined.
*/
func SerializeStdBytesTag(v []byte, writer io.Writer) error {
	return SerializeRawTag(tags.IL_BYTES_TAG_ID, v, writer)
}
