package sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testSequenceOne = []int{0, 3, 6, 9, 12, 15}
)

func TestNewSequence(t *testing.T) {
	seq := NewSequence(testSequenceOne)

	assert.Equal(t, 1, len(*seq))
}

func TestExtrapolate(t *testing.T) {
	seq := NewSequence(testSequenceOne)

	Extrapolate(seq, 0)

	assert.Equal(t, 2, len(*seq))
	assert.Equal(t, false, AllEqual(seq, 1, 0))
}

func TestExtrapolateAll(t *testing.T) {
	seq := NewSequence(testSequenceOne)

	ExtrapolateAll(seq)

	assert.Equal(t, 3, len(*seq))
	assert.Equal(t, true, AllEqual(seq, 2, 0))
}

func TestNextInSequence(t *testing.T) {
	seq := NewSequence(testSequenceOne)
	ExtrapolateAll(seq)

	next := NextInSequence(seq)
	assert.Equal(t, 18, next)
}

func TestPrevInSequence(t *testing.T) {
	seq := NewSequence(testSequenceOne)
	ExtrapolateAll(seq)

	prev := PrevInSequence(seq)
	assert.Equal(t, -3, prev)

	seq = NewSequence([]int{1, 3, 6, 10, 15, 21})
	ExtrapolateAll(seq)
	prev = PrevInSequence(seq)
	assert.Equal(t, 0, prev)
}
