// strconvx is a tiny package that helps converting values from/to a string.
package strconvx

// StringConverter defines a type to be able to convert from/to a string.
type StringConverter interface {
	CanToString
	CanFromString
}

// CanToString defines a type to be able to convert to a string.
type CanToString interface {
	ToString() (string, error)
}

// CanFromString defines a type to be able to convert from a string.
type CanFromString interface {
	FromString(string) error
}

// New creates a StringConverter instance from the given value. Note that
// this method is a wrapper around the default namespace's New method.
// Which means it doesn't support override/adapt existing types. Please
// read Namespace.New to learn more.
func New(v any) (StringConverter, error) {
	return defaultNS.New(v)
}
