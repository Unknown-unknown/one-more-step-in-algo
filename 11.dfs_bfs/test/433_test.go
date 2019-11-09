package test

import (
	"fmt"
	"testing"
)

func Test_433(t *testing.T) {
	// slice := ([]byte)("hello")
	// fmt.Printf("call: %v, %p\n", slice, slice)

	// AddOneToEachElement(slice)

	// No.443
	start, end := "AACCGGTT", "AACCGGTA"
	bank := []string{"AACCGGTA"}
	res := minMutationBFS(start, end, bank)
	fmt.Printf("1.res: %v\n", res)

	start, end = "AACCGGTT", "AAACGGTA"
	bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	res = minMutationBFS(start, end, bank)
	fmt.Printf("2.res: %v\n", res)

	start, end = "AAAAACCC", "AACCCCCC"
	bank = []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	res = minMutationBFS(start, end, bank)
	fmt.Printf("3.res: %v\n", res)
}

func AddOneToEachElement(slice []byte) {
	fmt.Printf("before: %v, %p\n", slice, slice)

	for i := range slice {
		fmt.Printf("in: %v, %p\n", slice, slice)
		slice[i]++
		slice = append(slice, 'a')
	}
	fmt.Printf("after: %v, %p\n", slice, slice)
}

func minMutationBFS(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	res := 0
	queue := []string{start}
	visited := make([]bool, len(bank))
	for len(queue) > 0 {
		size := len(queue) // !! pitfall !!
		for j := 0; j < size; j++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("j=%d, queue: %v, cur: %v\n", j, queue, cur)
			// fmt.Printf("dequeue: %v, cur: %v, end: %v\n", queue, cur, end)

			for i := 0; i < len(bank); i++ {
				if cur == end {
					// fmt.Printf("return: %v\n", res)
					return res
				}
				if !visited[i] && isSingleCharDiff(cur, bank[i]) {
					visited[i] = true
					queue = append(queue, bank[i])
					// fmt.Printf("append queue: %v, visited: %v\n", queue, visited)
				}
			}
		}
		// fmt.Printf("res: %v\n", res)
		res++
	}

	return -1
}

func isSingleCharDiff(s1, s2 string) bool {
	diff := 0
	for i, s := range s1 {
		if s != []rune(s2)[i] {
			diff++
		}
	}
	return diff == 1
}
