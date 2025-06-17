package internal

type StringMarshaler interface {
	ToString() (string, error)
}

type StringUnmarshaler interface {
	FromString(string) error
}

func marshalTextViaToString(v StringMarshaler) ([]byte, error) {
	s, err := v.ToString()
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func unmarshalTextViaFromString(v StringUnmarshaler, text []byte) error {
	return v.FromString(string(text))
}
