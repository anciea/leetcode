package main

import "slices"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	vis := make([]int, len(positions))
	for i := range vis {
		vis[i] = i
	}

	slices.SortFunc(vis, func(i, j int) int { return positions[i] - positions[j] })
	stack := make([]int, 0)

	for _, v := range vis {
		// no.v 号机器人
		if directions[v] == 'R' {
			stack = append(stack, v)
		} else {
			for len(stack) > 0 && healths[v] > 0 {
				top := stack[len(stack)-1]
				if healths[v] < healths[top] {
					healths[top]--
					healths[v] = 0
				} else if healths[v] > healths[top] {
					stack = stack[:len(stack)-1]
					healths[v]--
					healths[top] = 0
				} else {
					stack = stack[:len(stack)-1]
					healths[v] = 0
					healths[top] = 0
				}
			}
			// stack = append(stack, v)
		}
	}

	ans := make([]int, 0)
	for _, i := range healths {
		if i > 0 {
			ans = append(ans, i)
		}
	}

	return ans
}
