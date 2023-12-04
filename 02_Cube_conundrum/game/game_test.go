package game

import "testing"

func TestParseGameSummary(t *testing.T) {
	input := "Game 1: 4 green, 7 blue; 2 blue, 4 red; 5 blue, 2 green, 2 red; 1 green, 3 red, 9 blue; 3 green, 9 blue; 7 green, 2 blue, 2 red"
	expects := GameSummary{ID: 1, CubeCount: map[CubeColour]uint{ColourGreen: 7, ColourRed: 4, ColourBlue: 9}}

	actual := ParseGameSummary(input)

	if expects.ID != actual.ID {
		t.Errorf("Game ID's don't match - expected '%d' got '%d'\n", expects.ID, actual.ID)
		t.Fail()
	}

	if expects.CubeCount[ColourGreen] != actual.CubeCount[ColourGreen] {
		t.Errorf("Colour GREEN's don't match - expected '%d' got '%d'\n", expects.CubeCount[ColourGreen], actual.CubeCount[ColourGreen])
		t.Fail()
	}

	if expects.CubeCount[ColourRed] != actual.CubeCount[ColourRed] {
		t.Errorf("Colour RED's don't match - expected '%d' got '%d'\n", expects.CubeCount[ColourRed], actual.CubeCount[ColourRed])
		t.Fail()
	}

	if expects.CubeCount[ColourBlue] != actual.CubeCount[ColourBlue] {
		t.Errorf("Colour BLUE's don't match - expected '%d' got '%d'\n", expects.CubeCount[ColourBlue], actual.CubeCount[ColourBlue])
		t.Fail()
	}
}
