/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (54.04%)
 * Likes:    1956
 * Dislikes: 169
 * Total Accepted:    432.8K
 * Total Submissions: 800.2K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 * 
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,2,3]
 * Output: 3
 * 
 * Example 2:
 * 
 * 
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 * 
 * 
 */
// v1, O(n)
func majorityElement(nums []int) int {
	dict := make(map[int]int)	// key: num, value: count
	for i := 1; i <= len(nums); i++ {
		dict[nums[i-1]] = dict[nums[i-1]]+1		
	}
	for k, v := range dict {
		if v > len(nums) / 2 {
			return k
		}
	}
	return 0
}

// v2
func majorityElement(nums []int) int {
	dict := make(map[int]int)	// key: num, value: count
	for i := 1; i <= len(nums); i++ {
		dict[nums[i-1]] = dict[nums[i-1]]+1
		if dict[nums[i-1]] > len(nums) / 2 {
			return nums[i-1]
		}
	}
	return 0
}

// v3
// ** 看到此法，震惊……虽然，在golang中并没有提高太多效率
// ** More: https://leetcode-cn.com/problems/majority-element/solution/qiu-zhong-shu-by-leetcode-2/
import "sort"
func majorityElement(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

