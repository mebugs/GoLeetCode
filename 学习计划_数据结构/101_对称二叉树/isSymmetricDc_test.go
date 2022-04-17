package isSymmetricDc

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 是否对称
func isSymmetric(root *TreeNode) bool {
	var induction func(nodeL, nodeR *TreeNode) bool
	induction = func(nodeL, nodeR *TreeNode) bool {
		// 比对截止到叶子
		if nodeL == nil && nodeR == nil {
			return true
		}
		// 任意一侧到了叶子，另一边未到叶子
		if nodeL == nil || nodeR == nil {
			return false
		}
		// 相等继续向下，镜像对比
		return nodeL.Val == nodeR.Val && induction(nodeL.Left, nodeR.Right) && induction(nodeL.Right, nodeR.Left)
	}
	return induction(root, root)
}

func Test101(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		Right: &TreeNode{5, nil, &TreeNode{6, &TreeNode{7, nil, nil}, nil}},
	}
	t.Log(isSymmetric(root))
	root = &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{4, nil, nil}},
		Right: &TreeNode{2, &TreeNode{4, nil, nil}, nil},
	}
	t.Log(isSymmetric(root))
}
