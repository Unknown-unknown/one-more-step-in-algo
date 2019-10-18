package test

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_641(t *testing.T) {
	obj := Constructor(4)
	param_1 := obj.InsertFront(1)
	param_2 := obj.InsertLast(2)
	param_3 := obj.DeleteFront()
	param_4 := obj.DeleteLast()
	fmt.Printf("%v, %v, %v, %v\n", param_1, param_2, param_3, param_4)
	param_5 := obj.GetFront()
	param_6 := obj.GetRear()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsFull()
	fmt.Printf("%v, %v, %v, %v\n", param_5, param_6, param_7, param_8)
}

type MyCircularDeque struct {
	data     *list.List
	capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{new(list.List), k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// already full
	if this.data.Len() >= this.capacity {
		return false
	}
	newNode := this.data.PushFront(value)
	if newNode != nil {
		return true
	}
	return false
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// already full
	if this.data.Len() >= this.capacity {
		return false
	}
	newNode := this.data.PushBack(value)
	if newNode != nil {
		return true
	}
	return false
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	front := this.data.Front()
	if front != nil {
		node := this.data.Remove(front)
		if node != nil {
			return true
		}
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	back := this.data.Back()
	if back != nil {
		node := this.data.Remove(back)
		if node != nil {
			return true
		}
	}
	return false
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	front := this.data.Front()
	if front != nil {
		return this.data.Front().Value.(int)
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	back := this.data.Back()
	if back != nil {
		return this.data.Back().Value.(int)
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.data.Len() == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.data.Len() == this.capacity
}
