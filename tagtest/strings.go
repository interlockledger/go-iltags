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
	"math/rand"
	"strings"
	"unicode"
)

/*
Generates a random unicode string. It generates strings with 1 to 32 characteres
in length.
*/
func GenerateRandomString() string {
	return GenerateRandomStringWithLen(rand.Int()&0x1F + 1)
}

// Generates a random unicode string with n characters. Only graphical
// codepoints are used. It will return "" if n is 0 or less than 0.
func GenerateRandomStringWithLen(n int) string {
	if n <= 0 {
		return ""
	}
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		r := rune(rand.Int() & 0x1FFFFF)
		for !unicode.IsGraphic(r) {
			r = rune(rand.Int() & 0x1FFFFF)
		}
		b.WriteRune(r)
	}
	return b.String()
}

// Creates a list of unique random strings.
func CreateUniqueStringArray(n int) []string {

	l := make([]string, n)
	dl := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		s := GenerateRandomString()
		if _, ok := dl[s]; !ok {
			dl[s] = true
			l[i] = s
		}
	}
	return l
}
