package test

import (
	"fmt"
	"testing"
)

func Test_77(t *testing.T) {
	res := combineRecursion(4, 2)
	fmt.Printf("res of recursion: %v\n", res)
	res = combineBacktracking(4, 2)
	fmt.Printf("res of backtracking: %v\n", res)
}

func combineBacktracking(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}
	res := [][]int{}
	elems := []int{}
	backtracking(1, k, n, &elems, &res)
	return res
}

func backtracking(i, k, n int, elems *[]int, res *[][]int) {
	if k == 0 {
		tmp := append([]int{}, *elems...) // <- 需要 copy 出一份新的 elems，append 到 res 中
		fmt.Printf("before append: %v, elems: %v\n", *res, *elems)
		*res = append(*res, tmp)
		fmt.Printf("append: %v, elems: %v\n", *res, *elems)
		return
	}
	for j := i; j < n+1; j++ {
		*elems = append(*elems, j) // 回
		fmt.Printf("回 %v\n", elems)
		backtracking(j+1, k-1, n, elems, res)
		*elems = (*elems)[:len(*elems)-1] // 溯
		fmt.Printf("溯: %v\n", elems)
	}
	fmt.Printf("backtracking: %v, %v\n", res, elems)
}

func combineRecursion(n int, k int) [][]int {
	if k == n || k == 0 {
		res := []int{}
		for i := 1; i <= k; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}
	tmp := combineRecursion(n-1, k-1)
	res := [][]int{}
	for _, t := range tmp {
		res = append(res, append(t, n))
	}
	res = append(res, combineRecursion(n-1, k)...)
	return res
}
