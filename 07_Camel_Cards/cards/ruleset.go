package cards

type Ruleset []ScoreFunc

var DefaultRuleset = []ScoreFunc{HighCard, OnePair, TwoPair, ThreeOfAKind, FullHouse, FourOfAKind, FiveOfAKind}

func (h Hand) GroupWithJacks() (m map[Card]int) {
	m = h.Group()
	if jackCount, hasJacks := m[Jack]; hasJacks {
		for t := range m {
			if t != Jack {
				m[t] += jackCount
			}
		}
	}

	return m
}
