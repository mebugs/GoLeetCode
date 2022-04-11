package deleteDuplicates

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 缓存节点
	tmp := head
	for tmp.Next != nil {
		// 如果下一节点=当前值，移除下一节点，进行新一轮比对
		if tmp.Next.Val == tmp.Val {
			tmp.Next = tmp.Next.Next
		} else {
			// 指针后移
			tmp = tmp.Next
		}
	}
	return head
}

func Test83(t *testing.T) {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}}}
	dl := deleteDuplicates(head)
	fmt.Println("-------")
	for dl != nil {
		fmt.Println(dl)
		dl = dl.Next
	}
}
