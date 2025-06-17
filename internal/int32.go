package internal

import "strconv"

type Int32 int32

func (i Int32) ToString() (string, error) {
	return strconv.FormatInt(int64(i), 10), nil
}

func (i *Int32) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*i = Int32(v)
	return nil
}

func (i Int32) MarshalText() ([]byte, error) {
	return marshalTextViaToString(i)
}

func (i *Int32) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(i, text)
}
