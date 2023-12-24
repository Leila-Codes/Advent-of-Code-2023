package documents

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapFromFile(t *testing.T) {
	exampleMap := MapFromFile("../Puzzle08_Example_Input.txt")

	assert.Equal(t, 7, len(exampleMap.Nodes), "map is not correct size")
	assert.Equal(t, "AAA", exampleMap.Nodes["AAA"].Value, "first item is incorrect")
	assert.NotNil(t, exampleMap.Nodes["AAA"].Left, "node links for AAA are invalid")
	assert.NotNil(t, exampleMap.Nodes["AAA"].Right, "node links for AAA are invalid")
	assert.NotNil(t, exampleMap.Nodes["ZZZ"].Left, "node links for ZZZ are invalid")
	assert.NotNil(t, exampleMap.Nodes["ZZZ"].Right, "node links for ZZZ are invalid")
}
