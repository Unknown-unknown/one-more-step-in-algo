package linked_list_advance

/*
- 单链表反转
- 链表中环的检测
- 两个有序链表的合并
- 删除链表倒数第n个结点
- 求链表的中间结点
*/

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length int
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// func (l *LinkedList) Reverse() *LinkedList {

// }
