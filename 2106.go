package main

import (
)

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	const N = 200010
	a := make([]int, N)

	// 映射水果位置
	for _, t := range fruits {
		pos, amount := t[0], t[1]
		a[pos+1] = amount
	}

	startPos += 1
	n := 200001

	// 前缀和
	for i := 1; i <= n; i++ {
		a[i] += a[i-1]
	}

	ans := 0
	for i := 0; i <= k; i++ {
		right := startPos + i
		if right > n {
			right = n
		}

		j := k - 2*i
		if j < (k-i)/2 {
			j = (k - i) / 2
		}
		if j < 0 {
			j = 0
		}

		left := startPos - j
		if left < 1 {
			left = 1
		}

		cur := a[right] - a[left-1]
		if cur > ans {
			ans = cur
		}
	}

	return ans
}
