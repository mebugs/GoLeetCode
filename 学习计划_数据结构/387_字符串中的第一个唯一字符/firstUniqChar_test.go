package firstUniqChar

import (
	"testing"
)

func firstUniqChar(s string) int {
	// 记录下标
	strMap := make(map[int32]int, 26)
	// 判断存在
	for i, r := range s {
		_, ok := strMap[r]
		if ok {
			strMap[r] = -1
		} else {
			strMap[r] = i
		}
	}
	// 判断index
	for i, r := range s {
		if strMap[r] > -1 {
			return i
		}
	}
	return -1
}

func Test387(t *testing.T) {
	t.Log(firstUniqChar("leetcode"))
	t.Log(firstUniqChar("loveleetcode"))
	t.Log(firstUniqChar("aabb"))
}
