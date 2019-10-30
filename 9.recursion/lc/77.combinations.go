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
	// terminator
	if k == n || k == 0 {
		res := []int{}
		for i := 1; i <= k; i++ {
			res = append(res, i)
		}
		return [][]int{res}
	}

	// 将从 n 个数字中选择 k 个数字，分解成：从 n - 1 个数字中选择 k - 1 个、再给每组结果加上 n，和从 n - 1 个数字中选择 k 个数字两种情况。
	// 类比爬楼梯问题，最终分解为：最后一步是一个台阶，和两个台阶两种情况。
	// 归根到底，还是在分解问题，找重复性，最后组合所有可能性。递归中，渗透着分治的思想，可见是九九归一了。
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
func combine(n int, k int) [][]int {
	if n == 0 || k == 0 {
		return [][]int{}
	}
	res := [][]int{}
	elems := []int{}
	backtracking(1, k, n, &elems, &res)
	return res
}

func backtracking(i, k, n int, elems *[]int, res *[][]int) {
	if k == 0 {
		// !! 注意：
		// 这里需要 copy 出一份新的 elems，append 到 res 中，
		// 否则之后 *elems 的改变也会相应改变 *res 中的数组元素的值，最终会得到一串儿相同的元素组成的数组的数组 ┑(￣Д ￣)┍
		tmp := append([]int{}, *elems...)
		*res = append(*res, tmp)
		return
	}
	for j := i; j < n+1; j++ {
		*elems = append(*elems, j) // 回
		fmt.Printf("回 %v\n", elems)
		backtracking(j+1, k-1, n, elems, res)
		*elems = (*elems)[:len(*elems)-1] // 溯
		fmt.Printf("溯: %v\n", elems)
	}
}

// @lc code=end

