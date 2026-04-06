package main

import (
	"slices"
	"sort"
)
/*
写复杂了..
步数不超过9
*/

func robotSim(commands []int, obstacles [][]int) int {
	// 北=0, 东=1, 南=2, 西=3
	// XX := []int{0, 1, 0, -1}
	// YY := []int{1, 0, -1, 0}

	start := 0 
	mpX := make(map[int][]int)
	mpY := make(map[int][]int)

	for i := range obstacles {
		item := obstacles[i]
		// mpX 记录：同一 X 坐标下，存在障碍物的所有 Y 坐标
		mpX[item[0]] = append(mpX[item[0]], item[1])
		// mpY 记录：同一 Y 坐标下，存在障碍物的所有 X 坐标
		mpY[item[1]] = append(mpY[item[1]], item[0])
	}

	for _, v := range mpX {
		sort.Ints(v)
	}
	for _, v := range mpY {
		sort.Ints(v)
	}

	now := []int{0, 0}
	ans := 0 

	for _, item := range commands {
		if item == -2 {
			start = (start + 3) % 4
		} else if item == -1 {
			start = (start + 1) % 4
		} else if item >= 1 && item <= 9 {
			if start == 0 { // 面向北 (+Y方向)
				targetY := now[1] + item
				obs := mpX[now[0]] // 拿出同一X轴上的所有障碍物
				// 找第一个大于当前 Y 的障碍物
				idx, _ := slices.BinarySearch(obs, now[1]+1)
				if idx < len(obs) && obs[idx] <= targetY {
					now[1] = obs[idx] - 1 // 挡住了，停在障碍物前一格
				} else {
					now[1] = targetY      // 没挡住，到达目标
				}
			} else if start == 2 { // 面向南 (-Y方向)
				targetY := now[1] - item
				obs := mpX[now[0]]
				// 找当前位置后方的第一个障碍物
				idx, _ := slices.BinarySearch(obs, now[1])
				if idx > 0 && obs[idx-1] >= targetY {
					now[1] = obs[idx-1] + 1 // 挡住了，停在障碍物前一格
				} else {
					now[1] = targetY
				}
			} else if start == 1 { // 面向东 (+X方向)
				targetX := now[0] + item
				obs := mpY[now[1]] // 拿出同一Y轴上的所有障碍物
				idx, _ := slices.BinarySearch(obs, now[0]+1)
				if idx < len(obs) && obs[idx] <= targetX {
					now[0] = obs[idx] - 1
				} else {
					now[0] = targetX
				}
			} else if start == 3 { // 面向西 (-X方向)
				targetX := now[0] - item
				obs := mpY[now[1]]
				idx, _ := slices.BinarySearch(obs, now[0])
				if idx > 0 && obs[idx-1] >= targetX {
					now[0] = obs[idx-1] + 1
				} else {
					now[0] = targetX
				}
			}
			ans = max(ans, simNow(now[0], now[1]))
		}
	}

	return ans
}

func simNow(x, y int) int {
	return x*x + y*y
}

// 注：如果你的 Go 版本低于 1.21，用自己写的 max 即可
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
