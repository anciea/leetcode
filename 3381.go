package main

import "math"

func maxSubarraySum(nums []int, k int) int64 {
  sum := make([]int, len(nums)+1)

	for i,x := range nums {
		sum[i+1] = sum[i] + x
	}


	minS := make([]int, k)
	for i := range minS {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	for j, s := range sum {
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}
