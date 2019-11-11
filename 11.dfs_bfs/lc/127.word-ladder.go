/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (25.43%)
 * Likes:    1967
 * Dislikes: 931
 * Total Accepted:    317.2K
 * Total Submissions: 1.2M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 *
 *
 * Note:
 *
 *
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output: 5
 *
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: 0
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 *
 *
 */

// @lc code=start
// v1, bfs
func ladderLength(beginWord string, endWord string, wordList []string) int {
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
		}
		res++
	}
	return 0
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

// @lc code=end

