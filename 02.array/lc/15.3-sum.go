/*
/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (24.72%)
 * Likes:    4462
 * Dislikes: 510
 * Total Accepted:    639K
 * Total Submissions: 2.6M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate triplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */

 // v1, 暴力求解，时间复杂度 O(n^3)（未去重）
 func threeSum(nums []int) [][]int {
	 res := [][]int{}
	 for i := 0; i < len(nums) - 2; i++ {
		 for j := i+1; j < len(nums) - 1; j++ {
			 for k := j+1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			 }
		 }		 
	 }
	 return res
 }

// v2, 双指针，先排序，然后左右下标往中间推进
// ![two-pointer]
import "sort"
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// * 时间复杂度：O(n*log(n))，参 go/1.12.6/libexec/src/sort/sort.go 
	// * golang 对于 Sort() 的实现值得一看，特别是不同条件下对于不同排序方法的选择
	// ![source-code]
	// ![sort]
	sort.Ints(nums)	
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return [][]int{}
	}
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {	// !!去重
			continue
		}
		l := i+1
		r := len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1]	{	// !!去重
					l++
				}
				for l < r && nums[r] == nums[r-1] {	// !!去重
					r--
				}
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// v3, 哈希表实现？？？ // ![todo]
