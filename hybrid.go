package strconvx

import (
	"encoding"
	"reflect"
)

type hybrid struct {
	CanToString
	CanFromString
}

func (h *hybrid) ToString() (string, error) {
	if h.CanToString != nil {
		return h.CanToString.ToString()
	}
	return "", ErrCannotToString
}

func (h *hybrid) FromString(s string) error {
	if h.CanFromString != nil {
		return h.CanFromString.FromString(s)
	}
	return ErrCannotFromString
}

func (h *hybrid) IsValid() bool {
	return h.CanToString != nil || h.CanFromString != nil
}

func (h *hybrid) validateAsComplete() error {
	if h.CanToString == nil {
		return ErrCannotToString
	}
	if h.CanFromString == nil {
		return ErrCannotFromString
	}
	return nil
}

// createHybridStringConverter tries to create a hybrid StringConverter from a
// reflect.Value. It will make the most of the interfaces rv has implemented,
// including strconvx.CanToString, strconvx.CanFromString,
// encoding.TextMarshaler, and encoding.TextUnmarshaler. Returns nil if the
// reflect.Value does not implement any of the above.
func createHybridStringConverter(rv reflect.Value) StringConverter {
	h := &hybrid{}

	// Check strconvx.CanToString and encoding.TextMarshaler.
	if rv.Type().Implements(stringMarshalerType) {
		h.CanToString = rv.Interface().(CanToString)
	} else if rv.Type().Implements(textMarshalerType) {
		h.CanToString = &textMarshaler{
			rv.Interface().(encoding.TextMarshaler),
			nil,
		}
	}

	// Check strconvx.CanFromString and encoding.TextUnmarshaler.
	if rv.Type().Implements(stringUnmarshalerType) {
		h.CanFromString = rv.Interface().(CanFromString)
	} else if rv.Type().Implements(textUnmarshalerType) {
		h.CanFromString = &textMarshaler{
			nil,
			rv.Interface().(encoding.TextUnmarshaler),
		}
	}

	if h.IsValid() {
		return h
	}

	return nil
}

type textMarshaler struct {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

func (w textMarshaler) ToString() (string, error) {
	b, err := w.TextMarshaler.MarshalText()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (w textMarshaler) FromString(s string) error {
	return w.TextUnmarshaler.UnmarshalText([]byte(s))
}

var (
	stringMarshalerType   = typeOf[CanToString]()
	stringUnmarshalerType = typeOf[CanFromString]()
	textMarshalerType     = typeOf[encoding.TextMarshaler]()
	textUnmarshalerType   = typeOf[encoding.TextUnmarshaler]()
)
