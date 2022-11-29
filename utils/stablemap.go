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

/*
The key/value pair used by BaseStableMap.
*/
type BaseStableMapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

/*
This struct implements a generic map that preserves the insertion order of the
entries. This means that, when enumerated, the entries will be listed according
to the insertion order.

This implementation has been designed to be as efficient as possible when
compared to the built-in map.
*/
type BaseStableMap[K comparable, V any] struct {
	removed int
	entries []*BaseStableMapEntry[K, V]
	index   map[K]int
}

/*
Returns the number of entries in this map.
*/
func (m *BaseStableMap[K, V]) Size() int {
	return len(m.entries) - m.removed
}

/*
Returns true if this map is empty.
*/
func (m *BaseStableMap[K, V]) Empty() bool {
	return m.Size() == 0
}

/*
Returns a list of keys ordered by the insertion order.
*/
func (m *BaseStableMap[K, V]) Keys() []K {
	if m.Empty() {
		return make([]K, 0)
	}
	keys := make([]K, 0, len(m.entries))
	for _, e := range m.entries {
		if e != nil {
			keys = append(keys, e.Key)
		}
	}
	return keys
}

/*
Returns a list with all entries in this map. The order of the pairs are defined
by the key's first insertion order.
*/
func (m *BaseStableMap[K, V]) Entries() []BaseStableMapEntry[K, V] {
	if m.Empty() {
		return make([]BaseStableMapEntry[K, V], 0)
	}
	keys := make([]BaseStableMapEntry[K, V], 0, len(m.entries))
	for _, e := range m.entries {
		if e != nil {
			keys = append(keys, *e)
		}
	}
	return keys
}

/*
Puts a new value into this map. If a key is already in the map, the associated
value will be replaced but the original insertion order will remain unchanged.
*/
func (m *BaseStableMap[K, V]) Put(key K, value V) {
	if m.entries == nil {
		m.entries = make([]*BaseStableMapEntry[K, V], 0, 8)
		m.index = make(map[K]int)
	}
	if idx, found := m.index[key]; found {
		m.entries[idx].Value = value
	} else {
		m.entries = append(m.entries, &BaseStableMapEntry[K, V]{
			Key:   key,
			Value: value})
		m.index[key] = len(m.entries) - 1
	}
}

/*
Returns the value associated with the given key.
*/
func (m *BaseStableMap[K, V]) Get(key K) (value V, found bool) {
	if idx, found := m.index[key]; found {
		return m.entries[idx].Value, true
	}
	return
}

/*
Removes the given key from this map. If a removed key is inserted again, it will
be positioned at the end of the entry list.
*/
func (m *BaseStableMap[K, V]) Remove(key K) bool {
	if !m.Empty() {
		if idx, found := m.index[key]; found {
			delete(m.index, key)
			m.entries[idx] = nil
			m.removed++
			return true
		}
	}
	return false
}

/*
Clears this map.
*/
func (m *BaseStableMap[K, V]) Clear() {
	m.index = nil
	m.entries = nil
	m.removed = 0
}

/*
Rebuilds the internal index in order to free unused space that may be wasted
when multiple entries are removed from this map.

Returns true if the rebuild was performed or false otherwise.
*/
func (m *BaseStableMap[K, V]) Rebuild() bool {
	if m.removed == 0 {
		return false
	}
	entries := m.entries
	m.Clear()
	for _, e := range entries {
		if e != nil {
			m.Put(e.Key, e.Value)
		}
	}
	return true
}

/*
StableStringMap is a map of strings to strings that preserves the insertion
order.
*/
type StableStringMap = BaseStableMap[string, string]
type StableStringMapEntry = BaseStableMapEntry[string, string]
