package main


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

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, 0)
	var batchHead, preBack, ans *ListNode
	ans = head

	for head != nil {
		cnt := head.Next
		stack = append(stack, head)

		if len(stack) == k {
			for len(stack) > 0 {
				back := stack[len(stack)-1]
				if len(stack) == k {
					batchHead = back
					if preBack != nil {
						preBack.Next = batchHead
					} else {
						ans = batchHead
					}
				}
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					preBack = back
					preBack.Next = nil 
					break
				}
				back.Next = stack[len(stack)-1]
			}
		}
		head = cnt
	}

	for _, i := range stack {
		if preBack != nil {
			preBack.Next = i
			preBack = i
		}
	}

	return ans
}
