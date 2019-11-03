/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (32.06%)
 * Likes:    934
 * Dislikes: 1572
 * Total Accepted:    435.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

// @lc code=start
// v1, 二分查找
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// v2, 牛顿迭代法
// 来源及 Magic Number：https://www.beyond3d.com/content/articles/8/
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// @lc code=end

