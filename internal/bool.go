package internal

import "strconv"

type Bool bool

func (b Bool) ToString() (string, error) {
	return strconv.FormatBool(bool(b)), nil
}

func (b *Bool) FromString(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*b = Bool(v)
	return nil
}
