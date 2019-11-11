package test

import (
	"fmt"
	"testing"
)

func Test_127(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("1.res: %d\n", res)

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	res = ladderLength(beginWord, endWord, wordList)
	fmt.Printf("2.res: %d\n", res)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	res := 0
	queue := []string{beginWord}
	visited := make([]bool, len(wordList))

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur == endWord {
				return res + 1
			}
			for j := 0; j < len(wordList); j++ {
				if !visited[j] && isSingleCharDiff(cur, wordList[j]) {
					visited[j] = true
					queue = append(queue, wordList[j])
				}
			}
			fmt.Printf("cur: %v, queue: %v\n", cur, queue)
		}
		res++
		fmt.Printf("res: %d\n", res)
	}
	return 0
}

// func isSingleCharDiff(s1, s2 string) bool {
// 	diff := 0
// 	for i, s := range s1 {
// 		if s != []rune(s2)[i] {
// 			diff++
// 		}
// 	}
// 	return diff == 1
// }
