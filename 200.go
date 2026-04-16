package main

func numIslands(grid [][]byte) int {
	xx := []int{0, 0, 1, -1}
	yy := []int{1, -1, 0, 0}
	ans := 0
	vis := make([][]int, len(grid))
	for i := range vis {
		vis[i] = make([]int, len(grid[0]))
	}
	var dfs func(x, y int)

	dfs = func(x, y int) {
		if grid[x][y] == '0' || vis[x][y] == 1 {
			return
		}
		vis[x][y] = 1

		for i := 0; i < 4; i++ {
			px := x + xx[i]
			py := y + yy[i]
			if px >= 0 && px < len(grid) && py >= 0 && py < len(grid[0]) && grid[px][py] == '1'{
				dfs(px, py)
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && vis[i][j] == 0{
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}
