package main

// func maxValue(n int, index int, maxSum int) int {
// 	l,r := 0,maxSum+1
// 	for l+1 < r {
// 		m := l + (r-l)/2
// 		if checkMaxValue(n,index,maxSum,m) {
// 			l = m
// 		} else {
// 			r = m
// 		}
// 	}
// 	return l
// }


// func checkMaxValue(n int, index int, maxSum int, m int) bool {
// 	left := segmentSum(m-1,index)
// 	right := segmentSum(m-1,n-index-1)
// 	return left + right + m <= maxSum
// }

// func segmentSum (n, l int) int{
// 	ans := 0
// 	if l >= n {
// 		ans += ((1+n)*n)/2
// 		l -= n
// 		ans += l
// 	} else {
// 		ans += ((2*n-l+1)*l)/2
// 	}
// 	return ans
// }
