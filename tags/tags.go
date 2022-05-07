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
	} else if size > 0 {
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
