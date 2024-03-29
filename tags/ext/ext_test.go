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
	"testing"

	"github.com/stretchr/testify/assert"
)

// A dummy writer that can be used to simulate Write errors.
type DummyWriter struct {
	Limit int
}

// Implementation of io.Writer.Write().
func (w *DummyWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n > w.Limit {
		n = w.Limit
		w.Limit = 0
		return n, io.ErrShortWrite
	} else {
		w.Limit -= n
		return n, nil
	}
}

func TestDummyWriter(t *testing.T) {
	w := DummyWriter{0}

	n, err := w.Write([]byte{})
	assert.Nil(t, err)
	assert.Equal(t, 0, n)

	w.Limit = 5
	n, err = w.Write([]byte{1, 2, 3, 4, 5})
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, 0, w.Limit)

	w.Limit = 6
	n, err = w.Write([]byte{1, 2, 3, 4, 5})
	assert.Nil(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, 1, w.Limit)

	n, err = w.Write([]byte{1, 2, 3, 4, 5})
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.Equal(t, 1, n)
	assert.Equal(t, 0, w.Limit)
}
