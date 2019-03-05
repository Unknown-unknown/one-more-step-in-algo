package linked_list

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
	var prev *Node = nil
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
		fast := l.head
		slow := l.head
		for fast != nil && fast.next != nil {
			fast = fast.next.next
			slow = slow.next
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

func RemoveBackN(l *LinkedList, n int) *LinkedList {
	return nil
}

func MiddleNode(l *LinkedList) *Node {
	return nil
}
