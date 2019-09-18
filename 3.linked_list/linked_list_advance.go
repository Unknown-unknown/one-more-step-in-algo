package linked_list

import "fmt"

/* Reverse process step by step:
eg: head -> 1 -> 2 -> 3 -> nil
=) head(pre) -> 1(current) -> 2(tmp) -> 3 -> nil
=) nil(pre) <- 1(current) -> 2(tmp) -> 3 -> nil
=) nil <- 1(pre) <- 2(current) -> 3(tmp) -> nil
=) nil <- 1 <- 2(prev) -> 3(current) -> nil(tmp)
=) nil <- 1 <- 2(prev) <- 3(current) -> nil(tmp)
=) nil <- 1 <- 2 <- 3(prev) -> nil(current)
=) nil <- 1 <- 2 <- 3(prev) <- head
*/
func (l *LinkedList) Reversed() {
	if l.head == nil || l.head.next == nil {
		return
	}
	current := l.head.next
	var prev *Node
	for current != nil {
		tmp := current.next
		current.next = prev
		prev = current
		current = tmp
	}
	l.head.next = prev
}

func (l *LinkedList) HasRing() bool {
	if l.head != nil {
		slow, fast := l.head, l.head
		for fast != nil && fast.next != nil {
			slow, fast = slow.next, fast.next.next
			if fast == slow {
				return true
			}
		}
	}
	return false
}

func MergeSortedList(ll, lr *LinkedList) *LinkedList {
	if ll == nil || ll.head == nil || ll.head.next == nil {
		return lr
	}
	if lr == nil || lr.head == nil || lr.head.next == nil {
		return ll
	}

	l := NewLinkedList()
	llnext := ll.head.next
	lrnext := lr.head.next
	for llnext != nil && lrnext != nil {
		if llnext.GetValue().(int) < lrnext.GetValue().(int) {
			l.PushBack(llnext.GetValue())
			llnext = llnext.next
		} else {
			l.PushBack(lrnext.GetValue())
			lrnext = lrnext.next
		}
	}

	for llnext != nil {
		l.PushBack(llnext.GetValue())
		llnext = llnext.next
	}

	for lrnext != nil {
		l.PushBack(lrnext.GetValue())
		lrnext = lrnext.next
	}
	return l
}

func (l *LinkedList) RemoveBackN(n int) *LinkedList {
	if l == nil || l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head
	for i := 0; i < n; i++ {
		if fast == nil {
			fmt.Printf("n is out of range, since list length is %d", l.length)
			return nil
		}
		fast = fast.next
	}

	for fast != nil {
		slow = slow.next
		fast = fast.next
	}
	l.Remove(slow)
	return l
}

func (l *LinkedList) FindMiddleNode() []*Node {
	if l == nil || l.head == nil {
		return nil
	}

	slow := l.head
	fast := l.head
	var tmp *Node
	var middle []*Node
	for fast != nil && fast.next != nil {
		tmp = slow
		slow = slow.next
		fast = fast.next.next
	}
	if l.length%2 == 0 {
		middle = append(middle, tmp)
		middle = append(middle, slow)
	} else {
		middle = append(middle, slow)
	}
	return middle
}
