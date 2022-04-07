package generate

import "testing"

func generate(numRows int) [][]int {
	nums := make([][]int, numRows)
	// 遍历赋值
	for index := range nums {
		// 每一行的数据量是index + 1
		nums[index] = make([]int, index+1)
		// 每一行第一个和最后一个都是1
		nums[index][0] = 1
		nums[index][index] = 1
		// 第一行无需计算
		if index == 0 {
			continue
		}
		// 本行的第N个数字 = 上行第N+第(N-1)数字之和
		// 第一个和最后一个数字不变（由于index从0开始，所以无需-1）
		for rI := 1; rI < index; rI++ {
			nums[index][rI] = nums[index-1][rI] + nums[index-1][rI-1]
		}
	}
	return nums
}

func Test118(t *testing.T) {
	t.Log(generate(5))
	t.Log(generate(1))
	t.Log(generate(3))
}
