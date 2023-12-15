package seeds

import (
	"testing"
)

func TestReadInput(t *testing.T) {
	input := ReadInput("../Puzzle05_Example_Input.txt")

	if len(input.Seeds) != 4 {
		t.Errorf("did not read enough seeds. expected '%d' got '%d'", 4, len(input.Seeds))
		t.Fail()
	}

	if len(input.Mappings) != 7 {
		t.Errorf("did not read enough mappings. expected '%d' got '%d'", 7, len(input.Mappings))
		t.Fail()
	}

	if input.Mappings[0].From != "seed" || input.Mappings[0].To != "soil" {
		t.Logf("mapping #1 from '%s' to '%s'", input.Mappings[0].From, input.Mappings[0].To)
		t.Errorf("mapping #1 source/dest incorrect. expected '%s' got '%s'", MappingCollection{From: "water", To: "soil"}, input.Mappings[0])
	}

	if len(input.Mappings[0].Ranges) != 2 {
		t.Errorf("mapping #1 did not read enough ranges. expected '%d' but got '%d'", 2, len(input.Mappings[0].Ranges))
		t.Fail()
	}

	col := input.Mappings[0]

	if v := col.Map(79); v != 81 {
		t.Errorf("mapping (%s) failed. expected '%d' but got '%d'", col, 81, v)
		t.Fail()
	}

	input.Mappings[0].Map(79)
}

func TestPuzzleInput_LocationForSeed(t *testing.T) {
	almanac := ReadInput("../Puzzle05_Example_Input.txt")

	if location := almanac.LocationForSeed(79); location != 82 {
		t.Errorf("seed location is not correct. expected '%d' but got '%d'", 82, location)
	}
}

func TestReadInput2(t *testing.T) {
	input := ReadInput("../Puzzle05_Input.txt")

	if l := len(input.Seeds); l != 20 {
		t.Errorf("did not read input seeds correctly. expected '%d' but got '%d'", 20, l)
		t.Fail()
	}

	if l := len(input.Mappings); l != 7 {
		t.Errorf("did not read mappings correctly. expected '%d' but got '%d'", 7, l)
		t.Fail()
	}
}

func testSeed(t *testing.T, seed int, expects []int) {
	var input = ReadInput("../Puzzle05_Example_Input.txt")

	t.Logf("test with seed %d", seed)

	for i, mapping := range input.Mappings {
		out := mapping.Map(seed)
		t.Logf("convert (%s) -> [%d => %d]", mapping, seed, out)

		if out != expects[i] {
			t.Errorf("conversion failed. expected '%d' but got '%d'", expects[i], out)
			t.FailNow()
		}
		seed = out
	}
}

func TestSeed79(t *testing.T) {
	var (
		seed    = 14
		expects = []int{14, 53, 49, 42, 42, 43, 43}
	)

	testSeed(t, seed, expects)
}

func TestSeed14(t *testing.T) {
	var (
		seed    = 14
		expects = []int{14, 53, 49, 42, 42, 43, 43}
	)

	testSeed(t, seed, expects)
}

func TestSeed55(t *testing.T) {
	var (
		seed    = 55
		expects = []int{57, 57, 53, 46, 82, 82, 86}
	)

	testSeed(t, seed, expects)
}

func TestSeed13(t *testing.T) {
	var (
		seed    = 13
		expects = []int{13, 52, 41, 34, 34, 35, 35}
	)

	testSeed(t, seed, expects)
}
