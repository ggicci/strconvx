package internal

import "strconv"

type Int16 int16

func (i Int16) ToString() (string, error) {
	return strconv.FormatInt(int64(i), 10), nil
}

func (i *Int16) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return err
	}
	*i = Int16(v)
	return nil
}

func (i Int16) MarshalText() ([]byte, error) {
	return marshalTextViaToString(i)
}

func (i *Int16) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(i, text)
}
