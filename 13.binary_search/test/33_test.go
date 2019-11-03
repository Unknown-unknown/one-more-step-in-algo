package test

import (
	"fmt"
	"testing"
)

func Test_33(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	res := searchIndex(nums)
	fmt.Printf("res: %v", res)
}

func searchIndex(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
