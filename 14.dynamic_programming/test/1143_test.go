package test

import (
	"fmt"
	"testing"
)

func Test_1143(t *testing.T) {
	s1, s2 := "abazdc", "bacbad"
	res := longestCommonSubsequenceV1(s1, s2)
	fmt.Printf("1.res: %v\n", res)
	res = longestCommonSubsequenceV2(s1, s2)
	fmt.Printf("1.1.res: %v\n", res)

	s1, s2 = "abcde", "ace"
	res = longestCommonSubsequenceV1(s1, s2)
	fmt.Printf("2.res: %v\n", res)
	res = longestCommonSubsequenceV2(s1, s2)
	fmt.Printf("2.1.res: %v\n", res)
}

// 二维数组
func longestCommonSubsequenceV1(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if []rune(text1)[i-1] == []rune(text2)[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 一维数组
func longestCommonSubsequenceV2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cur, prev := make([]int, m+1), make([]int, m+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			maxLen := max(prev[j], cur[j-1])
			if []rune(text1)[j-1] == []rune(text2)[i-1] {
				cur[j] = max(maxLen, prev[j-1]+1)
			} else {
				cur[j] = maxLen
			}
		}
		cur, prev = prev, cur
	}
	return prev[m]
}
