package main

func permute(nums []int) [][]int {
  ans := make([][]int, 0)
  cnt := make([]int, 0)
  var dfs func()
  vis := make([]bool, len(nums))

  dfs = func() {
    if len(cnt) == len(nums) {
      tmp := append([]int{}, cnt...)
      ans = append(ans, tmp)
      return
    }

    for i := range len(nums) {
      if vis[i] {
        continue
      }
      cnt = append(cnt, nums[i])
      vis[i] = true

      dfs()

      vis[i] = false
      cnt = cnt[:len(cnt)-1]
    }
  }
  return ans
}
