package test

import (
	"fmt"
	"testing"
)

func Test_189Rotate(t *testing.T) {
	cases1 := map[int][]int{
		4: []int{1, 2, 3},
		3: []int{1, 2, 3, 4, 5, 6, 7},
		2: []int{1, 2, 3, 4, 5},
		1: []int{1, 2},
		// 1: []int{1},
		0: []int{1, 2},
	}
	fmt.Println("v1============")
	for k, nums := range cases1 {
		rotate1(nums, k)
		fmt.Printf("rotated: %v\n", nums)
	}

	cases2 := map[int][]int{
		4: []int{1, 2, 3},
		3: []int{1, 2, 3, 4, 5, 6, 7},
		2: []int{1, 2, 3, 4, 5},
		1: []int{1, 2},
		// 1: []int{1},
		0: []int{1, 2},
	}
	fmt.Println("v2============")
	for k, nums := range cases2 {
		rotate2(nums, k)
		fmt.Printf("rotated: %v\n", nums)
	}
}

func rotate1(nums []int, k int) {
	if len(nums) <= 1 || k <= 0 {
		return
	}
	final := make([]int, len(nums))
	priv := len(nums) - k
	if priv < 0 {
		priv = len(nums) + priv
	}
	priv = priv % len(nums)
	left, right := nums[:priv], nums[priv:]
	final = append(right, left...)
	copy(nums, final)
}

func rotate2(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}
	length := len(nums)
	final := make([]int, length)
	for i := 0; i < length; i++ {
		final[(k+i)%length] = nums[i]
	}
	copy(nums, final)
}
