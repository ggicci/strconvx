package strconvx

import (
	"fmt"
	"reflect"
)

type StringConverterAdaptor[T any] func(*T) (StringConverter, error)
type AnyStringConverterAdaptor func(any) (StringConverter, error)

func ToAnyStringConverterAdaptor[T any](adapt StringConverterAdaptor[T]) (reflect.Type, AnyStringConverterAdaptor) {
	return typeOf[T](), func(v any) (StringConverter, error) {
		if cv, ok := v.(*T); ok {
			return adapt(cv)
		} else {
			return nil, fmt.Errorf("%w: cannot convert %T to %s", ErrTypeMismatch, v, typeOf[*T]())
		}
	}
}

var builtinAdaptors = make(map[reflect.Type]AnyStringConverterAdaptor)

func builtinStringConverter[T any](adaptor StringConverterAdaptor[T]) {
	typ, anyAdaptor := ToAnyStringConverterAdaptor[T](adaptor)
	builtinAdaptors[typ] = anyAdaptor
}
