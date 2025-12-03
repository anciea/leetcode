package main

func numIdenticalPairs(nums []int) int {
	vis := map[int]int{}
	ans := 0

	for _,i := range nums {
		ans += vis[i]
		vis[i]++
	}
	return ans
}
