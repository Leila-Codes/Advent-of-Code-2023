package gears

import "testing"

func TestSchematicFromFile(t *testing.T) {
	grid, err := SchematicFromFile("../Puzzle03_Example_Input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(grid) != 10 {
		t.Errorf("schematic isn't correct length. expected %d got %d", 10, len(grid))
	}

	for _, row := range grid {
		if len(row) != 10 {
			t.Errorf("schematic isn't correct width. expected %d got %d", 10, len(row))
		}
	}
}
