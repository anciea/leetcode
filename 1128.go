package main

func numEquivDominoPairs(dominoes [][]int) int {
	type pair struct {
		a, b int
	}
	vis := map[pair]int{}
	
	for _, dom := range dominoes {
		a, b := dom[0], dom[1]
		if a > b {
			a, b = b, a
		}
		vis[pair{a, b}]++
	}
	
	ans := 0
	for _, cnt := range vis {
		ans += cnt * (cnt - 1) / 2
	}
	
	return ans
}
