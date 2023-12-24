package cards

type Hand [5]Card

func (h Hand) Group() map[Card]int {
	m := make(map[Card]int)

	for _, c := range h {
		if v, ok := m[c]; ok {
			m[c] = v + 1
		} else {
			m[c] = 1
		}
	}

	return m
}

func (h Hand) Score(ruleset Ruleset) int {
	for i := len(ruleset) - 1; i >= 0; i-- {
		if ruleset[i](h) {
			return i + 1
		}
	}

	return 0
}

func HandFromText(text string) (h Hand) {
	for i, c := range text {
		h[i] = CardFromText(c)
	}

	return h
}
