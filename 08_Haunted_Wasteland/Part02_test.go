package main

import (
	"08_Haunted_Wasteland/documents"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelMapNavigate(t *testing.T) {
	testMap := documents.MapFromFile("Puzzle08_Example2_Input.txt")
	result := ParallelMapNavigate(testMap)

	assert.Equal(t, 6, result)
}
