package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	rt := make([][]int, n)
	for i := range rt {
		rt[i] = make([]int, 0)
	}

	for _, e := range edges {
		rt[e[0]] = append(rt[e[0]], e[1])
		rt[e[1]] = append(rt[e[1]], e[0])
	}
	vis := make([]bool, n)
	var dfs func(i int) int

	ans := 0

	dfs = func(i int) int{
		vis[i] = true
		sum := values[i]
		for _, j := range rt[i] {
			if vis[j] {
				continue
			}
			sum += dfs(j)
		}
		if sum % k == 0 {
			ans++
		}
		return sum
	}
	dfs(0)

	return ans
}
