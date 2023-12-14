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

type MappingRangeType int

const (
	Source MappingRangeType = iota
	Destination
)

func (mr MappingRange) Contains(input int, rangeType MappingRangeType) bool {
	var (
		min int
		max int
	)

	switch rangeType {
	case Source:
		min, max = mr.SourceStart, mr.SourceStart+mr.Length
	case Destination:
		min, max = mr.DestinationStart, mr.DestinationStart+mr.Length
	}

	return input >= min && input < max
}

func (mr MappingRange) Map(input int) int {
	if mr.Contains(input, Source) {
		theta := mr.DestinationStart - mr.SourceStart
		return input + theta
	}

	return 0
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
