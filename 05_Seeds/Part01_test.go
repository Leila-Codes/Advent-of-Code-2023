package main

import (
	"05_Seeds/seeds"
	"testing"
)

func testLocation(t *testing.T, input seeds.PuzzleInput, seed, expectedLocation int) {
	if location := input.LocationForSeed(seed); location != expectedLocation {
		t.Errorf("seed mapping failed. expected '%d' but got '%d'", expectedLocation, location)
	}
}

func TestExamplePuzzle(t *testing.T) {
	var (
		mr      = seeds.ReadInput("Puzzle05_Example_Input.txt")
		expects = []int{82, 43, 86, 35}
	)

	for i, seed := range mr.Seeds {
		testLocation(t, mr, seed, expects[i])
	}
}

func TestMinimumLocation(t *testing.T) {
	var (
		mr          = seeds.ReadInput("Puzzle05_Example_Input.txt")
		minLocation = -1
	)

	for _, seed := range mr.Seeds {
		location := mr.LocationForSeed(seed)
		if minLocation == -1 {
			minLocation = location
			continue
		}

		minLocation = min(minLocation, location)
	}

	if minLocation != 35 {
		t.Errorf("min location not correct. expected '%d' but got '%d'", 35, minLocation)
	}
}
