package intersect

import "testing"

// HASH计数
// 排序后双指针方案也可以
func intersect1(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	// 计数对象
	// 阅读示例后，发现此处可以优化为一个int进行加减操作，减少最后一次map的循环
	type SumWork struct {
		OnePoint int
		TwoPoint int
	}
	// 最多计数器
	maxLength := len(nums1)
	if maxLength < len(nums2) {
		maxLength = len(nums2)
	}
	// HASH计数
	sumMap := make(map[int]SumWork, maxLength)
	for _, i1 := range nums1 {
		work, ok := sumMap[i1]
		if ok {
			work.OnePoint++
		} else {
			work = SumWork{
				OnePoint: 1,
				TwoPoint: 0,
			}
		}
		sumMap[i1] = work
	}
	for _, i2 := range nums2 {
		work, ok := sumMap[i2]
		if ok {
			work.TwoPoint++
		} else {
			work = SumWork{
				OnePoint: 0,
				TwoPoint: 1,
			}
		}
		sumMap[i2] = work
	}
	// 并集处理
	inArray := make([]int, 0)
	for index, item := range sumMap {
		if item.OnePoint > 0 && item.TwoPoint > 0 {
			// 取最小出现次数
			point := item.OnePoint
			if point > item.TwoPoint {
				point = item.TwoPoint
			}
			for i := 0; i < point; i++ {
				inArray = append(inArray, index)
			}
		}
	}
	return inArray
}

func Test350(t *testing.T) {
	t.Log(intersect1([]int{1, 2, 2, 1}, []int{2, 2}))
	t.Log(intersect1([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
