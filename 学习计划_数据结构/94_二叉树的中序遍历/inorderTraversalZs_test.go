package inorderTraversalZs

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 和144高度类似
// 递归法，区别在于读数的位置
func inorderTraversal(root *TreeNode) []int {
	intValue := make([]int, 0)
	var induction func(node *TreeNode)
	induction = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 先递归左
		induction(node.Left)
		// 递归完左，存中间的数据
		intValue = append(intValue, node.Val)
		// 后递归右
		induction(node.Right)
	}
	// 递归入口
	induction(root)
	return intValue
}

func Test94(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		Right: &TreeNode{5, nil, &TreeNode{6, &TreeNode{7, nil, nil}, nil}},
	}
	t.Log(inorderTraversal(root))
}
