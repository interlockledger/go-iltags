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

package wrapped

import (
	"io"

	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/direct"
)

/*
A WrappedValueTag is an extension of the tags.ILTag interface that defines tags
that use an arbitrary type to store its information.

Types that implement this interface must be able to use the wrapped value as a
read-only value for serialization.

On deserialization, the existing wrapped value, if any, is discarded (set to nil)
and a new value with the deserialized data must be created. This means that
the deserialization process will always create a new instance of the wrapped
value when the deserialization is successuful.

This behavior allows instances of WrappedValueTag to act as both plain tags and
also serializer/deserializers of other types.

This abstraction is specially useful when writing code to serialize external data
int ILTags format.

In order to simplify the usage of this abstraction, the implementation of the
methods from ILTagPayload are free to panic() if they are called when the
wrapped value is not set.

V may be any type as long as the read-only serialization, create new
deserialization rule holds.
*/
type WrappedValueTag[V any] interface {
	tags.ILTag
	/*
		Returns the wrapped value.
	*/
	Wrapped() *V
	/*
		Sets the wrapped value.
	*/
	SetWrapped(value *V)
}

/*
Computes the size of a sequence of WrappedValueTag tags using a list of wrapped
values.

The provided tag must be a implementation of the WrappedValueTag that
encapsulates the values. This instance will be reused to compute the size of the
tag for each value.

The wrapped value of tag will always be set to nil at the end of the process.
*/
func WrappedValueTagsSize[V any](tag WrappedValueTag[V], values ...*V) (size uint64) {
	defer func() {
		tag.SetWrapped(nil)
	}()
	for _, v := range values {
		tag.SetWrapped(v)
		size += tags.ILTagSize(tag)
	}
	return
}

/*
Computes the size of a sequence of WrappedValueTag tags using a list of wrapped
values. If the value is nil, the size of a NullTag will be used instead.

The provided tag must be a implementation of the WrappedValueTag that
encapsulates the values. This instance will be reused to compute the size of the
tag for each value.

The wrapped value of tag will always be set to nil at the end of the process.
*/
func WrappedValueTagsSizeOrNull[V any](tag WrappedValueTag[V], values ...*V) (size uint64) {
	defer func() {
		tag.SetWrapped(nil)
	}()
	for _, v := range values {
		if v != nil {
			tag.SetWrapped(v)
			size += tags.ILTagSize(tag)
		} else {
			size += 1
		}
	}
	return
}

func SerializeWrappedValueTags[V any](tag WrappedValueTag[V], writer io.Writer, values ...*V) error {
	defer func() {
		tag.SetWrapped(nil)
	}()
	for _, v := range values {
		tag.SetWrapped(v)
		if err := tags.ILTagSeralize(tag, writer); err != nil {
			return err
		}
	}
	return nil
}

func SerializeWrappedValueTagsOrNull[V any](tag WrappedValueTag[V], writer io.Writer, values ...*V) error {
	defer func() {
		tag.SetWrapped(nil)
	}()
	for _, v := range values {
		if v != nil {
			tag.SetWrapped(v)
			if err := tags.ILTagSeralize(tag, writer); err != nil {
				return err
			}
		} else {
			if err := direct.SerializeStdNullTag(writer); err != nil {
				return err
			}
		}
	}
	return nil
}

/*
Deserializes a sequence of WrappedValueTag tags and returns an array of pointers
to the values deserialized by the provided tag implementation. After all bytes
were used or an error occurs.

The wrapped value of tag will always be set to nil at the end of the process.
*/
func DeserializeWrappedValueTags[V any](factory tags.ILTagFactory, tag WrappedValueTag[V], size int, reader io.Reader) ([]*V, error) {
	defer func() {
		tag.SetWrapped(nil)
	}()
	r := &io.LimitedReader{N: int64(size), R: reader}
	l := make([]*V, 0)
	for r.N > 0 {
		if err := tags.ILTagDeserializeTagInTo(factory, r, tag); err != nil {
			return nil, err
		}
		l = append(l, tag.Wrapped())
	}
	return l, nil
}

/*
Deserializes a sequence of WrappedValueTag tags and returns an array of pointers
to the values deserialized by the provided tag implementation. It is similar to
DeserializeWrappedValueTags() but it also accept NullTags during the
deserialization.

Those NullTags are mapped to a nil in the list of the returned values.

The wrapped value of tag will always be set to nil at the end of the process.
*/
func DeserializeWrappedValueTagsOrNull[V any](factory tags.ILTagFactory, tag WrappedValueTag[V], size int, reader io.Reader) ([]*V, error) {
	defer func() {
		tag.SetWrapped(nil)
	}()
	r := &io.LimitedReader{N: int64(size), R: reader}
	l := make([]*V, 0)
	for r.N > 0 {
		isNull, err := tags.ILTagDeserializeIntoOrNull(factory, r, tag)
		if err != nil {
			return nil, err
		}
		if isNull {
			tag.SetWrapped(nil)
		}
		l = append(l, tag.Wrapped())
	}
	return l, nil
}
