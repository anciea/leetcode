package main

func numberOfPaths(grid [][]int, k int) int {
  vis := make([][][]int, len(grid))
  for i := range vis {
    vis[i] = make([][]int, len(grid[0]))
    for j := range vis[i] {
      vis[i][j] = make([]int, k)
    }
  }
	n,m := len(grid), len(grid[0])

	vis[0][0][grid[0][0]%k] = 1

	for i := 0;i < n; i++ {
		for j := 0;j < m; j++ {
			for p := 0;p < k; p++ {
				if i > 0 && vis[i-1][j][p] > 0 {
					vis[i][j][(p+grid[i][j])%k] = (vis[i][j][(p+grid[i][j])%k] + vis[i-1][j][p]) % 1000000007
				}
				if j > 0 && vis[i][j-1][p] > 0 {
					vis[i][j][(p+grid[i][j])%k] = (vis[i][j][(p+grid[i][j])%k] + vis[i][j-1][p]) % 1000000007
				}
			}
		}
	}

	return vis[n-1][m-1][0] % 1000000007
}
