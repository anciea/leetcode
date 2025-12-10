package main

const mod = int(1e9) + 7

func countPermutations(complexity []int) int {
	minn := complexity[0]
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= minn {
			return 0
		}
	}
	x := len(complexity) - 1
	return A(x, x) % mod
}

func A(n, m int) int {
    if m < 0 || m > n {
        return 0
    }
    res := int(1)
    for i := int(0); i < m; i++ {
        res = res * (n - i) % mod
    }
    return res
}
