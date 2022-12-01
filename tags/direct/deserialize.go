package direct

import (
	"io"
	"math"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
)

// Deserializes the tagId and verifies if it matches the expected tag ID.
func deserializeTagId(expectedId tags.TagID, reader io.Reader) (err error) {
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
	if err := deserializeTagId(expectedId, reader); err != nil {
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
	if err := deserializeTagId(expectedId, reader); err != nil {
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

//------------------------------------------------------------------------------

/*
Deserializes a standard NullTag.
*/
func DeserializeStdNullTag(reader io.Reader) error {
	if v, err := serialization.ReadUInt8(reader); err != nil {
		return err
	} else if v != 0 {
		return tags.ErrUnexpectedTagId
	}
	return nil
}

/*
Deserializes a standard NullTag with the given Id.
*/
func DeserializeNullTag(tagId tags.TagID, reader io.Reader) error {
	return deserializeSmallValueHeaderWithSize(tagId, 0, reader)
}

//------------------------------------------------------------------------------

/*
Deserializes a standard BoolTag.
*/
func DeserializeStdBoolTag(reader io.Reader) (bool, error) {
	v, err := deserializeStdUInt8TagCore(tags.IL_BOOL_TAG_ID, reader)
	if err != nil {
		return false, err
	}
	if v == 0 {
		return false, nil
	} else if v == 1 {
		return true, nil
	} else {
		return false, tags.ErrBadTagFormat
	}
}

/*
Deserializes a BoolTag.
*/
func DeserializeBoolTag(tagId tags.TagID, reader io.Reader) (bool, error) {
	v, err := DeserializeUInt8Tag(tagId, reader)
	if err != nil {
		return false, err
	}
	if v == 0 {
		return false, nil
	} else if v == 1 {
		return true, nil
	} else {
		return false, tags.ErrBadTagFormat
	}
}

//------------------------------------------------------------------------------

func deserializeStdUInt8TagCore(tagId tags.TagID, reader io.Reader) (uint8, error) {
	if err := deserializeTagId(tagId, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt8(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt8Tag.
*/
func DeserializeStdUInt8Tag(reader io.Reader) (uint8, error) {
	return deserializeStdUInt8TagCore(tags.IL_UINT8_TAG_ID, reader)
}

/*
Deserializes an UInt8Tag.
*/
func DeserializeUInt8Tag(tagId tags.TagID, reader io.Reader) (uint8, error) {
	if err := deserializeSmallValueHeaderWithSize(tagId, 1, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt8(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt8Tag.
*/
func DeserializeStdInt8Tag(reader io.Reader) (int8, error) {
	if v, err := deserializeStdUInt8TagCore(tags.IL_INT8_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return int8(v), nil
	}
}

/*
Deserializes an UInt8Tag.
*/
func DeserializeInt8Tag(tagId tags.TagID, reader io.Reader) (int8, error) {
	if v, err := DeserializeUInt8Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return int8(v), nil
	}
}

//------------------------------------------------------------------------------

func deserializeStdUInt16TagCore(tagId tags.TagID, reader io.Reader) (uint16, error) {
	if err := deserializeTagId(tagId, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt16(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt16Tag.
*/
func DeserializeStdUInt16Tag(reader io.Reader) (uint16, error) {
	return deserializeStdUInt16TagCore(tags.IL_UINT16_TAG_ID, reader)
}

/*
Deserializes an UInt16Tag.
*/
func DeserializeUInt16Tag(tagId tags.TagID, reader io.Reader) (uint16, error) {
	if err := deserializeSmallValueHeaderWithSize(tagId, 2, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt16(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt16Tag.
*/
func DeserializeStdInt16Tag(reader io.Reader) (int16, error) {
	if v, err := deserializeStdUInt16TagCore(tags.IL_INT16_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return int16(v), nil
	}
}

/*
Deserializes an UInt16Tag.
*/
func DeserializeInt16Tag(tagId tags.TagID, reader io.Reader) (int16, error) {
	if v, err := DeserializeUInt16Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return int16(v), nil
	}
}

//------------------------------------------------------------------------------

func deserializeStdUInt32TagCore(tagId tags.TagID, reader io.Reader) (uint32, error) {
	if err := deserializeTagId(tagId, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt32(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt32Tag.
*/
func DeserializeStdUInt32Tag(reader io.Reader) (uint32, error) {
	return deserializeStdUInt32TagCore(tags.IL_UINT32_TAG_ID, reader)
}

/*
Deserializes an UInt32Tag.
*/
func DeserializeUInt32Tag(tagId tags.TagID, reader io.Reader) (uint32, error) {
	if err := deserializeSmallValueHeaderWithSize(tagId, 4, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt32(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt32Tag.
*/
func DeserializeStdInt32Tag(reader io.Reader) (int32, error) {
	if v, err := deserializeStdUInt32TagCore(tags.IL_INT32_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

/*
Deserializes an UInt32Tag.
*/
func DeserializeInt32Tag(tagId tags.TagID, reader io.Reader) (int32, error) {
	if v, err := DeserializeUInt32Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return int32(v), nil
	}
}

//------------------------------------------------------------------------------

func deserializeStdUInt64TagCore(tagId tags.TagID, reader io.Reader) (uint64, error) {
	if err := deserializeTagId(tagId, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt64(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt64Tag.
*/
func DeserializeStdUInt64Tag(reader io.Reader) (uint64, error) {
	return deserializeStdUInt64TagCore(tags.IL_UINT64_TAG_ID, reader)
}

/*
Deserializes an UInt64Tag.
*/
func DeserializeUInt64Tag(tagId tags.TagID, reader io.Reader) (uint64, error) {
	if err := deserializeSmallValueHeaderWithSize(tagId, 8, reader); err != nil {
		return 0, err
	}
	if v, err := serialization.ReadUInt64(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a standard UInt64Tag.
*/
func DeserializeStdInt64Tag(reader io.Reader) (int64, error) {
	if v, err := deserializeStdUInt64TagCore(tags.IL_INT64_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return int64(v), nil
	}
}

/*
Deserializes an UInt64Tag.
*/
func DeserializeInt64Tag(tagId tags.TagID, reader io.Reader) (int64, error) {
	if v, err := DeserializeUInt64Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return int64(v), nil
	}
}

//------------------------------------------------------------------------------

/*
Deserializes a standard Float32Tag.
*/
func DeserializeStdFloat32Tag(reader io.Reader) (float32, error) {
	if v, err := deserializeStdUInt32TagCore(tags.IL_BIN32_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return math.Float32frombits(v), nil
	}
}

/*
Deserializes an Float32Tag.
*/
func DeserializeFloat32Tag(tagId tags.TagID, reader io.Reader) (float32, error) {
	if v, err := DeserializeUInt32Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return math.Float32frombits(v), nil
	}
}

//------------------------------------------------------------------------------

/*
Deserializes a standard Float64Tag.
*/
func DeserializeStdFloat64Tag(reader io.Reader) (float64, error) {
	if v, err := deserializeStdUInt64TagCore(tags.IL_BIN64_TAG_ID, reader); err != nil {
		return 0, err
	} else {
		return math.Float64frombits(v), nil
	}
}

/*
Deserializes an Float64Tag.
*/
func DeserializeFloat64Tag(tagId tags.TagID, reader io.Reader) (float64, error) {
	if v, err := DeserializeUInt64Tag(tagId, reader); err != nil {
		return 0, err
	} else {
		return math.Float64frombits(v), nil
	}
}

//------------------------------------------------------------------------------

/*
Deserializes a standard Float128Tag.
*/
func DeserializeStdFloat128Tag(reader io.Reader) ([]byte, error) {
	if err := deserializeTagId(tags.IL_BIN128_TAG_ID, reader); err != nil {
		return nil, err
	}
	v := make([]byte, 16)
	_, err := reader.Read(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

/*
Deserializes an Float128Tag.
*/
func DeserializeFloat128Tag(tagId tags.TagID, reader io.Reader) ([]byte, error) {
	if err := deserializeSmallValueHeaderWithSize(tagId, 16, reader); err != nil {
		return nil, err
	}
	v := make([]byte, 16)
	_, err := reader.Read(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

//------------------------------------------------------------------------------

/*
Deserializes a standard ILIntTag
*/
func DeserializeStdILIntTag(reader io.Reader) (uint64, error) {
	if err := deserializeTagId(tags.IL_ILINT_TAG_ID, reader); err != nil {
		return 0, err
	}
	if v, _, err := ilint.DecodeFromReader(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes an ILIntTag.
*/
func DeserializeILIntTag(tagId tags.TagID, reader io.Reader) (uint64, error) {
	s, err := deserializeExplicitHeader(tagId, reader)
	if err != nil {
		return 0, err
	}
	if s < 1 || s > 9 {
		return 0, tags.ErrBadTagFormat
	}
	tmp := make([]byte, int(s))
	if _, err := reader.Read(tmp); err != nil {
		return 0, err
	}
	if v, sr, err := ilint.Decode(tmp); err != nil {
		return 0, tags.ErrBadTagFormat
	} else if sr != len(tmp) {
		return 0, tags.ErrBadTagFormat
	} else {
		return v, nil
	}
}

//------------------------------------------------------------------------------

/*
Deserializes a standard SignedILIntTag
*/
func DeserializeStdSignedILIntTag(reader io.Reader) (int64, error) {
	if err := deserializeTagId(tags.IL_SIGNED_ILINT_TAG_ID, reader); err != nil {
		return 0, err
	}
	if v, _, err := ilint.DecodeSignedFromReader(reader); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

/*
Deserializes a SignedILIntTag.
*/
func DeserializeSignedILIntTag(tagId tags.TagID, reader io.Reader) (int64, error) {
	s, err := deserializeExplicitHeader(tagId, reader)
	if err != nil {
		return 0, err
	}
	if s < 1 || s > 9 {
		return 0, tags.ErrBadTagFormat
	}
	tmp := make([]byte, int(s))
	if _, err := reader.Read(tmp); err != nil {
		return 0, err
	}
	if v, sr, err := ilint.DecodeSigned(tmp); err != nil {
		return 0, tags.ErrBadTagFormat
	} else if sr != len(tmp) {
		return 0, tags.ErrBadTagFormat
	} else {
		return v, nil
	}
}

//------------------------------------------------------------------------------

/*
Deserializes a BytesTag.
*/
func DeserializeStdBytesTag(reader io.Reader) ([]byte, error) {
	return DeserializeRawTag(tags.IL_BYTES_TAG_ID, reader)
}

/*
Deserializes a RawTag.
*/
func DeserializeRawTag(tagId tags.TagID, reader io.Reader) ([]byte, error) {
	s, err := deserializeExplicitHeader(tagId, reader)
	if err != nil {
		return nil, err
	}
	tmp := make([]byte, int(s))
	if _, err := reader.Read(tmp); err != nil {
		return nil, err
	}
	return tmp, nil
}
