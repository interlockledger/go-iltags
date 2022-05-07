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
	"reflect"
	"testing"

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
	assert.Equal(t, []string{"b"}, m.keys)
	s, ok := m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "B", s)

	m.Put("a", "A")
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.keys)
	s, ok = m.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "A", s)

	m.Put("b", "C")
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, []string{"b", "a"}, m.keys)
	s, ok = m.Get("b")
	assert.True(t, ok)
	assert.Equal(t, "C", s)

	m.Remove("a")
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.keys)
	s, ok = m.Get("a")
	assert.False(t, ok)
	assert.Equal(t, "", s)

	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)
}

func TestStableILTagMap(t *testing.T) {
	var m StableILTagMap

	assert.Nil(t, m.entries)
	assert.Nil(t, m.keys)
	assert.Equal(t, 0, m.Size())

	v := NewStringTag(32)
	m.Put("b", v)
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, []string{"b"}, m.keys)
	s, ok := m.Get("b")
	assert.True(t, ok)
	assert.Same(t, v, s)

	/*
		m.Put("a", "A")
		assert.Equal(t, 2, m.Size())
		assert.Equal(t, []string{"b", "a"}, m.keys)
		s, ok = m.Get("a")
		assert.True(t, ok)
		assert.Equal(t, "A", s)

		m.Put("b", "C")
		assert.Equal(t, 2, m.Size())
		assert.Equal(t, []string{"b", "a"}, m.keys)
		s, ok = m.Get("b")
		assert.True(t, ok)
		assert.Equal(t, "C", s)

		m.Remove("a")
		assert.Equal(t, 1, m.Size())
		assert.Equal(t, []string{"b"}, m.keys)
		s, ok = m.Get("a")
		assert.False(t, ok)
		assert.Equal(t, "", s)

		m.Clear()
		assert.Nil(t, m.entries)
		assert.Nil(t, m.keys)
	*/
}
