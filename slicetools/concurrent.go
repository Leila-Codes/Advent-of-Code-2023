package slicetools

import "sync"

type ConcurrentSlice[T interface{}] struct {
	mutex  sync.Mutex
	values []T
}

func NewConcurrentSlice[T interface{}]() *ConcurrentSlice[T] {
	return &ConcurrentSlice[T]{
		mutex:  sync.Mutex{},
		values: make([]T, 0),
	}
}

func (c *ConcurrentSlice[T]) Append(value T) {
	c.mutex.Lock()
	c.values = append(c.values, value)
	c.mutex.Unlock()
}

func (c *ConcurrentSlice[T]) Values() []T {
	return c.values
}
