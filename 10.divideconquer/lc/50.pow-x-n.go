/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.53%)
 * Likes:    1011
 * Dislikes: 2431
 * Total Accepted:    373.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 * 
 * Example 1:
 * 
 * 
 * Input: 2.00000, 10
 * Output: 1024.00000
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 2.10000, 3
 * Output: 9.26100
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * Note:
 * 
 * 
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 * 
 * 
 */

// @lc code=start
// v1, 暴力, O(n)
// v2, 分治, O(logn)：
/* 
1.terminator 
2.process (split your big problem)
3.drill down(subproblems), merge(subresult)
4.revert states
*/
func myPow(x float64, n int) float64 {
	// terminator
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	// process current level
	// drill down
	res := myPow(x, n/2)
	// merge
	if n%2 != 0 {
		res = res * res * x
	} else {
		res = res * res
	}
	return res
}

// v3, 牛顿迭代法，时间复杂度 O(logn)，空间复杂度 O(1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	subRes := x
	res := 1.0
	for i := n; i > 0; i = i / 2 {
		if i%2 == 1 {
			res = subRes * res
		}
		subRes = subRes * subRes
	}
	return res
}
// @lc code=end

