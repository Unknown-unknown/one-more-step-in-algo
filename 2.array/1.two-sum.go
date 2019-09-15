/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
/************************************
- 用 golang 来实现有时候是会有一点别扭的，比如一些常用小函数是没有的，比如对应 js 中 array 的 includes
- 
*/
// v1, O(n2)
func twoSum(nums []int, target int) []int {
	indexes := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return indexes
}

// v2, O(n)
func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
	for i := 1; i <= len(nums); i++ {	// 因为 map 在 key 不存在时，value 为0。因此从1开始循环
		dict[nums[i-1]] = i
	}
	indexes := []int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if dict[complement] != 0 && dict[complement] - 1 != i {	// !! 注意此处的去重
			return []int{i, dict[complement] - 1}
		}
	}
	return indexes
}

// v3, O(n)
func twoSum(nums []int, target int) []int {
    dict := make(map[int]int)
    indexes := []int{}
	for i := 1; i <= len(nums); i++ {
        complement := target - nums[i-1]
        if dict[complement] != 0 {	// !! 在v2基础上改进，在循环过程中直接检查是否补数已经存在，因此无需考虑去重
			return []int{i-1, dict[complement]-1}
		}
		dict[nums[i-1]] = i
	}
	return indexes
}
