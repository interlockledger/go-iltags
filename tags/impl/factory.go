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
	"bytes"
	"io"

	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
)

/*
This function returns the size of implicit tags when they are predefined.

It returns -1 for tags that are not implicit or have variable sizes such as the
implicit ILInt tag family.
*/
func ImplicitPayloadSize(id TagID) int {
	IMPLICITY_TAG_SIZES := []int{
		0,  // IL_NULL_TAG_ID TagID
		1,  // IL_BOOL_TAG_ID TagID
		1,  // IL_INT8_TAG_ID TagID
		1,  // IL_UINT8_TAG_ID TagID
		2,  // IL_INT16_TAG_ID TagID
		2,  // IL_UINT16_TAG_ID TagID
		4,  // IL_INT32_TAG_ID TagID
		4,  // IL_UINT32_TAG_ID TagID
		8,  // IL_INT64_TAG_ID TagID
		8,  // IL_UINT64_TAG_ID TagID
		-1, // IL_ILINT_TAG_ID TagID
		4,  // IL_BIN32_TAG_ID TagID
		8,  // IL_BIN64_TAG_ID TagID
		16, // IL_BIN128_TAG_ID TagID
		-1, // IL_SIGNED_ILINT_TAG_ID TagID
		-1, // Reserved
	}
	if id < 16 {
		return IMPLICITY_TAG_SIZES[int(id)]
	} else {
		return -1
	}
}

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

// This is the type of the common interface for all ILTag creators.
type TagCreatorFunc func(TagID) ILTag

// Standard tag factory.
type StandardTagFactory struct {
	// Strict mode. If true, unknown tags will result in an Ir true
	Strict      bool
	tagCreators map[TagID]TagCreatorFunc
}

// Registers a custom tag creator for the given Tag ID. Only non reserved ids
// can be registered.
func (f *StandardTagFactory) RegisterTag(tagId TagID, tagCreator TagCreatorFunc) {
	if tagId.Reserved() {
		panic("Reserved tags cannot be overriden.")
	}
	f.tagCreators[tagId] = tagCreator
}

// Creates a non reserved tag using the creator map.
func (f *StandardTagFactory) createTagFromCreators(tagId TagID) (ILTag, error) {
	c := f.tagCreators[tagId]
	if c != nil {
		return c(tagId), nil
	} else if f.Strict {
		return nil, NewErrUnsupportedTagId(tagId)
	} else {
		return NewRawTag(tagId), nil
	}
}

// Creates an initialized tag that implements the given tag ID. Returns nil
// if the ID is not supported.
func (f *StandardTagFactory) CreateTag(tagId TagID) (ILTag, error) {
	if tagId.Reserved() {
		return newStandardTag(tagId)
	} else {
		return f.createTagFromCreators(tagId)
	}
}

func (f *StandardTagFactory) Deserialize(reader io.Reader) (ILTag, error) {
	tagId, size, err := readTagHeader(reader)
	if err != nil {
		return nil, err
	}
	t, err := f.CreateTag(tagId)
	if err != nil {
		return nil, err
	}
	if err = readTagPayload(f, reader, size, t); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

func (f *StandardTagFactory) DeserializeInto(reader io.Reader, tag ILTag) error {
	tagId, size, err := readTagHeader(reader)
	if err != nil {
		return err
	}
	if tagId != tag.Id() {
		return NewErrUnexpectedTagId(tagId, tag.Id())
	}
	if err = readTagPayload(f, reader, size, tag); err != nil {
		return err
	} else {
		return nil
	}
}

func (f *StandardTagFactory) FromBytes(b []byte) (ILTag, error) {
	r := io.LimitedReader{R: bytes.NewReader(b), N: int64(len(b))}
	t, err := f.Deserialize(&r)
	if err != nil {
		return nil, err
	} else if r.N == 0 {
		return t, nil
	} else {
		return nil, ErrBadTagFormat
	}
}
