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

// Utility functions
package utils

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShredBytes(t *testing.T) {
	src := make([]byte, 10)
	exp := make([]byte, 10)

	rand.Read(src)
	ShredBytes(src)
	assert.Equal(t, src, exp)

	rand.Read(src)
	ShredBytes(src[0:1]) // Test the cleanup of the entire capacity.
	assert.Equal(t, src, exp)

	ShredBytes(nil)
}

func TestFindFirstInStringList(t *testing.T) {
	sample := []string{"a", "b", "c", "A", "a"}

	assert.Equal(t, 0, FindFirstInStringSlice(sample, "a"))
	assert.Equal(t, 1, FindFirstInStringSlice(sample, "b"))
	assert.Equal(t, 2, FindFirstInStringSlice(sample, "c"))
	assert.Equal(t, 3, FindFirstInStringSlice(sample, "A"))
	assert.Equal(t, -1, FindFirstInStringSlice(sample, "B"))

	sample = []string{}
	assert.Equal(t, -1, FindFirstInStringSlice(sample, "a"))

	assert.Equal(t, -1, FindFirstInStringSlice(nil, "a"))
}

func TestRemoveFromStringList(t *testing.T) {
	sample := []string{"a", "b", "c", "A", "a"}

	s := RemoveFromStringSlice(sample, 4)
	assert.Equal(t, []string{"a", "b", "c", "A"}, s)
	assert.Same(t, &sample[0], &s[0])
	assert.Equal(t, []string{"a", "b", "c", "A", ""}, sample)

	s = RemoveFromStringSlice(s, 0)
	assert.Equal(t, []string{"b", "c", "A"}, s)
	assert.Same(t, &sample[0], &s[0])
	assert.Equal(t, []string{"b", "c", "A", "", ""}, sample)

	s = RemoveFromStringSlice(s, 1)
	assert.Equal(t, []string{"b", "A"}, s)
	assert.Same(t, &sample[0], &s[0])
	assert.Equal(t, []string{"b", "A", "", "", ""}, sample)

	s = RemoveFromStringSlice(s, 1)
	assert.Equal(t, []string{"b"}, s)
	assert.Same(t, &sample[0], &s[0])
	assert.Equal(t, []string{"b", "", "", "", ""}, sample)

	s = RemoveFromStringSlice(s, 0)
	assert.Equal(t, []string{}, s)
	assert.Equal(t, []string{"", "", "", "", ""}, sample)

	assert.Panics(t, func() {
		RemoveFromStringSlice([]string{"b", "A"}, -1)
	})
	assert.Panics(t, func() {
		RemoveFromStringSlice([]string{"b", "A"}, 3)
	})
	assert.Panics(t, func() {
		RemoveFromStringSlice(nil, 0)
	})
}
