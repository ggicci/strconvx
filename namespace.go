package strconvx

import (
	"fmt"
	"reflect"
	"time"

	"github.com/ggicci/strconvx/internal"
)

// defaultNS is the default Namespace instance held by this package, where
// we registered the adaptors for the builtin types, e.g. bool, int, string, etc.
var defaultNS = NewNamespace()

// Namespace is the place to register type adaptors (of AnyStringConverterAdaptor).
type Namespace struct {
	adaptors map[reflect.Type]AnyStringConverterAdaptor
}

// NewNamespace creates a namespace where you can register adaptors to
// override/adapt the converting behaviours of existing types.
func NewNamespace() *Namespace {
	return &Namespace{
		adaptors: make(map[reflect.Type]AnyStringConverterAdaptor),
	}
}

// New creates a StringConverter instance from the given value. If the given value itself
// is already a StringConverter, it will return directly. Otherwise, it will try to create
// a StringConverter instance by trying the following approaches:
//  1. check if there's a custom adaptor for the type of the given value,
//     if so, use it to adapt the given value to a StringConverter.
//  2. same as above, but check the builtin adaptors, which support the builtin types,
//     e.g. int, string, float64, etc.
//  3. try to create a "hybrid" instance, which makes use of the methods FromString,
//     ToString, MarshalText and UnmarshalText to fullfill the StringConverter interface.
//
// It has three options:
//
//	New(v)
//
// 1. with only default options, it will try all the 3 ways as listed above to
// create a StringConverter.
//
//	New(v, NoHybrid())
//
// 2. without hybrid, i.e. won't try the 3rd method, returns an
// ErrUnsupportedType error.
//
//	New(v, CompleteHybrid())
//
// 3. the hybrid must be a "complete" hybrid, which means it has to implement
// both FromString and ToString method that the StringConverter interface requires,
// while a "partial"/"incomplete" hybrid, one of theses two methods can be
// absent, and the absent one always returns an error, either
// ErrNotStringMarshaler or ErrNotStringUnmarshaler.
func (c *Namespace) New(v any, opts ...Option) (StringConverter, error) {
	if vs, ok := v.(StringConverter); ok {
		return vs, nil
	}

	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}
	return c.createStringConverter(v, options)
}

func (c *Namespace) createStringConverter(v any, opts *options) (StringConverter, error) {
	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}
	if rv.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("%w: value must be a non-nil pointer", ErrNotPointer)
	}
	if rv.IsNil() {
		return nil, fmt.Errorf("%w: value must be a non-nil pointer", ErrNilPointer)
	}

	baseType := rv.Type().Elem()

	// Check if there is a custom adaptor for the base type.
	if adapt, ok := c.adaptors[baseType]; ok {
		return adapt(rv.Interface())
	}

	// Check if there is a built-in adaptor for the base type.
	if adapt, ok := builtinAdaptors[baseType]; ok {
		return adapt(rv.Interface())
	}

	// Try to create a hybrid StringConverter from the reflect.Value.
	if !opts.Has(optionNoHybrid) {
		h := createHybridStringConverter(rv)
		if h != nil {
			if opts.Has(optionCompleteHybrid) {
				if err := h.(*hybrid).validateAsComplete(); err != nil {
					return nil, err
				}
			}
			return h, nil
		}
	}

	return nil, unsupportedType(baseType)
}

// Adapt registers a custom adaptor for the given type.
//
//  1. You must create a Namespace instance and register the adaptor there.
//  2. Call ToAnyStringConverterAdaptor to create an adaptor of a specific type.
//
// Example:
//
//	ns := strconvx.NewNamespace()
//	typ, adaptor := strconvx.ToAnyStringConverterAdaptor[bool](func(b *bool) (strconvx.StringConverter, error) {
//		// todo
//	})
//	ns.Adapt(typ, adaptor)
func (c *Namespace) Adapt(typ reflect.Type, adaptor AnyStringConverterAdaptor) {
	c.adaptors[typ] = adaptor
}

func unsupportedType(rt reflect.Type) error {
	return fmt.Errorf("%w: %v", ErrUnsupportedType, rt)
}

func init() {
	builtinStringConverter[string](func(v *string) (StringConverter, error) { return (*internal.String)(v), nil })
	builtinStringConverter[bool](func(v *bool) (StringConverter, error) { return (*internal.Bool)(v), nil })
	builtinStringConverter[int](func(v *int) (StringConverter, error) { return (*internal.Int)(v), nil })
	builtinStringConverter[int8](func(v *int8) (StringConverter, error) { return (*internal.Int8)(v), nil })
	builtinStringConverter[int16](func(v *int16) (StringConverter, error) { return (*internal.Int16)(v), nil })
	builtinStringConverter[int32](func(v *int32) (StringConverter, error) { return (*internal.Int32)(v), nil })
	builtinStringConverter[int64](func(v *int64) (StringConverter, error) { return (*internal.Int64)(v), nil })
	builtinStringConverter[uint](func(v *uint) (StringConverter, error) { return (*internal.Uint)(v), nil })
	builtinStringConverter[uint8](func(v *uint8) (StringConverter, error) { return (*internal.Uint8)(v), nil })
	builtinStringConverter[uint16](func(v *uint16) (StringConverter, error) { return (*internal.Uint16)(v), nil })
	builtinStringConverter[uint32](func(v *uint32) (StringConverter, error) { return (*internal.Uint32)(v), nil })
	builtinStringConverter[uint64](func(v *uint64) (StringConverter, error) { return (*internal.Uint64)(v), nil })
	builtinStringConverter[float32](func(v *float32) (StringConverter, error) { return (*internal.Float32)(v), nil })
	builtinStringConverter[float64](func(v *float64) (StringConverter, error) { return (*internal.Float64)(v), nil })
	builtinStringConverter[complex64](func(v *complex64) (StringConverter, error) { return (*internal.Complex64)(v), nil })
	builtinStringConverter[complex128](func(v *complex128) (StringConverter, error) { return (*internal.Complex128)(v), nil })
	builtinStringConverter[time.Time](func(v *time.Time) (StringConverter, error) { return (*internal.Time)(v), nil })
	builtinStringConverter[[]byte](func(b *[]byte) (StringConverter, error) { return (*internal.ByteSlice)(b), nil })
}
