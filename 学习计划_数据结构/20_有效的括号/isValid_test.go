package isValid

import (
	"testing"
)

func isValid(s string) bool {
	n := len(s)
	// 长度非偶数，不满足
	if n%2 == 1 || n == 0 {
		return false
	}
	// 需要反向匹配的括号
	p := map[byte]byte{')': '(', ']': '[', '}': '{'}
	// 创建一个栈,最长出现2分之一的长度
	stack := make([]byte, n/2)
	// 栈下标
	stackIndex := -1
	maxIndex := len(stack) - 1
	// 遍历字符串
	for _, item := range []byte(s) {
		// 匹配是否是反向括号
		l, ok := p[item]
		if ok {
			// 栈中无值或者栈的最新值和匹配值不对应
			if stackIndex < 0 || stack[stackIndex] != l {
				return false
			}
			// 匹配成功，数据出栈
			stackIndex--
		} else {
			// 数据入栈
			stackIndex++
			// 超出栈范围
			if stackIndex > maxIndex {
				return false
			}
			stack[stackIndex] = item
		}
	}
	// 栈中有剩余的值
	if stackIndex > -1 {
		return false
	}
	return true
}

func Test20(t *testing.T) {
	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("(]"))
	t.Log(isValid("([)]"))
	t.Log(isValid("{[]}"))
}
