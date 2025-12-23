package main

// func lengthOfLIS(nums []int) (ans int) {
// 	n := len(nums)
// 	memo := make([]int, n)

// 	var dfs func(i int) int
// 	dfs = func(i int) (res int) {
// 		if memo[i] > 0 {
// 			return memo[i]
// 		}
// 		for j,x := range nums[:i] {
// 			if x < nums[i] {
// 				res = max(res, dfs(j))
// 			}
// 		}
// 		res++
// 		memo[i] = res
// 		return
// 	}

// 	for i := range n {
// 		ans = max(ans, dfs(i))
// 	}
// 	return 
// }

func lengthOfLIS(nums []int) (ans int) {
	n := len(nums)
	memo := make([]int, n)
	
	for i := range n {
		res := 0
		p := &memo[i]
		for j, x := range nums[:i] {
			if x < nums[i] {
				res = max(res, memo[j])
			}
		}
		res++
		*p = res
	}

	for i := range n {
		ans = max(ans, memo[i])
	}

	return ans
}
