package main

func minDeletionSize(strs []string) (ans int) {
	next:
  for i := range len(strs[0]) {
		for j := range len(strs)-1 {
			if strs[j][i] > strs[j+1][i] {
				ans++
				continue next
			}
		}  
  }
	return
}
