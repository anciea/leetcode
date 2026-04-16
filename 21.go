package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for a != nil || b != nil {
		if a == nil {
			curr.Next = b
			break
		} else if b == nil {
			curr.Next = a
			break
		} else {
			if a.Val < b.Val {
				curr.Next = a
				a = a.Next
			} else {
				curr.Next = b
				b = b.Next
			}
		}
		curr = curr.Next
	}
	return head.Next
}

func main21() {
	// 构造链表 1->2->4
	a := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// 构造链表 1->3->4
	b := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	result := mergeTwoLists(a, b)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
