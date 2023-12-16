package races

import (
	"bufio"
	"github.com/leila-codes/advent-of-code-2023/textools"
	"os"
	"strings"
)

func FromFile(filePath string) []Race {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	var (
		times, distances []int
	)

	for reader.Scan() {
		parts := strings.Split(reader.Text(), ":")

		switch strings.ToLower(parts[0]) {
		case "time":
			times = textools.ParseNumberList(parts[1])
		case "distance":
			distances = textools.ParseNumberList(parts[1])
		}
	}

	races := make([]Race, len(times))

	for i := 0; i < len(times); i++ {
		races[i] = Race{
			Duration:       times[i],
			RecordDistance: distances[i],
		}
	}

	return races
}
