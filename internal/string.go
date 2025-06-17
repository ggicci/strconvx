package internal

type String string

func (s String) ToString() (string, error) {
	return string(s), nil
}

func (sv *String) FromString(s string) error {
	*sv = String(s)
	return nil
}

func (s String) MarshalText() ([]byte, error) {
	return marshalTextViaToString(s)
}

func (s *String) UnmarshalText(text []byte) error {
	return unmarshalTextViaFromString(s, text)
}
