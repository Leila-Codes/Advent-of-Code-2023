package main

import (
	"07_Camel_Cards/cards"
	"bufio"
	"fmt"
	"os"
	"sort"
)

func RankGames(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	var games []cards.Game

	for reader.Scan() {
		games = append(games, cards.GameFromText(reader.Text(), cards.DefaultRuleset))
	}

	sort.Slice(games, func(i, j int) bool {
		if games[j].Score == games[i].Score {
			fmt.Printf("Games [%d] & [%d] hold same score - checking cards...\n", j, i)
			for k := 0; k < 5; k++ {
				fmt.Printf("Card [%d] - (Player%d = %d) - (Player%d = %d)\n", k, j, games[j].Hand[k], i, games[i].Hand[k])
				if games[j].Hand[k] != games[i].Hand[k] {
					return games[j].Hand[k] > games[i].Hand[k]
				}
			}
		}

		return games[j].Score > games[i].Score
	})

	var totalScore int

	for i := 0; i < len(games); i++ {
		totalScore += (i + 1) * games[i].Bet
	}

	return totalScore
}

func main() {
	fmt.Println("Answer is: ", RankGames("Puzzle07_Input.txt"))
}
