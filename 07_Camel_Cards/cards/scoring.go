package cards

type ScoreFunc func(Hand) bool

func FiveOfAKind(h Hand) bool {
	groups := h.Group()

	if len(groups) == 1 {
		return true
	}

	for _, count := range groups {
		if count == 5 {
			return true
		}
	}

	return false
}

func FourOfAKind(h Hand) bool {
	groups := h.Group()

	for _, count := range groups {
		if count == 4 {
			return true
		}
	}

	return false
}

func FullHouse(h Hand) bool {
	groups := h.Group()

	var (
		threeOfAKind, twoOfAKind bool
	)

	for _, count := range groups {
		if count == 3 {
			threeOfAKind = true
		}

		if count == 2 {
			twoOfAKind = true
		}
	}

	return threeOfAKind && twoOfAKind
}

func ThreeOfAKind(h Hand) bool {
	groups := h.Group()

	for _, count := range groups {
		if count == 3 {
			return true
		}
	}

	return false
}

func pairs(h Hand) int {
	groups := h.Group()

	var pairCount int

	for _, count := range groups {
		if count == 2 {
			pairCount++
		}
	}

	return pairCount
}

func TwoPair(h Hand) bool {
	pairCount := pairs(h)
	return pairCount == 2
}

func OnePair(h Hand) bool {
	pairCount := pairs(h)
	return pairCount == 1
}

func HighCard(h Hand) bool {
	return len(h.Group()) == 5
}
