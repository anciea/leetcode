package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	stack := make([]*ListNode, 0)

	root := head
	for root != nil {
		stack = append(stack, root)
		root = root.Next
	}
	len := len(stack)

	ans := &ListNode{}
	root = ans
	for i := range len {
		if i%2 == 0 {
			root.Next = stack[i/2]
		} else {
			root.Next = stack[len-i/2]
		}
		root = root.Next
		root.Next = nil
	}

	head = ans.Next
}