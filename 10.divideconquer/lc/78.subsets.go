/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (54.95%)
 * Likes:    2487
 * Dislikes: 60
 * Total Accepted:    435.7K
 * Total Submissions: 778.8K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠[3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
// v1, 可以转换成 左右括号组合 的思路，n 个格子，可选可不选
// 比如，将 [1,2,3] 变成 [1,0.0] 这样选或是不选的数组
func subsets(nums []int) [][]int {
	res := [][]int{}
	if nums == nil {
		return res
	}
	dfs(&res, nums, []int{}, 0)
	return res
}

func dfs(res *[][]int, nums, list []int, i int) { // i:第几层
	// terminator
	if i == len(nums) {
		*res = append(*res, list)
		return
	}
	// process current logic
	// drill down
	dfs(res, nums, list, i+1) // not pick the number at this index
	list = append(list, nums[i])
	// !! next dfs, pass a copy of list, instead of list itself
	dfs(res, nums, append([]int{}, list...), i+1) // pick the number at this index

	// restore state
	list = append([]int{}, list[:0]...)
}

// v2, 思路：依次将 nums 中的数字添加到已有的结果数组中
/*
 * 这里遇到的问题不是算法上的，而是 golang 语言本身对于 `append` 方法实现的理解上。
 * 对于 slice 的 `append` 后，返回的是一个新的 slice，同时也会更改底层的 array。
 *
 * Ref:
 * - https://stackoverflow.com/a/35276346/1594792
 * - https://blog.golang.org/slices
 */
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		newset := [][]int{}
		for _, subset := range res {
			// v2.1: explain step by step
			// newSubset := append([]int{}, subset...)
			// newSubset = append(newSubset, num)
			// newset = append(newset, newSubset)

			// v2.2: all in one, simple and ease
			res = append(res, append([]int{num}, subset...)) //!! `append` directly causes unexpected change of underlying array, so create a clean one
		}
		res = append(res, newset...)
	}
	return res
}

// @lc code=end

