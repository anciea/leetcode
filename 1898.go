package main

func maximumRemovals(s string, p string, removable []int) int {
	l,r := 0, len(removable)+1
	for l+1 < r {
		m := l + (r-l) / 2
		if checkMaximun(s, p, removable[:m]) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func checkMaximun(s string, p string, removable []int) bool {
	mp := make([]bool, len(s)+1)
	for _, r := range removable {
		mp[r] = true
	}
	i,n := 0,len(s)
	for _,x := range p {
		for (i < n && mp[i]) || (i < n && rune(s[i]) != x) {
			i++
		}
		if i == n {
			return false
		}
		i++
	}
	return true
}
