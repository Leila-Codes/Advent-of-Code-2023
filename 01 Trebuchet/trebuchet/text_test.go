package trebuchet

import "testing"

func slicesMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func TestFindAllNumbers(t *testing.T) {

	inputs := []string{
		"two9tfvjqsgqsixnine",
		"plxjdxghsix17",
		"sevenonesevennine4three2seven",
		"7bvdgpghzhpeight512vxbnfqjctb",
		"xqptzkfive4xqbjzpqfkfspqv5kgqbdtfive",
		"ninenineone5two",
		"llv5",
		"94csjjgl42three",
		"4bkxxv6",
		"onetwothreefourfiveightsixsevenine",
	}

	expects := [][]int{
		{2, 9, 6, 9},
		{6, 1, 7},
		{7, 1, 7, 9, 4, 3, 2, 7},
		{7, 8, 5, 1, 2},
		{5, 4, 5, 5},
		{9, 9, 1, 5, 2},
		{5},
		{9, 4, 4, 2, 3},
		{4, 6},
		{1, 2, 3, 4, 5, 8, 6, 7, 9},
	}

	testFindAllNumbersHelper(t, inputs, expects)
}

func testFindAllNumbersHelper(t *testing.T, inputs []string, expects [][]int) {
	for i, in := range inputs {
		actual := FindAllNumbers(in)
		if !slicesMatch(expects[i], actual) {
			t.Errorf("test[%d] - slices did not match - expected '%v' got '%v'\n", i, expects[i], actual)
		}
	}
}
