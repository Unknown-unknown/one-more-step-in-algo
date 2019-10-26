package test

import (
	"fmt"
	"testing"
)

func Test_Tree(t *testing.T) {
	root := TreeNode{0, nil, nil}
	left := TreeNode{1, nil, nil}
	right := TreeNode{2, nil, nil}
	left1 := TreeNode{3, nil, nil}
	right1 := TreeNode{4, nil, nil}
	left2 := TreeNode{5, nil, nil}
	left3 := TreeNode{6, nil, nil}
	root.Left = &left
	root.Right = &right
	root.Left.Left = &left1
	root.Left.Right = &right1
	root.Right.Left = &left2
	root.Left.Left.Left = &left3

	// !! 拒绝人肉递归，大脑容易死机……┑(￣Д ￣)┍
	res := levelOrderBFS(&root)
	fmt.Printf("level order of bfs: %v\n", res)

	fmt.Println("-----------------------")

	res = levelOrderRecursion(&root)
	fmt.Printf("level order of recursion: %v\n", res)
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

// No.102 level traversal of bfs
func levelOrderBFS(root *TreeNode) [][]int {
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
			fmt.Printf("i = %d, currentLen = %d, queue = %v, res = %v\n", i, currentLen, queue, res)
		}
		level++
	}

	return res
}

// No.102 level traversal of recursion
func levelOrderRecursion(root *TreeNode) [][]int {
	res := [][]int{}
	helper(root, 1, &res)
	return res
}

func helper(root *TreeNode, level int, res *[][]int) {
	fmt.Printf("helper called: level = %d, res = %v\n", level, res)
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
