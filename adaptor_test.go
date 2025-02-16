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

func TestToAnyStringConverterAdaptor(t *testing.T) {
	typ, adaptor := ToAnyStringConverterAdaptor[bool](func(b *bool) (StringConverter, error) {
		return (*YesNo)(b), nil
	})
	assert.Equal(t, typ, typeOf[bool]())

	var validBoolean bool = true
	converter, err := adaptor(&validBoolean)
	assert.NoError(t, err)
	v, err := converter.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "yes", v)
	assert.NoError(t, converter.FromString("no"))
	assert.False(t, validBoolean)

	var invalidType int = 0
	converter, err = adaptor(&invalidType)
	assert.ErrorIs(t, err, ErrTypeMismatch)
	assert.Nil(t, converter)
	assert.ErrorContains(t, err, "cannot convert *int to *bool")
}
