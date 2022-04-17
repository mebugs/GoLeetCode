package maxDepthBi

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth1(root *TreeNode) int {
	depth := 0
	var induction func(node *TreeNode, level int)
	induction = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 记录深度
		if depth < level+1 {
			depth = level + 1
		}
		// 先递归左
		induction(node.Left, level+1)
		// 后递归右
		induction(node.Right, level+1)

	}
	// 递归入口
	induction(root, 0)
	return depth
}

// 优化计算
func maxDepth(root *TreeNode) int {
	var induction func(root *TreeNode) int
	induction = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(induction(root.Left), induction(root.Right)) + 1

	}
	// 递归入口
	return induction(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test104(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		Right: &TreeNode{5, nil, &TreeNode{6, &TreeNode{7, nil, nil}, nil}},
	}
	t.Log(maxDepth(root))
}
