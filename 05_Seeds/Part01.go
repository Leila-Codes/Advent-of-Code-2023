package main

import (
	"05_Seeds/seeds"
	"fmt"
)

func main() {
	input := seeds.ReadInput("Puzzle05_Input.txt")

	var (
		minLocation = -1
	)

	for _, seed := range input.Seeds {
		out := input.LocationForSeed(seed)
		if minLocation == -1 {
			minLocation = out
			continue
		}

		minLocation = min(minLocation, out)
	}

	fmt.Println("Answer is: ", minLocation)
}
