package main

import (
	"05_Seeds/seeds"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		input   = seeds.ReadInput("Puzzle05_Input.txt")
		seedOut = make(chan int)
		errChan = make(chan error)
	)

	go input.ExpandSeeds(seedOut, errChan)

	var (
		minLocation = -1
	)

	for {
		select {
		case seed := <-seedOut:
			out := input.LocationForSeed(seed)
			if minLocation == -1 {
				minLocation = out
				continue
			}

			minLocation = min(minLocation, out)
		case err := <-errChan:
			if err == io.EOF {
				fmt.Println("Answer is: ", minLocation)
				os.Exit(0)
			} else {
				panic(err)
			}
		}
	}
}
