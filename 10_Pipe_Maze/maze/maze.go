package maze

import (
	"bufio"
	"github.com/leila-codes/advent-of-code-2023/matrix"
	"os"
)

type Maze struct {
	*matrix.Matrix[Pipe]
}

// FromFile - Loads a maze from the given filepath into memory.
func FromFile(filePath string) *Maze {
	maze := new(Maze)

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		*maze.Matrix = append(*maze.Matrix, []Pipe(reader.Text()))
	}

	return maze
}
