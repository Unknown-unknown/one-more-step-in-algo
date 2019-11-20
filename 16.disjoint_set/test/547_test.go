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
	fmt.Printf("init parent: %v\n", parent)

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && i != j {
				fmt.Printf("%d + %d\n", i, j)
				union(i, j, parent)
				fmt.Printf("parent: %v\n", parent)
			}
		}
	}

	fmt.Printf("final parent: %v\n", parent)

	count := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			count++
		}
	}
	return count
}

func find(i int, parent []int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent[i], parent)
}

func union(i, j int, parent []int) {
	xset, yset := find(i, parent), find(j, parent)
	if xset != yset {
		parent[xset] = yset
	}
}
