package scratchcard

import "testing"

func TestCardFromLine(t *testing.T) {
	testInput := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := CardFromLine(testInput)

	if card.winners.Length() != 5 {
		t.Errorf("winning numbers length doesn't match; expected '%d' got '%d'", 5, card.winners.Length())
		t.Fail()
	}

	if card.numbers.Length() != 8 {
		t.Errorf("selected numbers length doesn't match; expected '%d' got '%d'", 8, card.numbers.Length())
		t.Fail()
	}
}

func TestCard_Score(t *testing.T) {
	card := CardFromLine("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	score := card.Score()
	expects := 8

	if score != expects {
		t.Errorf("Score doesn't match. expected '%d' but got '%d'", expects, score)
		t.Fail()
	}
}
