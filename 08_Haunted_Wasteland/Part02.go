package main

import (
	"08_Haunted_Wasteland/documents"
	"fmt"
	"strings"
)

func endsWithA(input string) bool {
	return strings.HasSuffix(input, "A")
}

func parallelWorkerThread(current *documents.MapNode, instructionChan <-chan rune, resultChan chan<- bool, test func(string) bool) {
	for {
		step := <-instructionChan

		switch step {
		case 'L':
			current = current.Left
		case 'R':
			current = current.Right
		}

		resultChan <- test(current.Value)
	}
}

func ParallelMapNavigate(wm *documents.WastelandMap) int {
	var (
		currentNodes = wm.FindAll(endsWithA)
		numOfSteps   int
		arrived      bool
	)

	for !arrived {
		nextStep := wm.Instructions[numOfSteps%len(wm.Instructions)]

		var allArrived = true
		for i, node := range currentNodes {
			switch nextStep {
			case 'L':
				currentNodes[i] = node.Left
			case 'R':
				currentNodes[i] = node.Right
			}

			if currentNodes[i].Value[2] == 'Z' {
				fmt.Printf("Ghost #%d - arrived in %d steps\n", i, numOfSteps)
			}

			if allArrived && currentNodes[i].Value[2] != 'Z' {
				allArrived = false
			}
		}

		numOfSteps++

		if allArrived {
			break
		}
	}

	return numOfSteps
}

func main() {
	myMap := documents.MapFromFile("Puzzle08_Input.txt")
	numOfSteps := ParallelMapNavigate(myMap)
	fmt.Println("Answer is: ", numOfSteps)
}
