package main

// import (
//     "math"
// )

// func maximumAmount(coins [][]int) int {
// 	n := len(coins[0])
// 	f := make([][3]int, n+1)
// 	for j := range f {
// 		f[j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
// 	}
// 	f[1] = [3]int{}
// 	for _, row := range coins {
// 		for j, x := range row {
// 			f[j+1][2] = max(f[j][2]+x, f[j+1][2]+x, f[j][1], f[j+1][1])
// 			f[j+1][1] = max(f[j][1]+x, f[j+1][1]+x, f[j][0], f[j+1][0])
// 			f[j+1][0] = max(f[j][0], f[j+1][0]) + x
// 		}
// 	}
// 	return f[n][2]
// }


/**
DFS超时了。。。
**/

// func maximumAmount(coins [][]int) int {
// 	n, m := len(coins), len(coins[0])
// 	var dfs func(i, j, coin int, min1, min2 int)
// 	xx := []int{1, 0}
// 	yy := []int{0, 1}

// 	ans := math.MinInt64
// 	min1, min2 := math.MaxInt, math.MaxInt
// 	dfs = func(i, j, coin int, min1, min2 int) {
// 		x := coins[i][j]
// 		coin += x
// 		if x < 0 {
// 			if x < min1 {
// 				min2 = min1
// 				min1 = x
// 			} else if x < min2 {
// 				min2 = x
// 			}
// 		}
// 		if i == n-1 && j == m-1 {
// 			if min1 != math.MaxInt {
// 				coin += -min1
// 			}
// 			if min2 != math.MaxInt {
// 				coin += -min2
// 			}
// 			ans = max(ans, coin)
// 			return
// 		}
// 		// if ans <= coins {
// 		//     return
// 		// }
// 		for p := 0; p < 2; p++ {
// 			ix := i + xx[p]
// 			iy := j + yy[p]
// 			if ix < n && iy < m {
// 				dfs(ix, iy, coin, min1, min2)
// 			}
// 		}
// 	}
// 	dfs(0, 0, 0, min1, min2)
// 	return ans

// }
