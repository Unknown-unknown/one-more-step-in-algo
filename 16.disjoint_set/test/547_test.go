package test

import (
	"fmt"
	"testing"
)

func Test_547(t *testing.T) {
	M := [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 0},
		[]int{0, 0, 1},
	}
	res := findCircleNum(M)
	fmt.Printf("1.res: %v\n", res)

	M = [][]int{
		[]int{1, 1, 0},
		[]int{1, 1, 1},
		[]int{0, 1, 1},
	}
	res = findCircleNum(M)
	fmt.Printf("2.res: %v\n", res)
}

func findCircleNum(M [][]int) int {
	parent := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		parent[i] = -1
	}

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && i != j {
				union(i, j, parent)
			}
		}
	}

	count := 0
	for _, v := range parent {
		if v == -1 {
			count++
		}
	}
	return count
}

func union(i, j int, parent []int) {
	xset, yset := find(i, parent), find(j, parent)
	if xset != yset {
		parent[xset] = yset
	}
}

func find(i int, parent []int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent[i], parent)
}
