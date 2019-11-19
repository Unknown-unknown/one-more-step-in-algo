/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (53.43%)
 * Likes:    888
 * Dislikes: 118
 * Total Accepted:    411.8K
 * Total Submissions: 764.2K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 * 
 * Example 1:
 * 
 * 
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "rat", t = "car"
 * Output: false
 * 
 * 
 * Note:
 * You may assume the string contains only lowercase alphabets.
 * 
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 * 
 */

// @lc code=start
// anagram: 每个字母出现的频次相同
// v1, 暴力破解, sort的方式选择先把 string split，排序后再 join，最后比较已排序的string, 时间复杂度 O(nlogn)
import (
	"sort"
	"strings"
)
func isAnagram(s string, t string) bool {
	return sortedString(s) == sortedString(t)
}

func sortedString(s string) string {
	bs := strings.Split(s, "")
	sort.Strings(bs)
	return strings.Join(bs, "")
}

// v2, 暴力破解, sort的方式选择将 string 转成 []byte，排序后再比较, O(nlogn)，但运行时间要比前一种方法短得多
import "sort"
func isAnagram(s string, t string) bool {
	return sortedString(s) == sortedString(t)
}

func sortedString(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	return string(bs)
}

// v3, hashmap, 统计每个字母出现的频次，分别存入 map 中，再利用 golang 的 reflect.DeepEqual 比较两个 map，效率最高
/* 
* 对于 DeepEqual 方法的详细说明：（deeply equal实在翻译无力，保持原文好了……）
- 类型不同的值绝不会 deeply equal
- 数组：对应的每个元素都 deeply equal
- struct：包括导出和非导出字段在内的全部字段，对应值都相等
- 函数：都为 nil，否则不是 deeply equal
- interface：如果二者持有 deeply equal 的具体值
- map：满足以下全部条件：都为 nil或都不为 nil，长度相同，要么是完全相同的 map 对象、要么是对应 key 的 值是 deeply equal 的。
- pointer：Go的 == 操作符计算为 true，或是指向 deeply equal 的值
- slice：满足以下全部条件：都为 nil或都不为 nil，长度相同，要么是指向相同的底层 array、要么是他们对应的元素都是 deeply equal。注意：非 nil 的空的 slice，和 nil 的 slice（比如，[]byte{} vs []byte(nil)）不是 deeply equal 的
- 其他值：使用 Go 的 == 操作符比较
*/
import "reflect"
func isAnagram(s string, t string) bool {
	cache1 := make(map[rune]int)
	cache2 := make(map[rune]int)
	for _, r := range s {
		cache1[r]++
	}
	for _, r := range t {
		cache2[r]++
	}
	return reflect.DeepEqual(cache1, cache2)
}

// @lc code=end

