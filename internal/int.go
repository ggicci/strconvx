package internal

import "strconv"

type Int int

func (i Int) ToString() (string, error) {
	return strconv.Itoa(int(i)), nil
}

func (i *Int) FromString(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = Int(v)
	return nil
}
