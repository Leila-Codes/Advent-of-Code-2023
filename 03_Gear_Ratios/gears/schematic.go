package gears

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Schematic [][]rune

func SchematicFromFile(filePath string) (Schematic, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	g := Schematic{}

	for scanner.Scan() {
		g = append(g, []rune(scanner.Text()))
	}

	return g, err
}

type RuneCheckFunc func(y, x int) bool

func (s Schematic) Get(x, y int) rune {
	return s[y][x]
}

func (s Schematic) IsNumber(y, x int) bool {
	return unicode.IsDigit(s[y][x])
}

func (s Schematic) IsSpecial(y, x int) bool {
	return IsSpecial(s[y][x])
}

func (s Schematic) AdjacentTo(y, x int, test RuneCheckFunc) bool {
	var (
		minX = max(x-1, 0)
		minY = max(y-1, 0)
		maxX = min(x+1, len(s[y])-1)
		maxY = min(y+1, len(s)-1)
	)

	fmt.Printf("checking adjacency of [%d,%d] ('%s')\n", x, y, string(s[y][x]))

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			res := test(y, x)
			fmt.Printf("test [%d,%d] - %t\n", x, y, res)
			if res {
				return true
			}
		}
	}

	return false

}
