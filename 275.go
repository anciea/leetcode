package main

func hIndex(citations []int) int {
	check := func(mid int) bool {
		ans := 0
		for _,i := range citations {
			if i >= mid {
				ans++
			}
		}
		return ans >= mid
	}

	l, r := 0, len(citations)+1
	for l+1 < r {
		mid := l + (r - l) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}
