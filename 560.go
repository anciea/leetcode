package main

func subarraySum(nums []int, k int) int {
	presum,ans := 0,0
	presumArray := make(map[int]int, 0)
	presumArray[0]=1
	for _,x := range nums {
		presum += x
		if y, ok := presumArray[presum-k]; ok {
			ans += y
		}
		presumArray[presum]++

	}
	return ans
}