package main

func minimizedMaximum(n int, quantities []int) int {
  l, r := 0, 1000007
  for l+1 < r {
    m := l + (r-l)/2
    if checkMinimizedMaximum(n, quantities, m) {
      r = m
    } else {
      l = m
    }
  }
  return r
}

func checkMinimizedMaximum(n int, quantities []int, m int) bool {
	ans := 0
	for _, quantity := range quantities {
		ans += (quantity + m - 1) / m
	}
	return ans <= n
}
