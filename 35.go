package main

import "sort"

func searchInsert(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return sort.SearchInts(nums, target)
}
