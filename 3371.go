package main

import "math"

func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		cnt[x]--
		if (total-x)%2 == 0 && cnt[(total-x)/2] > 0 {
			ans = max(ans, x)
		}
		cnt[x]++
	}
	return ans
}
