// strconvx is a tiny package that helps converting values from/to text.
package strconvx

// StringMarshaler is implemented by types that can represent themselves as a string.
// It is similar in spirit to encoding.TextMarshaler, but returns a string directly.
type StringMarshaler interface {

	// ToString returns the string representation of the value.
	ToString() (string, error)
}

// StringUnmarshaler is implemented by types that can parse themselves from a string.
// It is similar in spirit to encoding.TextUnmarshaler, but accepts a string directly.
type StringUnmarshaler interface {
	// FromString parses the value from its string representation.
	FromString(string) error
}

// StringCodec is implemented by types that support bidirectional conversion
// between their value and a string representation.
//
// It combines both StringMarshaler and StringUnmarshaler interfaces.
type StringCodec interface {
	StringMarshaler
	StringUnmarshaler
}

// New creates a StringCodec instance for the given value.
// This is a shortcut to the default namespace's New method.
// Note: since this uses the default namespace, it does not support
// overriding or customizing existing type bindings.
// For more advanced usage, see Namespace.New.
func New(v any) (StringCodec, error) {
	return defaultNS.New(v)
}
