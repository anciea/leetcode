package main

func maxOperations(nums []int, k int) int {
	count := make(map[int]int)
	ans := 0
	
	for _, num := range nums {
		comp := k - num
		if count[comp] > 0 {
			ans++
			count[comp]--
		} else {
			count[num]++
		}
	}
	
	return ans
}
