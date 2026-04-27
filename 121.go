package main

import "math"

func maxProfit(prices []int) int {
  minn := math.MaxInt32
  ans := 0
  for _, x := range prices {
    ans = max(ans, x-minn)
    minn = min(minn, x)
  }
  
  return ans
}
