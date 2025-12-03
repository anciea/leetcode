package main

func containsNearbyDuplicate(nums []int, k int) bool {
	vis := map[int]int{}
	for i,x := range nums {
		if e, ok := vis[x]; ok && i-e <= k {
			return true
		}
		vis[x] = i
	}
	return false
}
