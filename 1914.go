package main

var dirs = [4][2]int{{0,1}, {1,0}, {0,-1},{-1,0}}

func rotateGrid(grid [][]int, k int) [][]int {
  m0, n0 := len(grid), len(grid[0])
  a := make([]int, (m0+n0-2)*2)

  for i := range min(m0, n0)/2{
    m, n := m0-i*2, n0-i*2
    // 开始坐标
    x, y := i, i
    a = a[:0]
    for _, dir := range dirs {
      for range n-1 {
        a = append(a, grid[x][y])
        x += dir[0]
        y += dir[1]
      }
      n, m = m, n //拐弯
    }

    shift := k % len(a)
    a = append(a[shift:], a[:shift]...)

    j := 0
    for _, dir := range dirs {
      for range n-1 {
        grid[x][y] = a[j]
        j++
        x += dir[0]
        y += dir[1]
      }
      n, m = m, n
    }
  }
  return grid
}


/*
0,0 0,1 0,2 0,3
1,0 1,1 1,2 1,3
2,0 2,1 2,2 2,3
3,0 3,1 3,2 3,3
*/
