package main

import "sort"

/*
时间复杂度：	O(N^2)
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	mp := make(map[int]int, 0)
	for _, i := range nums {
		mp[i]++
	}
	ans := make([][]int, 0)
	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			continue
		}
		mp[x]--

		subNums := nums[i+1:]
		for j, y := range subNums {
			if j > 0 && y == subNums[j-1] {
				continue
			}

			z := -x - y
			if y > z {
				break
			}

			mp[y]--
			if mp[z] > 0 {
				ans = append(ans, []int{x, y, z})
			}
			mp[y]++
		}
		mp[x]++
	}

	return ans
}
*/

func threeSum(a []int) [][]int {
	sort.Ints(a)
	n := len(a)
	ans := make([][]int, 0)
	for i := 0; i < len(a)-2; i++ {
		if a[i] > 0 {
			break
		}
		if i > 0 && a[i] == a[i-1] {
			continue
		}

		cnt := -a[i]
		l, r := i+1, n-1
		for l < r {
			p := a[l] + a[r]
			if cnt == p {
				ans = append(ans, []int{a[i], a[l], a[r]})
				l++
				r--
				for l < r && a[l] == a[l-1] {
					l++
				}
				for l < r && a[r] == a[r+1] {
					r--
				}
			} else if cnt < p {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
