package main

import "math"

func minimumCardPickup(cards []int) int {
	vis := map[int]int{}
	ans := math.MaxInt
	for i,x := range cards {
		if e, ok := vis[x]; ok {
			ans = min(ans, i-e+1)
		}
		vis[x] = i
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
