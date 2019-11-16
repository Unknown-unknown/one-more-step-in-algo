/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (40.68%)
 * Likes:    1409
 * Dislikes: 164
 * Total Accepted:    208.4K
 * Total Submissions: 504.1K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */

// @lc code=start

// v1, brute-force，递归，n层：left or right，2^n，42/43，timeout
func minimumTotal(triangle [][]int) int {
	total := len(triangle)
	return recursion(0, 0, total, triangle)
}

func recursion(level, i, total int, triangle [][]int) int {
	if level == total-1 {
		return triangle[level][i]
	}
	left := recursion(level+1, i, total, triangle)
	right := recursion(level+1, i+1, total, triangle)
	return min(left, right) + triangle[level][i]
}

// v2, 记忆化搜索，自顶向下
func minimumTotal(triangle [][]int) int {
	total := len(triangle)
	cache := make(map[int]map[int]int)
	for i := 0; i < total; i++ {
		cache[i] = make(map[int]int)
	}
	return recursion(0, 0, total, cache, triangle)
}

func recursion(level, i, total int, cache map[int]map[int]int, triangle [][]int) int {
	if v, ok := cache[level][i]; ok {
		return v
	}
	if level == total-1 {
		cache[level][i] = triangle[level][i]
		return triangle[level][i]
	}
	left := recursion(level+1, i, total, cache, triangle)
	right := recursion(level+1, i+1, total, cache, triangle)
	cache[level][i] = min(left, right) + triangle[level][i]
	return min(left, right) + triangle[level][i]
}

// v3, DP: 存在最优子结构，考虑使用 DP，自底向上
// a. 重复性（分治）: problem(i,j) = min(sub(i+1, j), sub(i+1, j+1)) + a[i,j]
// b. 定义状态方程: f[i,j]
// c. DP方程: f[i,j] = min(f[i+1, j], f[i+1, j+1]) + a[i,j]
func minimumTotal(triangle [][]int) int {
	dp := triangle
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] += min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

