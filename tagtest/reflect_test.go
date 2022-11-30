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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertStructEmbeds(t *testing.T) {
	type A struct {
		a int
	}
	type B struct {
		b int
	}
	type C struct {
		A
	}
	type D struct {
		B
	}
	type E struct {
		A
		B
	}
	type F struct {
		C
	}

	var a A
	a.a = 1
	var b B
	b.b = 1
	var c C
	var d D
	var e E
	var f F

	assert.False(t, StructEmbeds(a, a))
	assert.False(t, StructEmbeds(a, b))
	assert.False(t, StructEmbeds(a, c))
	assert.False(t, StructEmbeds(a, d))
	assert.False(t, StructEmbeds(a, e))

	assert.False(t, StructEmbeds(b, a))
	assert.False(t, StructEmbeds(b, b))
	assert.False(t, StructEmbeds(b, c))
	assert.False(t, StructEmbeds(b, d))
	assert.False(t, StructEmbeds(b, e))

	assert.True(t, StructEmbeds(c, a))
	assert.False(t, StructEmbeds(c, b))
	assert.False(t, StructEmbeds(c, c))
	assert.False(t, StructEmbeds(c, d))
	assert.False(t, StructEmbeds(c, e))

	assert.False(t, StructEmbeds(d, a))
	assert.True(t, StructEmbeds(d, b))
	assert.False(t, StructEmbeds(d, c))
	assert.False(t, StructEmbeds(d, d))
	assert.False(t, StructEmbeds(c, e))

	assert.True(t, StructEmbeds(e, a))
	assert.True(t, StructEmbeds(e, b))
	assert.False(t, StructEmbeds(e, c))
	assert.False(t, StructEmbeds(e, d))
	assert.False(t, StructEmbeds(c, e))

	assert.False(t, StructEmbeds(f, a))
	assert.False(t, StructEmbeds(f, b))
	assert.True(t, StructEmbeds(f, c))
	assert.False(t, StructEmbeds(f, d))
	assert.False(t, StructEmbeds(f, e))
}
