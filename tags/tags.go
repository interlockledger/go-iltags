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

// Maximum tag size that can be handled by this library. It in this version it
// is set to 512MB. This limit may be revised in the future.
const MAX_TAG_SIZE uint64 = 1024 * 1024 * 512

// This is the interface of all ILTags.
type ILTag interface {
	ILTagPayload
	ILTagHeader
}

// Returns the size of the header of the tag.
func tagHeaderSize(tag ILTag) uint64 {
	size := uint64(ilint.EncodedSize(tag.Id().UInt64()))
	if !tag.Id().Implicit() {
		size += uint64(ilint.EncodedSize(tag.ValueSize()))
	}
	return size
}

// Serializes the tag header.
func seralizeTagHeader(tag ILTag, writer io.Writer) error {
	err := serialization.WriteILInt(writer, tag.Id().UInt64())
	if err != nil {
		return err
	}
	if !tag.Id().Implicit() {
		err = serialization.WriteILInt(writer, tag.ValueSize())
	}
	return err
}

// Returns the size of the tag in bytes.
func ILTagSize(tag ILTag) uint64 {
	return tagHeaderSize(tag) + tag.ValueSize()
}

// Serializes the the tag.
func ILTagSeralize(tag ILTag, writer io.Writer) error {
	err := seralizeTagHeader(tag, writer)
	if err != nil {
		return err
	}
	err = tag.SerializeValue(writer)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Helper function that converts the tag into a byte array.
func ILTagToBytes(tag ILTag) ([]byte, error) {
	size := ILTagSize(tag)
	if size > MAX_TAG_SIZE {
		return nil, ErrTagTooLarge
	}
	writer := bytes.NewBuffer(make([]byte, 0, int(size)))
	if err := ILTagSeralize(tag, writer); err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}

// Converts the given byte array into a ILTag using the specified tag factory.
// This function fails if the format does not contain a tag or if the data is not
// fully used by the tag.
func NewILTagFromBytes(f ILTagFactory, b []byte) (ILTag, error) {
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
