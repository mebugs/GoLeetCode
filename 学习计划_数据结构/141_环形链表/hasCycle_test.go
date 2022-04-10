package hasCycle

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 定义一个内存地址的Hash
	addsMap := map[string]bool{fmt.Sprint(head): true}
	node := head.Next
	for {
		if node == nil {
			return false
		}
		_, ok := addsMap[fmt.Sprint(node)]
		if ok {
			return true
		} else {
			addsMap[fmt.Sprint(node)] = true
			node = node.Next
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 快指针比慢指针先出发
	slow, fast := head, head.Next
	// 未相遇
	for fast != slow {
		// 快指针是不是到头了
		if fast == nil || fast.Next == nil {
			return false
		}
		// 慢指针走一步
		slow = slow.Next
		// 快指针走两步
		fast = fast.Next.Next
	}
	return true
}

func Test141(t *testing.T) {
	round := &ListNode{4, &ListNode{}}
	nodes := &ListNode{1, &ListNode{2, &ListNode{3, round}}}
	round.Next = nodes
	t.Log(hasCycle(nodes))
}
