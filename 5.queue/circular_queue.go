package queue

import "fmt"

type CircularQueue struct {
	data     []interface{}
	capacity int
	top      int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

func (q *CircularQueue) IsEmpty() bool {
	if q.top == q.tail {
		return true
	}
	return false
}

func (q *CircularQueue) IsFull() bool {
	if (q.tail+1)%q.capacity == q.top {
		return true
	}
	return false
}

func (q *CircularQueue) Enqueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return false
	}
	v := q.data[q.top]
	q.top = (q.top + 1) % q.capacity
	return v
}

func (q *CircularQueue) Print() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	res := "head"
	i := q.top
	for {
		res += fmt.Sprintf("<- %v", q.data[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	res += "<-tail"
	return res
}
