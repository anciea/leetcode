package main
import "math"

func countTriples(n int) int {
	ans := 0
	for i := 1;i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := int(math.Sqrt(float64(i*i + j*j+1)))
			if k*k == i*i + j*j && k <= n {
				ans++
			}
		}
	}
	return ans
}
