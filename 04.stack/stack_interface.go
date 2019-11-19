package stack

type Stack interface {
	Top() interface{}
	Empty() bool
	Size() int
	Push(v interface{})
	Pop()
}
