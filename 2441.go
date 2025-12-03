package main

func findMaxK(nums []int) int {
	vis := make(map[int]bool)
	maxK := -1
	
	for _, num := range nums {
		vis[num] = true
		if num > 0 && vis[-num] {
			if num > maxK {
				maxK = num
			}
		} else if num < 0 && vis[-num] {
			absNum := -num
			if absNum > maxK {
				maxK = absNum
			}
		}
	}
	
	return maxK
}
