package test

import (
	"fmt"
	"testing"
)

func Test_200(t *testing.T) {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	res := numIslands(grid)
	fmt.Printf("1.res: %v\n", res)

	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	res = numIslands(grid)
	fmt.Printf("2.res: %v\n", res)
}

var g [][]byte
var dx []int
var dy []int

func numIslands(grid [][]byte) int {
	g = grid
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				sink(i, j)
				num++
			}
		}
	}
	return num
}

func sink(i, j int) {
	// terminator
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[0]) || g[i][j] == '0' {
		return
	}
	// current logic
	g[i][j] = '0'

	// drill down
	for k := 0; k < len(dx); k++ {
		x, y := i+dx[k], j+dy[k]
		sink(x, y)
	}
}
