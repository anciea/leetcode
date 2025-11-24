package main


func rangeAddQueries(n int, queries [][]int) [][]int {
	ans := make([][]int, n+10)
	for i := range ans {
		ans[i] = make([]int, n+10)
	}
	for _, q := range queries {
		x1,y1,x2,y2 := q[0]+1,q[1]+1,q[2]+1,q[3]+1
		ans[x1][y1]++
		ans[x1][y2+1]--
		ans[x2+1][y1]--
		ans[x2+1][y2+1]++
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ans[i][j] = ans[i][j-1] + ans[i-1][j] - ans[i-1][j-1] + ans[i][j]
		}
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = ans[i+1][j+1]
		}
	}

	return result
}
