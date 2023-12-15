package scratchcard

import (
	"bufio"
	"fmt"
	"github.com/leila-codes/advent-of-code-2023/hashset"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	winners  []int
	numbers  hashset.Set[int]
	Replicas int
}

var reWhiteSpace = regexp.MustCompile("\\s+")

func parseNumberList(list string) []int {
	nums := reWhiteSpace.Split(list, -1)

	var retVal = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var err error
		retVal[i], err = strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}
	}

	return retVal
}

func CardFromLine(line string) *Card {
	idx := strings.IndexRune(line, ':')
	line = line[idx+2:]

	fmt.Println(line)

	idx = strings.IndexRune(line, '|')

	winnersStr := line[0 : idx-1]
	inputStr := line[idx+1:]

	winners := parseNumberList(strings.TrimSpace(winnersStr))
	input := parseNumberList(strings.TrimSpace(inputStr))

	return &Card{
		winners,
		hashset.NewSetWithValues(input),
		1,
	}
}

func CardsFromFile(filePath string) ([]*Card, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	cards := make([]*Card, 0)

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		card := CardFromLine(reader.Text())
		cards = append(cards, card)
	}

	return cards, nil
}

func (c *Card) String() string {
	return fmt.Sprintf("Winners: %v, Numbers: %v", c.winners, c.numbers)
}

func (c *Card) Score() int {
	score := 0

	for _, num := range c.winners {
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

func (c *Card) NumberOfWinners() (winners int) {
	for _, num := range c.winners {
		if c.numbers.Contains(num) {
			winners++
		}
	}

	return winners
}

func (c *Card) Copy() {
	c.Replicas += 1
}
