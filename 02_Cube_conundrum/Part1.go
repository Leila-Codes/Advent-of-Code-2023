package main

import (
	"02_Cube_conundrum/game"
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("Puzzle02_Input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	limits := map[game.CubeColour]uint{
		game.ColourRed:   12,
		game.ColourGreen: 13,
		game.ColourBlue:  14,
	}

	var sum uint

	for scanner.Scan() {
		summary := game.ParseGameSummary(scanner.Text())

		if summary.Possible(limits) {
			sum += summary.ID
		}
	}

	fmt.Println("Answer is: ", sum)
}
