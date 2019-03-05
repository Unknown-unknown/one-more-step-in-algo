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
