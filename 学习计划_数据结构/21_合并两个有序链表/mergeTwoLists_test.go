package mergeTwoLists

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归切换指针
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func Test21(t *testing.T) {
	ln := mergeTwoLists(&ListNode{1, &ListNode{3, &ListNode{4, nil}}}, &ListNode{1, &ListNode{2, &ListNode{2, nil}}})
	fmt.Println("-------")
	for ln != nil && ln.Next != nil {
		fmt.Println(ln)
		ln = ln.Next
	}
	ln = mergeTwoLists(nil, &ListNode{1, &ListNode{2, &ListNode{2, nil}}})
	fmt.Println("-------")
	for ln != nil && ln.Next != nil {
		fmt.Println(ln)
		ln = ln.Next
	}
	ln = mergeTwoLists(nil, nil)
	fmt.Println("-------")
	for ln != nil && ln.Next != nil {
		fmt.Println(ln)
		ln = ln.Next
	}
	ln = mergeTwoLists(&ListNode{1, &ListNode{3, &ListNode{4, nil}}}, nil)
	fmt.Println("-------")
	for ln != nil && ln.Next != nil {
		fmt.Println(ln)
		ln = ln.Next
	}
}
