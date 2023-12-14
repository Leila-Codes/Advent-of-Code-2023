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

	for i, card := range cards {
		winners := card.NumberOfWinners()

		//fmt.Printf("Card %d - Has %d matching numbers.\n", i, winners)

		for r := 0; r < card.Replicas; r++ {
			for j := 1; j <= winners; /* && (i+j) < len(cards)*/ j++ {
				//fmt.Printf("Copying Card %d\n", i+j)
				cards[i+j].Copy()
			}
		}
	}

	total := 0

	for i, card := range cards {
		fmt.Printf("Card %d\tCopies: %d\n", i+1, card.Replicas)
		total += card.Replicas
	}

	fmt.Println("Answer is: ", total)
}
