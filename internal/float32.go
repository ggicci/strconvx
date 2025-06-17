package internal

import "strconv"

type Float32 float32

func (f Float32) ToString() (string, error) {
	return strconv.FormatFloat(float64(f), 'f', -1, 32), nil
}

func (f *Float32) FromString(s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*f = Float32(v)
	return nil
}

func (f Float32) MarshalText() ([]byte, error) {
	return marshalTextViaToString(f)
}

func (f *Float32) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(f, text)
}
