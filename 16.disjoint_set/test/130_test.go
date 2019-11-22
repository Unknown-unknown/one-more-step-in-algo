package test

import (
	"fmt"
	"testing"
)

func Test_130(t *testing.T) {
	grid := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	surronded(grid)
	fmt.Printf("res: %v\n", grid)
}

func surronded(board [][]byte) {
	if len(board) == 0 {
		return
	}
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		check(&board, i, 0)     // first column
		check(&board, i, col-1) // last column
	}
	for j := 1; j < col-1; j++ {
		check(&board, 0, j)     // first row
		check(&board, row-1, j) // last row
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func check(board *[][]byte, i, j int) {
	if (*board)[i][j] == 'O' { // turn 0 into 1 if on the border
		(*board)[i][j] = '1'
		if i-1 > 0 {
			check(board, i-1, j)
		}
		if j-1 > 0 {
			check(board, i, j-1)
		}
		if i+1 < len(*board) {
			check(board, i+1, j)
		}
		if j+1 < len((*board)[0]) {
			check(board, i, j+1)
		}
	}
}
