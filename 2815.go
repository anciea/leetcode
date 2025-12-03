package main

import (
	"math"
	"strconv"
)

func maxSum(nums []int) int {
	ans := math.MinInt
	vis := [10]int{} // digits 0-9

	for _, x := range nums {
		str := strconv.Itoa(x)
		maxDigit := 0
		for i := range str {
			digit := int(str[i] - '0')
			if digit > maxDigit {
				maxDigit = digit
			}
		}
		if prev := vis[maxDigit]; prev != 0 {
			sum := x + prev
			ans = max(ans, sum)
		}
		vis[maxDigit] = max(vis[maxDigit], x)
	}
	if ans == math.MinInt {
		return -1
	}
	return ans
}
