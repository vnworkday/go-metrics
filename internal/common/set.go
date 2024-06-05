package common

import "strings"

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

func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for value := range s.set {
		if !other.Contains(value) {
			return false
		}
	}

	return true
}

func StringToSet(str, sep string) Set[string] {
	m := make(map[string]struct{})
	parts := strings.Split(str, sep)
	for i := 0; i < len(parts); i++ {
		if parts[i] == "" {
			continue
		}

		m[parts[i]] = struct{}{}
	}
	return Set[string]{set: m}
}
