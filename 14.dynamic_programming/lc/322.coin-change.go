/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (31.75%)
 * Likes:    2474
 * Dislikes: 89
 * Total Accepted:    270.2K
 * Total Submissions: 831.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * Example 1:
 *
 *
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 *
 */

// @lc code=start
/*
 * v1，暴力，递归，指数级
 * v2，BFS，状态树，第一次遇到叶子节点为0
 * v3，DP，f(n) = min(f(n-k), for k in [1, 2, 5]) + 1
 * 类比 0-1 背包问题：
 *	- 背包的容量 ~ amount
 *	- 可以放入背包的每种物品的价值 ~ coins
 *	- 每种 coin 的花费是 1
 *	- 最终问题转化为：填满背包的最小花费是多少
 * - F[i, v] = max{F[i − 1, v], F[i − 1, v − C i ] + W i }
 * 前 i 种方案，花费金额为 v，所能获得的加大价值 = max(不使用第i种方案获得价值，使用第i种方案获得的总价值)
 */
func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

