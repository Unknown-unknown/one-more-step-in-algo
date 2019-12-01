package test

import "testing"

func Test_52(t *testing.T) {
}

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	count := 0
	dfs(n, 0, 0, 0, 0, &count)
	return count
}

func dfs(n, row, cols, pie, na int, count *int) {
	if row >= n {
		*count++
		return
	}
	bits := (^(cols | pie | na)) & ((1 << n) - 1)
	for bits > 0 {
		p := bits & -bits        // 取到最低位的1
		bits = bits & (bits - 1) // 在 p 位置放上皇后
		dfs(n, row+1, cols|p, (pie|p)<<1, (na|p)>>1, count)
	}
}
