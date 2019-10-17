/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (46.24%)
 * Likes:    1452
 * Dislikes: 129
 * Total Accepted:    365.4K
 * Total Submissions: 782K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 * 
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// v1, iterative，添加头部哨兵。对于一次检查，需要更新的就是三个 Next 指针。
// 更新的顺序是先把后面的 node 都勾连上，在返回来勾连前面的。
/* prev[->]1[->]2[->]3->4
=> prev->2->1->3->4
=> prev->2->1->4->3
*/
// ![sentry]
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{0, head}
	current := pre
	for current.Next != nil && current.Next.Next != nil {
		start := current.Next
		end := current.Next.Next
		current.Next = end
		start.Next = end.Next
		end.Next = start
		current = start
	}
	return pre.Next
}

// v2, recursion
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	next := head.Next
// }
// @lc code=end

