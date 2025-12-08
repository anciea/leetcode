package main 

func countBadPairs(nums []int) int64 {
    mp := make(map[int]int)
    var res int64 = 0
    for i := 0; i < len(nums); i++ {
        key := nums[i] - i
        res += int64(i - mp[key])
        mp[key]++
    }
    return res
}
