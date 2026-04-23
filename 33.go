package main

import ()

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			//left sorted
			if target <= nums[mid] && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// r sorted
			if target >= nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
  return -1
}
