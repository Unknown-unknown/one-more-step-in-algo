package test

import (
	"fmt"
	"testing"
)

func Test_51(t *testing.T) {
	res := solveNQueens(4)
	fmt.Printf("res: %v", res)
}

func solveNQueens(n int) [][]string {
	res := [][]string{} // res 用来存放每种解法中 Queen 的列位置的数组
	shu, pie, na := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	DFS(n, []int{}, shu, pie, na, &res)
	return res
}

func DFS(n int, rows []int, shu map[int]bool, pie map[int]bool, na map[int]bool, res *[][]string) {
	row := len(rows)
	if row == n {
		sol := print(rows, n)
		*res = append(*res, sol)
		return
	}

	for col := 0; col < n; col++ {
		if !shu[col] && !pie[row-col] && !na[row+col] {
			shu[col] = true
			pie[row-col] = true
			na[row+col] = true
			DFS(n, append(rows, col), shu, pie, na, res)
			shu[col] = false
			pie[row-col] = false
			na[row+col] = false
		}
	}
}

func print(cols []int, n int) []string {
	rows := []string{}
	for j := 0; j < len(cols); j++ { // 每种解法中每行 Queen 位于哪列
		row := ""
		for k := 0; k < n; k++ { // 每种解法的每一行
			if k == cols[j] {
				row += "Q"
			} else {
				row += "."
			}
		}
		rows = append(rows, row)
	}
	return rows
}
