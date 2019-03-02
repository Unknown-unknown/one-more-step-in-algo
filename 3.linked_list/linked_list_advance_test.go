package linked_list

import (
	"fmt"
	"testing"
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
func TestIsRing(t *testing.T) {
	l2 := NewLinkedList()
	n1 := l2.InsertAfter(l2.head, 1)
	n2 := l2.InsertAfter(n1, 2)
	n3 := l2.InsertAfter(n2, 3)
	n4 := l2.InsertAfter(n3, 4)
	n4.next = n1
	fmt.Println(l2.IsRing())
}
