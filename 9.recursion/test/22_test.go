package test

import (
	"fmt"
	"testing"
)

func Test_22(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Printf("res: %v", res)
}

func generateParenthesis(n int) []string {
	res := []string{}
	return generate(0, 0, n, "", res)
}

// generate all possibles
// * how to check valid when processing: left 随时可以加，只要不超标；right 左个数>右个数
// * 剪枝
func generate(left, right, n int, s string, res []string) []string {
	// terminator
	if left == n && right == n {
		fmt.Println(s)
		res = append(res, s)
		return res
	}

	// process current logic: left, right
	// drill down
	if left <= n {
		l := generate(left+1, right, n, s+"(", res)
		res = l
	}
	if left > right {
		r := generate(left, right+1, n, s+")", res)
		res = r
	}

	// reverse states
	return res
}
