package main

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	l, r := 0, int(2e9)
	for l+1 < r {
		m := l + (r-l)/2
		if checkMaxNumberOfAlloys(n, k, budget, composition, stock, cost, m) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func checkMaxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int, m int) bool {
	for i := 0; i < k; i++ {
		sum := 0
		flag := true
		for j := 0; j < n; j++ {
			p := stock[j] - composition[i][j]*m
			if p < 0 {
				sum += (-p * cost[j])
			}
			if sum > budget {
				flag = false
			}
		}
		if flag {
			return true
		}
	}
	return false
}
