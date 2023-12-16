package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWinnerNumber(t *testing.T) {
	score := GetWinnerNumber("Puzzle06_Example_Input.txt")

	assert.Equal(t, 288, score, "score did not match")
}
