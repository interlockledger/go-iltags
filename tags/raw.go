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
	"io"

	"github.com/interlockledger/go-iltags/serialization"
)

//------------------------------------------------------------------------------

// Implementation of the raw payload.
type RawPayload struct {
	Payload []byte
}

// Implementation of ILTagPayload.ValueSize().
func (p *RawPayload) ValueSize() uint64 {
	if p.Payload != nil {
		return uint64(len(p.Payload))
	} else {
		return 0
	}
}

// Implementation of ILTagPayload.SerializeValue()
func (p *RawPayload) SerializeValue(writer io.Writer) error {
	if p.Payload != nil {
		return serialization.WriteBytes(writer, p.Payload)
	}
	return nil
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *RawPayload) DeserializeValue(factory ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 0 {
		return ErrBadTagFormat
	}
	p.Payload = make([]byte, valueSize)
	if valueSize > 0 {
		return serialization.ReadBytes(reader, p.Payload)
	} else {
		return nil
	}
}

// Implementation of the raw tag.
type RawTag struct {
	ILTagHeaderImpl
	RawPayload
}

// Create a new RawTag.
func NewRawTag(id TagID) *RawTag {
	var t RawTag
	t.SetId(id)
	return &t
}
