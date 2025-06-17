package internal

import "strconv"

type Int64 int64

func (i Int64) ToString() (string, error) {
	return strconv.FormatInt(int64(i), 10), nil
}

func (i *Int64) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = Int64(v)
	return nil
}

func (i Int64) MarshalText() ([]byte, error) {
	return marshalTextViaToString(i)
}

func (i *Int64) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(i, text)
}
