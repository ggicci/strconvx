package strconvx

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type YesNo bool

func (yn YesNo) ToString() (string, error) {
	if yn {
		return "yes", nil
	} else {
		return "no", nil
	}
}

func (yn *YesNo) FromString(s string) error {
	switch strings.ToLower(s) {
	case "yes":
		*yn = true
	case "no":
		*yn = false
	default:
		return errors.New("invalid value")
	}
	return nil
}

func TestToAnyAdaptor(t *testing.T) {
	typ, adaptor := ToAnyAdaptor(func(b *bool) (StringCodec, error) {
		return (*YesNo)(b), nil
	})
	assert.Equal(t, typ, typeOf[bool]())

	var validBoolean bool = true
	codec, err := adaptor(&validBoolean)
	assert.NoError(t, err)
	v, err := codec.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "yes", v)
	assert.NoError(t, codec.FromString("no"))
	assert.False(t, validBoolean)

	var invalidType int = 0
	codec, err = adaptor(&invalidType)
	assert.ErrorIs(t, err, ErrTypeMismatch)
	assert.Nil(t, codec)
	assert.ErrorContains(t, err, "cannot convert *int to *bool")
}
