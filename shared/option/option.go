package option

type Option[T any] struct {
	val T
	ok bool
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		val: val,
		ok: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		ok: false,
	}
}

// ---

// Use this with caution, as the value may be invalid.
func (o Option[T]) Unwrap() T {
	return o.val
}

func (o Option[T]) TryUnwrap() (T, bool) {
	return o.val, o.ok
}

func (o Option[T]) IsSome() bool {
	return o.ok
}

func (o Option[T]) IsNone() bool {
	return !o.ok
}
