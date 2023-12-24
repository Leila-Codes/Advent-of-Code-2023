package slicetools

type List[T interface{}] []T

func ForEach[T interface{}](list List[T], task func(in T)) {
	for _, item := range list {
		task(item)
	}
}

func Map[IN, OUT interface{}](list List[IN], mapper func(in IN) OUT) []OUT {
	output := make([]OUT, len(list))
	for i, item := range list {
		output[i] = mapper(item)
	}
	return output
}
