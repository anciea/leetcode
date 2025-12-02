package main

func countTrapezoids(points [][]int) int {
	pointNum := make(map[int]int)
	mod := 1000000007

	for _, point := range points {
		pointNum[point[1]]++
	}
	ans, sum := 0, 0
	for _, num := range pointNum {
		edge := num * (num - 1) / 2
		ans = (ans + sum * edge) % mod
		sum = (sum + edge) % mod
	}
	return ans
}
