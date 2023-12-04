package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	puzzleInput string
)

func init() {
	f, err := os.Open("Part1_Input.txt")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	puzzleInput = string(data)
}

func main() {

	// 1) Split the text into lines
	lines := strings.Split(puzzleInput, "\n")

	// prepare to sum the lines together
	var sum int

	// 2) for each line of text
	for _, line := range lines {
		// 2a) find all the "digit" characters (numbers)
		nums := make([]rune, 0)
		for _, char := range line {
			if unicode.IsDigit(char) {
				nums = append(nums, char)
			}
		}

		if len(nums) > 0 {
			// 2b) concatenate first and last number and parse as an integer
			val, err := strconv.Atoi(string(nums[0]) + string(nums[len(nums)-1]))
			if err != nil {
				panic(err)
			}

			// 2c) sum the number with the previous
			sum += val
		}
	}

	fmt.Println("Answer is: ", sum)
}
