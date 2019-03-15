package queue

import "fmt"

type ArrayQueue struct {
	capacity int
	top      int
	tail     int
	data     []interface{}
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{capacity, 0, 0, make([]interface{}, 0, capacity)}
}

func (q *ArrayQueue) Front() interface{} {
	if q.top == 0 {
		return nil
	}
	return q.data[q.top]
}

func (q *ArrayQueue) Back() interface{} {
	if q.tail == 0 {
		return nil
	}
	return q.data[q.tail]
}

func (q *ArrayQueue) Empty() bool {
	return q.top == q.tail
}

func (q *ArrayQueue) Size() int {
	return q.tail - q.top
}

func (q *ArrayQueue) Push(v interface{}) {
	if q.tail+1 > q.capacity {
		return
	}
	q.data = append(q.data, v)
	q.tail++
}

func (q *ArrayQueue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	ele := q.data[q.top]
	q.data[q.top] = nil
	q.top++
	return ele
}

func (q *ArrayQueue) Print() {
	if q.Empty() {
		fmt.Println("empty queue")
		return
	}
	for i := 0; i < len(q.data); i++ {
		if q.data[i] != nil {
			fmt.Println(q.data[i])
		}
	}
}
