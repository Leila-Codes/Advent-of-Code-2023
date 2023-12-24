package sequence

type Sequence *[][]int

func NewSequence(nums []int) Sequence {
	seq := make([][]int, 1)
	seq[0] = nums

	return &seq
}

//func (s *Sequence) Last() []int {
//	return s[len(s)-1]
//}

func Extrapolate(s Sequence, line int) {
	length := len((*s)[line])

	nums := make([]int, length-1)

	for i := 0; i < length-1; i++ {
		var (
			curr = (*s)[line][i]
			next = (*s)[line][i+1]
		)

		nums[i] = next - curr
	}

	*s = append(*s, nums)
}

func ExtrapolateAll(s Sequence) {
	iter := 0
	for !AllEqual(s, len(*s)-1, 0) {
		Extrapolate(s, iter)
		iter++
	}
}

func NextInSequence(s Sequence) int {
	var (
		currLine  = (*s)[0]
		currValue = currLine[len(currLine)-1]
	)

	// Bottom-up processing
	var delta int
	for i := len(*s) - 2; i > 0; i-- {
		line := (*s)[i]
		delta = delta + line[len(line)-1]
	}

	return currValue + delta
}

func PrevInSequence(s Sequence) int {
	var (
		currValue = (*s)[0][0]
	)

	var delta int
	for i := len(*s) - 2; i > 0; i-- {
		line := (*s)[i]
		delta = line[0] - delta
	}

	return currValue - delta
}

func AllEqual(s Sequence, line, value int) bool {
	for _, v := range (*s)[line] {
		if value != v {
			return false
		}
	}

	return true
}
