package main

import "10_Pipe_Maze/maze"

func main() {
	m := maze.FromFile("Puzzle10_Input.txt")

	start := m.FindAll(maze.Animal)

	for _, s := range start {
		bounds := m.GetAdjacencyBounds(s)

		paths := make([]maze.Pipe, 9)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if m.Get(x, y) != maze.Ground {
					paths = append(paths, m.Get(x, y))
				}
			}
		}
	}
}
