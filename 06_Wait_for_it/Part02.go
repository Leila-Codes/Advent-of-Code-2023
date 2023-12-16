package main

import (
	"06_Wait_for_it/races"
	"fmt"
)

func main() {
	race := races.Race{
		Duration:       41667266,
		RecordDistance: 244104712281040,
	}

	fmt.Println("Answer is: ", race.WinnerHoldsMS())
}
