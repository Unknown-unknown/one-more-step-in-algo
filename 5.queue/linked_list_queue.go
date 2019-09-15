package queue

type Node struct {
	next  *Node
	value interface{}
}

type LLQueue struct {
	top *Node
}

func NewLinkedListQueue() *LLQueue {
	return &LLQueue{&Node{nil, nil}}
}

func (q *LLQueue) Front() interface{} {
	return q.top.value
}

func (q *LLQueue) Back() interface{} {
	cur := q.top
	for cur != nil && cur.next != nil {
		cur = cur.next
	}
	return cur.value
}

func (q *LLQueue) Empty() bool {
	return q.top.next == nil
}

func (q *LLQueue) Size() int {
	cur := q.top
	size := 0
	for cur != nil && cur.next != nil {
		cur = cur.next
		size++
	}
	return size
}

func (q *LLQueue) Push(v interface{}) {

}

func (q *LLQueue) Pop() interface{} {
	return nil
}

func (q *LLQueue) Print() {
}
