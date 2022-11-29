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
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

/*
This function asserts if the given structure embeds another directly. This is based on the
code described here https://stackoverflow.com/questions/61585699/check-if-a-struct-has-struct-embedding-at-run-time
*/
func AssertStructEmbeds(actual interface{}, embedded interface{}) bool {
	embeddedType := reflect.TypeOf(embedded)
	actualType := reflect.TypeOf(actual)
	if actualType.Kind() != reflect.Struct {
		return false
	}
	if embeddedType == actualType {
		return false
	}
	for i := 0; i < actualType.NumField(); i++ {
		f := actualType.Field(i)
		if f.Anonymous && f.Type == embeddedType {
			return true
		}
	}
	return false
}

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
	var b B
	var c C
	var d D
	var e E
	var f F

	assert.False(t, AssertStructEmbeds(a, a))
	assert.False(t, AssertStructEmbeds(a, b))
	assert.False(t, AssertStructEmbeds(a, c))
	assert.False(t, AssertStructEmbeds(a, d))
	assert.False(t, AssertStructEmbeds(a, e))

	assert.False(t, AssertStructEmbeds(b, a))
	assert.False(t, AssertStructEmbeds(b, b))
	assert.False(t, AssertStructEmbeds(b, c))
	assert.False(t, AssertStructEmbeds(b, d))
	assert.False(t, AssertStructEmbeds(b, e))

	assert.True(t, AssertStructEmbeds(c, a))
	assert.False(t, AssertStructEmbeds(c, b))
	assert.False(t, AssertStructEmbeds(c, c))
	assert.False(t, AssertStructEmbeds(c, d))
	assert.False(t, AssertStructEmbeds(c, e))

	assert.False(t, AssertStructEmbeds(d, a))
	assert.True(t, AssertStructEmbeds(d, b))
	assert.False(t, AssertStructEmbeds(d, c))
	assert.False(t, AssertStructEmbeds(d, d))
	assert.False(t, AssertStructEmbeds(c, e))

	assert.True(t, AssertStructEmbeds(e, a))
	assert.True(t, AssertStructEmbeds(e, b))
	assert.False(t, AssertStructEmbeds(e, c))
	assert.False(t, AssertStructEmbeds(e, d))
	assert.False(t, AssertStructEmbeds(c, e))

	assert.False(t, AssertStructEmbeds(f, a))
	assert.False(t, AssertStructEmbeds(f, b))
	assert.True(t, AssertStructEmbeds(f, c))
	assert.False(t, AssertStructEmbeds(f, d))
	assert.False(t, AssertStructEmbeds(f, e))
}

func TestRemoveKeyEntry(t *testing.T) {
	l := []string{"a", "b"}

	l = removeKeyEntry(l, "a")
	assert.Equal(t, []string{"b"}, l)

	l = removeKeyEntry(l, "a")
	assert.Equal(t, []string{"b"}, l)

	l = removeKeyEntry(l, "b")
	assert.Equal(t, []string{}, l)

	l = removeKeyEntry(l, "b")
	assert.Equal(t, []string{}, l)

	l = nil
	l = removeKeyEntry(l, "b")
	assert.Nil(t, l)
}

func TestStableStringMap(t *testing.T) {
	var m StableStringMap

	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)
	assert.Equal(t, 0, m.Size())

	m.Put("b", "B")
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.Keys())
	s, ok := m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "B", s)

	m.Put("a", "A")
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.Keys())
	s, ok = m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "A", s)

	m.Put("b", "C")
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.Keys())
	s, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "C", s)

	m.Remove("a")
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.Keys())
	s, ok = m.Get("a")
	assert.False(t, ok)
	assert.Equal(t, "", s)

	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)
}

func TestStableILTagMap(t *testing.T) {
	var m StableILTagMap

	a := NewStringTag(32)
	b := NewStringTag(33)
	c := NewStringTag(34)

	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)
	assert.Equal(t, 0, m.Size())

	m.Put("b", b)
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.Keys())
	s, ok := m.Get("b")
	assert.True(t, ok)
	assert.Same(t, b, s)

	m.Put("a", a)
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.Keys())
	s, ok = m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, a, s)

	m.Put("b", c)
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.Keys())
	s, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, c, s)

	m.Remove("a")
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.Keys())
	s, ok = m.Get("a")
	assert.False(t, ok)
	assert.Nil(t, s)

	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)

}

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

// Generates a random unicode string.
func GenerateRandomString() string {

	n := rand.Int()&0x1F + 1
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		r := rune(rand.Int() & 0x7FFF)
		for !utf8.ValidRune(r) {
			r = rune(rand.Int() & 0xFFFF)
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

// Creates a list of unique random strings and its serialization as a sequence
// of standard string tags.
func CreateSampleStringArray(n int) ([]string, []byte) {

	b := bytes.NewBuffer(nil)
	l := CreateUniqueStringArray(n)
	for _, s := range l {
		if SerializeStdStringTag(s, b) != nil {
			panic("Unable to serialize the String")
		}
	}
	return l, b.Bytes()
}
