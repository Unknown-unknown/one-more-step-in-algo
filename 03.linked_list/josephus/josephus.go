package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	tmp := new(ListNode)
	tmp.Next = tmp
	tmp.Val = val
	return tmp
}

// v1: circular linked list, time complexity: O(m*n)
// m: start from m-th
// n: n nodes in total
func GetJosephusPosition_1(m, n int) int {
	// 1.make a ring
	head := NewNode(1)
	prev := head
	for i := 2; i <= n; i++ {
		prev.Next = NewNode(i)
		prev = prev.Next
	}
	prev.Next = head

	// 2.begin to kill
	ptr1, ptr2 := head, head
	for ptr1.Next != ptr1 {
		count := 1
		for count != m {
			ptr2 = ptr1
			ptr1 = ptr1.Next
			count ++
		}
		// ptr2 -> ptr1 -> ..., remove ptr1
		fmt.Printf("%d got killed\n", ptr1.Val)
		ptr2.Next = ptr1.Next
		ptr1 = ptr2.Next
	}
	return ptr1.Val
}

// v2: recursion， time complexity: O(n)
func GetJosephusPosition_2(m, n int) int {
	// 1,2,3,4,5,6,7 | 2 (m = 2, n = 7)
	// -> 1,0,1,0,1,0,1
	// -> 0,0,1,0,0,0,1
	// -> 0,0,0,0,0,0,1

	// GetJosephusPosition_2(m, n)
	// -> m-th killed, turns to be GetJosephusPosition_2(m, n-1)
	// -> GetJosephusPosition_2(m, n-1), consider the position starting from m%n +1

	if n == 1 {
		return 1
	} 
	return (GetJosephusPosition_2(m, n - 1) + m-1) % n + 1
}

func main(){
	// m, n := 3, 41
	m, n := 2, 7
	fmt.Printf("Start: kill %d-th person of all %d, start from first...\n", m, n)
	res1 := GetJosephusPosition_1(m, n)
	fmt.Printf("Last person left is %d\n", res1)

	res2 := GetJosephusPosition_2(m, n)
	fmt.Printf("Last person left is %d\n", res2)
}
