package direct

import (
	"io"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
)

// Deserializes
func deserializeHeader(expectedId tags.TagID, reader io.Reader) (err error) {
	tagId, err := serialization.ReadILInt(reader)
	if err != nil {
		return err
	}
	if tagId != expectedId.UInt64() {
		return tags.NewErrUnexpectedTagId(expectedId, tags.TagID(tagId))
	}
	return nil
}

/*
Deserializes a small value header and checks if it matches the expected TagID.
*/
func deserializeSmallValueHeader(expectedId tags.TagID, reader io.Reader) (uint64, error) {
	if err := deserializeHeader(expectedId, reader); err != nil {
		return 0, err
	}
	s, err := serialization.ReadUInt8(reader)
	if err != nil {
		return 0, err
	}
	if s > 0xF7 {
		return 0, tags.ErrBadTagFormat
	}
	return uint64(s), nil
}

/*
Deserializes a small value header and checks if it matches the expected TagID and
size.
*/
func deserializeSmallValueHeaderWithSize(expectedId tags.TagID, size int, reader io.Reader) error {
	s, err := deserializeSmallValueHeader(expectedId, reader)
	if err != nil {
		return err
	}
	if s != uint64(size) {
		return tags.ErrBadTagFormat
	}
	return nil
}

/*
Deserializes an explicit header and checks if it matches the expected TagID.
*/
func deserializeExplicitHeader(expectedId tags.TagID, reader io.Reader) (uint64, error) {
	if err := deserializeHeader(expectedId, reader); err != nil {
		return 0, err
	}
	s, err := serialization.ReadILInt(reader)
	if err != nil {
		return 0, err
	}
	if s > tags.MAX_TAG_SIZE {
		return 0, tags.ErrTagTooLarge
	}
	return s, nil
}
