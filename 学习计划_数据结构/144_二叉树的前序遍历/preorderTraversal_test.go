package preorderTraversalDd

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func preorderTraversal(root *TreeNode) []int {
	intValue := make([]int, 0)
	var induction func(node *TreeNode)
	induction = func(node *TreeNode) {
		if node == nil {
			return
		}
		intValue = append(intValue, node.Val)
		// 先递归左
		induction(node.Left)
		// 后递归右
		induction(node.Right)
	}
	// 递归入口
	induction(root)
	return intValue
}

func Test144(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		Right: &TreeNode{5, nil, &TreeNode{6, &TreeNode{7, nil, nil}, nil}},
	}
	t.Log(preorderTraversal(root))
}
