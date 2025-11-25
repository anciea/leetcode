package main

func smallestRepunitDivByK(k int) int {
	vis := map[int]bool{}
	x := 1 % k
	for x > 0 && !vis[x] {
		vis[x] = true
		x = (x * 10 + 1) % k
	}
	if x == 0 {
		return len(vis) + 1
	}
	return -1
}
