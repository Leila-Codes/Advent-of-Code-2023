package main

import (
	"03_Gear_Ratios/gears"
	"fmt"
)

func GetGears(filePath string) (ratios []int) {
	schematic, err := gears.SchematicFromFile("Puzzle03_Example_Input.txt")
	if err != nil {
		panic(err)
	}

	// Find all asterisks
	locations := schematic.FindAll('*')

	for _, loc := range locations {
		fmt.Printf("Found '*' at [%d,%d]\n", loc.X, loc.Y)
		numberLocations := schematic.FindAdjacent(loc.Y, loc.X, schematic.IsNumber)
		fmt.Printf("Found %d adjacent numbers...\n", len(numberLocations))

		if isGear := len(numberLocations) >= 2; isGear {
			var (
				numOne = schematic.ReadInt(numberLocations[0])
				numTwo = schematic.ReadInt(numberLocations[1])
			)

			ratios = append(ratios, numOne*numTwo)
		}
	}

	return ratios
}

func main() {
	ratios := GetGears("Puzzle03_Input.txt")
	fmt.Println("Answer is: ", gears.Sum(ratios...))
}
