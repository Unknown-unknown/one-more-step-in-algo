package linked_list

import (
	"fmt"
	"testing"
)

func TestIsPalindromeByLinkedList(t *testing.T) {
	cases := []string{"", "i", "hello", "worlddlrow", "worldlrow"}
	for _, str := range cases {
		l := NewLinkedList()
		for _, c := range str {
			l.PushBack(string(c))
		}
		l.Print()
		fmt.Printf("%s: %+v", str, isPalindromeByLinkedList(l))
		fmt.Println()
	}
}

func TestIsPalindromeByArray(t *testing.T) {
	cases := []string{"", "i", "hello", "worlddlrow", "worldlrow"}
	for _, str := range cases {
		fmt.Printf("%s: %+v", str, isPalindromeByArray(str))
		fmt.Println()
	}
}
