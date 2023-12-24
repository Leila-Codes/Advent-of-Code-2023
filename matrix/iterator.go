package matrix

type Iterator[T interface{}] interface {
	Next() T
	HasNext() bool
}

type matrixIterator[T interface{}] struct {
	m                Matrix[T]
	currCol, currRow int
}

type columnIterator[T interface{}] matrixIterator[T]

func ColumnIterator[T interface{}](m Matrix[T], column int) *columnIterator[T] {
	return &columnIterator[T]{
		m:       m,
		currCol: column,
		currRow: 0,
	}
}

func RowIterator[T interface{}](m Matrix[T], row int) *rowIterator[T] {
	return &rowIterator[T]{
		m:       m,
		currRow: row,
		currCol: 0,
	}
}

type rowIterator[T interface{}] matrixIterator[T]

func (c *columnIterator[T]) HasNext() bool {
	return c.currRow+1 < c.m.Height()
}

func (c *columnIterator[T]) Next() T {
	c.currRow++
	return c.m[c.currRow][c.currCol]
}

func (r *rowIterator[T]) HasNext() bool {
	return r.currCol+1 < r.m.Width(r.currCol)
}

func (r *rowIterator[T]) Next() T {
	r.currCol++
	return r.m[r.currRow][r.currCol]
}
