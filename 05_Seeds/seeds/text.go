package seeds

import (
	"bufio"
	"fmt"
	"github.com/Leila-Codes/Advent-of-Code-2023/textools"
	"io"
	"os"
	"strings"
)

type PuzzleInput struct {
	Seeds    []int
	Mappings []MappingCollection
}

func (pz PuzzleInput) LocationForSeed(seed int) int {
	for _, m := range pz.Mappings {
		out := m.Map(seed)

		//fmt.Printf("%s\t%d ➡️ %d\n", m, seed, out)

		seed = out
	}

	return seed
}

// ExpandSeeds - Processes the seed pairs and expands to a full list
// puzzle05 - part 2/2
func (pz PuzzleInput) ExpandSeeds(seedOut chan int, errChan chan error) {
	for i := 0; i < len(pz.Seeds); i += 2 {
		var (
			start  = pz.Seeds[i]
			length = pz.Seeds[i+1]
		)

		fmt.Printf("Expanding from %d➡️%d (length=%d)\n", start, start+length, length)

		for j := start; j < start+length; j++ {
			seedOut <- j
		}
	}

	errChan <- io.EOF
}

func ReadInput(filePath string) PuzzleInput {
	pz := PuzzleInput{}

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)
	reader.Scan()

	seedList := reader.Text()
	pz.Seeds = textools.ParseNumberList(seedList[strings.IndexRune(seedList, ':')+1:])

	var mapBuffer string
	for reader.Scan() {
		if len(reader.Text()) == 0 {
			if len(mapBuffer) > 0 {
				pz.Mappings = append(pz.Mappings, MappingCollectionFromText(mapBuffer[1:]))
				mapBuffer = ""
			}
			continue
		}

		mapBuffer += "\n" + reader.Text()
	}

	if len(mapBuffer) > 0 {
		pz.Mappings = append(pz.Mappings, MappingCollectionFromText(mapBuffer))
	}

	return pz
}
