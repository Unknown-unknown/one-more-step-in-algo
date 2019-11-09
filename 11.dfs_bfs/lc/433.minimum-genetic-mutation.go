/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 *
 * https://leetcode.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (39.09%)
 * Likes:    308
 * Dislikes: 37
 * Total Accepted:    24.9K
 * Total Submissions: 62.9K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * A gene string can be represented by an 8-character long string, with choices
 * from "A", "C", "G", "T".
 *
 * Suppose we need to investigate about a mutation (mutation from "start" to
 * "end"), where ONE mutation is defined as ONE single character changed in the
 * gene string.
 *
 * For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.
 *
 * Also, there is a given gene "bank", which records all the valid gene
 * mutations. A gene must be in the bank to make it a valid gene string.
 *
 * Now, given 3 things - start, end, bank, your task is to determine what is
 * the minimum number of mutations needed to mutate from "start" to "end". If
 * there is no such a mutation, return -1.
 *
 * Note:
 *
 *
 * Starting point is assumed to be valid, so it might not be included in the
 * bank.
 * If multiple mutations are needed, all mutations during in the sequence must
 * be valid.
 * You may assume start and end string is not the same.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * return: 1
 *
 *
 *
 *
 * Example 2:
 *
 *
 * start: "AACCGGTT" -> AACCGGTA -> AAACGGTA
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * return: 2
 *
 *
 *
 *
 * Example 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * return: 3
 *
 *
 *
 *
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	res := 0
	queue := []string{start}
	visited := make([]bool, len(bank))
	for len(queue) > 0 {
		size := len(queue)
		for j := 0; j < size; j++ { // !! pitfall !!如果直接用 len(queue) 做 for 循环，则每次该层的循环次数都以第一次 queue 的长度为准，不会重新计算
			cur := queue[0]
			queue = queue[1:]

			if cur == end {
				return res
			}
			for i := 0; i < len(bank); i++ {
				if !visited[i] && isSingleCharDiff(cur, bank[i]) {
					visited[i] = true
					queue = append(queue, bank[i])
				}
			}
		}
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

// @lc code=end

