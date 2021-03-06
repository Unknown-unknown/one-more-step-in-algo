/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (41.27%)
 * Likes:    1235
 * Dislikes: 57
 * Total Accepted:    164.6K
 * Total Submissions: 391K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1: [1,3,0,2]
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2: [2,0,3,1]
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

// @lc code=start
// ![backtracking]
func solveNQueens(n int) [][]string {
	res := [][]string{} // res 用来存放每种解法中 Queen 的列位置的数组
	shu, pie, na := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	DFS(n, []int{}, shu, pie, na, &res)
	return res
}

func DFS(n int, rows []int, shu map[int]bool, pie map[int]bool, na map[int]bool, res *[][]string) {
	row := len(rows)
	if row == n {
		sol := print(rows, n)
		*res = append(*res, sol)
		return
	}

	for col := 0; col < n; col++ {
		if !shu[col] && !pie[row-col] && !na[row+col] {
			shu[col] = true
			pie[row-col] = true
			na[row+col] = true
			DFS(n, append(rows, col), shu, pie, na, res)
			shu[col] = false
			pie[row-col] = false
			na[row+col] = false
		}
	}
}

func print(cols []int, n int) []string {
	rows := []string{}
	for j := 0; j < len(cols); j++ { // 每种解法中每行 Queen 位于哪列
		row := ""
		for k := 0; k < n; k++ { // 每种解法的每一行
			if k == cols[j] {
				row += "Q"
			} else {
				row += "."
			}
		}
		rows = append(rows, row)
	}
	return rows
}

// @lc code=end

