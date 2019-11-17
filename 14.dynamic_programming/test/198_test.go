package test

import (
	"fmt"
	"testing"
)

func Test_198(t *testing.T) {
	case1 := []int{1, 2, 3, 1}
	res := rob(case1)
	fmt.Printf("1.res: %v\n", res)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
	}
	a[0][0] = 0
	a[0][1] = nums[0]

	for i := 1; i < n; i++ {
		a[i][0] = max(a[i-1][0], a[i-1][1])
		a[i][1] = a[i-1][0] + nums[i]
	}
	return max(a[n-1][0], a[n-1][1])
}
