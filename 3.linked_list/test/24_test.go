package test

import (
	"fmt"
	"testing"
)

func Test_24(t *testing.T) {
	head := ListNode{1, nil}
	first := ListNode{2, nil}
	second := ListNode{3, nil}
	third := ListNode{4, nil}
	fourth := ListNode{5, nil}
	head.Next = &first
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	swapped := swapPairsIterative(&head)
	fmt.Printf("swapped: %v", swapped)
}

/* prev->1->2->3->4
=> prev->2->1->3->4
=> prev->2->1->4->3
*/
func swapPairsIterative(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{0, head}
	cur := prev
	for cur.Next != nil && cur.Next.Next != nil {
		start := cur.Next
		end := cur.Next.Next
		cur.Next = end
		start.Next = end.Next
		end.Next = start
		cur = start
	}
	return prev.Next
}
