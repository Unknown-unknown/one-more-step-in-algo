package test

import (
	"fmt"
	"sort"
	"testing"
)

func Test_49(t *testing.T) {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	res := groupAnagrams(strs)
	fmt.Printf("res: %v\n", res)

	res = groupAnagramsPrecise(strs)
	fmt.Printf("res: %v\n", res)
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
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

func groupAnagramsPrecise(strs []string) [][]string {
	groups := make(map[[26]byte]int) // 记录共有哪些种模式，key 为 26 个字母的 []byte，value 为该 key 对应的 index
	res := make([][]string, 0)
	for i := range strs {
		list := [26]byte{}
		for j := range strs[i] { // for 循环构造好 26 个字母出现频次的 key
			list[strs[i][j]-'a']++
		}

		if idx, ok := groups[list]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			groups[list] = len(res) - 1
		}
		fmt.Printf("i = %d, list = %v, res = %v, groups = %v\n", i, list, res, groups)
	}
	fmt.Printf("groups = %v, res = %v", groups, res)
	return res
}
