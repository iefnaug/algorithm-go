package string

import (
	"fmt"
	"testing"
)

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

*/

//动态规划
func longestPalindrome(s string) string {
	b := []byte(s)
	if len(b) <= 1 {
		return s
	}
	maxLen := 1
	begin := 0
	dp := make([][]bool, len(b))
	for i := 0; i < len(b); i++ {
		dp[i] = make([]bool, len(b))
		dp[i][i] = true
	}

	for step := 2; step <= len(b); step++ {
		for i := 0; i <= len(b)-step; i++ {
			if b[i] == b[i+step-1] {
				if step <= 3 {
					dp[i][i+step-1] = true
				} else {
					dp[i][i+step-1] = dp[i+1][i+step-2]
				}
				if dp[i][i+step-1] && step > maxLen {
					maxLen = step
					begin = i
				}
			} else {
				dp[i][i+step-1] = false
			}
		}
	}
	return string(b[begin : begin+maxLen])
}

func TestLongestPalindrome(t *testing.T) {
	//s := "babad"
	//s := "bb"
	s := "aacabdkacaa"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}
