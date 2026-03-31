package main

func checkStrings(s1 string, s2 string) bool {
	mp1 := make(map[byte]int)
	mp2 := make(map[byte]int)
	for i, j := 0, 1; j < len(s1); i, j = i+2, j+2 {
		mp1[s1[j]]++
		mp1[s2[j]]--
		mp2[s1[i]]++
		mp2[s2[i]]--
	}

	if len(s1)%2 == 1  {
		mp2[s1[len(s1)-1]]++
		mp2[s2[len(s1)-1]]--
	}

	for _, v := range mp1 {
		if v != 0 {
			return false
		}
	}

	for _, v := range mp2 {
		if v != 0 {
			return false
		}
	}

	return true
}
