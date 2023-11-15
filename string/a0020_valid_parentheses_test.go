package string

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	l := len(s)
	if l&1 == 1 {
		return false
	}
	mappings := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < l; i++ {
		if v, result := mappings[s[i]]; result {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	valid := isValid("{[[()]]}}")
	fmt.Println(valid)
	valid = isValid("{[[()]]}")
	fmt.Println(valid)
}
