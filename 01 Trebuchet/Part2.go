package main

import (
	"01_Trebuchet/trebuchet"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numMatch = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")

var puzzleInput string

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
	lines := strings.Split(puzzleInput, "\n")

	var sum int

	for _, line := range lines {
		nums := trebuchet.FindAllNumbers(line)

		if len(nums) < 1 {
			continue
		}

		first, last := trebuchet.FirstAndLast(nums)

		intStr := strconv.Itoa(first) + strconv.Itoa(last)

		value, _ := strconv.Atoi(intStr)

		sum += value

		//fmt.Printf("numbers '%v' -> '%s'\n", nums, intStr)
	}

	fmt.Println("Answer is: ", sum)
}
