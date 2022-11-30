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
	"bytes"
	"io"
)

/*
LimitedWriter is a struct that implements the io.Writer and io.ByteWriter
interface but can simulate IO errors along the way when a certain number of
bytes is written to it.
*/
type LimitedWriter struct {
	// If not nil, it will record all write operations.
	W *bytes.Buffer
	// Number of bytes avaiable before the write error.
	N int
}

/*
Creates a new LimitedWriter with the given limit. If record is true, it will
record the bytes written into the buffer W.
*/
func NewLimitedWriter(n int, record bool) *LimitedWriter {
	w := LimitedWriter{N: n}
	if record {
		w.W = bytes.NewBuffer(make([]byte, 0, n))
	}
	return &w
}

/*
Implements Writer.Write(). Returns io.ErrShortWrite if the number of bytes to
write is smaller than the number of bytes to write.
*/
func (w *LimitedWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n <= w.N {
		w.N -= n
	} else {
		n = w.N
		w.N = 0
		err = io.ErrShortWrite
	}
	if w.W != nil && p != nil && n > 0 {
		if _, err2 := w.W.Write(p[:n]); err2 != nil {
			panic(err2)
		}
	}
	return n, err
}

/*
Implements ByteWriter.WriteByte(). Returns io.ErrShortWrite if the write limit
has been reached.
*/
func (w *LimitedWriter) WriteByte(c byte) error {
	if w.N > 0 {
		w.N -= 1
		if w.W != nil {
			if err := w.W.WriteByte(c); err != nil {
				panic(err)
			}
		}
		return nil
	} else {
		return io.ErrShortWrite
	}
}
