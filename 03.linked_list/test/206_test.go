package test

import (
	"fmt"
	"testing"
)

func Test_206(t *testing.T) {
	head := ListNode{0, nil}
	first := ListNode{1, nil}
	second := ListNode{2, nil}
	head.Next = &first
	first.Next = &second
	reversed := reverseList(&head)
	fmt.Printf("reversed: %v", reversed)
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	current := head
	var prev, next *ListNode
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
