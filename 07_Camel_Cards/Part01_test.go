package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRankGames(t *testing.T) {
	score := RankGames("Puzzle07_Example_Input.txt")

	assert.Equal(t, 6440, score, "score doesn't match")
}
