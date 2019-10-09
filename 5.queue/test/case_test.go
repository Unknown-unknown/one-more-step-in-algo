package test

import (
	"fmt"
	"testing"
)

func Test_239SlidingWindow(t *testing.T) {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	res1 := maxSlidingWindowBruteForce(nums1, k1)
	fmt.Printf("1.maxes: %v\n", res1)

	nums2 := []int{1, -1}
	k2 := 1
	res2 := maxSlidingWindowBruteForce(nums2, k2)
	fmt.Printf("2.maxes: %v\n", res2)

}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func maxSlidingWindowBruteForce(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	maxes := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		max := MinInt
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		maxes = append(maxes, max)
	}
	return maxes
}
