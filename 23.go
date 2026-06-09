package main

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type DumpList []*ListNode
func(d *DumpList) Push(k any) {*d = append(*d, k.(*ListNode))}
func(d *DumpList) Pop() any{
	n := len(*d)
	x := (*d)[n-1]
	*d = (*d)[:n-1]
	return x
}
func(d DumpList) Len()int {return len(d)}
func(d DumpList) Less(i, j int) bool{ return d[i].Val < d[j].Val}
func(d DumpList) Swap(i, j int) {d[i], d[j] = d[j], d[i]}

func mergeKLists(lists []*ListNode) *ListNode {
	d := &DumpList{}
	heap.Init(d)
	
	for i := range lists {
		if lists[i] != nil {
			heap.Push(d, lists[i])
		}
	}

	ans := &ListNode{}
	root := ans

	for d.Len() > 0 {
		node := heap.Pop(d).(*ListNode)
		root.Next = node
		root = root.Next
		if node.Next != nil {
			heap.Push(d, node.Next)
		}
	}
	return ans.Next
}
