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

// Namespace is the place to register type adaptors (of AnyAdaptor).
type Namespace struct {
	adaptors map[reflect.Type]AnyAdaptor
}

// NewNamespace creates a namespace where you can register adaptors to
// override/adapt the converting behaviours of existing types.
func NewNamespace() *Namespace {
	return &Namespace{
		adaptors: make(map[reflect.Type]AnyAdaptor),
	}
}

// New creates a StringCodec instance from the given value. If the given value itself
// is already a StringCodec, it will return directly. Otherwise, it will try to create
// a StringCodec instance by trying the following approaches:
//  1. check if there's a custom adaptor for the type of the given value,
//     if so, use it to adapt the given value to a StringCodec.
//  2. same as above, but check the builtin adaptors, which support the builtin types,
//     e.g. int, string, float64, etc.
//  3. try to create a "hybrid" instance, which makes use of the methods FromString,
//     ToString, MarshalText and UnmarshalText to fullfill the StringCodec interface.
//
// It has three options:
//
//	New(v)
//
// 1. with only default options, it will try all the 3 ways as listed above to
// create a StringCodec.
//
//	New(v, NoHybrid())
//
// 2. without hybrid, i.e. won't try the 3rd method, returns an
// ErrUnsupportedType error.
//
//	New(v, CompleteHybrid())
//
// 3. the hybrid must be a "complete" hybrid, which means it has to implement
// both FromString and ToString method that the StringCodec interface requires,
// while a "partial"/"incomplete" hybrid, one of theses two methods can be
// absent, and the absent one always returns an error, either
// ErrNotStringMarshaler or ErrNotStringUnmarshaler.
func (c *Namespace) New(v any, opts ...Option) (StringCodec, error) {
	if vs, ok := v.(StringCodec); ok {
		return vs, nil
	}

	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}
	return c.createStringCodec(v, options)
}

func (c *Namespace) createStringCodec(v any, opts *options) (StringCodec, error) {
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

	// Try to create a hybrid StringCodec from the reflect.Value.
	if !opts.Has(optionNoHybrid) {
		h := createHybridStringCodec(rv)
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
//  2. Call ToAnyAdaptor to create an adaptor of a specific type.
//
// Example:
//
//	ns := strconvx.NewNamespace()
//	typ, adaptor := strconvx.ToAnyAdaptor[bool](func(b *bool) (strconvx.StringCodec, error) {
//		// todo
//	})
//	ns.Adapt(typ, adaptor)
func (c *Namespace) Adapt(typ reflect.Type, adaptor AnyAdaptor) {
	c.adaptors[typ] = adaptor
}

// UndoAdapt removes the custom adaptor for the given type. It's a reverse operation of Adapt.
func (c *Namespace) UndoAdapt(typ reflect.Type) {
	delete(c.adaptors, typ)
}

func unsupportedType(rt reflect.Type) error {
	return fmt.Errorf("%w: %v", ErrUnsupportedType, rt)
}

func init() {
	builtinAdaptor(func(v *string) (StringCodec, error) { return (*internal.String)(v), nil })
	builtinAdaptor(func(v *bool) (StringCodec, error) { return (*internal.Bool)(v), nil })
	builtinAdaptor(func(v *int) (StringCodec, error) { return (*internal.Int)(v), nil })
	builtinAdaptor(func(v *int8) (StringCodec, error) { return (*internal.Int8)(v), nil })
	builtinAdaptor(func(v *int16) (StringCodec, error) { return (*internal.Int16)(v), nil })
	builtinAdaptor(func(v *int32) (StringCodec, error) { return (*internal.Int32)(v), nil })
	builtinAdaptor(func(v *int64) (StringCodec, error) { return (*internal.Int64)(v), nil })
	builtinAdaptor(func(v *uint) (StringCodec, error) { return (*internal.Uint)(v), nil })
	builtinAdaptor(func(v *uint8) (StringCodec, error) { return (*internal.Uint8)(v), nil })
	builtinAdaptor(func(v *uint16) (StringCodec, error) { return (*internal.Uint16)(v), nil })
	builtinAdaptor(func(v *uint32) (StringCodec, error) { return (*internal.Uint32)(v), nil })
	builtinAdaptor(func(v *uint64) (StringCodec, error) { return (*internal.Uint64)(v), nil })
	builtinAdaptor(func(v *float32) (StringCodec, error) { return (*internal.Float32)(v), nil })
	builtinAdaptor(func(v *float64) (StringCodec, error) { return (*internal.Float64)(v), nil })
	builtinAdaptor(func(v *complex64) (StringCodec, error) { return (*internal.Complex64)(v), nil })
	builtinAdaptor(func(v *complex128) (StringCodec, error) { return (*internal.Complex128)(v), nil })
	builtinAdaptor(func(v *time.Time) (StringCodec, error) { return (*internal.Time)(v), nil })
	builtinAdaptor(func(b *[]byte) (StringCodec, error) { return (*internal.ByteSlice)(b), nil })
}
