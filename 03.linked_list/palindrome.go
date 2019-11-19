package linked_list

func isPalindromeByLinkedList(l *LinkedList) bool {
	len := l.length
	if len == 0 || len == 1 {
		return true
	}

	stack := make([]string, len/2)
	current := l.head
	for i := 1; i <= len; i++ {
		current = current.next
		if len%2 != 0 && i == len/2+1 {
			continue
		}
		if i <= len/2 {
			stack[i-1] = current.GetValue().(string)
		} else {
			if stack[len-i] != current.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

func isPalindromeByArray(a string) bool {
	if len(a) == 0 || len(a) == 1 {
		return true
	}
	length := len(a)
	for i := 0; i < length; i++ {
		if a[i] != a[length-i-1] {
			return false
		}
	}
	return true
}
