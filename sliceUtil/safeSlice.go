package sliceUtil

import (
	"sync"
)

type SafeSlice[T any] struct {
	mu    sync.RWMutex
	slice []*T
}

func (s *SafeSlice[T]) Append(t *T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.slice = append(s.slice, t)
}

func (s *SafeSlice[T]) Remove(index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.slice) {
		return false
	}
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return true
}

func (s *SafeSlice[T]) Get(index int) (*T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if index < 0 || index >= len(s.slice) {
		return nil, false
	}
	return s.slice[index], true
}

func (s *SafeSlice[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.slice)
}

// Range 遍历 slice 中的所有元素
func (s *SafeSlice[T]) Range(f func(index int, value *T)) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for i, v := range s.slice {
		f(i, v)
	}
}

// Snapshot 返回 slice 的拷贝
func (s *SafeSlice[T]) Snapshot() []*T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	snapshot := make([]*T, len(s.slice))
	copy(snapshot, s.slice)
	return snapshot
}
