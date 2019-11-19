/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (29.68%)
 * Likes:    2041
 * Dislikes: 644
 * Total Accepted:    241K
 * Total Submissions: 811.5K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,0]
 * Output: 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,4,-1,1]
 * Output: 2
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [7,8,9,11,12]
 * Output: 1
 * 
 * 
 * Note:
 * 
 * Your algorithm should run in O(n) time and uses constant extra space.
 * 
 */

/*
* 考虑去重，比如：[0,1,1,2,2]
*/

// v1, memory usage: 2.0M
import "sort"
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)		// !!虽然提交通过了，但是怀疑这里的复杂度并不是 O(n)...

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count ++
			if i>=1 && nums[i] == nums[i-1] {
				count --
			}
			if nums[i] == count {
				if i == len(nums) - 1 {
					return nums[i] + 1
				}
				continue
			} else {
				return count
			}
		}
	}
	return count+1
}

// v2, use hashmap, memory usage: 2.4M
func firstMissingPositive(nums []int) int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = i
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := dict[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}

// v3, 
// * hasn't understood it yet ┑(￣Д ￣)┍
func firstMissingPositive(nums []int) int {
    n := len(nums)
    i := 0
    for i < n {
        if nums[i] > 0 && nums[i] <= n && nums[i] != i + 1 && nums[nums[i] -1] != nums[i] {
            nums[i], nums[nums[i] -1] = nums[nums[i] -1], nums[i]
        } else {
            i++
        }
    }
    
    for i := 0; i < n; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    
    return n + 1
}
