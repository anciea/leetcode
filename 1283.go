package main

// func smallestDivisor(nums []int, threshold int) int {
// 	l, r := 1, 1000000000
// 	for l < r {
// 		m := (l + r) >> 1
// 		if check(nums, threshold, m) {
// 			r = m
// 		} else {
// 			l = m + 1
// 		}
// 	}
// 	return l
// }

// func check(nums []int, threshold int, m int) bool {
// 	sum := 0
// 	for _, x := range nums {
// 		sum += (x + m - 1) / m
// 	}
// 	return sum <= threshold
// }
