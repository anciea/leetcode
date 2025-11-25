package main

// type pair struct {
// 	x, y int
// }

// func findMaxForm(strs []string, m int, n int) int {
// 	pairs := make([]pair, 0)
// 	for _, s := range strs {
// 		x, y := 0, 0
// 		for _, i := range s {
// 			if i == '0' {
// 				x++
// 			} else {
// 				y++
// 			}
// 		}
// 		pairs = append(pairs, pair{x, y})
// 	}

// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}

// 	for _, k := range pairs {
// 		for i := m; i >= k.x; i-- {
// 			for j := n; j >= k.y; j-- {
// 				dp[i][j] = max(dp[i][j], dp[i-k.x][j-k.y]+1)
// 			}
// 		}
// 	}
// 	return dp[m][n]
// }
