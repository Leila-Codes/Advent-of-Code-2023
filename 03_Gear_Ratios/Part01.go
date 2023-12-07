package main

import (
	"03_Gear_Ratios/gears"
	"fmt"
	"strconv"
	"unicode"
)

func GetEngineParts(filePath string) (partNumbers []int) {
	schematic, err := gears.SchematicFromFile(filePath)
	if err != nil {
		panic(err)
	}

	var (
		buff         []rune
		intBuff      []rune
		isPartNumber bool
	)

	for y := 0; y < len(schematic); y++ {
		for x := 0; x < len(schematic[y]); x++ {
			char := schematic[y][x]
			buff = append(buff, char)

			if unicode.IsDigit(char) {
				intBuff = append(intBuff, char)

				//fmt.Printf("'%s' is number, intBuff = %v\n", string(char), string(intBuff))

				if !isPartNumber && schematic.AdjacentTo(y, x, schematic.IsSpecial) {
					isPartNumber = true
				}
			} else {
				//fmt.Printf("'%s' is NOT number\n", string(char))
				if len(intBuff) > 0 {
					if isPartNumber {
						v, err := strconv.Atoi(string(intBuff))
						if err != nil {
							panic(err)
						}

						fmt.Printf("part no. detected %d\n", v)

						partNumbers = append(partNumbers, v)
						isPartNumber = false
					}

					intBuff = make([]rune, 0)
				}
			}
		}
	}

	return partNumbers
}

func main() {
	fmt.Println("answer is: ", gears.Sum(GetEngineParts("Puzzle03_Input.txt")...))
}
