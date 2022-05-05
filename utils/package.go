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

// Utility functions.
package utils

// Shreds the given byte array. It will zeroes the entire underlying byte array
// including its full capacity. It does nothing if b is nil.
func ShredBytes(b []byte) {
	if b != nil {
		c := cap(b)
		b := b[0:c:c]
		for i := range b {
			b[i] = 0
		}
	}
}

// Finds the first index of the string in the given array. Returns -1 if the
// string is not present or l is nil.
//
// This function considers that l is not ordered, thus it runs a sequential scan
// to find the index of s.
func FindFirstInStringSlice(l []string, s string) int {
	for i, c := range l {
		if c == s {
			return i
		}
	}
	return -1
}

// Removes the element at index i from a slice of strings, shifting the elements
// whenever required. It always returns a new subslice of the same slice with the
// last element removed. As a potential optimization, it also sets the last
// element of the slice to "" in order to give the garbage collector a chance to
// delete unused strings.
//
// It will panic if l is nil or i is out of range.
func RemoveFromStringSlice(l []string, i int) []string {
	_ = l[i] // Force panic on bad parameters
	// Shifting elements using copy(). See https://go.dev/blog/slices for the
	// details about the safety of this operation.
	size := len(l)
	if size > 1 {
		copy(l[i:], l[i+1:])
	}
	l[size-1] = "" // Required because of https://go.dev/blog/slices-intro
	return l[:size-1]
}
