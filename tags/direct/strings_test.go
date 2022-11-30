package direct

import (
	"bytes"
	"io"
	"math/rand"
	"testing"
	"unicode/utf8"

	"github.com/interlockledger/go-iltags/ilint"
	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/interlockledger/go-iltags/tagtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringTagSize(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := tagtest.GenerateRandomString()
		require.True(t, utf8.ValidString(s))
		size := ilint.EncodedSize(id.UInt64()) +
			ilint.EncodedSize(uint64(len(s))) +
			len(s)
		assert.Equal(t, uint64(size), StringTagSize(id, s))
	}
}

func TestSerializeStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := tagtest.GenerateRandomString()
		b := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, b))

		exp := bytes.NewBuffer(nil)
		assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
		assert.Nil(t, serialization.WriteILInt(exp, uint64(len(s))))
		assert.Nil(t, serialization.WriteBytes(exp, []byte(s)))

		assert.Equal(t, exp.Bytes(), b.Bytes())
	}

	id := tags.TagID(256)
	s := "123456"
	w := tagtest.NewLimitedWriter(1, false)
	assert.NotNil(t, SerializeStringTag(id, s, w))

	w = tagtest.NewLimitedWriter(2, false)
	assert.NotNil(t, SerializeStringTag(id, s, w))

	w = tagtest.NewLimitedWriter(3, false)
	assert.NotNil(t, SerializeStringTag(id, s, w))
}

func TestDeserializeStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.TagID(rand.Uint64())
		s := tagtest.GenerateRandomString()

		exp := bytes.NewBuffer(nil)
		assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
		assert.Nil(t, serialization.WriteILInt(exp, uint64(len(s))))
		assert.Nil(t, serialization.WriteBytes(exp, []byte(s)))
		assert.Nil(t, serialization.WriteUInt8(exp, 0))
		serialized := exp.Bytes()

		r := &io.LimitedReader{R: bytes.NewReader(serialized), N: int64(len(serialized) - 1)}
		a, err := DeserializeStringTag(id, r)
		assert.Nil(t, err)
		assert.Equal(t, s, a)
		assert.Equal(t, int64(0), r.N)

		r = &io.LimitedReader{R: bytes.NewReader(serialized), N: int64(len(serialized))}
		a, err = DeserializeStringTag(id, r)
		assert.Nil(t, err)
		assert.Equal(t, s, a)
		assert.Equal(t, int64(1), r.N)

		// Errors
		idSize := ilint.EncodedSize(id.UInt64())
		sizeSize := ilint.EncodedSize(uint64(len(s)))

		re := bytes.NewReader(serialized[:idSize-1])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		// Truncate the last byte
		re = bytes.NewReader(serialized[:idSize+sizeSize-2])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		re = bytes.NewReader(serialized[:idSize+sizeSize])
		a, err = DeserializeStringTag(id, re)
		assert.Error(t, err)
		assert.Equal(t, "", a)

		re = bytes.NewReader(serialized)
		a, err = DeserializeStringTag(tags.TagID(id+1), re)
		assert.ErrorIs(t, err, tags.ErrUnexpectedTagId)
		assert.Equal(t, "", a)
	}

	exp := bytes.NewBuffer(nil)
	id := tags.TagID(1234)
	assert.Nil(t, serialization.WriteILInt(exp, id.UInt64()))
	assert.Nil(t, serialization.WriteILInt(exp, uint64(tags.MAX_TAG_SIZE+1)))
	re := bytes.NewReader(exp.Bytes())
	a, err := DeserializeStringTag(id, re)
	assert.ErrorIs(t, err, tags.ErrTagTooLarge)
	assert.Equal(t, "", a)
}

func TestStdStringTagSize(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := tagtest.GenerateRandomString()
		require.True(t, utf8.ValidString(s))
		assert.Equal(t, StdStringTagSize(s), StringTagSize(id, s))
	}
}

func TestSerializeStdStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := tagtest.GenerateRandomString()
		exp := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, exp))

		b := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStdStringTag(s, b))
		assert.Equal(t, exp.Bytes(), b.Bytes())
	}

	b := tagtest.NewLimitedWriter(2, false)
	assert.Error(t, SerializeStdStringTag("1234", b))
}

func TestDeserializeStdStringTag(t *testing.T) {

	for i := 0; i < 10; i++ {
		id := tags.IL_STRING_TAG_ID
		s := tagtest.GenerateRandomString()
		serialized := bytes.NewBuffer(nil)
		assert.Nil(t, SerializeStringTag(id, s, serialized))

		a, err := DeserializeStdStringTag(bytes.NewReader(serialized.Bytes()))
		assert.Nil(t, err)
		assert.Equal(t, s, a)
	}

	// Bad String tag or anything else.
	id := tags.IL_STRING_TAG_ID + 1
	s := tagtest.GenerateRandomString()
	serialized := bytes.NewBuffer(nil)
	assert.Nil(t, SerializeStringTag(id, s, serialized))
	a, err := DeserializeStdStringTag(bytes.NewReader(serialized.Bytes()))
	assert.Error(t, err)
	assert.Equal(t, "", a)
}
