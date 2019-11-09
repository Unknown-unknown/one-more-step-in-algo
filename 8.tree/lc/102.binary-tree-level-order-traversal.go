/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (50.05%)
 * Likes:    1878
 * Dislikes: 47
 * Total Accepted:    452.1K
 * Total Submissions: 892.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
// v1, recursion
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	helper(root, 1, &res)
	return res
}

func helper(root *TreeNode, level int, res *[][]int) {
	// terminator
	if root == nil {
		return
	}

	// current level logic
	if len(*res) < level {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[level-1] = append((*res)[level-1], root.Val)
	}

	// drill down
	helper(root.Left, level+1, res)
	helper(root.Right, level+1, res)
	// reverse if needed
}

// v2, bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		currentLen := len(queue)
		for i := 0; i < currentLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if len(res) <= level {
				res = append(res, []int{node.Val})
			} else {
				res[level] = append(res[level], node.Val)
			}
		}
		level++
	}

	return res
}

// @lc code=end

