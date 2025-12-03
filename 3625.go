package main

import "math"

func countTrapezoids(points [][]int) int {
	cnt := map[float64]map[float64]int{} //斜率 截距 个数
	type xyP struct {
		x,y int
	}
	cnt2 := map[xyP]map[float64]int{} // 中点 斜率 个数
	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		for j := i+1; j < len(points); j++ {
			x2, y2 := points[j][0], points[j][1]
			dx, dy := x-x2, y-y2
			k := math.MaxFloat64
			b := float64(x)
			if dx != 0 {
				k = float64(dy)/float64(dx)
				b = float64(y*dx-x*dy)/float64(dx)
			}
			if _, ok := cnt[k]; !ok {
				cnt[k] = map[float64]int{}
			}
			cnt[k][b]++
			
			mid := xyP{x: x+x2, y: y+y2}
			if _, ok := cnt2[mid]; !ok {
				cnt2[mid] = map[float64]int{}
			}
			cnt2[mid][k]++
		}
	}

	ans := 0

	for _, v := range cnt {
		s := 0
		for _, v2 := range v {
			ans += v2 * s
			s += v2
		}
	}

	for _, v := range cnt2 {
		s := 0
		for _, v2 := range v {
			ans -= s * v2
			s += v2
		}
	}

	return ans
}
