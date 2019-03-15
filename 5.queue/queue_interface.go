package queue

type Queue interface {
	Front() interface{}
	Back() interface{}
	Empty() bool
	Size() int
	Push(v interface{})
	Pop()
}
