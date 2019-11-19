package test

import "testing"

func Test_155(t *testing.T) {
}

type Node struct {
	value int
	min   int
}
type MinStack struct {
	items []Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{items: []Node{}}
}

func (this *MinStack) Push(x int) {
	if len(this.items) == 0 {
		this.items = append(this.items, Node{x, x})
		return
	}
	prev := this.items[len(this.items)-1].min
	if x < prev {
		this.items = append(this.items, Node{x, x})
	} else {
		this.items = append(this.items, Node{x, prev})
	}
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].value
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1].min
}
