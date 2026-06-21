package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	vis := make(map[*ListNode]bool)
	for head != nil {
		if vis[head] {
			return true
		}
		vis[head] = true
		head = head.Next
	}
	return false
}