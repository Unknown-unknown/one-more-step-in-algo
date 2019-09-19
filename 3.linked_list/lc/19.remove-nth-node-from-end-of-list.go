/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.59%)
 * Likes:    2136
 * Dislikes: 160
 * Total Accepted:    450.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 * 
 * Example:
 * 
 * 
 * Given linked list: 1->2->3->4->5, and n = 2.
 * 
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 * 
 * 
 * Note:
 * 
 * Given n will always be valid.
 * 
 * Follow up:
 * 
 * Could you do this in one pass?
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ![two-pointer]
/* 
* 这里的 pitfall 是，如果直接用 head 作为备份的节点，在 linked list 为 1，n = 1 时，fast.Next 就会成为野指针。
* 所以，还是加一个空节点，Next 指向 head，这样就不占用原链表的长度了。这算不算是，哨兵/带头链表的应用？🤔
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    priv := &ListNode{Next: head,}
    slow, fast, i := priv, priv, 1
	for i <= n {
		fast = fast.Next
		i++
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return priv.Next
}

