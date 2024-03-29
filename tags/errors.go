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

import "fmt"

var (
	// The tag is too large to be manipulated by this library.
	ErrTagTooLarge = fmt.Errorf("the given tag is too large to be handled by this library")
	// Unsupported/unknown tag id.
	ErrUnsupportedTagId = fmt.Errorf("unsupported tag ID")
	// Bad tag format
	ErrBadTagFormat = fmt.Errorf("bad tag format")
	// Unexpected tag id.
	ErrUnexpectedTagId = fmt.Errorf("unexpected tag ID")
)

// Create a new UnsupportedTagIdError with the specified tag id.
func NewErrUnsupportedTagId(id TagID) error {
	return fmt.Errorf("unsupported tag with id %d: %w", id, ErrUnsupportedTagId)
}

// Create a new NewErrUnexpectedTagId with the specified tag id.
func NewErrUnexpectedTagId(expected, id TagID) error {
	return fmt.Errorf("expecting tag with id %d but got the id %d: %w", expected, id, ErrUnexpectedTagId)
}
