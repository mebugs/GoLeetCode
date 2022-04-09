package canConstruct

import "testing"

func canConstructOld(ransomNote string, magazine string) bool {
	// ransomNote 字母构成
	rA := make([]int, 26)
	// magazine 字母构成
	mA := make([]int, 26)
	// 转换byte
	rB := []byte(ransomNote)
	mB := []byte(magazine)
	// 通过下标寄存26个字母出现的此处
	for _, r := range rB {
		// 取的字母对应的下标
		rA[r-'a']++
	}
	for _, m := range mB {
		mA[m-'a']++
	}
	for i := 0; i < 26; i++ {
		// ransomNote 出现该字母次数更多 无法构成
		if rA[i] > mA[i] {
			return false
		}
	}
	return true
}

// 优化版
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	// 字母构成
	mA := make([]int, 26)
	// 通过下标寄存26个字母出现的此处,资源构成
	for _, m := range magazine {
		mA[m-'a']++
	}
	for _, r := range ransomNote {
		// 取的字母对应的下标
		mA[r-'a']--
		if mA[r-'a'] < 0 {
			return false
		}
	}
	return true
}

func Test383(t *testing.T) {
	t.Log(canConstruct("a", "b"))
	t.Log(canConstruct("a", "ab"))
	t.Log(canConstruct("aba", "aab"))
}
