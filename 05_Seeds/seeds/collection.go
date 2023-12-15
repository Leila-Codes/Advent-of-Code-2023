package seeds

import (
	"fmt"
	"strings"
	"unicode"
)

type MappingCollection struct {
	From, To string
	Ranges   []MappingRange
}

func (mc MappingCollection) String() string {
	return fmt.Sprintf("%11s ➡️ %-11s", mc.From, mc.To)
}

func (mc MappingCollection) Map(input int) int {
	for _, m := range mc.Ranges {
		//fmt.Printf("mapping %s")
		if m.Contains(input) {
			return m.Map(input)
		}
	}

	return input
}

func MappingCollectionFromText(textBlock string) MappingCollection {
	var (
		lines = strings.Split(textBlock, "\n")
		col   = MappingCollection{}
	)

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		if unicode.IsDigit(rune(l[0])) {
			col.Ranges = append(col.Ranges, MappingRangeFromText(l))
		} else {
			parts := strings.Split(l, "-to-")
			col.From = parts[0]
			col.To = parts[1][0:strings.Index(parts[1], " map:")]
		}
	}

	return col
}
