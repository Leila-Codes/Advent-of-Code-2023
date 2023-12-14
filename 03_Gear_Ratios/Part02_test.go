package main

import (
	"testing"
)

func TestGetGears(t *testing.T) {
	ratios := GetGears("Puzzle03_Example_Input.txt")
	total := Sum(ratios...)

	if total != 467835 {
		t.Errorf("puzzle 03 answer is not correct. expected '%d' but got '%d'", 467835, total)
	}
}
