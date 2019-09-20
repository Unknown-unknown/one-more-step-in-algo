/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (26.27%)
 * Likes:    2249
 * Dislikes: 104
 * Total Accepted:    215.6K
 * Total Submissions: 819.1K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 * 
 * Example 1:
 * 
 * 
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 * 
 * 
 */

// v1, brute force, time complexity: O(n³), as expected...  Time Limit Exceeded... 2333333
func longestValidParentheses(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i+2; j <= len(s); j+=2 {
			if isValid(s[i:j]) {
				if j - i > max {
					max = j - i
				}
			}
		}
	}
	return max
}

func isValid (s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	stack := []rune{}
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
			continue
		}
		if c == '(' {
			stack = append(stack, c)
		} else {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// v2, dynamic programming.... skip... 🤣 for now...
func longestValidParentheses(s string) int {

}

// v3, stack, time complexity: O(n), space complexity: O(n)
func longestValidParentheses(s string) int {
	max := 0
	stack := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if i - stack[len(stack)-1] > max {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}

// v4, two-pointer, 防止漏解，“((()”
func longestValidParentheses(s string) int {
	left, right, max := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left ++
		} else {
			right ++
		}
		if left == right {
			if right * 2 > max {
				max = right * 2
			}
		} else if right >= left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left * 2 > max {
				max = left * 2
			}
		} else if left >= right {
			left, right = 0, 0
		}
	}
	return max
}
