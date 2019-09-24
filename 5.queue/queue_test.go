package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	res := q.Enqueue(1)
	assert.True(t, res)
	res = q.Enqueue(2)
	assert.True(t, res)
	res = q.Enqueue(3)
	assert.True(t, res)
	res = q.Enqueue(4)
	assert.True(t, res)
	res = q.Enqueue(5)
	assert.True(t, res)
	res = q.Enqueue(6)
	assert.False(t, res)
	fmt.Println(q.Print())
	v := q.Dequeue()
	assert.True(t, v == 1)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	v = q.Dequeue()
	assert.True(t, v == 5)
	v = q.Dequeue()
	assert.Nil(t, v)
	fmt.Println(q.Print())
}

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	fmt.Println(q.Print())
	q.Enqueue(2)
	fmt.Println(q.Print())
	q.Enqueue(3)
	fmt.Println(q.Print())
	q.Enqueue(4)
	fmt.Println(q.Print())
	q.Enqueue(5)
	fmt.Println(q.Print())
	v := q.Dequeue()
	assert.True(t, v == 1)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	v = q.Dequeue()
	assert.True(t, v == 5)
	v = q.Dequeue()
	assert.Nil(t, v)
	fmt.Println(q.Print())
}

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	res := q.Enqueue(1)
	assert.True(t, res)
	res = q.Enqueue(2)
	assert.True(t, res)
	res = q.Enqueue(3)
	assert.True(t, res)
	res = q.Enqueue(4)
	assert.True(t, res)
	res = q.Enqueue(5)
	assert.False(t, res)
	fmt.Println(q.Print())
	fmt.Println(q.IsFull())
	v := q.Dequeue()
	assert.True(t, v == 1)
	q.Dequeue()
	q.Dequeue()
	v = q.Dequeue()
	assert.True(t, v == 4)
	fmt.Println(q.Print())
}
