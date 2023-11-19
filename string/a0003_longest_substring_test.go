package string

import (
	"fmt"
	"testing"
)

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
提示：
0 <= s.length <= 5 * 10^4
s 由英文字母、数字、符号和空格组成
*/
func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	mappings := map[byte]int{}
	maxLength, length, minIndex := 0, 0, 0
	for i := 0; i < len(b); i++ {
		if v, ok := mappings[b[i]]; ok && v >= minIndex {
			//存在重复的值
			length = i - v
			minIndex = v + 1
		} else {
			length++
		}
		mappings[b[i]] = i
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "pwwkew"
	max := lengthOfLongestSubstring(s)
	fmt.Println(max)
}
