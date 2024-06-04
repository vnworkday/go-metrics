package common

type Set[T comparable] struct {
	set map[T]struct{}
}

func NewSet[T comparable](items ...T) Set[T] {
	s := Set[T]{set: make(map[T]struct{})}
	s.Add(items...)
	return s
}

func (s Set[T]) Add(values ...T) Set[T] {
	for _, value := range values {
		s.set[value] = struct{}{}
	}

	return s
}

func (s Set[T]) Remove(values ...T) Set[T] {
	for _, value := range values {
		delete(s.set, value)
	}

	return s
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s.set[value]
	return ok
}

func (s Set[T]) Len() int {
	return len(s.set)
}

func (s Set[T]) IsEmpty() bool {
	return s.set == nil || s.Len() == 0
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, s.Len())
	for value := range s.set {
		values = append(values, value)
	}

	return values
}
