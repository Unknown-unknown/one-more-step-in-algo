/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (49.66%)
 * Likes:    976
 * Dislikes: 56
 * Total Accepted:    233.9K
 * Total Submissions: 464.9K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 * 
 * Example:
 * 
 * 
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 * 
 */

// @lc code=start
// v1, recursion
func combine(n int, k int) [][]int {
	if k == n || k == 0 {
		res := []int{}
		for i := 1; i <= k; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}
	tmp := combine(n-1, k-1)
	res := [][]int{}
	for _, t := range tmp {
		res = append(res, append(t, n))
	}
	res = append(res, combine(n-1, k)...)
	return res
}

// v2, backtracking
// ![backtracking]
// todo:
// @lc code=end

