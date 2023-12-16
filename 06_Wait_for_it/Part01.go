package main

import (
	"06_Wait_for_it/races"
	"fmt"
)

func GetWinnerNumber(filePath string) int {
	boatRaces := races.FromFile(filePath)

	var score = -1

	for _, race := range boatRaces {
		winners := race.WinnerHoldsMS()
		if score == -1 {
			score = winners
			continue
		}

		score *= winners
	}

	return score
}

func main() {
	fmt.Println("Answer is: ", GetWinnerNumber("Puzzle06_Input.txt"))
}
