package internal

import "strconv"

type Complex128 complex128

func (c Complex128) ToString() (string, error) {
	return strconv.FormatComplex(complex128(c), 'f', -1, 128), nil
}

func (c *Complex128) FromString(s string) error {
	v, err := strconv.ParseComplex(s, 128)
	if err != nil {
		return err
	}
	*c = Complex128(v)
	return nil
}

func (c Complex128) MarshalText() ([]byte, error) {
	return marshalTextViaToString(c)
}

func (c *Complex128) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(c, text)
}
