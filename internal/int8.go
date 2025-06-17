package internal

import "strconv"

type Int8 int8

func (i Int8) ToString() (string, error) {
	return strconv.FormatInt(int64(i), 10), nil
}

func (i *Int8) FromString(s string) error {
	v, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return err
	}
	*i = Int8(v)
	return nil
}
