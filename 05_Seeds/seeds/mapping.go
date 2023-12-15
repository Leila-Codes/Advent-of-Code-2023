package seeds

import (
	"strconv"
	"strings"
)

type MappingRange struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func (mr MappingRange) Contains(input int) bool {

	return input >= mr.SourceStart &&
		input < mr.SourceStart+mr.Length
}

func (mr MappingRange) Map(input int) int {
	if mr.Contains(input) {
		theta := mr.DestinationStart - mr.SourceStart
		return input + theta
	}

	return input
}

func MappingRangeFromText(line string) MappingRange {
	parts := strings.Split(line, " ")

	dest, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	src, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	length, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	return MappingRange{dest, src, length}
}
