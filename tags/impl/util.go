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
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/utils"
)

func removeKeyEntry(l []string, k string) []string {
	i := utils.FindFirstInStringSlice(l, k)
	if i == -1 {
		return l
	} else {
		return utils.RemoveFromStringSlice(l, i)
	}
}

// This type is a specialization of the map that preserves the insertion order
// of the keys.
type StableStringMap struct {
	entries map[string]string
	keys    []string
}

// Puts a new value into this map
func (m *StableStringMap) Put(key, value string) {
	if m.entries == nil {
		m.entries = make(map[string]string)
		m.keys = make([]string, 0, 8)
	}
	_, ok := m.entries[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
	m.entries[key] = value
}

// Return the number of entries in this map.
func (m *StableStringMap) Size() int {
	return len(m.keys)
}

// Returns a list of the keys ordered by the first insertion order. Never modify
// modify or cache the contents of this slice.
func (m *StableStringMap) Keys() []string {
	return m.keys
}

// Returns the value associated with the given key and a flag that indicates if
// the value is present or not, just like a standard map.
func (m *StableStringMap) Get(key string) (string, bool) {
	s, ok := m.entries[key]
	return s, ok
}

// Removes the given value from this map.
func (m *StableStringMap) Remove(key string) {
	_, ok := m.entries[key]
	if ok {
		delete(m.entries, key)
		m.keys = removeKeyEntry(m.keys, key)
	}
}

// Remove all entries from this map. It will release all associated resources.
func (m *StableStringMap) Clear() {
	m.entries = nil
	m.keys = nil
}

// This type is a specialization of the map that preserves the insertion order
// of the keys.
type StableILTagMap struct {
	entries map[string]tags.ILTag
	keys    []string
}

// Puts a new value into this map
func (m *StableILTagMap) Put(key string, value tags.ILTag) {
	if m.entries == nil {
		m.entries = make(map[string]tags.ILTag)
		m.keys = make([]string, 0, 8)
	}
	_, ok := m.entries[key]
	if !ok {
		m.keys = append(m.keys, key)
	}
	m.entries[key] = value
}

// Return the number of entries in this map.
func (m *StableILTagMap) Size() int {
	return len(m.keys)
}

// Returns a list of the keys ordered by the first insertion order. Never modify
// modify or cache the contents of this slice.
func (m *StableILTagMap) Keys() []string {
	return m.keys
}

// Returns the value associated with the given key and a flag that indicates if
// the value is present or not, just like a standard map.
func (m *StableILTagMap) Get(key string) (tags.ILTag, bool) {
	s, ok := m.entries[key]
	return s, ok
}

// Removes the given value from this map.
func (m *StableILTagMap) Remove(key string) {
	_, ok := m.entries[key]
	if ok {
		delete(m.entries, key)
		m.keys = removeKeyEntry(m.keys, key)
	}
}

// Remove all entries from this map. It will release all associated resources.
func (m *StableILTagMap) Clear() {
	m.entries = nil
	m.keys = nil
}
