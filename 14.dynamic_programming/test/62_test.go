package test

import (
	"fmt"
	"testing"
)

func Test_62(t *testing.T) {
	m, n := 7, 3
	res := uniquePathsV1(m, n)
	fmt.Printf("1.res: %v\n", res)

	res = uniquePathsV2(m, n)
	fmt.Printf("2.res: %v\n", res)
}

func uniquePathsV1(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsV2(m int, n int) int {
	cur := make([]int, m)
	for i := 0; i < m; i++ {
		cur[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[m-1]
}
