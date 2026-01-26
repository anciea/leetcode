package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	ans, minDiff := make([][]int, 0), math.MaxInt32
	n := len(arr)
	mp := make(map[int]bool, 0)
	sort.Ints(arr)
	for i := range n-1 {
		if abs(arr[i]-arr[i+1]) < minDiff {
			minDiff = abs(arr[i]-arr[i+1])
		}
		mp[arr[i]] = true
		mp[arr[i+1]] = true
	}
	for k,_ := range mp {
		if ok := mp[k+minDiff]; ok {
			ans = append(ans, []int{k, k+minDiff})
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

