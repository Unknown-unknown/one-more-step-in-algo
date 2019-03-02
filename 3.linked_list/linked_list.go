package linked_list

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length int
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

func (l *LinkedList) InsertAfter(n *Node, v interface{}) *Node {
	if n == nil {
		return nil
	}
	newNode := NewNode(v)
	newNode.next = n.next
	n.next = newNode
	l.length++
	return newNode
}

func (l *LinkedList) InsertBefore(n *Node, v interface{}) *Node {
	if n == nil || n == l.head {
		return n
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current == n {
			break
		} else {
			prev = current
			current = current.next
		}
	}
	if current == nil {
		fmt.Println("node not found")
		return nil
	}
	newNode := NewNode(v)
	prev.next = newNode
	newNode.next = current
	l.length++
	return n
}

func (l *LinkedList) PushBack(v interface{}) *Node {
	prev := l.head
	current := l.head.next
	for current != nil {
		prev = current
		current = current.next
	}
	return l.InsertAfter(prev, v)
}

func (l *LinkedList) PushFront(v interface{}) *Node {
	return l.InsertAfter(l.head, v)
}

func (l *LinkedList) Remove(n *Node) {
	if n == nil || l.head == n {
		return
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current == n {
			prev.next = current.next
			current.next = nil
			break
		} else {
			prev = current
			current = current.next
		}
	}
	if current == nil {
		fmt.Printf("node to remove not found")
		return
	}
	l.length--
}

func (l *LinkedList) Print() {
	if l == nil || l.length == 0 {
		fmt.Printf("empty list")
		return
	}
	current := l.head.next
	desc := ""
	for current != nil {
		desc += fmt.Sprintf("%+v", current.value)
		current = current.next
		if current != nil {
			desc += " -> "
		}
	}

	fmt.Println(desc)
}
