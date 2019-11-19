package queue

import "fmt"

type ArrayQueue struct {
	capacity int
	top      int
	tail     int
	data     []interface{}
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{capacity, 0, 0, make([]interface{}, capacity)}
}

func (q *ArrayQueue) Enqueue(v interface{}) bool {
	if q.tail == len(q.data) {
		return false
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.top == q.tail {
		return nil
	}
	v := q.data[q.top]
	q.top++
	return v
}

func (q *ArrayQueue) Print() string {
	if q.top == q.tail {
		return "empty queue"
	}
	res := "head"
	for i := q.top; i < q.tail; i++ {
		res += fmt.Sprintf("<- %v", q.data[i])
	}
	res += "<-tail"
	return res
}
