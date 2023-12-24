package maze

type Pipe rune

const (
	Ground     Pipe = '.'
	Vertical   Pipe = '|'
	Horizontal Pipe = '-'
	NorthEast  Pipe = 'L'
	NorthWest  Pipe = 'J'
	SouthWest  Pipe = '7'
	SouthEast  Pipe = 'F'
	Animal     Pipe = 'S'
)

//var Pipes = []Pipe{Ground, Vertical, Horizontal, NorthEast, NorthWest, SouthWest}
