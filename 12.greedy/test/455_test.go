package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_455(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	res := findContentChildren(g, s)
	fmt.Printf("1.res: %v\n", res)

	g = []int{1, 2}
	s = []int{1, 2, 3}
	res = findContentChildren(g, s)
	fmt.Printf("2.res: %v\n", res)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j, res := 0, 0, 0
	gLen, sLen := len(g), len(s)
	for i < gLen && j < sLen {
		if g[i] <= s[j] {
			res++
			i++
			j++
		} else {
			j++
		}
	}
	return res
}
