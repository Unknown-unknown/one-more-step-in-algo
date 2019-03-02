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

func Merge(ll, lr *LinkedList) *LinkedList {
	return nil
}

func RemoveBackN(l *LinkedList, n int) *LinkedList {
	return nil
}

func MiddleNode(l *LinkedList) *Node {
	return nil
}
