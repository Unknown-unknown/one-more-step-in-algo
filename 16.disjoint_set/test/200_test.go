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
	res1 := numIslandsV1(grid)

	grid = [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	res2 := numIslandsV2(grid)
	fmt.Printf("1.res: %v, %v\n", res1, res2)

	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	res1 = numIslandsV1(grid)
	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	res2 = numIslandsV2(grid)
	fmt.Printf("2.res: %v, %v\n", res1, res2)
}

// v1, flood fill
var g [][]byte
var dx, dy []int

func numIslandsV1(grid [][]byte) int {
	g = grid
	dx, dy = []int{-1, 1, 0, 0}, []int{0, 0, 1, -1}
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if g[i][j] == '1' {
				sink(i, j)
				num++
			}
		}
	}
	return num
}

func sink(i, j int) {
	if i < 0 || i >= len(g) || j < 0 || j >= len(g[0]) || g[i][j] == '0' {
		return
	}

	g[i][j] = '0'

	for k := 0; k < len(dx); k++ {
		sink(i+dx[k], j+dy[k])
	}
}

// v2, disjoint set
var count int

func numIslandsV2(grid [][]byte) int {
	g = grid
	dx, dy := []int{1, 0, -1, 0}, []int{0, 1, 0, -1}

	row := len(g)
	if row == 0 {
		return 0
	}
	col := len(g[0])
	parent := make([]int, row*col)
	count = 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if g[i][j] == '1' {
				parent[i*col+j] = i*col + j
				count++
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if g[i][j] == '1' {
				for k := 0; k < len(dx); k++ {
					tmpX, tmpY := i+dx[k], j+dy[k]
					if tmpX >= 0 && tmpX < row && tmpY >= 0 && tmpY < col && grid[tmpX][tmpY] == '1' {
						unionIsland(tmpX*col+tmpY, i*col+j, parent)
					}
				}
			}
		}
	}

	return count
}

func unionIsland(i, j int, parent []int) {
	xset, yset := findIsland(i, parent), findIsland(j, parent)
	if xset != yset {
		parent[xset] = yset
		count--
	}
}

func findIsland(i int, parent []int) int {
	if parent[i] == i {
		return i
	}
	return findIsland(parent[i], parent)
}
