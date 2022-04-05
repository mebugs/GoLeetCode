package merge

import (
	"fmt"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 双数组指针
	mi := m - 1
	ni := n - 1
	// 总数组指针不断前移
	for ai := m + n - 1; mi >= 0 || ni >= 0; ai-- {
		// m==0
		if mi == -1 { // 数组1没值
			nums1[ai] = nums2[ni]
			ni-- // 数组2指针前移
		} else if ni == -1 {
			// 无需对数组1赋值
			nums1[ai] = nums1[mi] // 数组2没值
			mi--                  // 数组1指针前移
		} else if nums1[mi] > nums2[ni] {
			nums1[ai] = nums1[mi] // 数组1的末尾值被后一到总数组
			mi--                  // 数组1指针前移
		} else {
			nums1[ai] = nums2[ni] // 数组2的末尾值被后一到总数组
			ni--                  // 数组2指针前移
		}
	}
	// 打印测试
	fmt.Println(nums1)
}

func Test88(t *testing.T) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
	merge([]int{2, 0}, 1, []int{1}, 1)
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
