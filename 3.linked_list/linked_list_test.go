package linked_list

import (
	"log"
	"testing"
)

func TestInsertAfter(t *testing.T) {
	l := NewLinkedList()
	n1 := l.InsertAfter(l.head, 1)
	n2 := l.InsertAfter(n1, 2)
	l.InsertAfter(n2, 3)
	l.InsertAfter(nil, 4)
	l.Print()
}

func TestInserBefore(t *testing.T) {
	l := NewLinkedList()
	n1 := l.InsertAfter(l.head, 1)
	for i := 0; i < 5; i++ {
		l.InsertBefore(n1, i+1)
	}
	l.Print()

	var n *Node = NewNode(6)
	l.InsertBefore(n, 999)
	l.Print()
}

func TestPushFront(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.PushFront(i + 1)
	}
	l.Print()
}

func TestPushBack(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.PushBack(i + 1)
	}
	l.Print()
}

func TestRemoveNormal(t *testing.T) {
	l := NewLinkedList()
	var n *Node
	for i := 0; i < 10; i++ {
		v := l.PushBack(i + 1)
		if i == 6 {
			n = v
		}
	}
	l.Print()

	l.Remove(n)
	l.Print()
}

func TestRemoveNotFound(t *testing.T) {
	l := NewLinkedList()
	var n *Node = NewNode(6)
	for i := 0; i < 10; i++ {
		l.PushBack(i + 1)
	}
	l.Print()

	l.Remove(n)
	l.Print()
}

func TestRemoveBetter(t *testing.T) {
	l := NewLinkedList()
	var n *Node
	for i := 0; i < 10; i++ {
		v := l.PushBack(i + 1)
		if i == 6 {
			n = v
		}
	}
	l.Print()
	log.Printf("%d to be removed", n.GetValue())
	l.RemoveBetter(n)
	l.Print()
}
