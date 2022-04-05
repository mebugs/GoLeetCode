package maxSubArray

import "testing"

func Test53(t *testing.T) {
	t.Log(maxSubArray([]int{0, -1, 3, 6, -3, 4, 7, -1}))
}

func maxSubArray(nums []int) int {
	// 动态规划
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果前一位大于0，进行求和，否则丢弃
		if nums[i-1] > 0 {
			// 取得当前位的求和
			nums[i] = nums[i] + nums[i-1]
		}
		// 当前节点最大求和 大于最大值
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}
