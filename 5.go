package main

func longestPalindrome(s string) string {
  maxLen, ans := 1, 0
  n := len(s)
  dp := make([][]bool, n)
  
  for i := range dp {
    dp[i] = make([]bool, n)
    dp[i][i] = true
  }

  for i := 1; i < n; i++ {
    for j := 0; i+j < n; j++ {
      if s[j] == s[i+j] {
        if i <= 2 {
          dp[j][i+j] = true
        } else {
          dp[j][i+j] = dp[j+1][i+j-1]
        }
      } else {
        dp[j][i+j] = false
      }
      if dp[j][i+j] && maxLen < i+1 {
        maxLen = i+1
        ans = j
      }
    }
  }

  return s[ans:ans+maxLen]
}
