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
	"io"
	"time"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tags/direct"
	"github.com/interlockledger/go-iltags/tags/impl"
)

/*
TimestapTag is a generic timestamp tag that stores the timestamp as a signed
ILInt value in microseconds since 1970-01-01T00:00:00.000000Z.

Since it is not a standard tag it does not have a Standard tag ID associated
with it.
*/
type TimestapTag struct {
	impl.SignedILIntTag
}

/*
Create a new TimestapTag.

This function panics if the provided id is reserved for implicit tags.
*/
func NewTimestapTag(id tags.TagID) *TimestapTag {
	if id.Implicit() {
		panic("This tag cannot have an implicit tag id.")
	}
	var t TimestapTag
	t.SetId(id)
	return &t
}

/*
Sets the timestamp value. It will update TimestapTag.Payload with the
appropriate value.
*/
func (t *TimestapTag) SetTimestamp(ts time.Time) {
	t.Payload = ts.UnixMicro()
}

/*
Recovers the stored value as a time.Time instance. The returned value will be
at the local time zone.
*/
func (t *TimestapTag) GetTimestamp() time.Time {
	return time.UnixMicro(t.Payload)
}

/*
Recovers the stored value as a time.Time instance. The returned value will be
at UTC.
*/
func (t *TimestapTag) GetTimestampUTC() time.Time {
	return t.GetTimestamp().UTC()
}

/*
Serializes a TimestapTag directly into a writer. The provided tagId must belong
to an explicit tag.
*/
func SerializeTimestapTag(tagId tags.TagID, v time.Time, writer io.Writer) error {
	return direct.SerializeSignedILIntTag(tagId, v.UnixMicro(), writer)
}

/*
Deserializes a TimestapTag directly from a reader. The provided tagId must match
the expected tagId.

The returned time will be at the local timezone.
*/
func DeserializeTimestapTag(tagId tags.TagID, reader io.Reader) (time.Time, error) {
	if v, err := direct.DeserializeSignedILIntTag(tagId, reader); err != nil {
		return time.Time{}, err
	} else {
		return time.UnixMicro(v), nil
	}
}

//------------------------------------------------------------------------------

// Payload of a timestamp with timezone information. This payload does store the
// offset but not the original name of the timezone.
type TimestampTZPayload struct {
	impl.SignedILIntPayload
	/*
		Timezone offset in minutes from UTC. For example a UTC-03:00 will be
		represented as -180 while UTC+04:00 will be represented as 240.
	*/
	Offset int16
}

// Implementation of ILTagPayload.ValueSize().
func (p *TimestampTZPayload) ValueSize() uint64 {
	return p.SignedILIntPayload.ValueSize() + 2
}

// Implementation of ILTagPayload.SerializeValue()
func (p *TimestampTZPayload) SerializeValue(writer io.Writer) error {
	if err := p.SignedILIntPayload.SerializeValue(writer); err != nil {
		return err
	}
	return serialization.WriteInt16(writer, p.Offset)
}

/*
Implementation of ILTagPayload.DeserializeValue().
*/
func (p *TimestampTZPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	if valueSize < 3 {
		return tags.ErrBadTagFormat
	}
	r := &io.LimitedReader{R: reader, N: int64(valueSize)}
	if err := p.SignedILIntPayload.DeserializeValue(factory, valueSize, r); err != nil {
		return tags.ErrBadTagFormat
	}
	if r.N != 2 {
		return tags.ErrBadTagFormat
	}
	if offset, err := serialization.ReadInt16(reader); err != nil {
		return tags.ErrBadTagFormat
	} else {
		p.Offset = offset
	}
	return nil
}

//------------------------------------------------------------------------------

/*
TimestapTag is a generic timestamp tag that stores the timestamp as a signed
ILInt value in microseconds since 1970-01-01T00:00:00.000000Z and the timezone
original timezone information (except for the name of the timezone).

Since it is not a standard tag it does not have a Standard tag ID associated
with it.
*/
type TimestapTZTag struct {
	tags.ILTagHeaderImpl
	TimestampTZPayload
}

/*
Create a new TimestapTZTag.

This function panics if the provided id is reserved for implicit tags.
*/
func NewTimestapTZTag(id tags.TagID) *TimestapTZTag {
	if id.Implicit() {
		panic("This tag cannot have an implicit tag id.")
	}
	var t TimestapTZTag
	t.SetId(id)
	return &t
}

/*
Sets the timestamp value. It will update TimestapTag.Payload with the
appropriate value.
*/
func (t *TimestapTZTag) SetTimestamp(ts time.Time) {
	t.Payload = ts.UnixMicro()
	_, offs := ts.Zone()
	t.Offset = int16(offs / 60)
}

/*
Returns the timezone. Since the name is not stored by this tag, it will be left
blank.
*/
func (t *TimestapTZTag) GetTimeZone() *time.Location {
	return time.FixedZone("", int(t.Offset)*60)
}

/*
Recovers the stored value as a time.Time instance. The returned value will be
at the local time zone.
*/
func (t *TimestapTZTag) GetLocalTimestamp() time.Time {
	return time.UnixMicro(t.Payload)
}

/*
Recovers the stored value as a time.Time instance. The returned value will be
at the original timezone.
*/
func (t *TimestapTZTag) GetTimestamp() time.Time {
	return t.GetLocalTimestamp().In(t.GetTimeZone())
}

/*
Recovers the stored value as a time.Time instance. The returned value will be
at UTC.
*/
func (t *TimestapTZTag) GetTimestampUTC() time.Time {
	return t.GetTimestamp().UTC()
}
