/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (56.92%)
 * Likes:    3458
 * Dislikes: 206
 * Total Accepted:    408.8K
 * Total Submissions: 709.4K
 * Testcase Example:  '3'
 *
 * 
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 * 
 * 
 * 
 * For example, given n = 3, a solution set is:
 * 
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	return generate(0, 0, n, "", res)
}

// generate all possibles -> filter out invalid ones
// * how to check valid when processing: left 随时可以加，只要不超标；right 左个数>右个数
// * 剪枝
func generate(left, right, n int, s string, res []string) []string {
	// terminator
	if left == n && right == n {
		res = append(res, s)
		return res
	}

	// process current logic: left, right
	// drill down
	if left <= n {
		l := generate(left+1, right, n, s+"(", res)
		res = l
	}
	if left > right {
		r := generate(left, right+1, n, s+")", res)
		res = r
	}

	// reverse states
	return res
}
// @lc code=end

