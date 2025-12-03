package main

import "sort"

func findKthNumber(m int, n int, k int) int {
	return sort.Search(n*m, func(x int) bool {
		cnt := 0
		for i := 1; i <= m; i++ {
			cnt += min(n, x/i)
		}
		return cnt >= k
	})
}
