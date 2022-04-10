package removeElements

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归法
func removeElements1(head *ListNode, val int) *ListNode {
	// 这是最后一个节点
	if head == nil {
		return head
	}
	// 递归进入后续处理
	head.Next = removeElements(head.Next, val)
	// 移除当前节点
	if head.Val == val {
		head = head.Next
	}
	return head
}

// 迭代法
func removeElements(head *ListNode, val int) *ListNode {
	// 哨兵节点
	tmpHead := &ListNode{Next: head}
	// 便利链表
	for tmp := tmpHead; tmp.Next != nil; {
		// 判断next指针是否后移
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return tmpHead.Next
}

func Test203(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}}
	dl := removeElements(head, 3)
	fmt.Println("-------")
	for dl != nil {
		fmt.Println(dl)
		dl = dl.Next
	}
}
