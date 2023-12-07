package scratchcard

import (
	"04_Scratchcards/hashset"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winners hashset.Set[int]
	numbers hashset.Set[int]
}

var reWhiteSpace = regexp.MustCompile("\\s+")

func parseNumberList(list string) []int {
	nums := reWhiteSpace.Split(list, -1)

	var retVal = make([]int, len(nums))

	fmt.Printf("%+v\n", nums)

	for i := 0; i < len(nums); i++ {
		var err error
		retVal[i], err = strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}
	}

	return retVal
}

func CardFromLine(line string) Card {
	idx := strings.IndexRune(line, ':')
	line = line[idx+2:]

	fmt.Println(line)

	idx = strings.IndexRune(line, '|')

	winnersStr := line[0 : idx-1]
	inputStr := line[idx+1:]

	fmt.Printf("Winners: %s\nInput: %s\n", winnersStr, inputStr)

	winners := parseNumberList(strings.TrimSpace(winnersStr))
	input := parseNumberList(strings.TrimSpace(inputStr))

	return Card{
		hashset.NewSetWithValues(winners),
		hashset.NewSetWithValues(input),
	}
}

func CardsFromFile(filePath string) ([]Card, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	cards := make([]Card, 0)

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		card := CardFromLine(reader.Text())
		cards = append(cards, card)
	}

	return cards, nil
}

func (c Card) Score() int {
	score := 0

	for num := range c.winners {
		if c.numbers.Contains(num) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	return score
}
