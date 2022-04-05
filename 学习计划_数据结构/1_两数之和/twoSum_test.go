package twoSum

import "testing"

func Test1(t *testing.T) {
	t.Log(twoSum1([]int{2,7,11,15},9))
	t.Log(twoSum1([]int{3,2,4},6))
	t.Log(twoSum1([]int{2,5,5,11},10))
	t.Log(twoSum2([]int{2,7,11,15},9))
	t.Log(twoSum2([]int{3,2,4},6))
	t.Log(twoSum2([]int{2,5,5,11},10))
}
// 暴力算法
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}
	for i:=0;i<length-1;i++ {
		for j:=i+1;j<length;j++{
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}

// 通过target - i 获取 目标值
func twoSum2(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return nil
	}
	// 通过haspMap记录数据
	sumMap := make(map[int]int,length-1)
	for i,item := range nums{
		// 尝试获取与当前值匹配的index
		oIndex,ok := sumMap[item]
		if ok {
			// 返回上个index当前index
			return []int{oIndex,i}
		} else {
			// 求出减值（当前节点需要匹配的值）
			decNo := target -item
			// 记录目标值和当前节点
			sumMap[decNo] = i
		}
	}
	return nil
}