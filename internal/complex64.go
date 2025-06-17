package internal

import "strconv"

type Complex64 complex64

func (c Complex64) ToString() (string, error) {
	return strconv.FormatComplex(complex128(c), 'f', -1, 64), nil
}

func (c *Complex64) FromString(s string) error {
	v, err := strconv.ParseComplex(s, 64)
	if err != nil {
		return err
	}
	*c = Complex64(v)
	return nil
}

func (c Complex64) MarshalText() ([]byte, error) {
	return marshalTextViaToString(c)
}

func (c *Complex64) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(c, text)
}
