package internal

import "strconv"

type Uint16 uint16

func (u Uint16) ToString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u *Uint16) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return err
	}
	*u = Uint16(v)
	return nil
}
