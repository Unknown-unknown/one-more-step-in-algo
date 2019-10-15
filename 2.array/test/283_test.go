package test

import (
	"fmt"
	"testing"
)

func Test_283(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroesBruteForce(nums)
	fmt.Printf("brute force result: %v\n", nums)

	nums = []int{0, 1, 0, 3, 12}
	moveZeroesOptimal(nums)
	fmt.Printf("optimal result: %v\n", nums)
}

func moveZeroesBruteForce(nums []int) {
	res := []int{}
	countOfZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res = append(res, nums[i])
		} else {
			countOfZero++
		}
	}
	for j := 0; j < countOfZero; j++ {
		res = append(res, 0)
	}
	for k := 0; k < len(nums); k++ {
		nums[k] = res[k]
	}
}

func moveZeroesOptimal(nums []int) {
	for lastNonZeroFoundAt, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}
