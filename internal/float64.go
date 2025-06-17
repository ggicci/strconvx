package internal

import "strconv"

type Float64 float64

func (f Float64) ToString() (string, error) {
	return strconv.FormatFloat(float64(f), 'f', -1, 64), nil
}

func (f *Float64) FromString(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = Float64(v)
	return nil
}
