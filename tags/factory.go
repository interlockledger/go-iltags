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

package tags

import (
	"io"
)

/*
This is the interface of all ILTag factories. Factories are used to create
instances of concrete implementations of ILTags based on the given tag id.

It can also be used to deserialize tags from Readers and byte arrays.
*/
type ILTagFactory interface {
	/*
		Creates an unitialized tag that implements the given tag ID. It returns
		an error if the tag cannot be created by any reason (unsupported,
		unknown, etc).
	*/
	CreateTag(tagId TagID) (ILTag, error)

	/*
		Deserializes the tag found in the current position of the reader.
	*/
	Deserialize(reader io.Reader) (ILTag, error)

	/*
		Helper function that tries to deserialize the current tag into the given
		tag implementation. It fails if the tag id doesn't match or if the data
		is corrupted.
	*/
	DeserializeInto(reader io.Reader, tag ILTag) error

	/*
		Converts the given byte array into a ILTag using the given tag factory.

		This function fails if the format does not contain a tag or if the data is not
		fully used by the tag.
	*/
	FromBytes(b []byte) (ILTag, error)
}
