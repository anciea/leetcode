package main

import "sort"

func maximumHappinessSum(happiness []int, k int) (ans int64) {
	sort.Slice(happiness, func (i, j int)  bool {
		return happiness[i] > happiness[j]
	})
	vis := 0
	for i := range happiness {
		if vis >= happiness[i] {
			break
		}
		ans += int64(happiness[i])
		ans -= int64(vis)
		vis++
		if vis >= k {
			break
		}
	}
	return ans
}
