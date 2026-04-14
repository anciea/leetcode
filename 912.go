package main

// func quick_sort(a []int) []int {
// 	if len(a) <= 1 {
// 		return a
// 	}
// 	l, r, k := 0, len(a)-1, a[0]
// 	for l <= r {
// 		for l < len(a) && a[l] < k {
// 			l++
// 		}
// 		for r > 0 && a[r] > k {
// 			r--
// 		}
// 		if l <= r {
// 			a[r], a[l] = a[l], a[r]
// 			l++
// 			r--
// 		}
// 	}

// 	left := a[:r+1]
// 	right := a[l:]

// 	quick_sort(left)
// 	quick_sort(right)
// 	return a
// }


/*Hoare*/

func quick_sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	l, r := -1, len(a)
	k := a[(len(a)-1)/2]

	for l < r {
		l++ 
		for a[l] < k {
			l++
		}
		r--
		for a[r] > k {
			r--
		}

		if l < r {
			a[l], a[r] = a[r] , a[l]
		}
	}

	left := a[:r+1]
	right := a[r+1:]

	quick_sort(left)
	quick_sort(right)
	return a
}


func sortArray(nums []int) []int {
	return quick_sort(nums)
}
