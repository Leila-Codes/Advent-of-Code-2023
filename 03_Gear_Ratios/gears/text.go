package gears

import (
	"fmt"
	"strconv"
	"unicode"
)

// for each character
//     append to buffer
//     if a digit
//		   is previous a digit?
//             append to number buffer
//         is previous a special char?
//            yes: continue
//            no: clear buffers

var specialChars = []rune{'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}

func IsSpecial(char rune) bool {
	for _, r := range specialChars {
		if char == r {
			return true
		}
	}

	if char >= 128 {
		return true
	}
	//return (char >= 33 && char <= 45) ||
	//	(char >= 58 && char <= 64) ||
	//	(char >= 91 && char <= 96) ||
	//	(char >= 123 && char <= 126)

	//
	return false
}

func ParseLine(line string) (partNumbers []int) {

	var (
		buff    []rune
		intBuff []rune
	)

	for _, char := range line {
		//buff += string(char)
		buff = append(buff, char)

		// separate the numbers
		if unicode.IsDigit(char) /*|| char == '-'*/ {
			intBuff = append(intBuff, char)
		}

		//fmt.Printf("buff = %s\nintBuff = %s\n", string(buff), string(intBuff))

		if char == '.' || IsSpecial(char) {

			// was there a number in the buffer?
			var lengthGreatherThan = 0
			//if strings.HasPrefix("-", string(intBuff)) {
			//	lengthGreatherThan = 1
			//}

			if len(intBuff) > lengthGreatherThan {

				var (
					intStart = len(buff) - len(intBuff)
					intEnd   = len(buff) - 1
				)

				if len(buff) > len(intBuff)+1 {
					intStart -= 2
				} else {
					intStart -= 1
				}

				// does it start/end with special?
				if len(intBuff) == len(buff) || IsSpecial(buff[intStart]) || IsSpecial(buff[intEnd]) {
					val, err := strconv.Atoi(string(intBuff))

					if err != nil {
						panic(err)
					}

					fmt.Printf("Found int %d\n", val)

					partNumbers = append(partNumbers, val)
				} else {
					//fmt.Printf("Skipped int %s as no special char\n", string(intBuff))
				}

				buff = buff[len(buff)-1:]
				intBuff = make([]rune, 0)
			}
		}
	}

	return partNumbers
}
