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

package ext

import (
	"io"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/impl"
)

/*
This is the payload of the ChainNameBlockRefTag.
*/
type ChainNameBlockRefPayload struct {
	// The chain name tag.
	ChainNameTag impl.StringTag
	// The block id tag.
	BlockIdTag impl.ILIntTag
}

// Implements ILTagPayload.ValueSize().
func (p *ChainNameBlockRefPayload) ValueSize() uint64 {
	return tags.ILTagSequenceSize(&p.ChainNameTag, &p.BlockIdTag)
}

// Implements ILTagPayload.SerializeValue().
func (p *ChainNameBlockRefPayload) SerializeValue(writer io.Writer) error {
	return tags.ILTagSerializeTags(writer, &p.ChainNameTag, &p.BlockIdTag)
}

// Implements ILTagPayload.DeserializeValue().
func (p *ChainNameBlockRefPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	r := io.LimitedReader{R: reader, N: int64(valueSize)}
	if err := tags.ILTagDeserializeTagsInto(factory, &r, &p.ChainNameTag, &p.BlockIdTag); err != nil {
		return err
	}
	if r.N != 0 {
		return tags.ErrBadTagFormat
	}
	return nil
}

// Returns the chain name.
func (p *ChainNameBlockRefPayload) ChainName() string {
	return p.ChainNameTag.Payload
}

// Sets the chain name.
func (p *ChainNameBlockRefPayload) SetChainName(name string) {
	p.ChainNameTag.Payload = name
}

// Returns the block ID.
func (p *ChainNameBlockRefPayload) BlockId() uint64 {
	return p.BlockIdTag.Payload
}

// Sets the block ID.
func (p *ChainNameBlockRefPayload) SetBlockId(id uint64) {
	p.BlockIdTag.Payload = id
}

//------------------------------------------------------------------------------

/*
ChainNameBlockRefTag is a generic chain name/block id tag that stores the values
as a tag sequence composed by an StringTag followed by an ILIntTag, both with
their own IDs.

Since it is not a standard tag it does not have a Standard tag ID associated
with it.
*/
type ChainNameBlockRefTag struct {
	tags.ILTagHeaderImpl
	ChainNameBlockRefPayload
}

/*
Creates a new ChainNameBlockRefTag. It takes the id and 2 optional parameters.
The first optional parameter is the nameTagId and the second optional parameter
is the blockIdTagId. By default, they assume IL_STRING_TAG_ID and IL_ILINT_TAG_ID
respectively.

This function panics if nameTagId is equal to blockIdTagId.
*/
func NewChainNameBlockRefTag(id tags.TagID, innerTagIds ...tags.TagID) *ChainNameBlockRefTag {
	var nameTagId tags.TagID = tags.IL_STRING_TAG_ID
	var blockIdTagId tags.TagID = tags.IL_ILINT_TAG_ID
	if len(innerTagIds) > 0 {
		nameTagId = innerTagIds[0]
	}
	if len(innerTagIds) > 1 {
		blockIdTagId = innerTagIds[1]
	}
	if nameTagId == blockIdTagId {
		panic("nameTagId and blockIdTagId cannot have the same ID")
	}
	t := &ChainNameBlockRefTag{}
	t.SetId(id)
	t.ChainNameTag.SetId(nameTagId)
	t.BlockIdTag.SetId(blockIdTagId)
	return t
}
