package sliceUtil

import (
	"sync"
)

func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{}
}

type SafeSlice[T any] struct {
	mu    sync.RWMutex
	slice []*T
}

func (s *SafeSlice[T]) Append(elems ...*T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, elems...)
}

func (s *SafeSlice[T]) RemoveAt(index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.slice) {
		return false
	}
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return true
}

func (s *SafeSlice[T]) Remove(fn func(t *T) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := len(s.slice) - 1; i >= 0; i-- {
		if fn(s.slice[i]) {
			s.slice = append(s.slice[:i], s.slice[i+1:]...)
		}
	}
}

func (s *SafeSlice[T]) Set(index int, value *T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.slice) {
		return false
	}
	s.slice[index] = value
	return true
}

func (s *SafeSlice[T]) Insert(index int, value *T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index > len(s.slice) {
		return false
	}
	s.slice = append(s.slice[:index], append([]*T{value}, s.slice[index:]...)...)
	return true
}

func (s *SafeSlice[T]) Init(values []*T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = values
}

func (s *SafeSlice[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = []*T{}
}

func (s *SafeSlice[T]) Contains(fn func(t *T) bool) bool {
	for _, v := range s.slice {
		if fn(v) {
			return true
		}
	}
	return false
}

func (s *SafeSlice[T]) IndexOf(fn func(t *T) bool) int {
	for i, v := range s.slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

func (s *SafeSlice[T]) Find(fn func(t *T) bool) (int, *T) {
	for i, v := range s.slice {
		if fn(v) {
			return i, v
		}
	}
	return -1, nil
}

func (s *SafeSlice[T]) Len() int {
	return len(s.slice)
}

func (s *SafeSlice[T]) Values() []*T {
	return s.slice
}

func (s *SafeSlice[T]) Get(index int) *T {
	if index < 0 || index >= len(s.slice) {
		return nil
	}
	return s.slice[index]
}

func (s *SafeSlice[T]) Copy() *SafeSlice[T] {
	sliceCopy := make([]*T, len(s.slice))
	copy(sliceCopy, s.slice)
	return &SafeSlice[T]{slice: sliceCopy}
}
