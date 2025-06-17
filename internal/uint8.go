package internal

import "strconv"

type Uint8 uint8

func (u Uint8) ToString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u *Uint8) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return err
	}
	*u = Uint8(v)
	return nil
}
