package slicetools

import "io"

type CloseablePromise[T interface{}] struct {
	Value T
	Error error
}

type Promise[T interface{}] chan CloseablePromise[T]

func (p Promise[T]) pop() (T, error) {
	d := <-p
	return d.Value, d.Error
}

func (p Promise[T]) ForEach(task func(in T)) {
	for {
		val, err := p.pop()
		if err != nil && err != io.EOF {
			panic(err)
		}

		task(val)
	}
}

func (p Promise[T]) Map(mapper func(in T) interface{}) {

}
