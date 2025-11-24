package main

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	vis := make([]int, 0)

	for i := range items {
		if i > 0 {
			items[i][1] = max(items[i][1], items[i-1][1])
		}
		vis = append(vis, items[i][0])
	}

	for i,x := range queries {
		l := sort.SearchInts(vis, x+1)
		if l == len(vis) {
			queries[i] = items[len(items)-1][1]
		} else if l == 0 {
			queries[i] = 0
		} else {
			queries[i] = items[l-1][1]
		}
	}

	return queries
}
