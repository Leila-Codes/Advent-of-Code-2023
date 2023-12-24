package main

import (
	"09_Mirage_maintenance/sequence"
	"bufio"
	"fmt"
	"github.com/leila-codes/advent-of-code-2023/textools"
	"os"
)

func summariseSequences(filePath string) (sum int) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		seq := sequence.NewSequence(textools.ParseNumberList(reader.Text()))
		sequence.ExtrapolateAll(seq)
		sum += sequence.NextInSequence(seq)
	}

	return sum
}

func main() {
	fmt.Println("Answer is: ", summariseSequences("Puzzle09_Input.txt"))
}
