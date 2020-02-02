/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (22.84%)
 * Likes:    1893
 * Dislikes: 2144
 * Total Accepted:    321.2K
 * Total Submissions: 1.4M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * Example 1:
 *
 *
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 */

// @lc code=start
import "strconv"

func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	dp := make([]rune, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	for i := 2; i <= n; i++ {
		a, b := strconv.Atoi(string(s[i-1:i])), strconv.Atoi(string(s[i-2:i]))
		if a >= 1 && a <= 9 {
			dp[i] += dp[i-1]
		}
		if b >= 10 && b <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// @lc code=end

