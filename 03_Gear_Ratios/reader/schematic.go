package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Schematic []string

func SchematicFromFile(filePath string) *Schematic {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	schematic := new(Schematic)

	for reader.Scan() {
		*schematic = append(*schematic, reader.Text())
	}

	return schematic
}

type Point struct {
	X, Y int
}

type Bounds struct {
	Min, Max Point
}

func (p Point) String() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}

func (s Schematic) FindAll(what rune) []Point {
	locs := make([]Point, 0)

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if rune(s[y][x]) == what {
				locs = append(locs, Point{x, y})
			}
		}
	}

	return locs
}

func (s Schematic) GetAdjacencyBounds(p1 Point) Bounds {
	return Bounds{
		Point{max(p1.X-1, 0), max(p1.Y-1, 0)},
		Point{min(p1.X+1, len(s[p1.Y])), min(p1.Y+1, len(s))},
	}
}

func (s Schematic) ReadInt(p1 Point, visited *Set[Point]) int {
	var x = p1.X
	// 1) read backwards from the point to the start of the int
	for x = p1.X; x > 0 && unicode.IsDigit(rune(s[p1.Y][x-1])); x-- {
	}

	fmt.Printf("Reading int from %s -> %s\n", Point{x, p1.Y}, s[p1.Y][x:x+3])

	// 2) read the integer
	var intBuff string
	for x < len(s[p1.Y]) && unicode.IsDigit(rune(s[p1.Y][x])) {
		intBuff += string(s[p1.Y][x])
		visited.Add(Point{x, p1.Y})
		x++
	}

	val, err := strconv.Atoi(intBuff)
	if err != nil {
		panic(err)
	}

	return val
}

func (s Schematic) ReadAdjacentIntegers(p1 Point) []int {
	var (
		bounds  = s.GetAdjacencyBounds(p1)
		visited = NewSet[Point]()
		nums    = make([]int, 0)
	)

	for y := bounds.Min.Y; y <= bounds.Max.Y; y++ {
		for x := bounds.Min.X; x <= bounds.Max.X; x++ {

			p := Point{x, y}

			if !visited.Contains(p) {

				if unicode.IsDigit(rune(s[y][x])) {
					fmt.Printf("Found number at: [%d,%d]\n", x, y)

					nums = append(nums, s.ReadInt(p, &visited))
				}
			}
		}
	}

	return nums
}

//visited[Point{x, y}] = struct{}{}

/*// Seek backwards to ensure ENTIRE number read
	var x2 = x
	for x > 0 && unicode.IsDigit(rune(s[y][x2-1])) {
		x--
	}

	fmt.Printf("Number starts at: [%d,%d]\n", x2, y)

	for x3 := x2; x3 < len(s[y]) && unicode.IsDigit(rune(s[y][x3])); x3++ {
		intBuff = append(intBuff, rune(s[y][x3]))

		p2 := Point{x3, y}
		visited.Add(p2)
	}

} else if len(intBuff) > 0 {

	val, err := strconv.Atoi(string(intBuff))
	fmt.Printf("Number is: %d\n", val)
	nums = append(nums, val)
	intBuff = make([]rune, 0)
	if err != nil {
		panic(err)
	}*/
