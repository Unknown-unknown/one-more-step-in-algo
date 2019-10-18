/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 *
 * https://leetcode.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (38.63%)
 * Likes:    2183
 * Dislikes: 237
 * Total Accepted:    359.4K
 * Total Submissions: 917K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum
 * element in constant time.
 * 
 * 
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 * 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> Returns -3.
 * minStack.pop();
 * minStack.top();      --> Returns 0.
 * minStack.getMin();   --> Returns -2.
 * 
 * 
 * 
 * 
 */

// v1, array 或 linked list 实现 stack：缺点是 GetMin() 时需要遍历整个数组，时间复杂度 O(n)；优点是无需额外存储空间
// v2, 空间换时间，额外存储当前最小值：可以将当前最小值存在单独的 stack 中，也可以存在每个 node 中，时间复杂度为 O(1)
// @lc code=start
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



/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

