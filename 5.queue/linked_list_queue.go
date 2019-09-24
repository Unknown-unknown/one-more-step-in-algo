package queue

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LLQueue struct {
	top    *Node
	tail   *Node
	length int
}

func NewLinkedListQueue() *LLQueue {
	return &LLQueue{nil, nil, 0}
}

func (q *LLQueue) Enqueue(v interface{}) {
	node := &Node{nil, v}
	if q.tail == nil {
		q.top = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *LLQueue) Dequeue() interface{} {
	if q.top == nil {
		return nil
	}
	v := q.top.value
	q.top = q.top.next
	q.length--
	return v
}

func (q *LLQueue) Print() string {
	if q.top == nil {
		return "empty queue"
	}
	res := "head"
	for cur := q.top; cur != nil; cur = cur.next {
		res += fmt.Sprintf("<- %v", cur.value)
	}
	res += "<-tail"
	return res
}
