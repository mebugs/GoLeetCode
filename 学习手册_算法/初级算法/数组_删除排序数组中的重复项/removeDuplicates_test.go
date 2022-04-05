package removeDuplicates

import (
	"testing"
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	// 定义双指针 In Read
	letI := 0
	for letR := 1; letR < length; letR++ {
		// 如果Read下标的值不等于In下标
		if nums[letR] != nums[letI] {
			// 计数下标后移
			letI++
			// 记录数据
			nums[letI] = nums[letR]
		}
	}
	// fmt.Println(nums)
	return letI + 1
}

func Test(t *testing.T) {
	t.Log(removeDuplicates([]int{1, 1, 2}))
	t.Log(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
