package linked_list

import (
	"testing"

	"gotest.tools/assert"
)

var l *LinkedList

func init() {
	l = NewLinkedList()
	for i := 0; i < 10; i++ {
		l.PushBack(i + 1)
	}
}

func TestReversed(t *testing.T) {
	l.Print()
	l.Reversed()
	l.Print()
}

// head -> 1 -> 2 -> 3 -> 4 -> 1
func TestHasRing(t *testing.T) {
	assert.Assert(t, !l.HasRing())

	l2 := NewLinkedList()
	n1 := l2.InsertAfter(l2.head, 1)
	n2 := l2.InsertAfter(n1, 2)
	n3 := l2.InsertAfter(n2, 3)
	n4 := l2.InsertAfter(n3, 4)
	n4.next = n1
	// impossible to print a linked list with a ring, since it will loop over and over until stack over flow :P
	assert.Assert(t, l2.HasRing())
}

// ll: 1 -> 3 -> 5 -> 7 -> 9
// lr: 2 -> 4 -> 5 -> 7
func TestMergeSortedList(t *testing.T) {
	ll := NewLinkedList()
	ll.PushBack(1)
	ll.PushBack(3)
	ll.PushBack(5)
	ll.PushBack(7)
	ll.PushBack(9)
	ll.Print()

	lr := NewLinkedList()
	lr.PushBack(2)
	lr.PushBack(4)
	lr.PushBack(5)
	lr.PushBack(7)
	lr.Print()

	merged := MergeSortedList(ll, lr)
	merged.Print()
}

func TestRemoveBackN(t *testing.T) {
	l.RemoveBackN(3)
	l.Print()

	l.RemoveBackN(11)
}

func TestFindMiddleNode(t *testing.T) {
	// odd number of nodes
	l2 := NewLinkedList()
	l2.PushBack(1)
	l2.PushBack(3)
	l2.PushBack(5)
	l2.PushBack(7)
	l2.PushBack(9)
	l2.Print()
	m2 := l2.FindMiddleNode()
	assert.Assert(t, len(m2) == 1, "middle number count fail for odd")
	assert.Assert(t, m2[0].GetValue() == 5, "fail to find for odd")

	// even number of nodes
	m1 := l.FindMiddleNode()
	assert.Assert(t, len(m1) == 2, "middle number count fail for even")
	assert.Assert(t, m1[0].GetValue() == 4, "fail to find for even")
	assert.Assert(t, m1[1].GetValue() == 5, "fail to find for even")
}
