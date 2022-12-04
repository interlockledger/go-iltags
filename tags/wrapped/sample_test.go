package wrapped

import (
	"bytes"
	"io"
	"testing"

	"github.com/interlockledger/go-iltags/serialization"
	"github.com/interlockledger/go-iltags/tags"
	"github.com/stretchr/testify/assert"
)

/*
SampleWrappedPayload implements the payload of the example of the
WrappedValueTag
*/
type SampleWrappedPayload struct {
	Value *uint32
}

// Implements ILTagPayload.ValueSize().
func (p *SampleWrappedPayload) ValueSize() uint64 {
	return 4
}

// Implements ILTagPayload.SerializeValue().
func (p *SampleWrappedPayload) SerializeValue(writer io.Writer) error {
	if p.Value == nil {
		panic("Value is nil.")
	}
	return serialization.WriteUInt32(writer, *p.Value)
}

// Implements ILTagPayload.DeserializeValue().
func (p *SampleWrappedPayload) DeserializeValue(factory tags.ILTagFactory, valueSize int, reader io.Reader) error {
	v, err := serialization.ReadUInt32(reader)
	if err != nil {
		return err
	}
	value := new(uint32)
	*value = v
	p.Value = value
	return nil
}

// Creates a new SampleWrappedTag.
func NewSampleWrappedTag() *SampleWrappedTag {
	tag := &SampleWrappedTag{}
	tag.SetId(1234)
	return tag
}

/*
SampleWrappedTag is an example of a WrappedValueTag that wraps a uint32 value.
It is used to demonstrate h
*/
type SampleWrappedTag struct {
	tags.ILTagHeaderImpl
	SampleWrappedPayload
}

// Implements WrappedValueTag.Wrapped().
func (p *SampleWrappedTag) Wrapped() *uint32 {
	return p.Value
}

// Implements WrappedValueTag.SetWrapped().
func (p *SampleWrappedTag) SetWrapped(v *uint32) {
	p.Value = v
}

//------------------------------------------------------------------------------

func TestSampleWrappedTag(t *testing.T) {
	tag := NewSampleWrappedTag()

	exp := []byte{0xf9, 0x3, 0xda, 0x4, 0x0, 0x0, 0x12, 0x34}

	v1 := new(uint32)
	*v1 = 0x1234
	tag.SetWrapped(v1)
	w := bytes.NewBuffer(nil)
	assert.Nil(t, tags.ILTagSeralize(tag, w))
	assert.Equal(t, exp, w.Bytes())

	r := bytes.NewReader(exp)
	assert.Nil(t, tags.ILTagDeserializeTagInTo(nil, r, tag))
	assert.Equal(t, *v1, *tag.Value)
	assert.NotSame(t, v1, tag.Value)

	r = bytes.NewReader(exp[1:])
	assert.Error(t, tags.ILTagDeserializeTagInTo(nil, r, tag))
}
