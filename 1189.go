package main

import "math"

func maxNumberOfBalloons(text string) int {
	mp := make(map[byte]int)
	for i := range text {
		mp[text[i]]++
	}

	balloon := "balon"
	ans := math.MaxInt32
	for i := range balloon {
		if balloon[i] == 'b' ||balloon[i] == 'a' ||balloon[i] == 'n'{
			ans = min(mp[balloon[i]], ans)
		}
		if balloon[i] == 'l' || balloon[i] =='o' {
			ans = min(mp[balloon[i]]/2, ans)
		}
	}
	return ans
}
