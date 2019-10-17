package test

import "testing"

func Test_141(t *testing.T) {
}

func hasCycle141(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle142(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != fast {
				head, fast = head.Next, fast.Next
			}
			return head
		}
	}
	return nil
}
