package cards

import (
	"strconv"
	"strings"
)

type Game struct {
	Hand
	Bet, Score int
}

func GameFromText(line string, ruleset Ruleset) (game Game) {
	parts := strings.Split(line, " ")

	game.Hand = HandFromText(parts[0])

	game.Bet, _ = strconv.Atoi(strings.TrimSpace(parts[1]))

	game.Score = game.Hand.Score(ruleset)

	return game
}
