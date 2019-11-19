package stack

import (
	"log"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack(5)
	log.Print(s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print()
	s.Pop()
	s.Pop()
	log.Print(s.Size())
	s.Print()
}
