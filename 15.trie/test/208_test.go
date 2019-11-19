package test

import (
	"fmt"
	"testing"
)

func Test_208(t *testing.T) {
	// equivalent to 14/15
	// ["Trie","insert","insert","search"]\n[[],["a"],["pa"],["p"]]
	obj := Constructor()
	obj.Insert("a")
	obj.Insert("pa")
	res := obj.Search("p")
	fmt.Printf("res: %v\n", res)
}

type Node struct {
	next      map[string]*Node
	endOfWord bool
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{next: make(map[string]*Node)}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: NewNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, v := range []rune(word) {
		str := string(v)
		if cur.next[str] == nil {
			cur.next[str] = NewNode()
		}
		cur = cur.next[str]
	}
	cur.endOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, v := range []rune(word) {
		str := string(v)
		if cur.next[str] == nil {
			return false
		}
		cur = cur.next[str]
	}
	return cur.endOfWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, v := range []rune(prefix) {
		str := string(v)
		if cur.next[str] == nil {
			return false
		}
		cur = cur.next[str]
	}
	return true
}
