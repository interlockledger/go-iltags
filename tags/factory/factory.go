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

	. "github.com/interlockledger/go-iltags/tags"
	. "github.com/interlockledger/go-iltags/tags/base"
)

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
		return NewStandardTag(tagId)
	} else {
		return f.createTagFromCreators(tagId)
	}
}

func (f *StandardTagFactory) Deserialize(reader io.Reader) (ILTag, error) {
	tagId, size, err := ReadTagHeader(reader)
	if err != nil {
		return nil, err
	}
	t, err := f.CreateTag(tagId)
	if err != nil {
		return nil, err
	}
	if err = ReadTagPayload(f, reader, size, t); err != nil {
		return nil, err
	} else {
		return t, nil
	}
}

func (f *StandardTagFactory) DeserializeInto(reader io.Reader, tag ILTag) error {
	tagId, size, err := ReadTagHeader(reader)
	if err != nil {
		return err
	}
	if tagId != tag.Id() {
		return NewErrUnexpectedTagId(tagId, tag.Id())
	}
	if err = ReadTagPayload(f, reader, size, tag); err != nil {
		return err
	} else {
		return nil
	}
}
