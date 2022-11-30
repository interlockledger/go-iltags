package direct

import (
	"io"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
)

/*
Returns the size of the string tag that will hold the given string.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func StringTagSize(tagId tags.TagID, s string) uint64 {
	return tags.GetExplicitTagSize(tagId, uint64(len(s)))
}

/*
Serializes a string directly using the StringTag format.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func SerializeStringTag(tagId tags.TagID, s string, writer io.Writer) error {
	if err := serialization.WriteILInt(writer, uint64(tagId)); err != nil {
		return err
	}
	if err := serialization.WriteILInt(writer, uint64(len(s))); err != nil {
		return err
	}
	if err := serialization.WriteString(writer, s); err != nil {
		return err
	}
	return nil
}

/*
Deserializes a string tag directly into a string.

This function exists as a faster and more efficient way to deal with string tags
without using StringTag instances.
*/
func DeserializeStringTag(expectedId tags.TagID, reader io.Reader) (string, error) {
	id, err := serialization.ReadILInt(reader)
	if err != nil {
		return "", err
	}
	if tags.TagID(id) != expectedId {
		return "", tags.NewErrUnexpectedTagId(expectedId, tags.TagID(id))
	}
	size, err := serialization.ReadILInt(reader)
	if err != nil {
		return "", err
	}
	if size > tags.MAX_TAG_SIZE {
		return "", tags.ErrTagTooLarge
	}
	if s, err := serialization.ReadString(reader, int(size)); err != nil {
		return "", err
	} else {
		return s, nil
	}
}

/*
Returns the size of the standard string tag that will hold the given string.
*/
func StdStringTagSize(s string) uint64 {
	// Although it is equivalent to StringTagSize(tags.IL_STRING_TAG_ID, s),
	// this implementation uses a more direct approach to save on call to
	// ilint.EncodedSize(). It may save a lot of resources because this function
	// is used by a lot of other components in this library.
	l := uint64(len(s))
	return 1 + uint64(ilint.EncodedSize(l)) + l
}

/*
Serializes a string directly using the standard StringTag format.
*/
func SerializeStdStringTag(s string, writer io.Writer) error {
	return SerializeStringTag(tags.IL_STRING_TAG_ID, s, writer)
}

/*
Deserializes a standard string tag directly into a string.
*/
func DeserializeStdStringTag(reader io.Reader) (string, error) {
	return DeserializeStringTag(tags.IL_STRING_TAG_ID, reader)
}
