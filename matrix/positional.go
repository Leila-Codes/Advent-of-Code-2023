package matrix

type Point struct {
	X, Y int
}

type Bounds struct {
	Min, Max Point
}

func (m Matrix[T]) FindAll(what T) []Point {
	locs := make([]Point, 0)

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == what {
				locs = append(locs, Point{x, y})
			}
		}
	}

	return locs
}

func (m Matrix[T]) GetAdjacencyBounds(p1 Point) Bounds {
	return Bounds{
		Point{max(p1.X-1, 0), max(p1.Y-1, 0)},
		Point{min(p1.X+1, m.Width(p1.Y)), min(p1.Y+1, m.Height())},
	}
}
