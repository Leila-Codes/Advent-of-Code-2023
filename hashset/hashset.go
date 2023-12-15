package hashset

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func NewSetWithValues[T comparable](vals []T) Set[T] {
	s := Set[T]{}
	s.AddAll(vals)
	return s
}

func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) AddAll(values []T) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s Set[T]) Remove(value T) {
	if s.Contains(value) {
		delete(s, value)
	}
}

func (s Set[T]) Length() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	var (
		values = make([]T, len(s))
		iter   = 0
	)

	for i := range s {
		values[iter] = i
		iter++
	}

	return values
}
