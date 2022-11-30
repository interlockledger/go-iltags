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

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------

func TestSerializeStandardInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt16Tag(0x0123, w))
	assert.Equal(t, []byte{
		0x4,
		0x1, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt16Tag(-0x0123, w))
	assert.Equal(t, []byte{
		0x4,
		0xfe, 0xdd}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardUInt16Tag(0xFA, w))
	assert.Equal(t, []byte{
		0x5,
		0x0, 0xfa}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt16Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x16, 0x0123, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x01, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt16Tag(0x1234567890ABCDEF, 0x0123, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0x01, 0x23}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt16Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt16Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x16, 0x0123, w))
	assert.Equal(t, []byte{
		0x16,
		0x2,
		0x01, 0x23}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt16Tag(0x1234567890ABCDEF,
		-0x0123, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x2,
		0xfe, 0xdd}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt16Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStandardInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt32Tag(0x01234567, w))
	assert.Equal(t, []byte{0x6,
		0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt32Tag(-0x01234567, w))
	assert.Equal(t, []byte{0x6,
		0xfe, 0xdc, 0xba, 0x99}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt32Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardUInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardUInt32Tag(0x01234567, w))
	assert.Equal(t, []byte{
		0x7,
		0x01, 0x23, 0x45, 0x67}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt32Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt32Tag(0x16, 0x01234567, w))
	assert.Equal(t, []byte{
		0x16,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt32Tag(0x1234567890ABCDEF, 0x01234567, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt32Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt32Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt32Tag(0x16, 0x01234567, w))
	assert.Equal(t, []byte{
		0x16,
		0x4,
		0x1, 0x23, 0x45, 0x67}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt32Tag(0x1234567890ABCDEF, -0x01234567, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x4,
		0xfe, 0xdc, 0xba, 0x99}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt32Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}

//------------------------------------------------------------------------------

func TestSerializeStandardInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt64Tag(0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardInt64Tag(-0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{0x8,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt64Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeStandardUInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStandardUInt64Tag(0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x9,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeStandardUInt64Tag(0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeUInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt64Tag(0x16, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x16,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeUInt64Tag(0x1234567890ABCDEF, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeUInt64Tag(0xFA, 0xFA, w1),
		io.ErrShortWrite)
}

func TestSerializeInt64Tag(t *testing.T) {

	w := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt64Tag(0x16, 0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0x16,
		0x8,
		0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, w.Bytes())

	w = bytes.NewBuffer(nil)
	assert.Nil(t, SerializeInt64Tag(0x1234567890ABCDEF, -0x0123456789ABCDEF, w))
	assert.Equal(t, []byte{
		0xff, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcc, 0xf7,
		0x8,
		0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}, w.Bytes())

	w1 := tagtest.NewLimitedWriter(1, false)
	assert.ErrorIs(t, SerializeInt64Tag(0xFA, 123, w1),
		io.ErrShortWrite)
}
