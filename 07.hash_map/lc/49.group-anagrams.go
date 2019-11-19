/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (49.17%)
 * Likes:    2134
 * Dislikes: 134
 * Total Accepted:    417.1K
 * Total Submissions: 832.9K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 * 
 * Example:
 * 
 * 
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 * 
 * Note:
 * 
 * 
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 * 
 * 
 */

// @lc code=start
// *小结: 归根到底，是思考如何表示唯一字符串，有的直接用 string，有的用 []byte，有的更精细一点用 [26]byte。
// v1, 暴力，存每个 string 对应的频次 map，然后比较
// v2, hashmap, 遍历所有 string，对每一个 string sort 后存入对应 key 的 map，最后把 map 转成 array，时间复杂度 O(nklog(k))
import "sort"
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	res := make([][]string)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		key := string(b)
		groups[key] = append(groups[key], str)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

// v2, hashmap 的精细版本，因为题目中限制了只有小写字母，因此一个长度为 26 的 []byte 就够用了，还同时减少了 []byte 和 string 之间相互转换的时间。beats 99%...
// PS：这里可以学到的一个思路是，可以将“是否存在”或是“出现频次”整体构造成一个 key，类似，bitmap？
// PPS：另一个思路是 map + array，可以实现 ordered map 的效果，这种方式在实际工程代码中还是比较常见的。
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte]int) // 记录共有哪些种模式，key 为 26 个字母的 []byte，value 为该 key 对应的 index
	res := make([][]string, 0)
	for i := range strs {
		list := [26]byte{}
		for j := range strs[i] {	 // for 循环构造好 26 个字母出现频次的 key
			list[strs[i][j]-'a']++
		}
		if idx, ok := groups[list]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			groups[list] = len(res) - 1
		}
	}
	return res
}

// v3, 利用质因子相乘写法只有一种，思路很巧妙，但容易 stack overflow…… ┑(￣Д ￣)┍
// 参：https://leetcode.wang/leetCode-49-Group-Anagrams.html

// @lc code=end

