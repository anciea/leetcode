package main

import "math"

func minEatingSpeed(piles []int, h int) int {
  l,r := 1, math.MaxInt32
	for l < r {
		m := (l + r) >> 1
		if check(piles, h, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func check(piles []int, h int, k int) bool {
	ans := 0
	for _, v := range piles {
		ans += (v+k-1) / k
	}
	return ans <= h
}
