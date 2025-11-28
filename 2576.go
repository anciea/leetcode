package main

import (
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	l, r := 0, len(nums)/2+1
	sort.Ints(nums)

	for l+1 < r {
		mid := l + (r-l)/2
		if checkMaxNum(nums, mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l * 2 
}

func checkMaxNum(nums []int, mid int) bool {
	// 考虑o(n) 计算出每次ans
	n := len(nums)
	for i := 0; i < mid; i++ {
		if nums[i] *2 > nums[n-mid+i] {
			return false
		}
	}
	return true
}
