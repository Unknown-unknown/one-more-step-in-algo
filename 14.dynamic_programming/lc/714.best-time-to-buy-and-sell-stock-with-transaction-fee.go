/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 *
 * algorithms
 * Medium (51.53%)
 * Likes:    1122
 * Dislikes: 33
 * Total Accepted:    50.8K
 * Total Submissions: 97.1K
 * Testcase Example:  '[1,3,2,8,4,9]\n2'
 *
 * Your are given an array of integers prices, for which the i-th element is
 * the price of a given stock on day i; and a non-negative integer fee
 * representing a transaction fee.
 * You may complete as many transactions as you like, but you need to pay the
 * transaction fee for each transaction.  You may not buy more than 1 share of
 * a stock at a time (ie. you must sell the stock share before you buy again.)
 * Return the maximum profit you can make.
 *
 * Example 1:
 *
 * Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
 * Output: 8
 * Explanation: The maximum profit can be achieved by:
 * Buying at prices[0] = 1Selling at prices[3] = 8Buying at prices[4] =
 * 4Selling at prices[5] = 9The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) =
 * 8.
 *
 *
 *
 * Note:
 * 0 < prices.length .
 * 0 < prices[i] < 50000.
 * 0 .
 *
 */

// @lc code=start
const MinInt = -int(^uint(0)>>1) - 1

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp_i_0, dp_i_1 := 0, MinInt
	for i := 0; i < n; i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, tmp-prices[i]-fee)
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

