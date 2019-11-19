package stack

import (
	"log"
	"testing"
)

func TestLLStack(t *testing.T) {
	s := NewLLStack()
	log.Print(s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print()
	log.Print(s.Top())
	log.Print(s.Pop())
	log.Print(s.Pop())
	log.Print(s.Size())
	s.Print()
}
