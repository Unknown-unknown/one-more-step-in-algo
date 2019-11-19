package test

import (
	"math"
	"testing"
)

func Test_226(t *testing.T) {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// No.226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// No.98
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int64) bool {
	// terminator
	if root == nil {
		return true
	}

	// process current logic
	if int64(root.Val) <= min || int64(root.Val) >= max {
		return false
	}

	// drill down
	return isValid(root.Left, min, int64(root.Val)) && isValid(root.Right, int64(root.Val), max)
	// reverse states
}

// No.104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// No.111
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := minDepth(root.Left), minDepth(root.Right)
	if left == 0 || right == 0 {
		return max(left, right) + 1
	}
	return min(left, right) + 1
}

func min(x, y int) int {
	if x >= y {
		return y
	}
	return x
}

func max(x, y int) int {
	if y >= x {
		return y
	}
	return x
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// case1: both left and right are nil
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left, right := minDepth2(root.Left), minDepth2(root.Right)
	// case2: either of two is nil
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}
	// case3: neither is nil
	return min(left, right) + 1
}
