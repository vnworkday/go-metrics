package common

type Empty struct{}

// Default returns the zero value of the type T.
func Default[T any]() T {
	var zero T
	return zero
}
