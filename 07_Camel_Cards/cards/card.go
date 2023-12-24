package cards

type Card uint8

const (
	Two Card = 2 + iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func CardFromText(char rune) Card {
	switch char {
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return Ten
	case 'J':
		return Jack
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		fallthrough
	default:
		return Ace
	}
}
