package test

import (
	"fmt"
	"testing"
)

func Test_17(t *testing.T) {
	res := letterCombinations("23")
	fmt.Printf("res: %v", res)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	mapDict := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	res := []string{}
	search("", digits, 0, &res, mapDict)
	return res
}

func search(s, digits string, i int, res *[]string, dict map[rune][]string) {
	// terminator
	if i == len([]rune(digits)) {
		*res = append(*res, s)
		return
	}
	// process current logicp
	key := []rune(digits)[i]
	for j := 0; j < len(dict[key]); j++ {
		// drill down
		search(s+dict[key][j], digits, i+1, res, dict)
	}
	// reverse states
}
