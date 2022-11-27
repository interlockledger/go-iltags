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

package ext

import (
	"fmt"
	"io"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/impl"
)

/*
This is the interface for the versioned ILTag payload. It must be able to
serialize and deserialize the value into/from a byte stream.
*/
type VersionedPayloadData interface {
	/*
		Returns the current serialization version of the payload. It is called
		by VersionedPayload.SerializeValue() to determine what version should be
		written in the version field.
	*/
	Version() uint16

	/*
		Returns true if the given version is supported by this instance. It is
		called by by VersionedPayload.DeserializeValue() to determine if the
		version is supported by the payload before DeserializeValue() is called.
	*/
	SupportedVersion(version uint16) bool

	/*
		Returns the size of the serialized value in bytes. It is called by
		VersionedPayload.Size() to determine the size of the serialization.

		The returned size should not include the size of the version field.
	*/
	Size() uint64

	/*
		Serializes the value into a Writer. It is called by
		VersionedPayload.Serialize() just after the version field is
		serialized so it can complete the serialization process.

		This process should not include the version of the serialization.
	*/
	Serialize(writer io.Writer) error

	/*
		Deserializes the value stored in a Reader. It is called by
		VersionedPayload.Deserialize() just after the version field is
		deserialized and tested in order to deserialize the data.

		Since the field version has already been tested when this method is
		called, there is no need to test it again unless it is relevant to the
		deserialization process.

		It is very important to notice that this method is responsible to
		correctly deserialize the data for all versions for which
		SupportedVersion() return true.

		This method is required to consume all data passed to it otherwise
		VersionedPayload.Deserialize() will return an error to the caller.
	*/
	Deserialize(version uint16, factory tags.ILTagFactory, valueSize int,
		reader *io.LimitedReader) error
}

/*
VersionedPayload is a generic Implementation of ILTagPayload that encapsulates a
strutct that implements the interface VersionedPayloadData. It can be used to
easly implement payloads that have multiple serialization versions.

The serialized format will always be a version represented by an uint16 followed
by the serialization of the struct that implements the VersionedPayloadData
interface.

It can be declared as:

	p := VersionedPayload[*MyVersionedPayload]{Data: &{MyVersionedPayload}}

where MyVersionedPayload is a struct that implements the VersionedPayloadData
interface.

This payload is than used implement a VersionedPayloadTag.
*/
type VersionedPayload[T VersionedPayloadData] struct {
	/*
		The instance of VersionedPayloadData that provides the actual data. It
		must be manually initialize of the methods of this struct will panic
		with access to a nil poiter.
	*/
	Data T
}

// Implementation of ILTagPayload.ValueSize()
func (p *VersionedPayload[T]) ValueSize() uint64 {
	return 2 + p.Data.Size()
}

// Implementation of ILTagPayload.SerializeValue()
func (p *VersionedPayload[T]) SerializeValue(writer io.Writer) error {
	if err := serialization.WriteUInt16(writer, p.Data.Version()); err != nil {
		return err
	}
	return p.Data.Serialize(writer)
}

// Implementation of ILTagPayload.DeserializeValue()
func (p *VersionedPayload[T]) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 2 {
		return tags.ErrBadTagFormat
	}
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	version, err := serialization.ReadUInt16(r)
	if err != nil {
		return err
	}
	if !p.Data.SupportedVersion(uint16(version)) {
		return fmt.Errorf("unsupported version %d: %w", version, tags.ErrBadTagFormat)
	}
	if err := p.Data.Deserialize(version, factory, valueSize-2, r); err != nil {
		return err
	}
	if r.N != 0 {
		return tags.ErrBadTagFormat
	}
	return nil
}

/*
VersionedPayloadTag is a generic tag that stores a versioned payload. The
version is stored as an uint16 value while the actual data serialization is
provided by the struct that implements the interface VersionedPayloadData.

Since it is not a standard tag it does not have a Standard tag ID associated
with it.

One of the ways to define its type is:

	type MyVersionedTag = VersionedPayloadTag[*MyVersionedPayload]

where MyVersionedPayload is a struct that implements the VersionedPayloadData
inteface.
*/
type VersionedPayloadTag[T VersionedPayloadData] struct {
	impl.ILTagHeaderImpl
	VersionedPayload[T]
}

/*
Create a new VersionedPayloadTag. It takes a tag id and a pointer to the
instance of the struct that implements the VersionedPayloadData interface.

This function panics if the provided id is reserved for implicit tags or data is
nil.
*/
func NewVersionedPayloadTag[T VersionedPayloadData](id tags.TagID, data T) *VersionedPayloadTag[T] {
	if id.Implicit() {
		panic("This tag cannot have an implicit tag id.")
	}
	var t VersionedPayloadTag[T]
	t.SetId(id)
	t.Data = data
	t.Data.Version()
	return &t
}
