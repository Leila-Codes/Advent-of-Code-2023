package game

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeColour string

const (
	ColourRed   CubeColour = "red"
	ColourBlue  CubeColour = "blue"
	ColourGreen CubeColour = "green"
)

type GameSummary struct {
	ID        uint
	CubeCount map[CubeColour]uint
}

func (s GameSummary) String() string {
	var description strings.Builder

	description.WriteString(fmt.Sprintf("Game %d: ", s.ID))
	for colour, number := range s.CubeCount {
		description.WriteString(fmt.Sprintf("%s Max. %d\t", colour, number))
	}

	return description.String()
}

func (s GameSummary) Possible(limits map[CubeColour]uint) bool {
	for colour, limit := range limits {
		if s.CubeCount[colour] > limit {
			return false
		}
	}

	return true
}

func ParseGameSummary(input string) GameSummary {
	parts := strings.Split(input, ":")
	id, _ := strconv.ParseUint(strings.Split(parts[0], " ")[1], 10, 32)

	summary := GameSummary{ID: uint(id), CubeCount: make(map[CubeColour]uint)}

	cubeCounts := strings.Split(parts[1], " ")
	for i := 1; i < len(cubeCounts); i += 2 {
		//fmt.Printf("RAW [%d='%s'] [%d='%s']\n", i, cubeCounts[i], i+1, cubeCounts[i+1])
		count, err := strconv.ParseUint(cubeCounts[i], 10, 32)
		if err != nil {
			panic(err)
		}

		var colour CubeColour
		colourStr := cubeCounts[i+1]

		if strings.HasPrefix(colourStr, string(ColourRed)) {
			colour = ColourRed
		} else if strings.HasPrefix(colourStr, string(ColourBlue)) {
			colour = ColourBlue
		} else if strings.HasPrefix(colourStr, string(ColourGreen)) {
			colour = ColourGreen
		}
		//fmt.Printf("parsing '%d' cubes as colour '%s'\n", count, colour)

		summary.CubeCount[colour] = Max(summary.CubeCount[colour], uint(count))
	}

	return summary
}

func Max(a, b uint) uint {
	if b > a {
		return b
	}

	return a
}
