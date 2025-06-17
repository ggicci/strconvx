package internal

import "strconv"

type Uint32 uint32

func (u Uint32) ToString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u *Uint32) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return err
	}
	*u = Uint32(v)
	return nil
}

func (u Uint32) MarshalText() ([]byte, error) {
	return marshalTextViaToString(u)
}

func (u *Uint32) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(u, text)
}
