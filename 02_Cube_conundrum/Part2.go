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

	//limits := map[game.CubeColour]uint{
	//	game.ColourRed:   12,
	//	game.ColourGreen: 13,
	//	game.ColourBlue:  14,
	//}

	var sum uint

	for scanner.Scan() {
		summary := game.ParseGameSummary(scanner.Text())

		var power uint = 0

		for _, value := range summary.CubeCount {
			if power == 0 {
				//fmt.Printf("power was 0, setting to %d\n", value)
				power = value
			} else {
				if value > 0 {
					power = power * value
				}
			}

			//fmt.Printf("Min. %d %s \t", value, colour)
		}

		//fmt.Println()
		fmt.Printf("Game %d - Power: %d\n", summary.ID, power)
		//break
		//fmt.Println()

		sum += power
		//if summary.Possible(limits) {
		//	sum += summary.ID
		//}
	}

	fmt.Println("Answer is: ", sum)
}
