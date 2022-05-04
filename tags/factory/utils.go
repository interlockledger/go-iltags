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

package factory

import (
	"io"

	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
	. "github.com/interlockledger/go-iltags/tags/base"
)

// Creates a new standard ILTag. It returns an error if the ID is not a
// standard tag or if the tag is not defined.
func newStandardTag(id TagID) (ILTag, error) {
	var t ILTag
	switch id {
	case IL_NULL_TAG_ID:
		t = NewNullTag(id)
	case IL_BOOL_TAG_ID:
		t = NewBoolTag(id)
	case IL_INT8_TAG_ID:
		t = NewInt8Tag(id)
	case IL_UINT8_TAG_ID:
		t = NewUInt8Tag(id)
	case IL_INT16_TAG_ID:
		t = NewInt16Tag(id)
	case IL_UINT16_TAG_ID:
		t = NewUInt16Tag(id)
	case IL_INT32_TAG_ID:
		t = NewInt32Tag(id)
	case IL_UINT32_TAG_ID:
		t = NewUInt32Tag(id)
	case IL_INT64_TAG_ID:
		t = NewInt64Tag(id)
	case IL_UINT64_TAG_ID:
		t = NewUInt64Tag(id)
	case IL_ILINT_TAG_ID:
		t = NewILIntTag(id)
	case IL_BIN32_TAG_ID:
		t = NewFloat32Tag(id)
	case IL_BIN64_TAG_ID:
		t = NewFloat64Tag(id)
	case IL_BIN128_TAG_ID:
		t = NewFloat128Tag(id)
	case IL_SIGNED_ILINT_TAG_ID:
		t = NewSignedILIntTag(id)
	default:
		return nil, NewErrUnsupportedTagId(id)
	}
	return t, nil
}

// Reads a TagID from the reader.
func readTagID(reader io.Reader) (TagID, error) {
	if v, err := serialization.ReadILInt(reader); err != nil {
		return 0, err
	} else {
		return TagID(v), nil
	}
}

// Reads the tag header and returns the tag id and the tag size.
func readTagHeader(reader io.Reader) (TagID, uint64, error) {
	tagId, err := readTagID(reader)
	if err != nil {
		return 0, 0, err
	}
	var size uint64
	if tagId.Implicit() {
		size = uint64(int64(ImplicitPayloadSize(tagId)))
	} else {
		size, err = serialization.ReadILInt(reader)
		if err != nil {
			return 0, 0, err
		}
		if size > MAX_TAG_SIZE {
			return 0, 0, ErrTagTooLarge
		}
	}
	return tagId, size, nil
}

// Reads the payload of a tag.
func readTagPayload(factory ILTagFactory, reader io.Reader, size uint64, tag ILTag) error {
	if tag.Id() == IL_ILINT_TAG_ID || tag.Id() == IL_SIGNED_ILINT_TAG_ID {
		return tag.DeserializeValue(factory, -1, reader)
	} else {
		r := io.LimitedReader{R: reader, N: int64(size)}
		err := tag.DeserializeValue(factory, int(size), &r)
		if err != nil {
			return nil
		}
		if r.N != 0 {
			return ErrBadTagFormat
		} else {
			return nil
		}
	}
}
