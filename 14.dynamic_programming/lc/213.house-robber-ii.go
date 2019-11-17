/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 *
 * https://leetcode.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (35.57%)
 * Likes:    1153
 * Dislikes: 39
 * Total Accepted:    137.4K
 * Total Submissions: 384.6K
 * Testcase Example:  '[2,3,2]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed. All houses at this place are
 * arranged in a circle. That means the first house is the neighbor of the last
 * one. Meanwhile, adjacent houses have security system connected and it will
 * automatically contact the police if two adjacent houses were broken into on
 * the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [2,3,2]
 * Output: 3
 * Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money
 * = 2),
 * because they are adjacent houses.
 *
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 */

// @lc code=start
// 由于第一家和最后一家相连，所以要么 0 <= i <=n-1，要么 1 <= i <=n，最后求二者最大值
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	dp1, dp2 := make([]int, n+1), make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i-2]) // 偷第一家
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i-1]) // 偷最后一家
	}
	return max(dp1[n], dp2[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

