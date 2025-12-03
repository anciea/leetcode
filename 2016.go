package main

import "math"

func maximumDifference(nums []int) int {
	minn := math.MaxInt
	ans := 0
	for _, i := range nums {
		ans = max(ans, i-minn)
		minn = min(minn, i)
	}
	if ans == 0 {
		ans = -1
	}
	return ans
}
