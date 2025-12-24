package main

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for i := range apple {
		sum += apple[i]
	}

	sort.Slice(capacity, func (i,j int) bool {
		return capacity[i] > capacity[j]
	})
	n := len(capacity)
	vis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		vis[i] = vis[i-1] + capacity[i-1]
	}

	idx := sort.Search(n, func (i int)  bool {
		return vis[i] >= sum
	})

	return idx
}
