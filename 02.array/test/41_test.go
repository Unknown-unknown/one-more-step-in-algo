package test

import (
	"fmt"
	"testing"
)

func Test_41firstMissingPositive(t *testing.T) {
	nums := []int{2, 0, -1, 2}
	missing1 := firstMissingPositive(nums)
	fmt.Println(missing1)

	nums = []int{0, 1, 1, 2, 2}
	missing2 := firstMissingPositive(nums)
	fmt.Println(missing2)
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	fmt.Printf("original : %v\n", nums)
	for i < n {
		if nums[i] > 0 && nums[i] <= n && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			fmt.Printf("%d <--> %d\n", nums[i], nums[nums[i]-1])
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	fmt.Printf("after 1st loop: %v\n", nums)
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
