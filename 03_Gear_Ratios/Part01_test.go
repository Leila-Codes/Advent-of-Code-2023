package main

import (
	"testing"
)

func TestGetEngineParts(t *testing.T) {
	parts := GetEngineParts("Puzzle03_Example_Input.txt")

	total := sum(parts...)

	if total != 4361 {
		t.Errorf("example sum not correct - expected '%d' got '%d'", 4361, total)
		t.Fail()
	}
}
