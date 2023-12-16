package textools

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseNumberList(line string) []int {
	var (
		parts = strings.Split(line, " ")
		nums  = make([]int, 0)
	)

	for _, p := range parts {
		p = strings.TrimSpace(p)

		if len(p) == 0 {
			continue
		}

		v, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println(err)
			continue
		}

		nums = append(nums, v)
	}

	return nums
}
