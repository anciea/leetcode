package main

import (
	"slices"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	var dfs func(t *TreeNode, tmp int)
	dfs = func(t *TreeNode, tmp int) {
    if t == nil {
      return
    }
		if len(ans) < tmp {
			ans = append(ans, make([]int, 0))
		}
    // fmt.Println(tmp, len(ans))
		ans[tmp-1] = append(ans[tmp-1], t.Val)
		dfs(t.Left, tmp+1)
		dfs(t.Right, tmp+1)
	}
	dfs(root, 1)
  for i := range ans {
    if i % 2 == 1 {
      slices.Reverse(ans[i])
    }
  }


	return ans
}
