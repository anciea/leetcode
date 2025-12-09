package main

func specialTriplets(nums []int) int {
	vis := [100010]int{}
	vis2 := [100010]int{}
	mod := 1000000007

	for _, i := range nums {
		vis[i]++
	}

	ans := 0
	for _, x := range nums {
		vis[x]--

		target := x * 2
		if target < len(vis) && vis2[target] > 0 && vis[target] > 0 {
			cnt := vis2[target] * vis[target]
			ans = (ans + cnt) % mod
		}

		vis2[x]++
	}
	return ans
}
