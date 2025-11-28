package main

func maximumCandies(candies []int, k int64) int {
	l, r := 0, 10000000+1
	for l+1 < r {
		mid := l + (r-l)/2
		if check(candies, mid, k) {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

func check(candies []int, mid int, k int64) bool {
	ans := int64(0)
	for _, can := range candies {
		ans += int64(can / mid)
	}	
	return ans >= k
}
