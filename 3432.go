package main


func countPartitions(nums []int) int {
	sum, pre, ans:= 0, 0, 0

	for i := range nums {
		sum += nums[i]
	}

	for i := 0; i < len(nums)-1; i++ {
		pre += nums[i]
		sum -= nums[i]
		if (pre-sum) % 2 == 0 {
			ans++
		}
	}
	return ans
}

