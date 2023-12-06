package gears

import (
	"fmt"
	"testing"
)

func sliceEquals[T comparable](actual, expects []T) error {
	if len(actual) != len(expects) {
		return fmt.Errorf("slices not equal - length's don't match. got '%d' expects '%d'", len(actual), len(expects))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expects[i] {
			return fmt.Errorf("slices don't match at index [%d] - got '%v' expected '%v'", i, actual[i], expects[i])
		}
	}

	return nil
}

func TestParseLine(t *testing.T) {

	testInputs := []string{
		".........823.835........%.........710.....749........134..%............................#812...&.....925.../..........276.......386..........",
		"............832*105..@........$..................*.........797.....535..932.........*....152...........123.........678.540...........-...6..",
		"743........634...*.................................................679...........*176......979....................928.......$..848..........",
		"......889........=.../.................*.............916.............201............-...............426&....627..............284............",
		"970.....+.-...701....#264....%....$.......113.......916........*..............620............809........$788....@.....446........904..*.....",
		".....785.......542.*............732...=........583...+....................424....$....702.367..993.......*....386..#.............*..........",
		".....................957.....542...731..941.......*....59..........571*554....214....108.@...............104*............=.............*....",
		"....771..500........868.......213......902....456........../.........255.....377.781......=348................133..@...367*696.....*........",
	}

	expects := [][]int{
		{812},
		{832, 105},
		{176},
		{426},
		{264, 788},
		{},
		{571, 554, 104},
		{348, 367, 696},
	}

	testParseLineAll(t, testInputs, expects)
}

func testParseLineAll(t *testing.T, testInputs []string, expects [][]int) {
	for i, input := range testInputs {
		err := testParseLine(input, expects[i])
		if err != nil {
			t.Errorf("test [%d] failed - %v", i, err)
			t.FailNow()
		}
	}
}

func testParseLine(testInput string, expects []int) error {
	actual := ParseLine(testInput)

	if err := sliceEquals(actual, expects); err != nil {
		return err
	}

	return nil
}
