package main

import (
	"cmp"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)
func minimumPairRemoval(nums []int) (ans int) {
	n := len(nums)
	type pair struct{ s, i int }
	pairs := redblacktree.NewWith[pair, struct{}](func(x, y pair) int {return cmp.Or(x.s-y.s, x.i-y.i)})
	dec := 0

	for i := range n-1 {
		x,y := nums[i], nums[i+1]
		if x > y {
			dec++
		}
		pairs.Put(pair{x+y, i}, struct{}{})
	}

	idx := redblacktree.New[int, struct{}]()
	for i := range n {
		idx.Put(i, struct{}{})
	}
	for dec > 0 {
		ans++

		it := pairs.Left()
		s := it.Key.s
		i := it.Key.i
		pairs.Remove(it.Key)
		
		node, _ := idx.Ceiling(i+1)
		nxt := node.Key
		if nums[i] > nums[nxt] {
			dec--
		}

		node, _ = idx.Floor(i-1)
		if node != nil {
			prev := node.Key
			if nums[prev] > nums[i] {
				dec--
			}
			if nums[prev] > s {
				dec++
			}
			pairs.Remove(pair{nums[prev] + nums[i], prev})
			pairs.Put(pair{nums[prev] + s, prev}, struct{}{})
		}

		node, _ = idx.Ceiling(nxt+1)
		if node != nil {
			nxt2 := node.Key
			if nums[nxt] > nums[nxt2] {
				dec--
			}
			if s > nums[nxt2] {
				dec++
			}
			pairs.Remove(pair{nums[nxt] + nums[nxt2], nxt})
			pairs.Put(pair{s + nums[nxt2], i}, struct{}{})
		}

		nums[i] = s
		idx.Remove(i)
	}
	return
}
