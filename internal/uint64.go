package internal

import "strconv"

type Uint64 uint64

func (u Uint64) ToString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u *Uint64) FromString(s string) error {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*u = Uint64(v)
	return nil
}
