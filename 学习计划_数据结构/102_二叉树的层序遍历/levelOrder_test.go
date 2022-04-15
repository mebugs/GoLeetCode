package levelOrder

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	valuesArray := make([][]int, 0)
	var induction func(node *TreeNode, level int)
	induction = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 如果长度超出等级，说明本层有有记录
		if len(valuesArray) > level {
			valuesArray[level] = append(valuesArray[level], node.Val)
		} else {
			valuesArray = append(valuesArray, []int{node.Val})
		}
		// 先递归左
		induction(node.Left, level+1)
		// 后递归右
		induction(node.Right, level+1)

	}
	// 递归入口
	induction(root, 0)
	return valuesArray
}

func Test102(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		Right: &TreeNode{5, nil, &TreeNode{6, &TreeNode{7, nil, nil}, nil}},
	}
	t.Log(levelOrder(root))
}
