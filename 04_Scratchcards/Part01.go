package main

import (
	"04_Scratchcards/scratchcard"
	"fmt"
)

func main() {
	cards, err := scratchcard.CardsFromFile("Puzzle04_Input.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, card := range cards {
		total += card.Score()
	}

	fmt.Println("Answer is: ", total)
}
