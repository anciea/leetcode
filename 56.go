package main

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i, j := 0, 1
	for j < n {
		if intervals[i][1] >= intervals[j][0] {
			intervals[i][1] = max(intervals[i][1], intervals[j][1])
		} else {
			i++
			intervals[i] = intervals[j]
		}
		j++
	}
	
	return intervals[:i+1]
}
