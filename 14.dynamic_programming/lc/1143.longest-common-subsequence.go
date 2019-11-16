/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 *
 * https://leetcode.com/problems/longest-common-subsequence/description/
 *
 * algorithms
 * Medium (57.67%)
 * Likes:    332
 * Dislikes: 8
 * Total Accepted:    19.7K
 * Total Submissions: 34.4K
 * Testcase Example:  '"abcde"\n"ace"'
 *
 * Given two strings text1 and text2, return the length of their longest common
 * subsequence.
 *
 * A subsequence of a string is a new string generated from the original string
 * with some characters(can be none) deleted without changing the relative
 * order of the remaining characters. (eg, "ace" is a subsequence of "abcde"
 * while "aec" is not). A common subsequence of two strings is a subsequence
 * that is common to both strings.
 *
 *
 *
 * If there is no common subsequence, return 0.
 *
 *
 * Example 1:
 *
 *
 * Input: text1 = "abcde", text2 = "ace"
 * Output: 3
 * Explanation: The longest common subsequence is "ace" and its length is 3.
 *
 *
 * Example 2:
 *
 *
 * Input: text1 = "abc", text2 = "abc"
 * Output: 3
 * Explanation: The longest common subsequence is "abc" and its length is 3.
 *
 *
 * Example 3:
 *
 *
 * Input: text1 = "abc", text2 = "def"
 * Output: 0
 * Explanation: There is no such common subsequence, so the result is 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= text1.length <= 1000
 * 1 <= text2.length <= 1000
 * The input strings consist of lowercase English characters only.
 *
 *
 */

// @lc code=start
/*
可能有以下几种情况：
1. s1 = ”“，s2= 任意字符串，LCS = 0
2. s1 = ”A“，s2 = 任意字符串，LCS = 1
3. s1 = ”.....A“，s2 = ”...A“，LCS = LCS（......，...）+ 1

从后往前对比，状态的方程就变为：
- if s1[-1] != s2[-1]: LCS[s1, s2] = max(LCS[s1-1, s2], LCS[s1, s2-1])
- if s1[-1] == s2[-1]: LCS[s1, s2] = LCS[s1-1, s2-1] + 1
*/
// v1, 二维数组存储
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if []rune(text1)[i-1] == []rune(text2)[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// v2, 两个一维数组存储
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	cur, prev := make([]int, m+1), make([]int, m+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			maxLen := max(prev[j], cur[j-1])
			if []rune(text1)[j-1] == []rune(text2)[i-1] {
				cur[j] = max(maxLen, prev[j-1]+1)
			} else {
				cur[j] = maxLen
			}
		}
		cur, prev = prev, cur
	}
	return prev[m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

