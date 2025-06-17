package internal

import "strconv"

type Uint uint

func (u Uint) ToString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u *Uint) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*u = Uint(v)
	return nil
}

func (u Uint) MarshalText() ([]byte, error) {
	return marshalTextViaToString(u)
}

func (u *Uint) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(u, text)
}
