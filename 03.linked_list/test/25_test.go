package test

import "testing"

func Test_25(t *testing.T) {
}

// todo: um....
func reverseKGroup(head *ListNode, k int) *ListNode {
	var next = head
	for i := 0; i < k; i++ {
		if next != nil {
			next = next.Next
		} else {
			return head
		}
	}

	var newHead *ListNode
	originHead := head
	nextElem := head.Next
	for head.Next != next {
		nextElem, head.Next = head.Next, newHead
		newHead, head = head, nextElem
	}
	head.Next = newHead
	originHead.Next = reverseKGroup(next, k)

	return head
}
