package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  var pre *ListNode
  if head == nil {
    return nil
  }
  for head.Next != nil {
    next := head.Next
    head.Next = pre
    pre = head
    head = next
  }
  head.Next = pre
  return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
