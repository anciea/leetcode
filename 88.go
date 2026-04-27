package main

func merge(a []int, m int, b []int, n int) {
	ans := make([]int, 0, m+n)
	i, j := 0, 0

	for i < m || j < n {
		if i == m {
			ans = append(ans, b[j:]...)
			break
		} else if j == n {
			ans = append(ans, a[i:]...)
			break
		}

		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}

	copy(a, ans)
}
