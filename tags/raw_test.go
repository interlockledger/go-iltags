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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawPayload(t *testing.T) {
	var _ ILTagPayload = (*RawPayload)(nil)
	sample := []byte("And so it begins. You have forgotten something, Commander.")

	var tag RawPayload

	// Size
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = []byte{}
	assert.Equal(t, uint64(0), tag.ValueSize())
	tag.Payload = sample
	assert.Equal(t, uint64(len(sample)), tag.ValueSize())

	// Serialize
	w := bytes.NewBuffer(nil)
	tag.Payload = nil
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = []byte{}
	assert.Nil(t, tag.SerializeValue(w))
	assert.Nil(t, w.Bytes())

	w = bytes.NewBuffer(nil)
	tag.Payload = sample
	assert.Nil(t, tag.SerializeValue(w))
	assert.Equal(t, sample, w.Bytes())

	// Deserialize
	r := bytes.NewReader([]byte{})
	f := &mockFactory{}
	tag.Payload = []byte{0xFF} // Just add garbage to ensure overwrite
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, []byte{}, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Nil(t, tag.DeserializeValue(f, 0, r))
	assert.Equal(t, []byte{}, tag.Payload)

	assert.Nil(t, tag.DeserializeValue(f, len(sample), r))
	assert.Equal(t, sample, tag.Payload)

	r = bytes.NewReader(sample)
	assert.Error(t, tag.DeserializeValue(f, len(sample)+1, r))

	assert.ErrorIs(t, tag.DeserializeValue(f, -1, r), ErrBadTagFormat)
}

// ------------------------------------------------------------------------------
func TestRawTag(t *testing.T) {
	var _ ILTag = (*RawTag)(nil)

	var tag RawTag
	assert.Equal(t, &ILTagHeaderImpl{}, &tag.ILTagHeaderImpl)
	assert.Equal(t, &RawPayload{}, &tag.RawPayload)
}

func TestNewRawTag(t *testing.T) {
	id := TagID(1234567)
	var tag *RawTag = NewRawTag(id)
	assert.Equal(t, id, tag.Id())
}
