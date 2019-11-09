/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 *
 * https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (58.48%)
 * Likes:    589
 * Dislikes: 48
 * Total Accepted:    74.5K
 * Total Submissions: 126.8K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * You need to find the largest value in each row of a binary tree.
 *
 * Example:
 *
 * Input:
 *
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \
 * ⁠     5   3   9
 *
 * Output: [1, 3, 9]
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// v1, bfs
const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0, 100)
	queue := make([]*TreeNode, 0, 100) // !! 可以尽量给一个比较大的 capacity，可以避免 array 容量不足时的扩容和拷贝，以此缩减 runtime 时间
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		largest := MinInt
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]

			// process current
			if cur.Val > largest {
				largest = cur.Val
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, largest)
	}
	return res
}

// @lc code=end

