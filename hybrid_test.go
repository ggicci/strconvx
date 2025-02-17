package strconvx

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHybridCoder_TestMarshalerOnly(t *testing.T) {
	apple := &TextMarshalerApple{}
	rv := reflect.ValueOf(apple)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "apple", text)

	assert.ErrorIs(t, sb.FromString("red apple"), ErrCannotFromString)
}

func TestHybridCoder_TextUnmarshalerOnly(t *testing.T) {
	banana := &TextUnmarshalerBanana{}
	rv := reflect.ValueOf(banana)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.ErrorIs(t, err, ErrCannotToString)
	assert.Empty(t, text)

	err = sb.FromString("yellow banana")
	assert.NoError(t, err)
	assert.Equal(t, "yellow banana", banana.Content)
}

func TestHybridCoder_TestMarshaler_and_TextUnmarshaler(t *testing.T) {
	orange := &TextMarshalerAndUnmarshalerOrange{Content: "orange"}
	rv := reflect.ValueOf(orange)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "orange", text)

	err = sb.FromString("red orange")
	assert.NoError(t, err)
	assert.Equal(t, "red orange", orange.Content)
}

func TestHybridCoder_StringMarshaler_TakesPrecedence(t *testing.T) {
	peach := &StringMarshalerAndTextMarshalerPeach{Content: "peach"}
	rv := reflect.ValueOf(peach)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "ToString:peach", text)

	err = sb.FromString("red peach")
	assert.ErrorIs(t, err, ErrCannotFromString)
}

func TestHybridCoder_StringUnmarshaler_TakesPrecedence(t *testing.T) {
	peach := &StringUnmarshalerAndTextUnmarshalerPeach{Content: "peach"}
	rv := reflect.ValueOf(peach)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.ErrorIs(t, err, ErrCannotToString)
	assert.Empty(t, text)

	err = sb.FromString("red peach")
	assert.NoError(t, err)
	assert.Equal(t, "FromString:red peach", peach.Content)
}

func TestHybridCoder_StringMarshaler_and_TextUnmarshaler(t *testing.T) {
	pineapple := &StringMarshalerAndTextUnmarshalerPineapple{Content: "pineapple"}
	rv := reflect.ValueOf(pineapple)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.NoError(t, err)
	assert.Equal(t, "ToString:pineapple", text)

	err = sb.FromString("red pineapple")
	assert.NoError(t, err)
	assert.Equal(t, "UnmarshalText:red pineapple", pineapple.Content)
}

func TestHybridCoder_MarshalText_Error(t *testing.T) {
	watermelon := &TextMarshalerSpoiledWatermelon{}
	rv := reflect.ValueOf(watermelon)
	sb := createHybridStringConverter(rv)
	assert.NotNil(t, sb)

	text, err := sb.ToString()
	assert.ErrorContains(t, err, "spoiled")
	assert.Empty(t, text)
}

func TestHybridCoder_ErrCannotInterface(t *testing.T) {
	type mystruct struct {
		unexportedName string
	}
	v := mystruct{unexportedName: "mystruct"}
	rv := reflect.ValueOf(v)

	sb := createHybridStringConverter(rv.Field(0))
	assert.Nil(t, sb)
}

func TestHybridCoder_NilOnNoInterfacesDetected(t *testing.T) {
	var zero zeroInterface
	rv := reflect.ValueOf(zero)

	sb := createHybridStringConverter(rv)
	assert.Nil(t, sb)
}

// TextMarshalerApple implements:
//   - CanToString - no
//   - CanFromString - no
//   - encoding.TextMarshaler - yes
//   - encoding.TextUnmarshaler - no
type TextMarshalerApple struct{}

func (t *TextMarshalerApple) MarshalText() ([]byte, error) {
	return []byte("apple"), nil
}

// TextUnmarshalerBanana implements:
//   - CanToString - no
//   - CanFromString - no
//   - encoding.TextMarshaler - no
//   - encoding.TextUnmarshaler - yes
type TextUnmarshalerBanana struct{ Content string }

func (t *TextUnmarshalerBanana) UnmarshalText(text []byte) error {
	t.Content = string(text)
	return nil
}

// TextMarshalerAndUnmarshalerOrange implements:
//   - CanToString - no
//   - CanFromString - no
//   - encoding.TextMarshaler - yes
//   - encoding.TextUnmarshaler - yes
type TextMarshalerAndUnmarshalerOrange struct{ Content string }

func (t *TextMarshalerAndUnmarshalerOrange) MarshalText() ([]byte, error) {
	return []byte(t.Content), nil
}

func (t *TextMarshalerAndUnmarshalerOrange) UnmarshalText(text []byte) error {
	t.Content = string(text)
	return nil
}

// StringMarshalerAndTextMarshalerPeach implements:
//   - CanToString - yes
//   - CanFromString - no
//   - encoding.TextMarshaler - yes
//   - encoding.TextUnmarshaler - no
type StringMarshalerAndTextMarshalerPeach struct{ Content string }

func (s *StringMarshalerAndTextMarshalerPeach) ToString() (string, error) {
	return "ToString:" + s.Content, nil
}

func (s *StringMarshalerAndTextMarshalerPeach) MarshalText() ([]byte, error) {
	return []byte("MarshalText:" + s.Content), nil
}

// StringUnmarshalerAndTextUnmarshalerPeach implements:
//   - CanToString - no
//   - CanFromString - yes
//   - encoding.TextMarshaler - no
//   - encoding.TextUnmarshaler - yes
type StringUnmarshalerAndTextUnmarshalerPeach struct{ Content string }

func (s *StringUnmarshalerAndTextUnmarshalerPeach) FromString(text string) error {
	s.Content = "FromString:" + text
	return nil
}

func (s *StringUnmarshalerAndTextUnmarshalerPeach) UnmarshalText(text []byte) error {
	s.Content = "UnmarshalText:" + string(text)
	return nil
}

// StringMarshalerAndTextUnmarshalerPineapple implements:
//   - CanToString - yes
//   - CanFromString - no
//   - encoding.TextMarshaler - no
//   - encoding.TextUnmarshaler - yes
type StringMarshalerAndTextUnmarshalerPineapple struct{ Content string }

func (s *StringMarshalerAndTextUnmarshalerPineapple) ToString() (string, error) {
	return "ToString:" + s.Content, nil
}

func (s *StringMarshalerAndTextUnmarshalerPineapple) UnmarshalText(text []byte) error {
	s.Content = "UnmarshalText:" + string(text)
	return nil
}

// StringMarshalerAndStringUnmarshalerCherry implements:
//   - CanToString - yes
//   - CanFromString - yes
//   - encoding.TextMarshaler - no
//   - encoding.TextUnmarshaler - no
type StringMarshalerAndStringUnmarshalerCherry struct {
	Content string
}

func (s *StringMarshalerAndStringUnmarshalerCherry) FromString(text string) error {
	s.Content = "FromString:" + text
	return nil
}

func (s *StringMarshalerAndStringUnmarshalerCherry) ToString() (string, error) {
	return "ToString:" + s.Content, nil
}

// TextMarshalerSpoiledWatermelon implements:
//   - CanToString - no
//   - CanFromString - no
//   - encoding.TextMarshaler - yes (but returns error)
//   - encoding.TextUnmarshaler - no
type TextMarshalerSpoiledWatermelon struct{}

func (t *TextMarshalerSpoiledWatermelon) MarshalText() ([]byte, error) {
	return nil, errors.New("spoiled")
}

type zeroInterface struct{}
