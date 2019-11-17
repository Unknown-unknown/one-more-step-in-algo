/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 *
 * https://leetcode.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (41.30%)
 * Likes:    3251
 * Dislikes: 111
 * Total Accepted:    394.5K
 * Total Submissions: 951.8K
 * Testcase Example:  '[1,2,3,1]'
 *
 * You are a professional robber planning to rob houses along a street. Each
 * house has a certain amount of money stashed, the only constraint stopping
 * you from robbing each of them is that adjacent houses have security system
 * connected and it will automatically contact the police if two adjacent
 * houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of
 * each house, determine the maximum amount of money you can rob tonight
 * without alerting the police.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: 4
 * Explanation: Rob house 1 (money = 1) and then rob house 3 (money =
 * 3).
 * Total amount you can rob = 1 + 3 = 4.
 *
 * Example 2:
 *
 *
 * Input: [2,7,9,3,1]
 * Output: 12
 * Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house
 * 5 (money = 1).
 * Total amount you can rob = 2 + 9 + 1 = 12.
 *
 *
 */

// @lc code=start
/*
 * 一维DP 进阶为 二维DP：
 *
 * DP 方程 v1：
 a[i]: 0..i 能偷到 max value：a[n-1]
 a[i][0,1]: 0: i不偷，1:偷

 a[i][0] = max(a[i-1][0], a[i-1][1])
 a[i][1] = a[i-1][0] + nums[i]

 * 类似上台阶问题中引申的问题：如果连续两步不同，有多少种走法？
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
	}
	a[0][0] = 0
	a[0][1] = nums[0]

	for i := 1; i < n; i++ {
		a[i][0] = max(a[i-1][0], a[i-1][1])
		a[i][1] = a[i-1][0] + nums[i]
	}
	return max(a[n-1][0], a[n-1][1])
}

/*
 * DP 方程 v2：
 1.分解子问题：dp[i]: 0..i天，max value
 2.状态空间：dp[i]
 3.DP方程：dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 { // ! 这里注意防止只有一个元素时越界
		return nums[0]
	}

	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

