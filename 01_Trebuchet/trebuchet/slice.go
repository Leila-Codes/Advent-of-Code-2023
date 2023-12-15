package trebuchet

// FirstAndLast - returns the first and last item from the given slice.
func FirstAndLast(input []int) (int, int) {
	return input[0], input[len(input)-1]
}
