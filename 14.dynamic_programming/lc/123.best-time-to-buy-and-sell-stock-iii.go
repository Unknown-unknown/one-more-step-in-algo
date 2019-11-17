/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/description/
 *
 * algorithms
 * Hard (34.74%)
 * Likes:    1462
 * Dislikes: 59
 * Total Accepted:    175.3K
 * Total Submissions: 496.3K
 * Testcase Example:  '[3,3,5,0,0,3,1,4]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete at most two
 * transactions.
 *
 * Note: You may not engage in multiple transactions at the same time (i.e.,
 * you must sell the stock before you buy again).
 *
 * Example 1:
 *
 *
 * Input: [3,3,5,0,0,3,1,4]
 * Output: 6
 * Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit
 * = 3-0 = 3.
 * Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 =
 * 3.
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit
 * = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you
 * are
 * engaging multiple transactions at the same time. You must sell before buying
 * again.
 *
 *
 * Example 3:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 */

// @lc code=start
/*
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
*/
const MinInt = -int(^uint(0)>>1) - 1

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp_i10, dp_i11 := 0, MinInt
	dp_i20, dp_i21 := 0, MinInt
	for i := 0; i < n; i++ {
		dp_i20 = max(dp_i20, dp_i21+prices[i])
		dp_i21 = max(dp_i21, dp_i10-prices[i])
		dp_i10 = max(dp_i10, dp_i11+prices[i])
		dp_i11 = max(dp_i11, -prices[i])
	}
	return dp_i20
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

