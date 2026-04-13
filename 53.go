package main

import "slices"

/*
前缀和
*/
// func maxSubArray(nums []int) int {
// 	minPre, preSum, ans := 0, 0, math.MinInt
// 	for _,x := range nums {
// 		preSum += x
// 		ans = max(ans, preSum-minPre)
// 		minPre = min(preSum, minPre)
// 	}
// 	return ans
// }

/*dp*/

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(0, dp[i-1]) + nums[i]
	}

	return slices.Max(dp)
}
