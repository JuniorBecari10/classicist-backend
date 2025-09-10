package option

import "encoding/json"

type Option[T any] struct {
	val T
	present bool
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		val: val,
		present: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		present: false,
	}
}

func FromSQL[T any](val T, valid bool) Option[T] {
	if valid {
		return Some(val)
	} else {
		return None[T]()
	}
}

// ---

// Use this with caution, as the value may be invalid.
// Not an unintialized value that leads to UB, but still not safe.
// The value may be the default value of T, or a past one if the Option was reused.
// Prefer TryUnwrap or UnwrapOr.
// This won't panic if you use it on a None Option.
func (o Option[T]) Unwrap() T {
	return o.val
}

func (o Option[T]) UnwrapOr(other T) T {
	if o.present {
		return o.val
	} else {
		return other
	}
}

// Use this in if statements.
// Example:
//
// if val, ok := option.TryUnwrap(); ok { ... }
func (o Option[T]) TryUnwrap() (T, bool) {
	return o.val, o.present
}

func (o Option[T]) IsPresent() bool {
	return o.present
}

func (o Option[T]) IsNotPresent() bool {
	return !o.present
}

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.present {
		return json.Marshal(o.val)
	} else {
		return []byte("null"), nil
	}
}
