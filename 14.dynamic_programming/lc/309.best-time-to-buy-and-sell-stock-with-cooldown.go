/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (44.77%)
 * Likes:    1680
 * Dislikes: 63
 * Total Accepted:    110.4K
 * Total Submissions: 244.5K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (ie, buy one and sell one share of the stock
 * multiple times) with the following restrictions:
 *
 *
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1
 * day)
 *
 *
 * Example:
 *
 *
 * Input: [1,2,3,0,2]
 * Output: 3
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 *
 */

// @lc code=start
const MinInt = -int(^uint(0)>>1) - 1

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, MinInt
	dp_pre_0 := 0 // dp[i-2][0]
	for i := 0; i < n; i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = tmp
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

