package gears

func Sum(vals ...int) (total int) {
	for _, v := range vals {
		total += v
	}

	return total
}
