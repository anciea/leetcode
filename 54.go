package main


// var dir = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// func spiralOrder(matrix [][]int) []int {
// 	n, m := len(matrix), len(matrix[0])
// 	ans := make([]int, 0)
// 	x, y, d := 0, 0, 0

// 	for i := 0; i < n*m; i++ {
// 		ans = append(ans, matrix[x][y])
// 		matrix[x][y] = 101
// 		if check(x+dir[d][0], y+dir[d][1], n, m, matrix) {
// 			x += dir[0][0]
// 			y += dir[0][1]
// 		} else {
// 			d = (d + 1) % 4
// 			x += dir[d][0]
// 			y += dir[d][1]
// 		}
// 	}
// 	return ans
// }

// func check(x, y, n, m int, matrix [][]int) bool {
// 	return x >= 0 && y >= 0 && x < n && y < m && matrix[x][y] <= 100
// }
