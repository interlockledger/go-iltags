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
	"bytes"
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
)

// This is the value of a raw null tag.
var rawNullTag = []byte{0}

/*
Maximum tag size that can be handled by this library. It in this version it
is set to 512MB. This limit may be revised in the future.
*/
const MAX_TAG_SIZE uint64 = 1024 * 1024 * 512

/*
This is the interface of all ILTags.
*/
type ILTag interface {
	ILTagPayload
	ILTagHeader
}

/*
Returns the size of the header of the tag.
*/
func tagHeaderSize(tag ILTag) uint64 {
	size := uint64(ilint.EncodedSize(tag.Id().UInt64()))
	if !tag.Id().Implicit() {
		size += uint64(ilint.EncodedSize(tag.ValueSize()))
	}
	return size
}

/*
Serializes the tag header.
*/
func seralizeTagHeader(tag ILTag, writer io.Writer) error {
	if err := serialization.WriteILInt(writer, tag.Id().UInt64()); err != nil {
		return err
	}
	if !tag.Id().Implicit() {
		return serialization.WriteILInt(writer, tag.ValueSize())
	}
	return nil
}

/*
Returns the size of the tag in bytes.
*/
func ILTagSize(tag ILTag) uint64 {
	return tagHeaderSize(tag) + tag.ValueSize()
}

/*
Serializes the the tag into a stream of bytes.
*/
func ILTagSeralize(tag ILTag, writer io.Writer) error {
	if err := seralizeTagHeader(tag, writer); err != nil {
		return err
	}
	if err := tag.SerializeValue(writer); err != nil {
		return err
	}
	return nil
}

/*
Helper function that converts the tag into a byte array directly by calling
ILTagSeralize().
*/
func ILTagToBytes(tag ILTag) ([]byte, error) {
	size := ILTagSize(tag)
	writer := bytes.NewBuffer(make([]byte, 0, int(size)))
	if err := ILTagSeralize(tag, writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

/*
This function returns the size of implicit tags when they are predefined.

It returns -1 for tags that are not implicit or have variable sizes such as the
implicit ILInt tag family.
*/
func implicitPayloadSize(id TagID) int {
	IMPLICITY_TAG_SIZES := []int{
		0,  // IL_NULL_TAG_ID
		1,  // IL_BOOL_TAG_ID
		1,  // IL_INT8_TAG_ID
		1,  // IL_UINT8_TAG_ID
		2,  // IL_INT16_TAG_ID
		2,  // IL_UINT16_TAG_ID
		4,  // IL_INT32_TAG_ID
		4,  // IL_UINT32_TAG_ID
		8,  // IL_INT64_TAG_ID
		8,  // IL_UINT64_TAG_ID
		-1, // IL_ILINT_TAG_ID
		4,  // IL_BIN32_TAG_ID
		8,  // IL_BIN64_TAG_ID
		16, // IL_BIN128_TAG_ID
		-1, // IL_SIGNED_ILINT_TAG_ID
		-1, // Reserved
	}
	if id < 16 {
		return IMPLICITY_TAG_SIZES[int(id)]
	} else {
		return -1
	}
}

// Reads a TagID from the reader.
func readTagID(reader io.Reader) (TagID, error) {
	if v, err := serialization.ReadILInt(reader); err != nil {
		return 0, err
	} else {
		return TagID(v), nil
	}
}

/*
Reads the tag header and returns the tag id and the tag size.

Some implicit tags will return the size as 0xFFFF_FFFF_FFFF_FFFF to denote that
the size is fixed.
*/
func readTagHeader(reader io.Reader) (TagID, uint64, error) {
	tagId, err := readTagID(reader)
	if err != nil {
		return 0, 0, err
	}
	var size uint64
	if tagId.Implicit() {
		size = uint64(int64(implicitPayloadSize(tagId)))
	} else {
		size, err = serialization.ReadILInt(reader)
		if err != nil {
			return 0, 0, err
		}
	}
	return tagId, size, nil
}

/*
Reads the payload of a tag. This function also verifies if the tag respects the
maximum size allowed by this library.
*/
func readTagPayload(factory ILTagFactory, reader io.Reader, size uint64, tag ILTag) error {
	if tag.Id() == IL_ILINT_TAG_ID || tag.Id() == IL_SIGNED_ILINT_TAG_ID {
		return tag.DeserializeValue(factory, -1, reader)
	} else if size > MAX_TAG_SIZE {
		return ErrTagTooLarge
	} else {
		r := io.LimitedReader{R: reader, N: int64(size)}
		err := tag.DeserializeValue(factory, int(size), &r)
		if err != nil {
			return err
		}
		if r.N != 0 {
			return ErrBadTagFormat
		}
	}
	return nil
}

/*
Deserializes the tag found in the current position of the reader.
*/
func ILTagDeserialize(factory ILTagFactory, reader io.Reader) (ILTag, error) {
	tagId, size, err := readTagHeader(reader)
	if err != nil {
		return nil, err
	}
	t, err := factory.CreateTag(tagId)
	if err != nil {
		return nil, err
	}
	if err = readTagPayload(factory, reader, size, t); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

/*
Helper function that tries to deserialize the current tag into the given
tag implementation. It fails if the tag id doesn't match or if the data
is corrupted.
*/
func ILTagDeserializeInto(factory ILTagFactory, reader io.Reader, tag ILTag) error {
	tagId, size, err := readTagHeader(reader)
	if err != nil {
		return err
	}
	if tagId != tag.Id() {
		return NewErrUnexpectedTagId(tagId, tag.Id())
	}
	if err = readTagPayload(factory, reader, size, tag); err != nil {
		return err
	} else {
		return nil
	}
}

/*
Converts the given byte array into a ILTag using the given tag factory.

This function fails if the format does not contain a tag or if the data is not
fully used by the tag.
*/
func ILTagFromBytes(factory ILTagFactory, b []byte) (ILTag, error) {
	if len(b) == 0 {
		return nil, ErrBadTagFormat
	}
	r := io.LimitedReader{R: bytes.NewReader(b), N: int64(len(b))}
	t, err := ILTagDeserialize(factory, &r)
	if err != nil {
		return nil, err
	} else if r.N == 0 {
		return t, nil
	} else {
		return nil, ErrBadTagFormat
	}
}

/*
This is a helper function that serialized the given tag into the writer or add a
ILNullTag is tag is nil.

Since 2022.11.28.
*/
func ILTagSeralizeWithNull(tag ILTag, writer io.Writer) error {
	if tag == nil {
		_, err := writer.Write(rawNullTag)
		return err
	} else {
		return ILTagSeralize(tag, writer)
	}
}

/*
This helper function reads the next tag in the stream and deserialize it into
the provided tag implementation unless an ILNullTag is found.

If the tag read matches the implementation of the provided tag, it will be
loaded into it and this function will return (false, nil). If the tag read is
an ILNullTag, it will left tag unmodified and return (true, nil). If the tag
read is neither an instance of the provided tag nor an ILNullTag, this function
will return (_, ErrUnexpectedTagId).

It will also return an error describing the issue if the tag could not be read.

Since 2022.11.28.
*/
func ILTagDeserializeIntoOrNull(factory ILTagFactory, reader io.Reader,
	tag ILTag) (bool, error) {
	tagId, size, err := readTagHeader(reader)
	if err != nil {
		return false, err
	}
	if tagId == IL_NULL_TAG_ID {
		return true, nil
	}
	if tagId != tag.Id() {
		return false, NewErrUnexpectedTagId(tagId, tag.Id())
	}
	if err = readTagPayload(factory, reader, size, tag); err != nil {
		return false, err
	} else {
		return false, nil
	}
}

/*
Returns the size of the header of the tag.
*/
func ComputeagHeaderSize(tag ILTag) uint64 {
	size := uint64(ilint.EncodedSize(tag.Id().UInt64()))
	if !tag.Id().Implicit() {
		size += uint64(ilint.EncodedSize(tag.ValueSize()))
	}
	return size
}

/*
Returns the total size of an explicit tag based on its ID and payload size.

The result of this function is meaningless if the given tag ID is not assigned
to an explicit tag. It will also happens if the valueSize is larger than
18446744073709551597 (because the total size will exceed 2^64-1).

Since 2022.11.30
*/
func GetExplicitTagSize(id TagID, valueSize uint64) uint64 {
	return uint64(ilint.EncodedSize(id.UInt64())) +
		uint64(ilint.EncodedSize(valueSize)) +
		valueSize
}

/*
Serializes one or more tags in sequenct to the given writer. It fails if one of
the tags cannot be serialized.

If a nil is provided, the given tag will be serialized as a NullTag.

The total number of bytes written can be precomputed by using the function
ILTagSequenceSize().

Since 2022.12.03
*/
func ILTagSerializeTags(writer io.Writer, tags ...ILTag) error {
	for _, t := range tags {
		if err := ILTagSeralizeWithNull(t, writer); err != nil {
			return err
		}
	}
	return nil
}

/*
Deserializes one or more tags into the provided tags. It fails on the first
deserialization failure.

Since 2022.12.03
*/
func ILTagDeserializeTagInTo(factory ILTagFactory, reader io.Reader, tags ...ILTag) error {
	for _, t := range tags {
		if err := ILTagDeserializeInto(factory, reader, t); err != nil {
			return err
		}
	}
	return nil
}

/*
Deserializes one or more tags into the provided tags. It will set the pointer to
the tag to nil if it founds a NullTag instead of the expected tag.

It fails on the first deserialization failure.

Since 2022.12.03
*/
func ILTagDeserializeTagInToOrNull(factory ILTagFactory, reader io.Reader, tags ...ILTag) ([]bool, error) {
	nullList := make([]bool, len(tags))
	for i, t := range tags {
		isNull, err := ILTagDeserializeIntoOrNull(factory, reader, t)
		if err != nil {
			return nil, err
		}
		nullList[i] = isNull
	}
	return nullList, nil
}

/*
Computes the total size of the given sequence of tags. It returns the sum of the
size of all tags passed to it. If a given entry is nil, a NullTag will be
considered instead.

This function returns the number of bytes that will be written by
ILTagSerializeTags() with the same parameter.
*/
func ILTagSequenceSize(tags ...ILTag) (size uint64) {
	for _, t := range tags {
		if t != nil {
			size += ILTagSize(t)
		} else {
			size++
		}
	}
	return
}
