/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (44.50%)
 * Likes:    5094
 * Dislikes: 192
 * Total Accepted:    628.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
/*
 * v1, 暴力，O(n^2)
 * v2, DP
 * 	a. 分治（子问题） max_sum(i) = max(max_sum(i-1), 0) + a[i]
 *  b. 状态数组定义：f[i]
 *	c. DP 方程：f[i] = max(f[i-1], 0) + a[i]
 *	最大子序和 = 当前元素自身最大，或 包含之前后最大
 */
func maxSubArray(nums []int) int {
	dp := nums
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	return max(dp...)
}

func max(arr ...int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// @lc code=end

