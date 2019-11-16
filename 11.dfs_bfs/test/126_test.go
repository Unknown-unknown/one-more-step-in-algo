package test

import (
	"fmt"
	"testing"
)

func Test_126(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := findLadders(beginWord, endWord, wordList)
	fmt.Printf("1.res: %v\n", res)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := [][]string{}
	// queue := []string{beginWord}

	// for len(queue) > 0 {
	// 	size := len(queue)
	// 	arr := []string{}
	// 	for i := 0; i < size; i++ {
	// 		cur := queue[0]
	// 		queue = queue[1:]

	// 		if cur == endWord {
	// 			arr = append(arr, cur)
	// 			continue
	// 		}
	// 		for j := 0; j < len(wordList); j++ {
	// 			if isSingleCharDiff(cur, wordList[j]) {
	// 				arr = append(arr, wordList[j])
	// 			}
	// 		}
	// 	}
	// 	res = append(res, arr)
	// }

	return res
}
