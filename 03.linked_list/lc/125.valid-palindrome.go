/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (32.27%)
 * Likes:    691
 * Dislikes: 2035
 * Total Accepted:    407.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 * 
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 * 
 * Example 1:
 * 
 * 
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "race a car"
 * Output: false
 * 
 * 
 */
// v1, use regexp lib
import "regexp"
func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return false
	}
	cleanStr := bytes.ToLower([]byte(reg.ReplaceAllString(s, "")))
	length := len(cleanStr)
	for i := 0; i < length; i++ {
		if cleanStr[i] != cleanStr[length - i - 1] {
			return false
		}
	}
	return true
}

// v2, use unicode lib
import (
	"unicode"
	"strings"
) 
func isPalindrome(s string) bool {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	arr := strings.FieldsFunc(strings.ToLower(s), f)
	cleanStr := strings.Join(arr, "")
	length := len(cleanStr)

	for i := 0; i < length; i++ {
		if cleanStr[i] != cleanStr[length - i - 1] {
			return false
		}
	}
	return true
}
