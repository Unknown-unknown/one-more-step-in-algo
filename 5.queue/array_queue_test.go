package queue

import (
	"log"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	log.Print(q.Empty())
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Print()
	log.Print(q.Pop())
	log.Print(q.Pop())
	log.Print(q.Pop())
	log.Print(q.Pop())
	log.Print(q.Size())
	q.Print()
}
