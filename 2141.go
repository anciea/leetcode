package main

import "sort"

func maxRunTime(n int, batteries []int) int64 {
	if len(batteries) < n {
		return 0
	}
	sort.Slice(batteries, func (a,b int) bool{
		return batteries[a] > batteries[b]
	})

	sum := 0
	vis := make([]int64, len(batteries))
	for i := range batteries {
		sum += batteries[i]
		if i == 0 {
			vis[i] = int64(batteries[i])
		} else {
			vis[i] = int64(batteries[i]) + vis[i-1]
		}
	}



	l, r := int64(0), int64((sum+n-1)/n+1)
	for l+1 < r {
		m := l+(r-l)/2
		if checkMaxRunTime(n, batteries, vis, m) {
			l = m
		} else {
			r = m
		}
	}

	return l
}

func checkMaxRunTime(n int, batteries []int, vis []int64, m int64) bool {
	sum := vis[len(vis)-1]-vis[n-1]
	for i := 0; i < n; i++ {
		if int64(batteries[i]) >= m {
			continue
		}
		if int64(batteries[i]) < m && sum + int64(batteries[i]) >= m {
			sum -= int64(m - int64(batteries[i]))
		} else {
			return false
		}
	}
	return true
}
