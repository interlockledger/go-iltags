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

package tagtest

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitedWriter(t *testing.T) {
	var w LimitedWriter
	var _ io.Writer = &w
	var _ io.ByteWriter = &w
}

func TestNewLimitedWriter(t *testing.T) {

	w := NewLimitedWriter(10, false)
	assert.Equal(t, 10, w.N)
	assert.Nil(t, w.W)

	w = NewLimitedWriter(10, true)
	assert.Equal(t, 10, w.N)
	assert.NotNil(t, w.W)
	assert.Equal(t, 10, w.W.Cap())
}

func TestLimitedWriter_Write(t *testing.T) {
	sample := FillSeq(make([]byte, 10))

	// No record
	w := NewLimitedWriter(10, false)
	assert.Equal(t, 10, w.N)
	assert.Nil(t, w.W)
	n, err := w.Write(sample[:1])
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, 9, w.N)
	n, err = w.Write(sample[1:9])
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, 1, w.N)
	n, err = w.Write(sample[9:10])
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, 0, w.N)
	n, err = w.Write(sample)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, 0, n)
	w.N = 1
	n, err = w.Write(sample)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, 1, n)
	assert.Equal(t, 0, w.N)

	// Recording 1
	w = NewLimitedWriter(10, true)
	assert.Equal(t, 10, w.N)
	assert.NotNil(t, w.W)
	n, err = w.Write(sample[:1])
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, 9, w.N)
	n, err = w.Write(sample[1:9])
	assert.Nil(t, err)
	assert.Equal(t, 8, n)
	assert.Equal(t, 1, w.N)
	n, err = w.Write(sample[9:10])
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, 0, w.N)
	assert.Equal(t, sample, w.W.Bytes())
	n, err = w.Write(sample[9:10])
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, 0, n)

	// Recording 2
	w = NewLimitedWriter(9, true)
	assert.Equal(t, 9, w.N)
	assert.NotNil(t, w.W)
	n, err = w.Write(sample)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, 9, n)
	assert.Equal(t, sample[:9], w.W.Bytes())
}

func TestLimitedWriter_WriteByte(t *testing.T) {
	sample := FillSeq(make([]byte, 10))

	// No record
	w := NewLimitedWriter(10, false)
	assert.Equal(t, 10, w.N)
	assert.Nil(t, w.W)
	for i, v := range sample {
		assert.Equal(t, 10-i, w.N)
		err := w.WriteByte(v)
		assert.Nil(t, err)
		assert.Equal(t, 10-i-1, w.N)
	}
	err := w.WriteByte(1)
	assert.ErrorIs(t, err, io.ErrShortWrite)

	// Recording
	w = NewLimitedWriter(10, true)
	assert.Equal(t, 10, w.N)
	assert.NotNil(t, w.W)
	for i, v := range sample {
		assert.Equal(t, 10-i, w.N)
		err := w.WriteByte(v)
		assert.Nil(t, err)
		assert.Equal(t, 10-i-1, w.N)
	}
	assert.Equal(t, sample, w.W.Bytes())
	err = w.WriteByte(1)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, sample, w.W.Bytes())
}
