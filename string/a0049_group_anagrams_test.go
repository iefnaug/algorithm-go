package string

import (
	"fmt"
	"sort"
	"testing"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。



示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]


提示：

1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func GroupAnagrams(strs []string) [][]string {
	mappings := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		b := []byte(strs[i])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		str := string(b)
		//if v, ok := mappings[str]; ok {
		//	mappings[str] = append(v, strs[i])
		//} else {
		//	ns := make([]string, 0)
		//	mappings[str] = append(ns, strs[i])
		//}
		mappings[str] = append(mappings[str], strs[i])
	}
	ret := make([][]string, 0, len(mappings))
	for _, v := range mappings {
		ret = append(ret, v)
	}
	return ret
}

//计数
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func TestGroupAnagrams(t *testing.T) {
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "", ""}
	anagrams := GroupAnagrams(strs)
	for _, s := range anagrams {
		fmt.Println(s)
	}
}
