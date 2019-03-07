package stack

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LLStack struct {
	top *Node
}

func NewLLStack() *LLStack {
	return &LLStack{nil}
}

func (s *LLStack) Top() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

func (s *LLStack) Empty() bool {
	return s.top == nil
}

func (s *LLStack) Size() int {
	size := 0
	// !! attention: if we loop through the stack by directly changing its top pointer,
	// we'll never get the right stack after calling this func outside.
	// So we make a copy of top pointer instead. So is `Print()` below.
	tmp := s.top
	for tmp != nil {
		tmp = tmp.next
		size++
	}
	return size
}

func (s *LLStack) Push(v interface{}) {
	s.top = &Node{s.top, v}
}

func (s *LLStack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	s.top = s.top.next
	return s.top.value
}

func (s *LLStack) Print() {
	if s.top == nil {
		fmt.Println("empty stack")
		return
	}
	tmp := s.top
	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
}
