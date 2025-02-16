package strconvx

import "errors"

var (
	ErrUnsupportedType      = errors.New("unsupported type")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrNotStringMarshaler   = errors.New("not a CanToString")
	ErrNotStringUnmarshaler = errors.New("not a CanFromString")
	ErrNotPointer           = errors.New("not a pointer")
	ErrNilPointer           = errors.New("nil pointer")
)
