package main

func numMagicSquaresInside(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 || n < 3 || m < 3 {
		return 0
	}
	ans := 0
	var check func(i, j int) bool
	check = func(a, b int) bool {
		if grid[a+1][b+1] != 5 {
			return false
		}
		seen := [10]bool{}
		for i := a; i <= a+2; i++ {
			for j := b; j <= b+2; j++ {
				v := grid[i][j]
				if v < 1 || v > 9 || seen[v] {
					return false
				}
				seen[v] = true
			}
		}
		if grid[a][b]+grid[a][b+1]+grid[a][b+2] != 15 {
			return false
		}
		if grid[a+1][b]+grid[a+1][b+1]+grid[a+1][b+2] != 15 {
			return false
		}
		if grid[a+2][b]+grid[a+2][b+1]+grid[a+2][b+2] != 15 {
			return false
		}
		if grid[a][b]+grid[a+1][b]+grid[a+2][b] != 15 {
			return false
		}
		if grid[a][b+1]+grid[a+1][b+1]+grid[a+2][b+1] != 15 {
			return false
		}
		if grid[a][b+2]+grid[a+1][b+2]+grid[a+2][b+2] != 15 {
			return false
		}
		if grid[a][b]+grid[a+1][b+1]+grid[a+2][b+2] != 15 {
			return false
		}
		if grid[a][b+2]+grid[a+1][b+1]+grid[a+2][b] != 15 {
			return false
		}
		return true
	}
	for i := 0; i < n-2; i++ {
		for j := 0; j < m-2; j++ {
			if check(i, j) {
				ans++
			}
		}
	}
	return ans
}
