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

package impl

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/direct"
	"github.com/interlockledger/go-iltags/tagtest"
)

// Creates a list of random uint64 values and its serialization as a sequence of
// ILInt values.
func CreateSampleILTagArray(n int) ([]tags.ILTag, []byte) {
	l := make([]tags.ILTag, n)
	b := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		var t tags.ILTag
		switch i % 3 {
		case 0:
			r := NewStdBoolTag()
			r.Payload = rand.Int()&0x1 == 0
			t = r
		case 1:
			r := NewStdFloat32Tag()
			r.Payload = rand.Float32()
			t = r
		case 2:
			r := NewStdStringTag()
			r.Payload = fmt.Sprintf("%d", rand.Uint64())
			t = r
		}
		l[i] = t
		if err := tags.ILTagSeralize(t, b); err != nil {
			panic("Unable to serialize the ILTag")
		}
	}
	return l, b.Bytes()
}

// Creates a list of random tags and its serialization.
func CreateSampleILInt64Array(n int) ([]uint64, []byte) {
	l := make([]uint64, n)
	b := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		l[i] = rand.Uint64()
		if err := serialization.WriteILInt(b, l[i]); err != nil {
			panic("Unable to serialize the ILInt")
		}
	}
	return l, b.Bytes()
}

// Creates a list of unique random strings and its serialization as a sequence
// of standard string tags.
func CreateSampleStringArray(n int) ([]string, []byte) {

	b := bytes.NewBuffer(nil)
	l := tagtest.CreateUniqueStringArray(n)
	for _, s := range l {
		if direct.SerializeStdStringTag(s, b) != nil {
			panic("Unable to serialize the String")
		}
	}
	return l, b.Bytes()
}
