package test

import (
	"fmt"
	"testing"
)

func Test_212(t *testing.T) {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	res := findWords(board, words)
	fmt.Printf("1.res: %v\n", res)
}

type Trie1 struct {
	next  map[rune]*Trie1
	isEnd bool
	word  string
}

func findWords(board [][]byte, words []string) []string {
	// build a trie
	root := &Trie1{next: make(map[rune]*Trie1)}
	for _, word := range words {
		cur := root
		for _, v := range []rune(word) {
			if cur.next[v] == nil {
				cur.next[v] = &Trie1{next: make(map[rune]*Trie1)}
			}
			cur = cur.next[v]
		}
		cur.word = word
		cur.isEnd = true
	}

	// dfs
	res := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, root, i, j, &res)
			if len(res) >= len(words) {
				return words
			}
		}
	}

	// print output
	return res
}

func dfs(board [][]byte, trie *Trie1, i, j int, res *[]string) {
	if trie.isEnd {
		*res = append(*res, trie.word)
		trie.isEnd = false
		trie.word = "" // de-duplicate
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return
	}
	letter := board[i][j]
	next, ok := trie.next[rune(letter)]
	if !ok {
		return
	}
	board[i][j] = '#'
	// 四连通
	dfs(board, next, i+1, j, res)
	dfs(board, next, i-1, j, res)
	dfs(board, next, i, j+1, res)
	dfs(board, next, i, j-1, res)

	board[i][j] = letter
}
