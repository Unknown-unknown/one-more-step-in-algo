package test

import (
	"fmt"
	"testing"
)

func Test_78(t *testing.T) {
	tmp := []int{1, 2, 3}
	fmt.Println(tmp[:0])

	res := subsetsV1([]int{9, 0, 3, 5, 7})
	fmt.Printf("1.res: %v\n", res)

	res = subsetsV2([]int{9, 0, 3, 5, 7})
	fmt.Printf("2.res: %v\n", res)

	// res = subsetsV3([]int{9, 0, 3, 5, 7})
	// fmt.Printf("3.res: %v\n", res)
}

func subsetsV1(nums []int) [][]int {
	res := [][]int{}
	if nums == nil {
		return res
	}
	dfs(&res, nums, []int{}, 0)
	return res
}

func dfs(res *[][]int, nums []int, list []int, i int) { // i:第几层
	// terminator
	if i == len(nums) {
		*res = append(*res, list)
		return
	}
	// process current logic
	// drill down
	// fmt.Printf("1.nums: %v, list: %v, i: %d\n", nums, list, i)
	dfs(res, nums, list, i+1) // not pick the number at this index
	list = append(list, nums[i])
	// fmt.Printf("2.nums: %v, list: %v, i: %d\n", nums, list, i)
	dfs(res, nums, append([]int{}, list...), i+1) // pick the number at this index

	// restore state
	list = list[:0]
	// fmt.Printf("3.list: %v, i: %d\n", list, i)
}

func subsetsV2(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		newset := [][]int{}
		for _, subset := range res {
			newSubset := cp(append(subset, num)) //!! `append` directly causes unexpected change of underlying array
			newset = append(newset, newSubset)
		}
		res = append(res, newset...)
	}
	return res
}

func cp(src []int) (dst []int) {
	for _, ele := range src {
		dst = append(dst, ele)
	}
	return
}

func subsetsV3(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		// newset := [][]int{}
		for _, subset := range res {
			// newSubset := append([]int{}, subset...)
			// newSubset = append(newSubset, num)
			// newset = append(newset, newSubset)

			res = append(res, append([]int{num}, subset...))
		}
	}
	return res
}
