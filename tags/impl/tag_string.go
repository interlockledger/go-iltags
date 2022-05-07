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
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	. "github.com/interlockledger/go-iltags/tags"
)

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

/*
Returns the size of the string tag that will hold the given string.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func StringTagSize(tagId TagID, s string) uint64 {
	l := len(s)
	return uint64(ilint.EncodedSize(uint64(tagId)) +
		ilint.EncodedSize(uint64(l)) + l)
}

/*
Serializes a string directly using the StringTag format.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func SerializeStringTag(tagId TagID, s string, writer io.Writer) error {
	if err := serialization.WriteILInt(writer, uint64(tagId)); err != nil {
		return err
	}
	if err := serialization.WriteILInt(writer, uint64(len(s))); err != nil {
		return err
	}
	if err := serialization.WriteString(writer, s); err != nil {
		return err
	}
	return nil
}

/*
Deserializes a string tag directly into a string.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func DeserializeStringTag(expectedId TagID, reader io.Reader) (string, error) {
	id, err := serialization.ReadILInt(reader)
	if err != nil {
		return "", err
	}
	if TagID(id) != expectedId {
		return "", NewErrUnexpectedTagId(expectedId, TagID(id))
	}
	size, err := serialization.ReadILInt(reader)
	if size > MAX_TAG_SIZE {
		return "", ErrTagTooLarge
	}
	if s, err := serialization.ReadString(reader, int(size)); err != nil {
		return "", err
	} else {
		return s, nil
	}
}

/*
Returns the size of the standard string tag that will hold the given string.
*/
func StdStringTagSize(s string) uint64 {
	return StringTagSize(IL_STRING_TAG_ID, s)
}

/*
Serializes a string directly using the standard StringTag format.
*/
func SerializeStdStringTag(s string, writer io.Writer) error {
	return SerializeStringTag(IL_STRING_TAG_ID, s, writer)
}

/*
Deserializes a standard string tag directly into a string.
*/
func DeserializeStdStringTag(reader io.Reader) (string, error) {
	return DeserializeStringTag(IL_STRING_TAG_ID, reader)
}
