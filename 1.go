package main

func twoSum(nums []int, target int) []int {
	vis := map[int]int{}
	for i,x := range nums {
		if j, ok := vis[target-x]; ok {
			return []int{i, j}
		}
		vis[x] = i
	}
	return []int{}
}
