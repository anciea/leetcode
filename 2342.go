package main

import (
	"math"
	"strconv"
)

func maximumSum(nums []int) int {
	vis := map[int]int{}
	ans := math.MinInt

	for _, x := range nums {
		digsum := 0
		str := strconv.Itoa(x)
		for i := range str {
			digsum += int(str[i] - '0')
		}
		if e, ok := vis[digsum]; ok {
			sum := x + e
			ans = max(ans, sum)
		}
		vis[digsum] = max(vis[digsum], x)
	}

	if ans == math.MinInt {
		return -1
	}
	return ans
}
