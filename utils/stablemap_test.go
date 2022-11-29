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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseStableMap(t *testing.T) {
	m := StableStringMap{}
	assert.Nil(t, m.entries)
	assert.Nil(t, m.index)
	assert.Equal(t, 0, m.removed)
}

func TestBaseStableMap_Size(t *testing.T) {
	m := StableStringMap{}

	assert.Equal(t, 0, m.Size())
	m.Put("k1", "v1")
	assert.Equal(t, 1, m.Size())

	m.Put("k2", "v2")
	assert.Equal(t, 2, m.Size())

	m.Remove("k1")
	assert.Equal(t, 1, m.Size())
}

func TestBaseStableMap_Empty(t *testing.T) {
	m := StableStringMap{}

	assert.True(t, m.Empty())
	m.Put("k1", "v1")
	assert.False(t, m.Empty())
	m.Put("k2", "v2")
	assert.False(t, m.Empty())
	m.Remove("k1")
	assert.False(t, m.Empty())
	m.Remove("k2")
	assert.True(t, m.Empty())
}

func TestBaseStableMap_Keys(t *testing.T) {
	m := StableStringMap{}

	assert.Equal(t, []string{}, m.Keys())
	m.Put("k2", "v2")
	assert.Equal(t, []string{"k2"}, m.Keys())

	m.Put("k1", "v1")
	assert.Equal(t, []string{"k2", "k1"}, m.Keys())

	m.Put("k3", "v3")
	assert.Equal(t, []string{"k2", "k1", "k3"}, m.Keys())

	m.Remove("k1")
	assert.Equal(t, []string{"k2", "k3"}, m.Keys())

	m.Remove("k2")
	assert.Equal(t, []string{"k3"}, m.Keys())

	m.Remove("k3")
	assert.Equal(t, []string{}, m.Keys())
}

func TestBaseStableMap_Entries(t *testing.T) {
	m := StableStringMap{}

	assert.Equal(t, []StableStringMapEntry{}, m.Entries())
	m.Put("k2", "v2")
	assert.Equal(t, []StableStringMapEntry{
		{"k2", "v2"},
	}, m.Entries())

	m.Put("k1", "v1")
	assert.Equal(t, []StableStringMapEntry{
		{"k2", "v2"},
		{"k1", "v1"},
	}, m.Entries())

	m.Put("k3", "v3")
	assert.Equal(t, []StableStringMapEntry{
		{"k2", "v2"},
		{"k1", "v1"},
		{"k3", "v3"},
	}, m.Entries())

	m.Remove("k1")
	assert.Equal(t, []StableStringMapEntry{
		{"k2", "v2"},
		{"k3", "v3"},
	}, m.Entries())

	m.Remove("k2")
	assert.Equal(t, []StableStringMapEntry{
		{"k3", "v3"},
	}, m.Entries())

	m.Remove("k3")
	assert.Equal(t, []StableStringMapEntry{}, m.Entries())
}

func TestBaseStableMap_Put(t *testing.T) {
	m := StableStringMap{}

	m.Put("k1", "v1")
	assert.Len(t, m.entries, 1)
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[0])
	assert.Equal(t, 0, m.index["k1"])

	m.Put("k3", "v3")
	assert.Len(t, m.entries, 2)
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[1])
	assert.Equal(t, 0, m.index["k1"])
	assert.Equal(t, 1, m.index["k3"])

	m.Put("k1", "nv1")
	assert.Len(t, m.entries, 2)
	assert.Equal(t, &StableStringMapEntry{"k1", "nv1"}, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[1])
	assert.Equal(t, 0, m.index["k1"])
	assert.Equal(t, 1, m.index["k3"])

	m.Put("k2", "v2")
	assert.Len(t, m.entries, 3)
	assert.Equal(t, &StableStringMapEntry{"k1", "nv1"}, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k2", "v2"}, m.entries[2])
	assert.Equal(t, 0, m.index["k1"])
	assert.Equal(t, 1, m.index["k3"])
	assert.Equal(t, 2, m.index["k2"])

	m.Remove("k1")
	assert.Len(t, m.entries, 3)
	assert.Nil(t, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k2", "v2"}, m.entries[2])
	_, found := m.index["k1"]
	assert.False(t, found)
	assert.Equal(t, 1, m.index["k3"])
	assert.Equal(t, 2, m.index["k2"])

	m.Put("k1", "v1")
	assert.Len(t, m.entries, 4)
	assert.Nil(t, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k2", "v2"}, m.entries[2])
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[3])
	assert.Equal(t, 1, m.index["k3"])
	assert.Equal(t, 2, m.index["k2"])
	assert.Equal(t, 3, m.index["k1"])
}

func TestBaseStableMap_Get(t *testing.T) {
	m := StableStringMap{}

	_, found := m.Get("k1")
	assert.False(t, found)

	m.Put("k1", "v1")
	v, found := m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "v1", v)

	m.Put("k3", "v3")
	v, found = m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "v1", v)
	v, found = m.Get("k3")
	assert.True(t, found)
	assert.Equal(t, "v3", v)

	m.Put("k2", "v2")
	v, found = m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "v1", v)
	v, found = m.Get("k3")
	assert.True(t, found)
	assert.Equal(t, "v3", v)
	v, found = m.Get("k2")
	assert.True(t, found)
	assert.Equal(t, "v2", v)

	m.Remove("k3")
	v, found = m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "v1", v)
	_, found = m.Get("k3")
	assert.False(t, found)
	v, found = m.Get("k2")
	assert.True(t, found)
	assert.Equal(t, "v2", v)

	m.Put("k3", "nv3")
	v, found = m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "v1", v)
	v, found = m.Get("k3")
	assert.True(t, found)
	assert.Equal(t, "nv3", v)
	v, found = m.Get("k2")
	assert.True(t, found)
	assert.Equal(t, "v2", v)

	m.Put("k1", "nv1")
	v, found = m.Get("k1")
	assert.True(t, found)
	assert.Equal(t, "nv1", v)
	v, found = m.Get("k3")
	assert.True(t, found)
	assert.Equal(t, "nv3", v)
	v, found = m.Get("k2")
	assert.True(t, found)
	assert.Equal(t, "v2", v)
}

func TestBaseStableMap_Remove(t *testing.T) {
	m := StableStringMap{}

	assert.False(t, m.Remove("k1"))
	m.Put("k1", "v1")
	m.Put("k2", "v2")
	m.Put("k3", "v3")
	assert.Len(t, m.entries, 3)
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k2", "v2"}, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[2])
	assert.Equal(t, 0, m.index["k1"])
	assert.Equal(t, 1, m.index["k2"])
	assert.Equal(t, 2, m.index["k3"])
	assert.Equal(t, 3, m.Size())

	assert.True(t, m.Remove("k2"))
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, 1, m.removed)
	assert.Len(t, m.entries, 3)
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[0])
	assert.Nil(t, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[2])
	assert.Equal(t, 0, m.index["k1"])
	assert.NotContains(t, m.index, "k2")
	assert.Equal(t, 2, m.index["k3"])

	assert.False(t, m.Remove("k2"))
	assert.Equal(t, 2, m.Size())
	assert.Equal(t, 1, m.removed)
	assert.Len(t, m.entries, 3)
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[0])
	assert.Nil(t, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[2])
	assert.Equal(t, 0, m.index["k1"])
	assert.NotContains(t, m.index, "k2")
	assert.Equal(t, 2, m.index["k3"])

	assert.True(t, m.Remove("k1"))
	assert.Equal(t, 1, m.Size())
	assert.Equal(t, 2, m.removed)
	assert.Len(t, m.entries, 3)
	assert.Nil(t, m.entries[0])
	assert.Nil(t, m.entries[1])
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[2])
	assert.NotContains(t, m.index, "k1")
	assert.NotContains(t, m.index, "k2")
	assert.Equal(t, 2, m.index["k3"])

	assert.True(t, m.Remove("k3"))
	assert.Equal(t, 0, m.Size())
	assert.Equal(t, 3, m.removed)
	assert.Len(t, m.entries, 3)
	assert.Nil(t, m.entries[0])
	assert.Nil(t, m.entries[1])
	assert.Nil(t, m.entries[2])
	assert.NotContains(t, m.index, "k1")
	assert.NotContains(t, m.index, "k2")
	assert.NotContains(t, m.index, "k3")

	assert.False(t, m.Remove("k3"))
	assert.Equal(t, 0, m.Size())
	assert.Equal(t, 3, m.removed)
	assert.Len(t, m.entries, 3)
	assert.Nil(t, m.entries[0])
	assert.Nil(t, m.entries[1])
	assert.Nil(t, m.entries[2])
	assert.NotContains(t, m.index, "k1")
	assert.NotContains(t, m.index, "k2")
	assert.NotContains(t, m.index, "k3")
}

func TestBaseStableMap_Clear(t *testing.T) {
	m := StableStringMap{}

	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.index)
	assert.Equal(t, 0, m.removed)

	m.Put("k1", "v1")
	m.Put("k2", "v2")
	m.Put("k3", "v3")
	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.index)
	assert.Equal(t, 0, m.removed)

	m.Put("k1", "v1")
	m.Put("k2", "v2")
	m.Put("k3", "v3")
	m.Remove("k1")
	m.Clear()
	assert.Nil(t, m.entries)
	assert.Nil(t, m.index)
	assert.Equal(t, 0, m.removed)
}

func TestBaseStableMap_Rebuild(t *testing.T) {
	m := StableStringMap{}

	assert.False(t, m.Rebuild())

	m.Put("k3", "v3")
	m.Put("k2", "v2")
	m.Put("k1", "v1")
	assert.False(t, m.Rebuild())

	m.Remove("k2")
	assert.True(t, m.Rebuild())
	assert.Len(t, m.entries, 2)
	assert.Equal(t, &StableStringMapEntry{"k3", "v3"}, m.entries[0])
	assert.Equal(t, &StableStringMapEntry{"k1", "v1"}, m.entries[1])
	assert.Equal(t, 0, m.index["k3"])
	assert.Equal(t, 1, m.index["k1"])
	assert.NotContains(t, m.index, "k2")
	assert.Equal(t, 2, m.Size())
}
