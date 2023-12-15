package trebuchet

import (
	"strconv"
	"strings"
	"unicode"
)

var resolvers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func FindAllNumbers(input string) []int {
	var (
		nums = make([]int, 0)
		tmp  string
	)

	for _, char := range input {
		if unicode.IsDigit(char) {
			//fmt.Printf("char '%s' is numeric\n", string(char))
			value, _ := strconv.Atoi(string(char))
			nums = append(nums, value)
		} else {
			tmp += string(char)

			for text, val := range resolvers {
				if strings.HasSuffix(tmp, text) {
					nums = append(nums, val)

					newTmp := tmp[len(tmp)-1:]

					tmp = newTmp
					break
				}
			}
		}
	}

	return nums
}
