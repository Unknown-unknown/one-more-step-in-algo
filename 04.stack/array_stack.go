package stack

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, capacity),
		top:  -1,
	}
}

func (s *ArrayStack) Top() interface{} {
	if s.top < 0 {
		return nil
	}
	return s.data[s.top]
}

func (s *ArrayStack) Empty() bool {
	if s.top < 0 {
		return true
	} else {
		return false
	}
}

func (s *ArrayStack) Size() int {
	if s.top < 0 {
		return 0
	} else {
		size := 0
		for _, v := range s.data {
			if v != nil {
				size++
			}
		}
		return size
	}
}

func (s *ArrayStack) Push(v interface{}) {
	s.data = append(s.data, v)
	s.top++
}

func (s *ArrayStack) Pop() {
	if s.top < 0 {
		return
	}
	s.data[s.top] = nil
	s.top--
}

func (s *ArrayStack) Print() {
	if s.top < 0 {
		fmt.Println("empty stack")
	}
	for i := len(s.data) - 1; i >= 0; i-- {
		if s.data[i] != nil {
			fmt.Println(s.data[i])
		}
	}
}
