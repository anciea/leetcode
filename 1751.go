package main

import (
	"sort"
)

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func (i,j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		st := events[i-1][0]
		value := events[i-1][2]
		idx := sort.Search(i-1, func(j int)bool {
			return events[j][1] >= st
		})

		for j := 1; j <= k; j++ {
			dp[i][j] = max(dp[i-1][j], dp[idx][j-1] + value)
		}

	}

	return dp[n][k]
}
