package races

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromFile(t *testing.T) {
	races := FromFile("../Puzzle06_Example_Input.txt")

	assert.Equal(t, 3, len(races), "did not load correct # of races")
	assert.Equal(t, 7, races[0].Duration, "durations do not match")
	assert.Equal(t, 9, races[0].RecordDistance, "distances do not match")
}
