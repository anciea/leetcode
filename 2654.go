package main

func minOperations(nums []int) int {
	n := len(nums)

	// 统计 1 的数量
	cnt1 := 0
	for _, x := range nums {
		if x == 1 {
			cnt1++
		}
	}

	// 若存在 1，传播即可
	if cnt1 > 0 {
		return n - cnt1
	}

	// 检查整体是否有公约数 > 1，如果有则永远无法得到 1
	g := nums[0]
	for _, x := range nums {
		g = gcd(g, x)
	}
	if g > 1 {
		return -1
	}

	// 寻找 gcd == 1 的最短子数组长度
	ans := n + 5
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				if j-i+1 < ans {
					ans = j - i + 1
				}
				break
			}
		}
	}

	// 第一个 1 需要 (ans-1) 步生成，之后传播 (n-1) 步
	return (ans - 1) + (n - 1)
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
