package gears

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Schematic [][]rune

type Coord struct {
	X, Y int
}

//type Coord [2]int

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

func (s Schematic) Get(location Coord) rune {
	return s[location.Y][location.X]
}

func (s Schematic) IsNumber(y, x int) bool {
	return unicode.IsDigit(s[y][x])
}

func (s Schematic) IsSpecial(y, x int) bool {
	return IsSpecial(s[y][x])
}

func (s Schematic) getAdjacentBounds(y, x int) (minX, minY, maxX, maxY int) {
	minX = max(x-1, 0)
	minY = max(y-1, 0)
	maxX = min(x+1, len(s[y])-1)
	maxY = min(y+1, len(s)-1)

	return minX, minY, maxX, maxY
}

func (s Schematic) AdjacentTo(y, x int, test RuneCheckFunc) bool {
	minX, minY, maxX, maxY := s.getAdjacentBounds(y, x)

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

func (s Schematic) FindAdjacent(startY, startX int, test RuneCheckFunc) []Coord {
	minX, minY, maxX, maxY := s.getAdjacentBounds(startY, startX)

	locs := make([]Coord, 0)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if test(y, x) {
				locs = append(locs, Coord{x, y})
			}
		}
	}

	return locs
}

func (s Schematic) ReadInt(location Coord) int {
	x := location.X

	for x > 0 && !s.IsNumber(location.Y, x) {
		x -= 1
	}

	var buff []rune
	for x := location.X; x < len(s[location.Y]); x++ {
		if s.IsNumber(location.Y, x) {
			buff = append(buff, s[location.Y][x])
		} else {
			break
		}
	}

	v, err := strconv.Atoi(string(buff))
	if err != nil {
		panic(err)
	}

	return v
}

func (s Schematic) FindAll(what rune) []Coord {
	locs := make([]Coord, 0)

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == what {
				locs = append(locs, Coord{x, y})
			}
		}
	}

	return locs
}
