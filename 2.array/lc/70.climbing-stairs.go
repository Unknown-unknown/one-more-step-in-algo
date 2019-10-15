/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (45.03%)
 * Likes:    2752
 * Dislikes: 97
 * Total Accepted:    488.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 * 
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 * 
 * Note: Given n will be a positive integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 * 
 * 
 */

// 懵逼状态：暴力？基本情况？
// !!找 最近重复子问题 
// f(3) = f(1) + f(2)
// f(4) = f(3) + f(2)
// @lc code=start

// v1, 暴力破解
func climbStairs(n int) int {
    if n == 1 || n == 2 {
		return n
	}

	return climbStairs(n - 1) + climbStairs(n-2)
}
q
// v2, 斐波那契数, 时间复杂度 O(n)，空间复杂度 O(1)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i < n+1; i++ {
		f3 := f1 + f2
		f1 = f2
		f2 = f3
	}
	return f2
}


// @lc code=end

