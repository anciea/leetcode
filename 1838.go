package main

import "slices"

// 1838. 最高频元素的频数
// https://leetcode.cn/problems/frequency-of-the-most-frequent-element/
func maxFrequency(nums []int, k int) int {
	slices.Sort(nums)
	ans := 1
	for l, r , total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}
