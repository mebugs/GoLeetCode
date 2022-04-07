package maxProfit

import "testing"

func maxProfit(prices []int) int {
	// 动态规划 取的截止目前的最小值，并获得今天售出的利润
	minPri := prices[0]
	maxLr := 0
	for i := 1; i < len(prices); i++ {
		// 今日减去最小买入 = 利润 > 最大利润则替换
		if prices[i]-minPri > maxLr {
			maxLr = prices[i] - minPri
		}
		// 今天的价格比最小买入低，计入新的买入价格
		if prices[i] < minPri {
			minPri = prices[i]
		}
	}
	return maxLr
}

func Test121(t *testing.T) {
	t.Log(maxProfit([]int{7, 2, 4, 3, 1, 6, 3}))
	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{7, 6, 4, 3, 1}))
}
