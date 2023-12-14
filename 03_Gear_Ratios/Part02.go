package main

import (
	"03_Gear_Ratios/reader"
	"fmt"
)

func GetGears(filePath string) (ratios []int) {
	schematic := reader.SchematicFromFile(filePath)

	// Find all asterisks
	locations := schematic.FindAll('*')

	fmt.Printf("Found %d '*' chars...\n", len(locations))

	for _, loc := range locations {
		nums := schematic.ReadAdjacentIntegers(loc)

		fmt.Printf("Location %s has '*' contains %d nums...\t %v\n", loc, len(nums), nums)

		if len(nums) == 2 {
			ratio := nums[0] * nums[1]

			ratios = append(ratios, ratio)
		}
	}

	return ratios
}

func Sum(vals ...int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

func main() {
	ratios := GetGears("Puzzle03_Input.txt")
	fmt.Println("Answer is: ", Sum(ratios...))
}
