package main

import (
	"08_Haunted_Wasteland/documents"
	"fmt"
)

func NavigateMap(wm *documents.WastelandMap) (numOfSteps int) {
	var (
		start       = "AAA"
		destination = "ZZZ"
		current     = wm.Nodes[start]
	)

	for current.Value != destination {
		var nextStep = wm.Instructions[numOfSteps%len(wm.Instructions)]
		switch nextStep {
		case 'L':
			current = current.Left
		case 'R':
			current = current.Right
		}
		numOfSteps++
	}

	return numOfSteps
}

func main() {
	puzzleMap := documents.MapFromFile("Puzzle08_Input.txt")
	numOfSteps := NavigateMap(puzzleMap)
	fmt.Println("Answer is: ", numOfSteps)
}
