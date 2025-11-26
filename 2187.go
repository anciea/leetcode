package main

func minimumTime(time []int, totalTrips int) int64 {
	l, r := 1, int(1e9)
	for l < r {
		m := (l + r) >> 1
		if check(time, totalTrips, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}

func check(time []int, totalTrips int, m int) bool {
	sum := 0
	for _, x := range time {
		sum += m / x
	}
	return sum >= totalTrips
}
