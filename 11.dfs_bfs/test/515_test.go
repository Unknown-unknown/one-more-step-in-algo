package test

import "testing"

func Test_515(t *testing.T) {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
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
