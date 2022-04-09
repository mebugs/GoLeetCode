package isAnagram

import "testing"

// 和383题类似，方法首要判定有所区别
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 资源比对寄存数组
	mA := make([]int, 26)
	// 通过下标寄存26个字母出现的此处,资源构成
	for _, m := range s {
		mA[m-'a']++
	}
	for _, r := range t {
		// 取的字母对应的下标
		mA[r-'a']--
		if mA[r-'a'] < 0 {
			return false
		}
	}
	return true
}

func Test242(t *testing.T) {
	t.Log(isAnagram("anagram", "anagram"))
	t.Log(isAnagram("rat", "car"))
	t.Log(isAnagram("ab", "abc"))
}
