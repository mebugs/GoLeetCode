package reverseList

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func reverseList1(head *ListNode) *ListNode {
	// 返回结尾节点
	if head == nil || head.Next == nil {
		return head
	}
	// 递归翻转
	newHead := reverseList(head.Next)
	// 下下个节点等于当前（实际翻转）
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 双指针迭代
func reverseList(head *ListNode) *ListNode {
	// 头指针
	var prev *ListNode
	// 执行指针
	curr := head
	// 缓存数据
	var tmp *ListNode
	// 执行指针后移
	for curr != nil {
		// 记录后续护具
		tmp = curr.Next
		// 当前节点指向前一节点
		curr.Next = prev
		// 头指针后移（头指针先移，否则执行指针会变为nil）
		prev = curr
		// 执行指针后移
		curr = tmp
	}
	return prev
}

func Test206(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}}
	dl := reverseList(head)
	fmt.Println("-------")
	for dl != nil {
		fmt.Println(dl)
		dl = dl.Next
	}
}
