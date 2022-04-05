package containsDuplicate

import "testing"

func Test217(t *testing.T) {
	t.Log(containsDuplicate([]int{1, 2, 3, 4, 5}))
	t.Log(containsDuplicate([]int{1, 2, 3, 3, 4, 5}))
}

func containsDuplicate(nums []int) bool {
	// Map 去重
	numMap := map[int]bool{}
	// 遍历数据
	for _, num := range nums {
		// 检查Map中是否存在
		if numMap[num] {
			return true
		} else {
			numMap[num] = true
		}
	}
	return false
}
