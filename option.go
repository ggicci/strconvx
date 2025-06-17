package strconvx

// Option adjusts the hybrid behaviour when creating a `StringCodec` instance with `New()` method.
type Option func(o *options)

// NoHybrid prevents `New()` from creating a hybrid `StringCodec` instance.
// Which means when calling `New()` with an object that doesn't implement
// `StringCodec` interface, `ErrUnsupportedType` will be returned.
//
// Example:
//
//	// if v implemented StringCodec, returns StringCodec
//	// if v didn't implement StringCodec, returns ErrUnsupportedType
//	New(v, NoHybrid())
func NoHybrid() Option {
	return func(o *options) {
		o.Opt(optionNoHybrid)
	}
}

// CompleteHybrid ensures that the hybrid instance created by `New()`
// has a full implementation of interface `StringCodec`.
func CompleteHybrid() Option {
	return func(o *options) {
		o.Opt(optionCompleteHybrid)
	}
}

type options struct {
	Value uint8
}

func defaultOptions() *options {
	return &options{}
}

func (o *options) Opt(v option) {
	o.Value |= uint8(v)
}

func (o *options) Has(v option) bool {
	return (o.Value & uint8(v)) > 0
}

type option int

const (
	optionNoHybrid option = 1 << iota
	optionCompleteHybrid
)
