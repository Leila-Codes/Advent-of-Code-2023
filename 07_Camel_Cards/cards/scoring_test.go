package cards

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiveOfAKind(t *testing.T) {
	hand := Hand{Ace, Ace, Ace, Ace, Ace}

	assert.Equal(t, 7, hand.Score(), "score doesn't match")
}

func TestFourOfAKind(t *testing.T) {
	hand := Hand{Ace, Ace, Ace, Ace, King}
	assert.Equal(t, 6, hand.Score(), "score doesn't match")
}

func TestFullHouse(t *testing.T) {
	hand := Hand{Ace, Ace, Ace, King, King}

	assert.Equal(t, 5, hand.Score(), "score doesn't match")
}

func TestThreeOfAKind(t *testing.T) {
	hand := Hand{Jack, Four, Six, Six, Six} // Devil's number

	assert.Equal(t, 4, hand.Score(), "score doesn't match")
}

func TestTwoPair(t *testing.T) {
	hand := Hand{Ace, Ace, King, King, Jack}

	assert.Equal(t, 3, hand.Score(), "score doesn't match")
}

func TestOnePair(t *testing.T) {
	hand := Hand{Queen, Five, Queen, Three, Ace}

	assert.Equal(t, 2, hand.Score(), "score doesn't match")
}

func TestHighCard(t *testing.T) {
	hand := Hand{Two, Three, Four, Five, Six}

	assert.Equal(t, 1, hand.Score(), "score doesn't match")
}
