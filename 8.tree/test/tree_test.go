package test

import "testing"

func Test_Tree(t *testing.T) {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)

	return res
}

// func preorderTraversal(root *TreeNode) []int {

// }

// func postorderTraversal(root *TreeNode) []int {

// }
