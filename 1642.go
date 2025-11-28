package main

import "sort"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	l,r := 0, len(heights)
	for l+1 < r {
		m := l + (r-l)/2
		if checkFurthestBuilding(heights,bricks,ladders,m) {
			l = m
		} else {
			r = m
		}
	}
	return l
}
type pr struct {
	first int
	second int
}

func checkFurthestBuilding(heights []int, bricks int, ladders int, m int) bool {
	top := []pr{}
	for i := 1;i <= m; i++ {
		if heights[i] > heights[i-1] {
			top = append(top,pr{heights[i]-heights[i-1],i})
		}
	}
	sort.Slice(top,func(a,b int) bool {
		return top[a].first > top[b].first
	})
	// fmt.Println(top)
	useLadder := make([]bool, len(heights)+1)
	for i := 0; i < min(ladders,len(top)); i++ {
		useLadder[top[i].second] = true
	}
	for i := 1; i <= m; i++ {
		if !useLadder[i] && heights[i] > heights[i-1]{
			bricks -= (heights[i] - heights[i-1])
		}
		if bricks < 0 {
			return false
		}
	}
	return true
}
