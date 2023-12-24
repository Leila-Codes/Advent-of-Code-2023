package matrix

type Matrix[T comparable] [][]T

func (m Matrix[T]) Height() int {
	return len(m)
}

func (m Matrix[T]) Width(line int) int {
	return len(m[line])
}

func (m Matrix[T]) Get(x, y int) T {
	return m[y][x]
}
