/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Add and Search Word - Data structure design
 *
 * https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
 *
 * algorithms
 * Medium (31.74%)
 * Likes:    1126
 * Dislikes: 64
 * Total Accepted:    138.6K
 * Total Submissions: 422.5K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * Design a data structure that supports the following two operations:
 *
 *
 * void addWord(word)
 * bool search(word)
 *
 *
 * search(word) can search a literal word or a regular expression string
 * containing only letters a-z or .. A . means it can represent any one
 * letter.
 *
 * Example:
 *
 *
 * addWord("bad")
 * addWord("dad")
 * addWord("mad")
 * search("pad") -> false
 * search("bad") -> true
 * search(".ad") -> true
 * search("b..") -> true
 *
 *
 * Note:
 * You may assume that all words are consist of lowercase letters a-z.
 *
 */

// @lc code=start
type WordDictionary struct {
	next  map[rune]*WordDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		next: make(map[rune]*WordDictionary),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, v := range []rune(word) {
		if this.next[v] == nil {
			this.next[v] = &WordDictionary{
				next: make(map[rune]*WordDictionary),
			}
		}
		this = this.next[v]
	}
	this.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for k, v := range word {
		if v == '.' {
			for _, v := range this.next {
				if v.Search(word[k+1:]) {
					return true
				}
			}
			return false
		} else {
			if this.next[v] == nil {
				return false
			}
			this = this.next[v]
		}
	}
	return this.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

