package main

import "slices"

func maxIceCream(costs []int, coins int) int {
	slices.Sort(costs)

	// 按照价格从低到高买
	for i, cost := range costs {
		if coins < cost { // 钱不够
			return i // 买 [0, i-1] 一共 i 根雪糕
		}
		coins -= cost
	}

	// 可以买所有雪糕
	return len(costs)
}
