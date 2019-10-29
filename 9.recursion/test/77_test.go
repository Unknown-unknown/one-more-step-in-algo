package test

import (
	"fmt"
	"testing"
)

func Test_77(t *testing.T) {
	res := combine(4, 2)
	fmt.Printf("res: %v\n", res)
}

// func combine(n int, k int) [][]int {
// 	if n == 0 || k == 0 || k > n {
// 		return [][]int{}
// 	}
// 	res := [][]int{}
// 	backtrack(1, k, n, []int{}, &res)
// 	return res
// }

// func backtrack(i, k, n int, tmp []int, res *[][]int) {
// 	if k == 0 {
// 		*res = append(*res, tmp)
// 		return
// 	}
// 	for j := i; j < n+1; j++ {
// 		tmp = append(tmp, j)
// 		backtrack(j+1, k-1, n, tmp, res)
// 	}
// 	fmt.Printf("res: %v\n", res)
// }

func combine(n int, k int) [][]int {
	if k == n || k == 0 {
		res := []int{}
		for i := 1; i <= k; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}
	tmp := combine(n-1, k-1)
	res := [][]int{}
	for _, t := range tmp {
		res = append(res, append(t, n))
	}
	res = append(res, combine(n-1, k)...)
	return res
}
