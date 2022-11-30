package direct

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tagtest"
	"github.com/stretchr/testify/assert"
)

var SAMPLE_TAG_IDS = []tags.TagID{
	0x12,
	0x1234,
	0x123456,
	0x12345678,
	0x1234567890,
	0x1234567890AB,
	0x1234567890ABCD,
	0x1234567890ABCDEF,
	0xFFFFFFFFFFFFFF0A,
}

var SAMPLE_TAG_SIZES = []uint64{
	0x12,
	0x1234,
	0x123456,
	0x12345678,
	0x1234567890,
	0x1234567890AB,
	0x1234567890ABCD,
	0x1234567890ABCDEF,
	0xFFFFFFFFFFFFFF0A,
}

func TestGetSmallValueExplicitTagSize(t *testing.T) {
	// Test borders
	assert.Equal(t, tags.GetExplicitTagSize(16, 0),
		getSmallValueExplicitTagSize(16, 0))
	assert.Equal(t, tags.GetExplicitTagSize(16, 0xF7),
		getSmallValueExplicitTagSize(16, 0xF7))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF7),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF7))

	// Test why it fails.
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF8),
		getSmallValueExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 0xF8)+1)
}

func TestExplicitXXTagSize(t *testing.T) {

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitBoolTagSize(16))

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitInt8TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 1),
		ExplicitInt8TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 1),
		ExplicitUInt8TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 1),
		ExplicitUInt8TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 2),
		ExplicitInt16TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 2),
		ExplicitInt16TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 2),
		ExplicitUInt16TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 2),
		ExplicitUInt16TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitInt32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitInt32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitUInt32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitUInt32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitInt64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitInt64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitUInt64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitUInt64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 4),
		ExplicitFloat32TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 4),
		ExplicitFloat32TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 8),
		ExplicitFloat64TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 8),
		ExplicitFloat64TagSize(0xFFFF_FFFF_FFFF_FFFF))

	assert.Equal(t, tags.GetExplicitTagSize(16, 16),
		ExplicitFloat128TagSize(16))
	assert.Equal(t, tags.GetExplicitTagSize(0xFFFF_FFFF_FFFF_FFFF, 16),
		ExplicitFloat128TagSize(0xFFFF_FFFF_FFFF_FFFF))
}

func TestILIntTagSize(t *testing.T) {

	v := uint64(0)
	tagId := tags.TagID(16)
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.EncodedSize(v))),
		ExplicitILIntTagSize(16, v))
	assert.Equal(t, uint64(ilint.EncodedSize(v))+1,
		StdILIntTagSize(v))

	v = uint64(0xFFFF_FFFF_FFFF_FF00)
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId,
		uint64(ilint.EncodedSize(v))),
		ExplicitILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.EncodedSize(v))+1,
		StdILIntTagSize(v))
}

func TestSignedILIntTagSize(t *testing.T) {

	v := int64(0)
	tagId := tags.TagID(16)
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = 1
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = -1
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = 9223372036854775807
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))

	v = -9223372036854775808
	tagId = 0xFFFF_FFFF_FFFF_FFFF
	assert.Equal(t, tags.GetExplicitTagSize(tagId, uint64(ilint.SignedEncodedSize(v))),
		ExplicitSignedILIntTagSize(tagId, v))
	assert.Equal(t, uint64(ilint.SignedEncodedSize(v))+1,
		StdSignedILIntTagSize(v))
}

func TestSerializeTagId(t *testing.T) {

	for _, id := range SAMPLE_TAG_IDS {
		w := bytes.NewBuffer(nil)
		assert.Nil(t, serializeTagId(id, w))
		assert.Equal(t, ilint.Encode(id.UInt64(), nil), w.Bytes())
	}
}

func TestSerializeSmallValueTagHeader(t *testing.T) {

	for _, id := range SAMPLE_TAG_IDS {
		w := bytes.NewBuffer(nil)
		assert.Nil(t, serializeSmallValueTagHeader(id, 0xf7, w))
		exp := bytes.NewBuffer(nil)
		_, err := ilint.EncodeToWriter(id.UInt64(), exp)
		assert.Nil(t, err)
		assert.Nil(t, exp.WriteByte(0xf7))
		assert.Equal(t, exp.Bytes(), w.Bytes())
	}

	w := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, serializeSmallValueTagHeader(0xF8, 0xf7, w), io.ErrShortWrite)
}

func TestSerializeStandardBoolTag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardBoolTag(false, w))
	assert.Equal(t, []byte{0x1, 0x0}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardBoolTag(true, w))
	assert.Equal(t, []byte{0x1, 0x1}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt8Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt8Tag(-4, w))
	assert.Equal(t, []byte{0x2, 0xfc}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt8Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardUInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardUInt8Tag(0xFA, w))
	assert.Equal(t, []byte{0x3, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt8Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt8Tag(0x16, 0xFA, w))
	assert.Equal(t, []byte{
		0x16,
		0x1,
		0xfa}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt8Tag(0x1234567890ABCDEF, 0xFA, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x1,
		0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt8Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt8Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt8Tag(0x16, 123, w))
	assert.Equal(t, []byte{
		0x16,
		0x1,
		0x7b}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt8Tag(0x1234567890ABCDEF, -123, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x1, 0x85}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt8Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt16Tag(-1234, w))
	assert.Equal(t, []byte{0x4, 0xfb, 0x2e}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardUInt16Tag(0xFA, w))
	assert.Equal(t, []byte{0x5, 0x0, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x16, 0xFA, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x0, 0xfa}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x1234567890ABCDEF, 0xFA, w))
	assert.Equal(t, []byte{0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0x0, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt16Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x16, 123, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x0, 0x7b}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x1234567890ABCDEF, -1234, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0xfb, 0x2e}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt16Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}
