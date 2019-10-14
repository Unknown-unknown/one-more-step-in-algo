package test

import (
	"fmt"
	"testing"
)

func Test_239SlidingWindow(t *testing.T) {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	res1 := maxSlidingWindowBruteForce(nums1, k1)
	fmt.Printf("1.brute force maxes: %v\n", res1)
	res1 = maxSlidingWindowDeque(nums1, k1)
	fmt.Printf("1.deque maxes: %v\n", res1)

	nums2 := []int{1, -1}
	k2 := 1
	res2 := maxSlidingWindowBruteForce(nums2, k2)
	fmt.Printf("2.brute force maxes: %v\n", res2)

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

func maxSlidingWindowDeque(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k)
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
