package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	p0 := dummy
	for i := 0; i < left-1; i++ {
		_ = i
		p0 = p0.Next
	}
	var pre *ListNode
	curr := p0.Next

	for i := 0; i < right-left+1; i++ {
		_ = i
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	p0.Next.Next = curr
	p0.Next = pre

	return dummy.Next
}
