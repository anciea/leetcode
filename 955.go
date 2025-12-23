package main


// func minDeletionSize(strs []string) (res int) {
// 	m,n := len(strs[0]), len(strs)
// 	a := make([]string, n) // 最终得到的字符串数组

// 	next:
// 	for i := range m {
// 		for j := range n-1 {
// 			if a[i] + string(strs[j][i]) < a[i+1]+string(strs[j+1][i]) {
// 				res++
// 				continue next
// 			}
// 		}

// 		for j, s := range strs {
// 			a[j] += string(s[i])
// 		}
// 	}
// 	return res
// }

