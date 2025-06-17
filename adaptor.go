package strconvx

import (
	"fmt"
	"reflect"
)

type Adaptor[T any] func(*T) (StringCodec, error)
type AnyAdaptor func(any) (StringCodec, error)

func ToAnyAdaptor[T any](adaptor Adaptor[T]) (reflect.Type, AnyAdaptor) {
	return typeOf[T](), func(v any) (StringCodec, error) {
		if cv, ok := v.(*T); ok {
			return adaptor(cv)
		} else {
			return nil, fmt.Errorf("%w: cannot convert %T to %s", ErrTypeMismatch, v, typeOf[*T]())
		}
	}
}

var builtinAdaptors = make(map[reflect.Type]AnyAdaptor)

func builtinAdaptor[T any](adaptor Adaptor[T]) {
	typ, anyAdaptor := ToAnyAdaptor(adaptor)
	builtinAdaptors[typ] = anyAdaptor
}
