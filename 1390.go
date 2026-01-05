package main

import "math"

func sumFourDivisors(nums []int) (ans int) {
	for i := range nums {
		n := nums[i]
		limit := int(math.Sqrt(float64(n)))
		found := false
		sum := 0
		ok := true
		for j := 2; j <= limit; j++ {
			if n%j != 0 {
				continue
			}
			other := n / j
			if other == j {
				ok = false
				break
			}
			if found {
				ok = false
				break
			}
			found = true
			sum = 1 + j + other + n
		}
		if ok && found {
			ans += sum
		}
	}
	return
}
